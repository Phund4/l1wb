package main

import "fmt"

// Тип данных для set
type void struct{};

// Значение по индексу в map
var member void;

//  Структура set
type Set struct {
	set map[string]void
}

// Добавление в set
func (s *Set) addToSet(key string) {
	s.set[key] = member;
}

// Массив в set
func arrayToSet(arr []string) Set {
	res := Set{set: map[string]void{}}
	for _, el := range arr {
		res.addToSet(el);
	}
	return res;
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := arrayToSet(arr)
	fmt.Println(set);
}