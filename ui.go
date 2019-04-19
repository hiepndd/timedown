package main

import (
	"fmt"
	"os"
	"unicode/utf8"

	termbox "github.com/nsf/termbox-go"
)

type Symbol []string

func (symbol Symbol) width() int {
	return utf8.RuneCountInString(symbol[0])
}

func (symbol Symbol) height() int {
	return len(symbol)
}

type Text []Symbol

func (text Text) width() int {
	w := 0
	for _, symbol := range text {
		w += symbol.width()
	}
	return w
}

func (text Text) height() int {
	return len(text[0])
}

func toText(str string) Text {
	symbols := make(Text, 0)
	for _, r := range str {
		if s, ok := defaultFont[r]; ok {
			symbols = append(symbols, s)
		}
	}
	return symbols
}

type Font map[rune]Symbol

func echo(symbol Symbol, pointA, pointB int) {
	x, y := pointA, pointB
	for _, line := range symbol {
		for _, r := range line {
			termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
			x++
		}
		x = pointA
		y++
	}
}

func clear() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}
}

func flush() {
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func stderr(s string, a ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, s, a...)
	if err != nil {
		panic(err)
	}
}
