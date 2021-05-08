package main

import "log"

func print(s []int) {
	// duljina -> koliko imam trenutno elemenata
	// kapacitet -> koliko memorije zauzimam trenutno, koliko ukupno mogu primiti elemenata
	log.Printf("elements: %v, len=%v, cap=%v\n", s, len(s), cap(s))
}

func main() {
	// statički array -> rijetko kad se koriste
	array := [5]int{1, 2, 3, 4, 5}
	print(array[:])

	// pointer na array ili dio arrayja
	slice := array[0:3]
	print(slice)

	// append vraća pointer na novi array
	// slice mijenja  i originalni array
	slice = append(slice, 6)
	print(slice)
	print(array[:])

	// dinamičko alociranje slicea, zadamo tip i veličinu kao parametre
	emptySlice := make([]int, 5)
	print(emptySlice)

	// kad nema dovoljno mjesta, alocira se još toliko memorije (veličine 5 -> dobijemo kapacitet 10)
	emptySlice = append(emptySlice, 1)
	print(emptySlice)

	// treći argument može biti kapacitet
	emptySliceWithExtraCap := make([]int, 5, 10)
	print(emptySliceWithExtraCap)

	//////////////////////////////////////////

	// dinamički slice jer nismo zadali veličinu u  []
	sliceToIterate := []int{50, 60, 70, 80}

	// uvijek vraća indeks, vrijednost
	for idx, val := range sliceToIterate {
		log.Printf("%v: %v\n", idx, val)
	}

	// točno -> kada nas nije briga za indeks
	for _, val := range sliceToIterate {
		log.Printf("%v\n", val)
	}

	// pazi, tu vraća samo indeks
	for val := range sliceToIterate {
		log.Printf("%v\n", val)
	}

	// za iteriranje od zadnjeg do prvog nema builtin načina -> preko for petlje
}
