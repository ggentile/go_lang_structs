package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func main() {
	list := List{}

	tiago := Pessoa{"Tiago", "Leão", 45}
	joao := Pessoa{"João", "Santos", 20}
	maria := Pessoa{"Maria", "dos Santos", 33}
	marcos := Pessoa{"Marcos", "Paulo", 21}
	daniel := Pessoa{"Daniel", "Marcutti", 10}

	list.Append(tiago)
	list.Append(joao)
	list.Append(maria)
	list.Append(marcos)
	list.Append(daniel)

	list.Display()

	fmt.Println("----------------------------------")

	pessoa := list.Search("Maria")
	fmt.Println(pessoa)

	list.Delete("Daniel")

	list.Display()
}
