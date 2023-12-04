package main

type Base struct{}

func (bs Base) String() string {
	return "base"
}

type Struct struct {
	*Base // ポインタとして埋め込む
}

func main() {
	s := &Struct{}

	println(s.String()) // panic
}
