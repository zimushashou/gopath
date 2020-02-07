package main

func main()  {
	e := NewQueue()
	e.Push(1)
	e.Push(5)
	e.Push(3)
	e.Push(7)
	e.Push(2)
	e.Print()
	e.Pop()
	e.Pop()
	e.Print()
}
