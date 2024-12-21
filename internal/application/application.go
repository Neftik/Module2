package application

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Andreyka-coder9192/calc_go/pkg/calculation"
)

type Config struct {
	Addr string
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Wrong Method"}`, http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var request struct {
		Expression string `json:"expression"`
	}
	err = json.Unmarshal(body, &request)
	if err != nil || request.Expression == "" {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		var errorMsg string
		statusCode := http.StatusUnprocessableEntity

		switch err {
		case calculation.ErrInvalidExpression:
			errorMsg = "Error calculation"
		case calculation.ErrDivisionByZero:
			errorMsg = "Division by zero"
		case calculation.ErrMismatchedParentheses:
			errorMsg = "Mismatched parentheses"
		case calculation.ErrInvalidNumber:
			errorMsg = "Invalid number"
		case calculation.ErrUnexpectedToken:
			errorMsg = "Unexpected token"
		case calculation.ErrNotEnoughValues:
			errorMsg = "Not enough values"
		case calculation.ErrInvalidOperator:
			errorMsg = "Invalid operator"
		case calculation.ErrOperatorAtEnd:
			errorMsg = "Operator at end"
		case calculation.ErrMultipleDecimalPoints:
			errorMsg = "Multiple decimal points"
		case calculation.ErrEmptyInput:
			errorMsg = "Empty input"
		default:
			errorMsg = "Error calculation"
			statusCode = http.StatusUnprocessableEntity
		}

		http.Error(w, fmt.Sprintf(`{"error":"%s"}`, errorMsg), statusCode)
		return
	}

	response := struct {
		Result string `json:"result"`
	}{
		Result: fmt.Sprintf("%v", result),
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error while marshaling response: %v", err)
		http.Error(w, `{"error":"Unknown error occurred"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJson)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

func (a *Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("application was successfully closed")
			return nil
		}
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, "calculation failed with error:", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error":"Not Found"}`, http.StatusNotFound)
	})

	return http.ListenAndServe(":"+a.config.Addr, nil)
}
