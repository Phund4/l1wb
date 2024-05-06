package main

import "fmt"

// Тип данных для set
type void struct{};

// Значение по индексу в map
var member void;

//  Структура set
type Set struct {
	set map[int]void
}

// Добавление в set
func (s *Set) addToSet(key int) {
	s.set[key] = member;
}

// Удаление из set
func (s *Set) removeFromSet(key int) {
	delete(s.set, key);
}


// Пересечение двух set
func Intersection(set1 Set, set2 Set) Set {
	// Объявление результата пересечения
	res := Set{set: make(map[int]void)};

	// Проходимся по однрму из set
	for key := range set1.set {
		_, ok := set2.set[key];
		if ok {
			res.addToSet(key);
		}
	}

	return res;
}

func main() {
	// Создание set1
	set1 := Set{set: make(map[int]void)}
	set1.addToSet(1);
	set1.removeFromSet(3);
	set1.addToSet(3);
	set1.removeFromSet(1);
	set1.addToSet(4);
	set1.addToSet(5);
	set1.addToSet(6);
	fmt.Println(set1)

	// Создание set2
	set2 := Set{set: make(map[int]void)}
	set2.addToSet(2);
	set2.addToSet(10);
	set2.addToSet(4);
	set2.addToSet(9);
	set2.addToSet(6);
	fmt.Println(set2)

	res := Intersection(set1, set2);
	fmt.Println(res);
}