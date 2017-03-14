package main

import (
	"fmt"
  "strings"
	"bufio"
	"os"
)
// arbol binario con valores enteros.
type Arbol struct {
	Izquierda  *Arbol
	Valor string
	Derecha *Arbol
}

func New() *Arbol {
	return &Arbol{}
}

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []string
	count int
}

func (s *Stack) Push(valor string) {
	s.nodes = append(s.nodes[:s.count], valor)
	s.count++
}

func (s *Stack) Pop() string {
	if s.count == 0 {
		return ""
	}
	s.count--
	return s.nodes[s.count]
}

func insertar(s *Stack, t *Arbol) *Arbol {
	for i:=0; i < len(s.nodes);i++{
		if t.Valor == "" {		
			t=&Arbol{nil, s.Pop(), nil}
			fmt.Println("raiz: ",t.Valor)
		} else if t.Izquierda == nil{
			t.Izquierda=&Arbol{nil, s.Pop(), nil}
			fmt.Println("izq: ",t.Izquierda.Valor)
		}else if t.Derecha == nil{
			t.Derecha=&Arbol{nil, s.Pop(), nil}
			fmt.Println("der: ",t.Derecha.Valor)
		}
	}
	return t
}

func RecorrerInorden(t *Arbol) {
	if t == nil {
		return
	}
	RecorrerInorden(t.Izquierda)
	fmt.Print(t.Valor)
  fmt.Print(" - ")
	RecorrerInorden(t.Derecha)
}

func RecorrerPreorden(t *Arbol) {
	if t == nil {
		return
	}
  fmt.Print(t.Valor)
  fmt.Print(" - ")
	RecorrerPreorden(t.Izquierda)
	RecorrerPreorden(t.Derecha)
}

func InsertarExp(arr []string) *Stack{
	st:=NewStack()
	for _, r := range arr{
		st.Push(r)
	}
	return st
}

func main() {
	scanner:= bufio.NewScanner(os.Stdin)
  scanner.Scan()
  op:=scanner.Text()
	ar:=strings.Split(op," ")
	fmt.Println(ar)
	t:=New()
	fmt.Println(t.Valor)
	s:=InsertarExp(ar)
	fmt.Println(s.nodes)
	insertar(s, t)


}
