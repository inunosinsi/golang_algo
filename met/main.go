package main

import "fmt"

type Character struct {
	hp      int
	attack  int
	defense int
}

//ダメージを受ける
func (c *Character) ReceiveDamage(n int) {
	c.hp = c.hp - n
}

func main() {
	yusha := &Character{hp: 100, attack: 5, defense: 5}
	yusha.ReceiveDamage(5)
	fmt.Println(yusha.hp)
}
