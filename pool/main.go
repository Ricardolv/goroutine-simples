package main

import (
	"fmt"
	"sync"
	"time"
)

// Você pode criar um pool de goroutines onde um número fixo de goroutines processa tarefas de um canal.
// O sync.WaitGroup é usado para garantir que todas as goroutines terminem antes de o programa principal continuar.

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Notifica que a goroutine terminou

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simula trabalho
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	const numJobs = 10
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Inicia o pool de workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// Envia tarefas para o canal de jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Fecha o canal para indicar que não há mais jobs

	wg.Wait() // Espera todas as goroutines terminarem
	fmt.Println("Todos os jobs foram processados")

	// Um canal (jobs) é usado para enviar tarefas para as goroutines.
	// Um número fixo de goroutines (numWorkers) é criado para processar as tarefas.
	// O sync.WaitGroup é usado para sincronizar a finalização das goroutines.
}
