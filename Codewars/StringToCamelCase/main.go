package main

import (
	"fmt"
	"strings"
)

// https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go

func main() {
	y := "Ninja_bridge_south_Rockstar_left_Yellow_Wall"
	if strings.ToUpper(string(y[0])) == string(y[0]) {
		// Starts with a cap letter
		if strings.ReplaceAll(y, "-", "") == y {
			g := strings.Split(strings.Title(y), "_")
			for i := range g {
				g[i] = strings.Title(g[i])
			}
			fmt.Println(strings.Join(g, ""))
			// s := strings.ReplaceAll(strings.Title(y), "_", "")
			// fmt.Println(strings.Title(s))
		} else {
			s := strings.ReplaceAll(strings.Title(y), "-", "")
			fmt.Println(strings.Title(s))
		}
	} else {
		m := strings.Split(strings.Title(y), "-")
		if len(m) == 1 {
			r := strings.Split(strings.Title(y), "_")
			r[0] = strings.ToLower(r[0])
			for i := range r {
				if i == 0 {
					continue
				}
				r[i] = strings.Title(r[i])
			}
			fmt.Println(strings.Join(r, ""))
		} else {
			fmt.Println(m)
			m[0] = strings.ToLower(m[0])
			fmt.Println(strings.Join(m, ""))

		}

	}

}

// Title everything, extract the first and lower, jpin all three back
