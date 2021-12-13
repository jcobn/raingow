package utils

import (
	"fmt"
	"math"
)

func RGB(i int) (int, int, int) {
	f := 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func PrintRgbText(text string) {
	rArr := []rune(text)
	for j := 0; j < len(rArr); j++ {
		r, g, b := RGB(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, rArr[j])
	}
}
