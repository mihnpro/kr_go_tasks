package main

import "fmt"

type PepeSchnele struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele {
	return &PepeSchnele{
		Speed:    speed,
		Charisma: charisma,
		Wisdom:   wisdom,
	}
}

func (p *PepeSchnele) GetRating() int {
	return (p.Speed * 2) + (p.Charisma * 3) + p.Wisdom
}

func (p *PepeSchnele) String() string {
	return fmt.Sprintf("Пепе Шнеле [Скорость: %d, Харизма: %d, Мудрость: %d] | Рейтинг: %d",
		p.Speed, p.Charisma, p.Wisdom, p.GetRating())
}

func main() {
	pepe1 := NewPepeSchnele(10, 20, 30)
	pepe2 := NewPepeSchnele(5, 15, 25)

	fmt.Println(pepe1.String())
	fmt.Println(pepe2.String())

	fmt.Println("Результат сравнения:")
	if pepe1.GetRating() > pepe2.GetRating() {
		fmt.Println(pepe1.String())
	} else {
		fmt.Println(pepe2.String())
	}
}

// 1.
// Создайте структуру PepeSchnele с тремя полями: Speed, Charisma, Wisdom (тип int для всех).

// 2.
// Напишите функцию NewPepeSchnele(speed, charisma, wisdom int) *PepeSchnele, которая создает и возвращает указатель на нового Пепе Шнеле с заданными характеристиками.

// 3.
// Напишите метод GetRating() для структуры PepeSchnele, который вычисляет и возвращает общий рейтинг по формуле: (Speed * 2) + (Charisma * 3) + Wisdom.

// 4.
// "Пепе Шнеле [Скорость: X, Харизма: Y, Мудрость: Z] | Рейтинг: R", где X, Y, Z — значения полей, а R — результат работы GetRating().

// В функции main() создайте как минимум двух разных Пепе Шнеле, выведите информацию о каждом с помощью fmt.Println и определите, у кого рейтинг выше.
