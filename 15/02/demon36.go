package main

type Dog struct {
	name string
}

func New(name string) Dog {
	return Dog{name}
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func main()  {
	//New("little pig").SetName("monster")
	//15\02\demon36.go:20:19: cannot call pointer method on New("little pig")
	//15\02\demon36.go:20:19: cannot take the address of New("little pig")
	map[string]int{"the": 0, "word": 0, "counter": 0}["word"]++
	map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
	map1["word"]++
}