package main

func main() {
	// Khoi tao mot doi tuong persion
	p := Persion{
		Name: "Nguyen Van A",
		Age:  20,
	}
	print(p.Age)
	p.SetName("Nguyen Van B")
	print(p.Name)

}

type Persion struct {
	Name string
	Age  int
}

func (p *Persion) SetName(name string) {
	p.Name = name
}

type Abser interface {
	Abs() float64
}
