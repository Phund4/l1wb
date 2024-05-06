package main;

import (
	"fmt"
	"math/big"
)

// Сложение
func Addition(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int).Add(a, b);
	return res;
}

// Вычитание
func Subtraction(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int).Sub(a, b);
	return res;
}

// Умножение
func Multiplication(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int).Mul(a, b);
	return res;
}

// Деление
func Division(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int).Div(a, b);
	return res;
}

func main() {
	// Определение переменных
	a := big.NewInt(0);
	b := big.NewInt(0);

	// Присваиваем значение через setString
	a.SetString("7835698327636734639284", 10);
	b.SetString("823475813867183944679254", 10);

	fmt.Println(Addition(a, b))
	fmt.Println(Subtraction(a, b))
	fmt.Println(Multiplication(a, b))
	fmt.Println(Division(a, b))
}