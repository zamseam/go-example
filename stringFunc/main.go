package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("count:", s.Count("test", "t"))

	p("HasSuffix:", s.HasSuffix("test", "st"))

	p("Index:", s.Index("test", "e"))

	p("join:", s.Join([]string{"a", "b"}, "-"))

	p("repeat:", s.Repeat("a", 5))

	p("Replace:", s.Replace("foo", "o", "0", -1))

	p("Replace:", s.Replace("foo", "o", "0", 1))

	p("split:", s.Split("a-b-c-d-e", "-"))

	p("ToLower:", s.ToLower("TEST"))

	p("ToUpper:", s.ToUpper("test"))

	p()
	p("len:", len("hello"))
	p("Char:", "hello"[1])
}