package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		z01.PrintRune('A')
		for ligne1 := 1; ligne1 < x-1; ligne1++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
		for ligne2 := 1; ligne2 < y-1; ligne2++ {
			z01.PrintRune('B')

			for espace := 2; espace <= x-1; espace++ {
				z01.PrintRune(' ')

			}
			z01.PrintRune('B')
			z01.PrintRune('\n')
		}

		z01.PrintRune('A')
		for ligne3 := 1; ligne3 < x-1; ligne3++ {
			z01.PrintRune('B')
		}

		z01.PrintRune('C')
		z01.PrintRune('\n')
	}
}
