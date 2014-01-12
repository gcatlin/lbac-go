package main

import (
	"fmt"
	"os"
	"unicode"
)

const TAB = '	'

type char rune

var Look char

func Read(c *char) {
	_, _ = fmt.Scanf("%c", c)
}

func Halt() {
	os.Exit(1)
}

func GetChar() {
	Read(&Look)
	// fmt.Printf("<<< %#v\n", string(Look))
}

func Error(s string) {
	fmt.Println()
	fmt.Println("Error: " + s + ".")
}

func Abort(s string) {
	Error(s)
	Halt()
}

func Expected(s string) {
	Abort(s + " Expected")
}

func Match(x char) {
	if Look == x {
		GetChar()
	} else {
		Expected("\"" + string(x) + "\"")
	}
}

func IsAlpha(c char) bool {
	return unicode.IsLetter(rune(c))
}

func IsDigit(c char) bool {
	return unicode.IsDigit(rune(c))
}

func GetName() char {
	if !IsAlpha(Look) {
		Expected("Name")
	}
	c := unicode.ToUpper(rune(Look))
	GetChar()
	return char(c)
}

func GetNum() char {
	if !IsDigit(Look) {
		Expected("Integer")
	}
	c := Look
	GetChar()
	return c

}

func Emit(s string) {
	fmt.Print(string(TAB) + s)
}

func EmitLn(s string) {
	Emit(s)
	fmt.Println()
}

func Init() {
	GetChar()
}

func main() {
	Init()
}
