package main

// ispis može i s "fmt", ali "log" bolji za ispis errora, pogotovo pri testiranju na serveru
import "log"

func main() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// kao while petlja
	i := 0
	for i <= 10 {
		log.Println(i)
		i++
	}

	// beskonačna petlja
	i = 0
	for {
		log.Println(i)
		i++
		if i == 10 {
			break
		}
	}
}
