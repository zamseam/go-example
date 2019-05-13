package main

import "flag"
import "fmt"

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")

	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	f := fmt.Println
	f("word:", *wordPtr)
	f("numb:", *numbPtr)
	f("fork:", *boolPtr)
	f("svar:", svar)
	f("tail:", flag.Args())
}