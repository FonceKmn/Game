package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

//-------------------------------Ilana Aid---------------------------------------

// Define the Snake struct with Health, Position_x, and Position_y fields
type Snake struct {
	Health     uint
	Position_x int
	Position_y int
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

/*
	func NewSnake(hp uint, p_x uint, p_y uint) *Snake {
		return &Snake{
			Health:     hp,
			Position_x: p_x,
			Position_y: p_y,
		}
	}

ilan := NewSnake(1,6,5)
*/
//-------------------------------Ilana Aid------------------------------------
//-------------------------------Meyve----------------------------------------
type Meyve struct {
	Position_x int
	Position_y int
}

func (meyve *Meyve) Relocation() {
	meyve.Position_x = rand.IntN(30)
	meyve.Position_y = rand.IntN(15)
}

//-------------------------------Meyve----------------------------------------
//-------------------------------Xerite---------------------------------------

type Map struct {
	width  int
	height int
}

func (m *Map) draw(x int, y int, fx int, fy int) {

	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			if (i == 0 || i == m.height-1) || (j == 0 || j == m.width-1) {
				fmt.Print("*")
			} else if i == y && j == x {
				fmt.Print("A")
			} else if i == fy && j == fx {
				fmt.Print("F")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

//-------------------------------Xerite---------------------------------------

func main() {

	//Deckleration og Fruit
	var Fruit Meyve
	Fruit.Position_x = 7
	Fruit.Position_y = 4
	//Deckleration of time
	interval := 1 * time.Second
	// Declare a variable of type Snake
	var ilan Snake
	// Initialize the Snake struct fields
	ilan.Health = 1
	ilan.Position_x = 6
	ilan.Position_y = 10
	// Print the coordinates and health of the snake
	fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)

	// Declare a map for drawing
	var xerite Map
	xerite.height = 15
	xerite.width = 30
	xerite.draw(ilan.Position_x, ilan.Position_y, Fruit.Position_x, Fruit.Position_y)
	var Game_loop bool
	Game_loop = true

	input_esas := make(chan rune)
	go func() {
		for {
			//time.Sleep(1 * time.Second)
			var keyboard_input rune
			fmt.Scanf("%c", &keyboard_input)
			input_esas <- keyboard_input
		}
	}()

	go func() {
		for Game_loop {
			//time.Sleep(1 * time.Second)
			switch <-input_esas {
			case 'a':
				Ilan_move_left(&ilan)
				if ilan.Position_x == 0 || ilan.Position_x == xerite.width-1 || ilan.Position_y == 0 || ilan.Position_y == xerite.height-1 {
					Game_loop = false
				} else {
					fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
					xerite.draw(ilan.Position_x, ilan.Position_y, Fruit.Position_x, Fruit.Position_y)
				}
			case 'w':
				Ilan_move_down(&ilan)
				if ilan.Position_x == 0 || ilan.Position_x == xerite.width-1 || ilan.Position_y == 0 || ilan.Position_y == xerite.height-1 {
					Game_loop = false
				} else {
					fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
					xerite.draw(ilan.Position_x, ilan.Position_y, Fruit.Position_x, Fruit.Position_y)
				}
			case 's':
				ilan.Ilan_move_up()
				if ilan.Position_x == 0 || ilan.Position_x == xerite.width-1 || ilan.Position_y == 0 || ilan.Position_y == xerite.height-1 {
					Game_loop = false
				} else {
					fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
					xerite.draw(ilan.Position_x, ilan.Position_y, Fruit.Position_x, Fruit.Position_y)
				}
			case 'd':
				Ilan_move_right(&ilan)
				if ilan.Position_x == 0 || ilan.Position_x == xerite.width-1 || ilan.Position_y == 0 || ilan.Position_y == xerite.height-1 {
					Game_loop = false
				} else {
					fmt.Printf("Ilanin Kordinatlari %d, %d Cani: %d\n", ilan.Position_x, ilan.Position_y, ilan.Health)
					xerite.draw(ilan.Position_x, ilan.Position_y, Fruit.Position_x, Fruit.Position_y)
				}
			case 'q':
				Game_loop = false
			}

		}

	}()
	time.Sleep(100 * interval)
	fmt.Println("Game Over")
}
