package main

import (
	"fmt"
	"sync"
	"time"
)

// Outra abordagem é usar um semáforo para limitar o número de goroutines ativas ao mesmo tempo.
// Isso pode ser implementado usando um canal com buffer.

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	sem <- struct{}{}        // Ocupa um slot no semáforo
	defer func() { <-sem }() // Libera o slot no semáforo

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second) // Simula trabalho
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	const numWorkers = 10
	const maxConcurrency = 3

	sem := make(chan struct{}, maxConcurrency) // Semáforo com limite de concorrência
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait() // Espera todas as goroutines terminarem
	fmt.Println("Todos os workers terminaram")
}

// O canal sem atua como um semáforo, limitando o número de goroutines ativas ao mesmo tempo.
// Cada goroutine ocupa um slot no semáforo antes de começar o trabalho e libera o slot ao terminar.
