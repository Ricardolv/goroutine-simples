package main

import (
	"fmt"
	"time"
)

// O select permite que você aguarde múltiplas operações de channels.
// É útil quando você precisa lidar com múltiplos channels simultaneamente.

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Mensagem do channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Mensagem do channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
