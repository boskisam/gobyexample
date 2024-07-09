// W Go, _array_ jest numerowaną sekwencją elementów o
// konkretnej długości. W typowym kodzie Go, [slices](slices) są
// dużo częściej spotykane; array jest użyteczny w pewnych
// specyficznych sytuacjach.

package main

import "fmt"

func main() {

	// Tutaj tworzymy array `a``, który będzie przechowywał
	// dokładnie 5 `int`ów. Typ elementów oraz ich długość są
	// częścią typu array. Domyślnie wartości w array są zerowe,
	// co dla `int`ów oznacza `0`.
	var a [5]int
	fmt.Println("emp:", a)

	// Możemy ustawić wartość konkretnego indexu używająć
	// syntaxu `array[index] = wartość`, i otrzymać wartość poprzez
	// `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// Wbudowana funkcja `len` zwraca długość array.
	fmt.Println("len:", len(a))

	// Możemy użyć poniższego syntaxu aby zdeklarować oraz
	// zainicjować array w jednej linii.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Możemy też wykorzystać kompilator aby zliczył ilość elementów
	// wpisując `...`.
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Jeśłi wyspecyfikujemy index z `:`, elementy pomiędzy
	// będą wyzerowane.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// Typy array'ów są jednowymiarowe, ale możemy
	// komponować typy aby stworzyć wielowymiarowe struktury
	// danych.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Możemy tworzyć i inicjalizować wielowymiarowe
	// array'e także za jednym razem.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
