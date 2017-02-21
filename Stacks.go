package main

import "fmt"

type Node struct{
  Nombre string
  Id int
  Hora string
  Diagnostico string
  Eps string
}

type Stack struct{
  nodes []*Node
  count int
}

func CrearStack() *Stack{
  return &Stack{}
}

func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (n *Node) String() {
	fmt.Println(n.Eps, "-> Nombre: " ,n.Nombre, "Diagnostico: ", n.Diagnostico)
}

func main(){
  var name,time,diagn,eps,epss,add string
  var id int
  s := CrearStack()
  fmt.Println("Agregar Nuevo Paciente? S/N")
  fmt.Scanln(&add)
  for add=="S" || add=="s"{
    fmt.Println("Agregar nuevo ")
    fmt.Println("Nombre: ")
    fmt.Scanln(&name)
    fmt.Println("ID: ")
    fmt.Scanln(&id)
    fmt.Println("Hora de Ingreso: ")
    fmt.Scanln(&time)
    fmt.Println("Diagnostico: ")
    fmt.Scanln(&diagn)
    fmt.Println("EPS: ")
    fmt.Scanln(&eps)
    s.Push(&Node{name, id, time, diagn, eps})
    fmt.Println("Agregar Nuevo Paciente? S/N")
    fmt.Scanln(&add)
  }
  fmt.Println("Mostrar Por Eps: ")
	fmt.Scanln(&epss)
  for s.count>0 {
    paciente:=s.Pop()
    if paciente.Eps==epss{
      paciente.String()
    }
  }
}
