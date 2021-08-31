package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	a := "abc"
	b := strings.ToUpper(a)
	strings.ToLower(b)

	sparrow := &Sparrow{}
	goose := &Goose{}
	phoenix := &Phoenix{}
	bird := &Bird{
		Sparrow: sparrow,
		Goose:   goose,
		Phoenix: phoenix,
	}

	fmt.Println("1.=============")
	Fly(bird, BirdCatSparrow)
	Fly(bird, BirdCatGoose)
	Fly(bird, BirdCatPhoenix)

	fmt.Println("2.=============")

	Fly1(sparrow)
	Fly1(goose)
	Fly1(phoenix)

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""

	}
}

type Bird struct {
	Sparrow *Sparrow
	Goose   *Goose
	Phoenix *Phoenix
}

const (
	BirdCatSparrow = iota
	BirdCatGoose
	BirdCatPhoenix
)

func Fly(bird *Bird, category int) {
	if category == BirdCatSparrow {
		bird.Sparrow.Fly()
	} else if category == BirdCatGoose {
		bird.Goose.Fly()
	} else {
		bird.Phoenix.Fly()
	}
}

type Flyer interface {
	Fly()
}

func Fly1(flyer Flyer) {
	flyer.Fly()
}

type Sparrow struct {
}

func (p *Sparrow) Fly() {
	fmt.Println("麻雀飞的太慢")
}

type Goose struct {
}

func (p *Goose) Fly() {
	fmt.Println("大雁飞的挺快")
}

type Phoenix struct {
}

func (p *Phoenix) Fly() {
	fmt.Println("凤凰飞的超级快")
}
