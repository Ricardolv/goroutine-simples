package main

import (
	"log/slog"
	"time"
)

// Para criar uma goroutine, basta usar a palavra-chave go
// antes da chamada de uma função. Isso faz com que a função seja executada em uma goroutine separada.

func printNumbers() {
	for i := 1; i <= 5; i++ {
		slog.Info("loop", "i:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Executa printNumbers em uma goroutine

	// Executa outra tarefa no main goroutine
	for i := 0; i < 2; i++ {
		slog.Info("Main goroutine")
		time.Sleep(1000 * time.Millisecond)
	}

	// Dá tempo para a goroutine terminar
	time.Sleep(3 * time.Second)
}
