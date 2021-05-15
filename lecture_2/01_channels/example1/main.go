package main

import "fmt"

func main() {
	// ako nemamo buffer -> dvije dretve moraju u isto vrijeme raditi, jedna čita, druga piše
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println()

	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println()

	// pri pokušaju pisanja četvrte poruke, dretva se blokira
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Print(<-ch)
	fmt.Println()
}
