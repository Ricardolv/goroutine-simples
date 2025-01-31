package main

import (
	"log/slog"
	"sync"
	"time"
)

// Goroutines podem precisar ser sincronizadas para garantir que uma goroutine termine antes que outra continue.
// Isso pode ser feito usando sync.WaitGroup.

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // Notifica que a goroutine terminou

	for i := 1; i <= 5; i++ {
		slog.Info("loop", "i:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // Adiciona uma goroutine ao WaitGroup
	go printNumbers(&wg)

	wg.Wait() // Espera todas as goroutines no WaitGroup terminarem
	slog.Info("Todas as goroutines terminaram")
}
