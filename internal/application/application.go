package application

import (
	"bufio"
	"encoding/json"
	"fmt"
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

	var request struct {
		Expression string `json:"expression"`
	}

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.Expression == "" {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		var errorMsg string
		switch {
		case err == calculation.ErrInvalidExpression:
			errorMsg = "Error calculation"
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, errorMsg), http.StatusUnprocessableEntity)
			return
		default:
			http.Error(w, `{"error":"Unknown error occurred"}`, http.StatusInternalServerError)
			return
		}
	}

	response := struct {
		Result string `json:"result"`
	}{
		Result: fmt.Sprintf("%v", result),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
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
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
