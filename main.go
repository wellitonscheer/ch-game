package main

import "fmt"

var screen [][]string

func main() {
	dp := "b"
	for i := 0; i < 30; i++ {
		if i < 1 {
			dp = "a"
		} else if i == 29 {
			dp = "b"
		}

		line := []string{}
		for i := 0; i < 100; i++ {
			line = append(line, dp)
		}
		screen = append(screen, line)
	}

	for {
		rs := ""
		for _, line := range screen {
			rn := ""
			for _, pixel := range line {
				rn = fmt.Sprintf("%s%s", rn, pixel)
			}
			rs = fmt.Sprintf("%s%s\n", rs, rn)
		}
		fmt.Printf("%s", rs)
	}
}
