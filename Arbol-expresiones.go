//Estudiantes:
//Laura Marcela Cucanchón - 20131020037
//Kevin parrado - 20131020033

package main

import (
  "fmt"
  "strconv"
  "bufio"
  "os"
)

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []*Arbol
	count int
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

type Arbol struct{
  Izquierda *Arbol
  Valor string
  Derecha *Arbol
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
  var i int
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
  i,_=strconv.Atoi(t.Valor)
  return i
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
  var c []string
  var st string = ""
  for _, r := range op {
    if(string(r) == " "){
        c = append (c,st)
        st = ""
    }else{
        st = st + string(r)
    }
  }
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
  valorFinal := (s.Pop().Valor)
  return (op + " = " + valorFinal)
}

func main(){
  fmt.Println("Inserte el árbol en posOrden")
  fmt.Println("Nota: Con espacios de por medio y al final")
  scanner:= bufio.NewScanner(os.Stdin)
  scanner.Scan()
  op:=scanner.Text()
  fmt.Println(operarPostfijo(op))
}
