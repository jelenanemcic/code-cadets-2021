package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// kontekst se zatvara nakon timeouta
	// kreiramo novi kontekst iz osnovnog (Background)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)

	// vraća funkciju koja kad se pozove zatvara kontekst -> eksplicitno ga zatvaramo
	// ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		t := time.Now()
		fmt.Println("separate goroutine started")
		// ctx.Done() vraća channel, ako je kontekst gotov channel se zatvara
		<-ctx.Done()
		fmt.Println("separate goroutine done", time.Since(t))
	}()

	// ovdje eksplicitno zatvaramo kontekst withCancel
	// cancel()

	fmt.Println("main goroutine sleeping")
	time.Sleep(time.Second * 5)
	fmt.Println("main goroutine wake up")
}
