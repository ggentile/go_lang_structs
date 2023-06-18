package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Sexo      string
}

func main() {
	pessoas := []Pessoa{
		{"Maria", "Da Graça", 50, "F"},
		{"Maria", "Da Silva", 30, "F"},
		{"Tiago", "Feiketi", 29, "M"},
		{"Paulo", "Ecom", 76, "M"},
	}

	table := HashTable{}

	keys := make([]int, len(pessoas))
	for i, pessoa := range pessoas {
		keys[i] = table.Put(pessoa)
	}

	for _, key := range keys {
		ps := table.Get(key)
		for _, p := range ps {
			fmt.Println(p.Nome, p.Sobrenome)
		}
	}
}
