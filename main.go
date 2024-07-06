package main

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//-------------------------------Ilana Aid------------------------------------

// Ilanin Health ve x,y pozisyalari
type Snake struct {
	Health     uint
	Position_x int
	Position_y int
}

func (snake *Snake) Ilan_move_up() {
	snake.Position_y--
}

func Ilan_move_down(ilan *Snake) {
	ilan.Position_y++
}

func Ilan_move_right(ilan *Snake) {
	ilan.Position_x++
}

func Ilan_move_left(ilan *Snake) {
	ilan.Position_x--
}

func (s *Snake) draw(cellSize float32) {
	rl.DrawRectangle(int32(float32(s.Position_x)*cellSize), int32(float32(s.Position_y)*cellSize), int32(cellSize), int32(cellSize), rl.Green)
}

//-------------------------------Ilana Aid------------------------------------

//-------------------------------Meyve----------------------------------------

// Define the Meyve struct with Position_x and Position_y fields
type Meyve struct {
	Position_x int
	Position_y int
}

func (m *Meyve) draw(cellSize float32) {
	rl.DrawRectangle(int32(float32(m.Position_x)*cellSize), int32(float32(m.Position_y)*cellSize), int32(cellSize), int32(cellSize), rl.Red)
}

func (m *Meyve) Relocation() {
	m.Position_x = rand.Intn(15)
	m.Position_y = rand.Intn(30)
}

//-------------------------------Meyve----------------------------------------

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rand.Seed(time.Now().UnixNano())

	cellSize := float32(rl.GetScreenHeight()) / 20.0

	ilan := Snake{Health: 0, Position_x: 10, Position_y: 10}
	meyve := Meyve{Position_x: 14, Position_y: 15}
	/*
		if ilan.Position_x == 0 || ilan.Position_x == xerite.width-1 || ilan.Position_y == 0 || ilan.Position_y == xerite.height-1 {
			Game_loop = false
		} else if ilan.Position_x == Fruit.Position_x && ilan.Position_y == Fruit.Position_y {
			Fruit.Relocation()
			ilan.Health++
		}
	*/
	var GameLoop bool = true
	for !rl.WindowShouldClose() && GameLoop {
		// Handle user input
		if rl.IsKeyPressed(rl.KeyUp) {
			ilan.Ilan_move_up()
			if ilan.Position_x == 0 || ilan.Position_x == rl.GetScreenWidth()-1 || ilan.Position_y == 0 || ilan.Position_y == rl.GetScreenHeight()-1 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			Ilan_move_down(&ilan)
			if ilan.Position_x == 0 || ilan.Position_x == rl.GetScreenWidth()-1 || ilan.Position_y == 0 || ilan.Position_y == rl.GetScreenHeight()-1 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			Ilan_move_right(&ilan)
			if ilan.Position_x == 0 || ilan.Position_x == rl.GetScreenWidth()-1 || ilan.Position_y == 0 || ilan.Position_y == rl.GetScreenHeight()-1 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			Ilan_move_left(&ilan)
			if ilan.Position_x == 0 || ilan.Position_x == rl.GetScreenWidth()-1 || ilan.Position_y == 0 || ilan.Position_y == rl.GetScreenHeight()-1 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// Meyve ve Ilani cekmek ucun.
		ilan.draw(cellSize)
		meyve.draw(cellSize)
		//rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	// Optional: Display a message before closing
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Oyun Bitdi Siz Meglub Oldunuz...", 190, 200, 20, rl.Red)
	rl.EndDrawing()
	time.Sleep(2 * time.Second) // Yazini gostermek ucun 2 saniye gozlemek ucun
