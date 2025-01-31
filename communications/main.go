package main

import (
	"fmt"
	"time"
)

// Goroutines podem se comunicar usando channels.
// Channels são canais de comunicação que permitem que goroutines enviem e recebam dados.

func sendData(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		ch <- fmt.Sprintf("Mensagem %d", i)
		time.Sleep(500 * time.Millisecond)
	}
	close(ch) // Fecha o channel após enviar todos os dados
}

func main() {
	ch := make(chan string)

	go sendData(ch)

	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("Channel fechado")
}
