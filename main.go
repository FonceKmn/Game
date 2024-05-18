package main

import "fmt"

// Define the Snake struct with Health, Position_x, and Position_y fields
type Snake struct {
	Health     int
	Position_x int
	Position_y int
}

func main() {
	// Declare a variable of type Snake
	var ilan Snake

	// Initialize the Snake struct fields
	ilan.Health = 1
	ilan.Position_x = 6
	ilan.Position_y = 10
	// Print the coordinates and health of the snake
	fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
	var Game_loop bool
	Game_loop = true
	for Game_loop {

		var keyboard_input rune
		fmt.Scanf("%c", &keyboard_input)
		switch keyboard_input {
		case 'a':
			Ilan_move_left(&ilan)
			fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
		case 's':
			Ilan_move_down(&ilan)
			fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
		case 'w':
			ilan.Ilan_move_up()
			fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
		case 'd':
			Ilan_move_right(&ilan)
			fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
		case 'q':
			Game_loop = false
		}

	}
	fmt.Println("Game Over")
}

// Belede Cagirmaq olar
func (snake *Snake) Ilan_move_up() {
	snake.Position_y++
}

// Ilanin Gezme Funksiyalari

func Ilan_move_down(ilan *Snake) {
	ilan.Position_y--
}

func Ilan_move_right(ilan *Snake) {
	ilan.Position_x++
}

func Ilan_move_left(ilan *Snake) {
	ilan.Position_x--
}
