package main

import "fmt"

type Entry struct {
	Name   string
	Suname string
	Year   int
}

func ZeroS() Entry {
	return Entry{}
}

func InitS(N, S string, y int) Entry {
	if y < 2000 {
		return Entry{Name: N, Suname: S, Year: 2000}
	}
	return Entry{Name: N, Suname: S, Year: y}
}

func zeroPos() *Entry {
	t := &Entry{}
	return t
}

func initPos(N, S string, y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Suname: "Unknown", Year: y}
	}
	return &Entry{Name: N, Suname: S, Year: y}
}

func main() {
	s1 := ZeroS()
	p1 := zeroPos()
	fmt.Println("s1: ", s1, "p1: ", p1)

	s2 := InitS("John", "Doe", 1990)
	p2 := initPos("John", "Doe", 1990)

	fmt.Println("s2: ", s2, "p2: ", p2)
	fmt.Println("year: ", s1.Year, p1.Year, s2.Year, p2.Year)

	ps := new(Entry)
	fmt.Println("ps: ", ps)
}
