package main

import (
	"log"
	"queue2stack/queue"
)

func main() {

	qu:=queue.NewQueue()

	qu.Enfileirar("10")
	qu.Enfileirar("20")

	log.Printf("1: %s | 2: %s ", qu.Desenfileirar(), qu.Desenfileirar())

	qu.Enfileirar("13")
	qu.Enfileirar("14")
	qu.Enfileirar("15")
	log.Printf("|3: %s", qu.Desenfileirar())
	qu.Enfileirar("16")
	log.Printf("Primeiro: %s", qu.Imprimir())

}
