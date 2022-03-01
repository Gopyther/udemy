package main

import "fmt"

type player struct {
	salary int
}

type pro struct {
	player
	bonus int
}

func (p player) getsalary() int {
	return p.salary
}

func main() {
	player1 := player{salary: 100}
	proPlayer := pro{player{salary: 120}, bonus: 30}

	fmt.Println(player1.getsalary())
	fmt.Println(proPlayer.getsalary() + proPlayer.bonus)

}
