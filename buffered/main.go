package main

import (
	"fmt"
	"time"
)

// Channels podem ser buffered, o que significa
// que eles podem armazenar um número limitado de valores sem precisar de um receptor imediato.

func main() {
	ch := make(chan string, 2) // Buffer de 2

	go func() {
		ch <- "Mensagem 1"
		ch <- "Mensagem 2"
		time.Sleep(1 * time.Second)
		ch <- "Mensagem 3"
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
