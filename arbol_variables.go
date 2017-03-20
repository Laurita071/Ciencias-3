//Laura Marcela Cucanchón - 20131020037
//Kevin parrado - 20131020033

package main

import (
  "fmt"
  "strconv"
  "bufio"
  "os"
  "strings"
)

var v *StackV

type Variable struct {
  notac string
  valor string
}

type Arbol struct{
  Izquierda *Arbol
  Valor string
  Derecha *Arbol
}

type Stack struct {
	nodes []*Arbol
	count int
}

type StackV struct {
	nodes []*Variable
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func NewStackV() *StackV {
	return &StackV{}
}

func (s *Stack) Push(n *Arbol) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Arbol {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func (s *StackV) PushV(n *Variable) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *StackV) PopV() *Variable {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

func RecorrerInorden(t *Arbol, inorden string) string{
  if t == nil{
    return inorden
  }else{
    inorden=RecorrerInorden(t.Izquierda, inorden)
    inorden=inorden+t.Valor
    inorden=RecorrerInorden(t.Derecha, inorden)
  }
  return inorden
}

func CalcularArbol(t *Arbol) int{
  if t.Izquierda!=nil && t.Derecha!=nil{
    switch t.Valor{
      case "+":
          return CalcularArbol(t.Izquierda) + CalcularArbol(t.Derecha)
      case "-":
          return CalcularArbol(t.Izquierda) - CalcularArbol(t.Derecha)
      case "*":
          return CalcularArbol(t.Izquierda) * CalcularArbol(t.Derecha)
      case "/":
          return CalcularArbol(t.Izquierda) / CalcularArbol(t.Derecha)
    }
  }
  i,e:=strconv.Atoi(t.Valor)
  if(e != nil){
    return valorVar(t.Valor)
  }
  return i
}

func valorVar(vlr string) int{
  for _, r:= range v.nodes {
    if r.notac==vlr {
      x,_:=strconv.Atoi(r.valor)
      return x
    }else{
      return 0
    }
  }
  return 0
}

func preevalComp(t *Arbol) string{
  if t == nil{
    return "No existe ningun arbol para evaluar\n"
  } else if t.Izquierda == nil && t.Derecha == nil{
    return "El nodo "+ t.Valor + " no tiene hijos\n"
  } else if t.Izquierda !=nil && t.Derecha ==nil {
    return "El arbol se encuentra incompleto, el nodo "+t.Valor+" no tiene hijo derecho\n"
  } else if t.Izquierda ==nil && t.Derecha !=nil{
    return "El arbol se encuentra incompleto, el nodo "+ t.Valor+" no tiene hijo izquierdo\n"
  } else {
    return preevalComp(t.Izquierda)+""+preevalComp(t.Derecha)
  }
  return " "
}

func operarPostfijo(op string) string{
  var x string
  c:=strings.Split(op, " ")
  x, c = c[len(c)-2], c[:len(c)-2]
  //v.PushV(&Variable{x,0})
  fmt.Println("C: ",c)
  if(len(c)<=2){
    return "Error, no existen suficientes datos"
  }
  s := NewStack()
  var contador int = 0
  for i:=0; i< len(c); i++{
    if(c[i]=="+" || c[i]=="-" || c[i]=="*" || c[i]=="/"){
      contador++
      if(s.count<2){
        return "Error, no existe la cantidad necesaria de valores para calcular"
      }else{
        s.Push(&Arbol{nil,c[i],nil})
        var arbol1 *Arbol = s.Pop()
        arbol1.Derecha = s.Pop()
        arbol1.Izquierda = s.Pop()
        var resultado = CalcularArbol(arbol1)
        s.Push(&Arbol{nil, strconv.Itoa(resultado),nil})
      }
    }else{
      s.Push(&Arbol{nil,c[i],nil})
    }
  }
  if(contador==0){
    return "Error, no existe expresión en el árbol dado"
  }
  vFinal:=s.Pop().Valor
  v.PushV(&Variable{x,vFinal})

  return (x + " = " + vFinal)
}

func main(){
  fmt.Println("Continuar? s/n")
  scanner:= bufio.NewScanner(os.Stdin)
  scanner.Scan()
  sgr:= scanner.Text()
  v=NewStackV()
  for sgr=="s" || sgr=="S" {
    fmt.Println("Inserte la expresión en posOrden")
    scanner.Scan()
    op:=scanner.Text()
    fmt.Println(operarPostfijo(op))
    for _,r:= range v.nodes{
      fmt.Println(r)
    }
    fmt.Println("Continuar? s/n")
    scanner.Scan()
    sgr=scanner .Text()
  }
}
