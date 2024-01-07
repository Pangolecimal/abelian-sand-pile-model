package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH  = 1000
	HEIGHT = WIDTH
	NUM    = 100
	SIZE   = WIDTH / NUM
	GAP    = 0
	SKIP   = 1
)

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "")
	rl.SetWindowPosition(200, 1440/2-HEIGHT/2)
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	var grid [NUM][NUM]int

	for j := 0; j < NUM; j++ {
		for i := 0; i < NUM; i++ {
			grid[j][i] = 0

		}
	}
	grid[NUM/2][NUM/2] = int(math.Pow(2, 12))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0x00_00_00_ff))
		for j := 0; j < NUM; j++ {
			for i := 0; i < NUM; i++ {
				x := int32(i * SIZE)
				y := int32(j * SIZE)

				v := grid[j][i]
				col := uint(0x00_00_00_ff)
				if v == 1 {
					col = 0x555555_ff
				}
				if v == 2 {
					col = 0xaaaaaa_ff
				}
				if v == 3 {
					col = 0xffffff_ff
				}
				if v > 3 {
					col = 0xf00000_ff
				}

				rl.DrawRectangle(x, y, SIZE, SIZE, rl.GetColor(col))
				// rl.DrawText(fmt.Sprint(value), x+GAP/4+SIZE/4, y+GAP/4+SIZE/4, SIZE/3, GetGray(uint8(col+128)))
			}
		}
		rl.EndDrawing()

		for i := 0; i < SKIP; i++ {
			update(&grid)
		}
	}
}

func update(grid *[NUM][NUM]int) {
	var grid_next [NUM][NUM]int
	for j := 1; j < NUM-1; j++ {
		for i := 1; i < NUM-1; i++ {
			if grid[j][i] == 0 {
				continue
			}

			if grid[j][i] <= 3 {
				grid_next[j][i] += grid[j][i]
				continue
			}

			add := grid[j][i] / 4
			grid_next[j][i] = grid[j][i] - add*4
			grid_next[j+0][i+1] += add
			grid_next[j+0][i-1] += add
			grid_next[j+1][i+0] += add
			grid_next[j-1][i+0] += add
		}
	}
	*grid = grid_next
}

func GetGray(color uint8) color.RGBA {
	return rl.GetColor(uint(color)*0x01010100 + 0xff)
}
