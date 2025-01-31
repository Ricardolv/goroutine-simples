package main

import (
	"fmt"
	"time"

	"github.com/gammazero/workerpool"
)

// Se você preferir uma solução pronta,
// existem bibliotecas externas que implementam pools de goroutines, como o workerpool:

func main() {
	wp := workerpool.New(3) // Cria um pool com 3 workers

	for i := 1; i <= 10; i++ {
		id := i
		wp.Submit(func() {
			fmt.Printf("Worker %d started\n", id)
			time.Sleep(time.Second) // Simula trabalho
			fmt.Printf("Worker %d finished\n", id)
		})
	}

	wp.StopWait() // Espera todas as tarefas terminarem
	fmt.Println("Todos os workers terminaram")
}

// A biblioteca workerpool gerencia o pool de goroutines para você.
// Você só precisa enviar tarefas para o pool usando Submit.
