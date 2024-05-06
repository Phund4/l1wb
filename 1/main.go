package main;

import (
	"fmt"
)

type Human struct {
	name string
}

func (h *Human) getName() string {
	return h.name;
}

type Action struct {
	// Embedding структуры Human в Action
	Human
}

func main() {
	h := Human{name: "hoho"};
	a := Action{Human: h}; // Необходимо указать используемый экземпляр Human
	fmt.Println(a.getName()); // hoho
	fmt.Println(h.getName()); // hoho
}