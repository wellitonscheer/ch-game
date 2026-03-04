package main

import (
	"fmt"
	"time"
)

type Screen struct {
	frame  [][]string
	height int
	width  int
}

func NewScreen() *Screen {
	return &Screen{
		frame:  [][]string{},
		height: 35,
		width:  150,
	}
}

func lineOfCharacters(times int, char string) []string {
	line := []string{}

	for i := 0; i < times; i++ {
		line = append(line, char)
	}

	return line
}

func (s *Screen) addLinesToFrame(lines int, char string) {
	for i := 0; i < lines; i++ {
		line := lineOfCharacters(s.width, char)
		s.frame = append(s.frame, line)
	}
}

func (s *Screen) Render() {
	s.addLinesToFrame(1, "-")
	s.addLinesToFrame(s.height, "a")
	s.addLinesToFrame(1, "-")
}

func (s *Screen) Display() {
	backFrame := ""
	for _, line := range s.frame {
		rn := ""
		for _, pixel := range line {
			rn = fmt.Sprintf("%s%s", rn, pixel)
		}
		backFrame = fmt.Sprintf("%s%s\n", backFrame, rn)
	}
	fmt.Print(backFrame)
}

func main() {
	ticker := time.NewTicker(time.Second / 60)

	screen := NewScreen()
	screen.Render()

	for {
		select {
		case <-ticker.C:
			screen.Display()
		}
	}
}
