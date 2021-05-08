package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

// defer stavlja na stack -> znači izvršavaju se obrnutim redoslijedom -> dobijemo obrnuti ispis
func countToTen(f *os.File) {
	for i := 10; i >= 0; i-- {
		defer f.WriteString(fmt.Sprint(i) + "\n")
	}
}

// defer može primiti i više naredbi -> pomoću lambde func() { ... }
// '_' -> kad nas nije briga za vraćenu vrijednost

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	// defer -> izvedi ovo prije nego što završi metoda (čak i ako se dogodi neki error)
	// najbolje pozvati odmah nakon otvaranja datoteke
	defer f.Close()
	countToTen(f)
	f.Sync()
}
