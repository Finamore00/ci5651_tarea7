package main

import "fmt"

/*
Implementación en Golang del precondicionamiento del algoritmo KMP para búsqueda
de patrones dentro de strings que nos permite encontrar la subcadena más larga de
un string S, T, tal que T es tanto prefijo como sufijo de S. 7ma tarea de Diseño
de Algoritmos I (CI5651). Universidad Simón Bolívar. Trimestre Ene-Mar 2024

Autor: Santiago Finamore
Carné: 18-10125
*/

func computeLPS(s string) string {
	lps := make([]int, len(s))
	lps[0] = 0
	lenV := 0

	i := 1
	for i < len(s) {
		if s[i] == s[lenV] {
			lenV += 1
			lps[i] = lenV
			i += 1
		} else {
			if lenV != 0 {
				lenV = lps[lenV-1]
			} else {
				lps[i] = 0
				i += 1
			}
		}
	}
	return s[:lps[len(lps)-1]]
}

func main() {
	s := "abracadabra"
	l := computeLPS(s)
	fmt.Println(l)
}
