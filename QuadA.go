package piscine

import (
	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x >= 0 && y >= 0 {
		z01.PrintRune('o')
		for ligne1 := 1; ligne1 < x-1; ligne1++ {
			z01.PrintRune(45)

			for ligne3 := 1; ligne3 < x-1; ligne3++ {
				z01.PrintRune(45)
			}
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
	}
}
