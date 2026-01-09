package main

import (
	"fmt"
	"strings"
)

// Некий певец оставляет послания, которые нужно защищать. Он использует специальный шифратор, работающий по следующим правилам:

// Фраза разбивается на отдельные слова (разделитель — пробел).

// Для каждого слова создается строка-шифр по алгоритму:

// Первый символ слова остается на своем месте.

// Все остальные символы слова реверсируются (переставляются в обратном порядке).

// Полученные зашифрованные слова собираются обратно в одну строку, разделенные пробелами.

// Пример:
// Исходная фраза: "Pepe Schnele is a legend"

// Шаг разбивки: ["Pepe", "Schnele", "is", "a", "legend"]

// Шифрование каждого слова:

// Pepe -> P + epe(реверс) -> P + epe -> Pepe (если символы повторяются, может совпасть)

// Schnele -> S + chnele(реверс) -> S + elenhc -> Selenhc

// is -> i + s(реверс) -> i + s -> is

// a -> a (слово из одного символа остается a)

// legend -> l + egend(реверс) -> l + dnege -> ldnege

// Итоговая строка: "Pepe Selenhc is a ldnege"
// Ваша задача:

// 1. Напишите функцию encryptWord(word string) string, которая шифрует одно слово по описанному алгоритму.

// 2. Напишите функцию encryptPhrase(phrase string) string, которая:

//     a) Разбивает фразу на слова (можно использовать strings.Fields).

//     b) Для каждого слова вызывает encryptWord.

//     c) Объединяет результаты в одну строку.

// В функции main() протестируйте шифратор на нескольких фразах, включая пример выше. Выведите исходные фразы и результат их шифрования.

func encryptWord(word string) string {
	r := []rune(word)

	if len(r) < 3 {
		return word
	}

	var withoutFirst []rune

	withoutFirst = append(withoutFirst, r[0])

	for i := len(r) - 1; i >= 1; i-- {
		withoutFirst = append(withoutFirst, r[i])
	}

	return string(withoutFirst)
}

func encryptPhrase(phrase string) string {

	words := strings.Split(phrase, " ")

	var encrypted []string
	for _, word := range words {
		encrypted = append(encrypted, encryptWord(word))
	}

	return strings.Join(encrypted, " ")
}

func main() {
	phrase1 := "Pepe Schnele is a legend"

	phrase2 := "Hello World"

	phrase3 := "I am a cat"

	phrase4 := "a b c d e f g h i j k l m n o p q r s t u v w x y z"

	phrase5 := "Programming is awesome"

	phrase6 := ""

	phrase7 := " "

	phrase8 := "supercalifragilisticexpialidocious"


	fmt.Println(phrase1)
	fmt.Println(encryptPhrase(phrase1))

	fmt.Println(phrase2)
	fmt.Println(encryptPhrase(phrase2))

	fmt.Println(phrase3)
	fmt.Println(encryptPhrase(phrase3))

	fmt.Println(phrase4)
	fmt.Println(encryptPhrase(phrase4))

	fmt.Println(phrase5)
	fmt.Println(encryptPhrase(phrase5))

	fmt.Println(phrase6)
	fmt.Println(encryptPhrase(phrase6))

	fmt.Println(phrase7)
	fmt.Println(encryptPhrase(phrase7))

	fmt.Println(phrase8)
	fmt.Println(encryptPhrase(phrase8))

}
