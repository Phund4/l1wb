/*

Стоит начать с неконстантной глобальной переменной, объявленной в коде. Такая практика не приветствуется

Далее в sumFunc мы ссылаемся на часть среза, тем самым наша ссылка
будет лежать в памяти, gb ее не уберет.

Необходимо при создании строки воспользоваться буфером

*/

package main

import (
	"bytes"
	"fmt"
)

var justString string

func createHugeString(stringLen int) string {
	// Инициализация буфера
	var buf bytes.Buffer

	// Заполнение буфера
	for i := 0; i < stringLen; i++ {
		buf.WriteString("a")
	}

	// Переводим буфер в string
	return buf.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
	fmt.Println(justString);
}