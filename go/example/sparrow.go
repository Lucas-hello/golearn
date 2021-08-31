package main

func SparrowFly() {
	sp := &Sparrow{}
	bird := &Bird{
		Sparrow: sp,
	}
	Fly(bird, BirdCatSparrow)
}

func GooseFly() {
	goose := &Goose{}
	bird := &Bird{
		Goose: goose,
	}
	Fly(bird, BirdCatGoose)
}
