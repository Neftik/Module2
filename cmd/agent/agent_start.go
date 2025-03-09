package main

import (
	"log"

	"github.com/Neftik/Module2/internal/agent"
)

func main() {
	agent := agent.NewAgent()
	log.Println("Запускаем агент")
	agent.Start()
}
