package main

import "sync"

type HashTable struct {
	Items map[int][]Pessoa
	lock  sync.RWMutex
}

func hash(nome string) (key int) {
	for _, letra := range nome {
		key = 31*key + int(letra)
	}

	return
}

func (ht *HashTable) Put(pessoa Pessoa) int {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	key := hash(pessoa.Nome)
	if ht.Items == nil {
		ht.Items = make(map[int][]Pessoa)
	}
	ht.Items[key] = append(ht.Items[key], pessoa)

	return key
}

func (ht *HashTable) Remove(nome string) {
	ht.lock.Lock()
	defer ht.lock.Unlock()

	key := hash(nome)
	delete(ht.Items, key)
}

func (ht *HashTable) Get(key int) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()

	return ht.Items[key]
}

func (ht *HashTable) Search(nome string) []Pessoa {
	ht.lock.RLock()
	defer ht.lock.RUnlock()

	key := hash(nome)

	return ht.Items[key]
}
