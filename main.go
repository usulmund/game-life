package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	rows = 30
	cols = 30
)

type Field [][]bool

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func EmptyField() Field {
	f := make(Field, rows+2)
	for i := range f {
		f[i] = make([]bool, cols+2)
	}
	return f
}

func (f Field) CreateRandomPopulation() Field {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			f[i][j] = rand.Intn(20) == 0
		}
		fmt.Println()
	}
	return f
}

func (f Field) FillBoard() Field {
	f[0] = f[rows]
	f[rows+1] = f[1]
	for i := 0; i < rows+2; i++ {
		f[i][0] = f[i][cols]
		f[i][cols+1] = f[i][1]
	}
	return f
}

func NewField() Field {
	f := EmptyField().CreateRandomPopulation()
	return f
}

func (f Field) Show() {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if f[i][j] {
				fmt.Printf("■ ")
			} else {
				fmt.Printf("□ ")
			}
		}
		fmt.Println()
	}
}

func StrLen(s string) int {
	strLen := 0
	for i := range s {
		strLen++
		i = i
	}
	return strLen
}

func (f Field) NeighborsCount(x, y int) int {
	neighborsCount := 0
	if f[x+1][y] {
		neighborsCount++
	}
	if f[x-1][y] {
		neighborsCount++
	}
	if f[x][y+1] {
		neighborsCount++
	}
	if f[x][y-1] {
		neighborsCount++
	}
	if f[x+1][y+1] {
		neighborsCount++
	}
	if f[x-1][y-1] {
		neighborsCount++
	}
	if f[x-1][y+1] {
		neighborsCount++
	}
	if f[x+1][y-1] {
		neighborsCount++
	}
	return neighborsCount
}

func (f Field) Alive(x, y int) bool {
	neighborsCount := f.NeighborsCount(x, y)
	return neighborsCount == 2 || neighborsCount == 3
}

func (f Field) NewLife(x, y int) bool {
	return f.NeighborsCount(x, y) == 3
}

func (f Field) NextStep() Field {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if !f[i][j] && f.NewLife(i, j) {
				f[i][j] = true
			}
		}
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if !f.Alive(i, j) {
				f[i][j] = false
			}
		}
	}

	return f
}

func (f Field) GetCntLiveCells() (cntLiveCells int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if f[i][j] {
				cntLiveCells++
			}
		}
	}
	return
}

func RunGame(iterCount int) {
	f := NewField()
	iter := 1
	for {
		ClearScreen()
		f.Show()
		f = f.NextStep().FillBoard()

		time.Sleep(100 * time.Millisecond)
		iter++
		if iterCount != -1 && iter > iterCount {
			fmt.Printf("Iter count: %d\nCount of living cells: %d\n", iterCount, f.GetCntLiveCells())
			break
		}
	}
}

func main() {
	var err error
	var iterCount int = -1
	f := NewField()
	f.Show()
	if len(os.Args) == 2 {
		iterCount, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Error: %s, enter other iter count\n", err)
			return
		}
	}
	RunGame(iterCount)
}
