package main

import (
	"log"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/rand"
)

const (
	Width          int32   = 800 // 20 ye bolende gelir dushur 40
	Height         int32   = 460 // 20 ye bolende gelir dushur 23
	GridSize       int32   = 20  // GridSize Teyin edirik
	UpdateInterval float64 = 0.1 // yarim saniye
)

//-------------------------------Ilana Aid------------------------------------

// Define the Snake struct with Health, Position_x, and Position_y fields

type Points struct { // Ilan qaqamizin kordinatlari ucun hoqqa veririk
	x int32
	y int32
}

type Snake struct {
	Health     int
	Position_x int32
	Position_y int32
	Direction  int32
	Body       []Points
}

func (s *Snake) RemoveTail() { // Ilan qaqamizin quyrugunu arrayinin 0. indexini silir Head Len(bodyden) - ile iterasiya olunmalidi
	s.Body = s.Body[1:len(s.Body)]
}

func (s *Snake) AddTail(x, y int32) { // Ilan qaqamizin quyrugu add olunur surushur effekti yaradir.
	s.Body = append(s.Body, Points{x, y}) // ilana body elave eliyir

	if s.Health < len(s.Body) { // eger ilanin bodysinin leni Healtshdan kicikdise silir.
		s.RemoveTail()
	}
}

func (snake *Snake) Ilan_move_up() {
	snake.Position_y--
	snake.AddTail(snake.Position_x, snake.Position_y)
}

func Ilan_move_down(ilan *Snake) {
	ilan.Position_y++
	ilan.AddTail(ilan.Position_x, ilan.Position_y)
}

func Ilan_move_right(ilan *Snake) {
	ilan.Position_x++
	ilan.AddTail(ilan.Position_x, ilan.Position_y)
}

func Ilan_move_left(ilan *Snake) {
	ilan.Position_x--
	ilan.AddTail(ilan.Position_x, ilan.Position_y)
}

func (s *Snake) move() {
	switch s.Direction {
	case 1:
		s.Ilan_move_up()
	case 2:
		Ilan_move_down(s)
	case 3:
		Ilan_move_right(s)
	case 4:
		Ilan_move_left(s)
	}
}

func (s *Snake) draw() {
	rl.DrawRectangle(s.Position_x*GridSize, s.Position_y*GridSize, GridSize, GridSize, rl.Red)

	for i := len(s.Body) - 1; i >= 0; i-- {
		rl.DrawRectangle(s.Body[i].x*GridSize, s.Body[i].y*GridSize, GridSize, GridSize, rl.Green)
	}
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

	rl.SetTargetFPS(30)

	ilan := Snake{Health: 0, Position_x: 10, Position_y: 10, Direction: 1}
	meyve := Meyve{Position_x: 14, Position_y: 15}

	lastMoveTime := time.Now() // updateInterval ucun
	var GameLoop bool = true

	for !rl.WindowShouldClose() && GameLoop {
		// Handle user input
		if rl.IsKeyPressed(rl.KeyUp) && ilan.Direction != 2 {
			ilan.Direction = 1
		}
		if rl.IsKeyPressed(rl.KeyDown) && ilan.Direction != 1 {
			ilan.Direction = 2
		}
		if rl.IsKeyPressed(rl.KeyRight) && ilan.Direction != 4 {
			ilan.Direction = 3
		}
		if rl.IsKeyPressed(rl.KeyLeft) && ilan.Direction != 3 {
			ilan.Direction = 4
		}

		if time.Since(lastMoveTime).Seconds() >= UpdateInterval {
			ilan.move()
			lastMoveTime = time.Now()
		}

		if rl.IsKeyPressed(rl.KeyQ) {
			GameLoop = false
		}

		// Ilan Meyve Yedikce Boyuyur
		if ilan.Position_x == meyve.Position_x && ilan.Position_y == meyve.Position_y {
			meyve.Relocation()
			ilan.Health++
		}

		// Ilanin Collisionu Divarla
		if ilan.Position_x < 0 || ilan.Position_x >= Width/GridSize || ilan.Position_y < 0 || ilan.Position_y >= Height/GridSize {
			GameLoop = false
		}

		// Ilanin collusionu Ozuyle (2 olanda sonuncunu nezere almiriq (saniyede 3 defe check eliyir FPS 30 du deye))
		if len(ilan.Body) > 2 {
			for i := 0; i < len(ilan.Body)-1; i++ {
				if ilan.Position_x == ilan.Body[i].x && ilan.Position_y == ilan.Body[i].y {
					GameLoop = false
					break
				}
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		// Draw the snake and Fruit
		ilan.draw()
		meyve.draw()
		//log.Println(ilan.Health) // Debugger
		log.Println(ilan.Position_x, ilan.Position_y, ilan.Body)
		rl.EndDrawing()
	}

	// Optional: Display a message before closing
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Oyun Bitdi Siz Meglub Oldunuz...", 190, 200, 30, rl.Red)
	rl.EndDrawing()
	time.Sleep(2 * time.Second) // Wait for 2 seconds to display the message
}
