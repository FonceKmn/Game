package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	Width    int32 = 800 // 20 ye bolende gelir dushur 40
	Height   int32 = 460 // 20 ye bolende gelir dushur 23
	GridSize int32 = 20
)

//-------------------------------Ilana Aid------------------------------------

// Define the Snake struct with Health, Position_x, and Position_y fields
type Snake struct {
	Health     uint
	Position_x int32
	Position_y int32
	Direction  int
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

func (s *Snake) draw() {
	rl.DrawRectangle(s.Position_x*GridSize, s.Position_y*GridSize, GridSize, GridSize, rl.Green)
}

//-------------------------------Ilana Aid------------------------------------

//-------------------------------Meyve----------------------------------------

// Define the Meyve struct with Position_x and Position_y fields
type Meyve struct {
	Position_x int32
	Position_y int32
}

func (m *Meyve) draw() {
	rl.DrawRectangle(m.Position_x*GridSize, m.Position_y*GridSize, 1*GridSize, 1*GridSize, rl.Red)
}

func (m *Meyve) Relocation() {
	m.Position_x = int32(rand.Intn(39))
	m.Position_y = int32(rand.Intn(22))
}

//-------------------------------Meyve----------------------------------------

func main() {
	rl.InitWindow(Width, Height, "raylib [core] example - Ilan Oyunu")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ilan := Snake{Health: 0, Position_x: 10, Position_y: 10}
	meyve := Meyve{Position_x: 14, Position_y: 15}

	var GameLoop bool = true
	for !rl.WindowShouldClose() && GameLoop {
		// Handle user input
		if rl.IsKeyPressed(rl.KeyUp) {
			ilan.Ilan_move_up()
			if ilan.Position_x < 0 || ilan.Position_x >= 40 || ilan.Position_y < 0 || ilan.Position_y >= 23 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			Ilan_move_down(&ilan)
			if ilan.Position_x < 0 || ilan.Position_x >= 40 || ilan.Position_y < 0 || ilan.Position_y >= 23 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyRight) {
			Ilan_move_right(&ilan)
			if ilan.Position_x < 0 || ilan.Position_x >= 40 || ilan.Position_y < 0 || ilan.Position_y >= 23 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyLeft) {
			Ilan_move_left(&ilan)
			if ilan.Position_x < 0 || ilan.Position_x >= 40 || ilan.Position_y < 0 || ilan.Position_y >= 23 {
				GameLoop = false
			} else if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
				meyve.Relocation()
				ilan.Health++
			}
		}
		if rl.IsKeyPressed(rl.KeyQ) {
			GameLoop = false
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// Draw the snake and Fruit
		ilan.draw()
		meyve.draw()
		//log.Println(meyve.Position_x, meyve.Position_y) //Debuger
		rl.EndDrawing()
	}

	// Optional: Display a message before closing
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Oyun Bitdi Siz Meglub Oldunuz...", 190, 200, 30, rl.Red)
	rl.EndDrawing()
	time.Sleep(2 * time.Second) // Wait for 2 seconds to display the message
}
