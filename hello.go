package main

import (
	f "fmt"
	q "rsc.io/quote"
	t "time"
)

func main() {
	f.Println(q.Hello())
	f.Println(q.Go())
	f.Println(q.Glass())
	f.Println(q.Opt())
	f.Println("Olá mundo")
	f.Println(t.Now())
}
