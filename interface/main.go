package main

import "fmt"

// インターフェースは設計図
type Animal interface {
	Bark()
	Sleep()
}

type Dog struct {
	Name string
}

// Dogの構造体がAnimalのインターフェースで指定しているBark()を持つことでAnimalとして振る舞える
func (d *Dog) Bark() {
	fmt.Println(d.Name + " Wan!")
}

//Animalインターフェースの条件を満たしてい場合に実行出来るコード
func DoBark(a Animal) {
	a.Bark()
	a.Sleep()
}

func main() {
	dog := &Dog{Name: "pochi"}

	/**
	 * Dog構造体はAnimalインターフェースの条件を満たしていたので、
	 * Animal型の変数に代入することが出来る
	 */
	var animal Animal
	animal = dog

	DoBark(animal)
}
