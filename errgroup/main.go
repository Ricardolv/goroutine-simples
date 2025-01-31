package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func worker(ctx context.Context, id int) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // Retorna se o contexto for cancelado
	default:
		fmt.Printf("Worker %d started\n", id)
		time.Sleep(time.Second) // Simula trabalho
		fmt.Printf("Worker %d finished\n", id)
		return nil
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := 1; i <= 5; i++ {
		id := i
		g.Go(func() error {
			return worker(ctx, id)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Erro: %v\n", err)
	} else {
		fmt.Println("Todos os workers terminaram com sucesso")
	}
}

// O errgroup gerencia um grupo de goroutines e permite que você lide com erros de forma centralizada.
// O contexto (ctx) pode ser usado para cancelar todas as goroutines em caso de erro.
