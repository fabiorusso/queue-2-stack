package queue

import "queue2stack/stack"

type Queue interface {
	Enfileirar(x string)
	Desenfileirar() string
	Imprimir() string
}

func NewQueue() Queue {
	return &queue{
		stack1: stack.Stack{},
		stack2: stack.Stack{},
	}
}

type queue struct {
	stack1 stack.Stack
	stack2 stack.Stack
}

func transfer(source, dest *stack.Stack) {
	for ; !source.IsEmpty(); {
		dest.Push(source.Pop())
	}
}

func (q *queue) Enfileirar(x string) {
	q.stack1.Push(x)
}

func (q *queue) Desenfileirar() string {
	if !q.stack2.IsEmpty() {
		return q.stack2.Pop()
	} else if !q.stack1.IsEmpty() {
		transfer(&q.stack1, &q.stack2)
		return q.Desenfileirar()
	} else {
		return ""
	}
}

func (q *queue) Imprimir() string {
	if !q.stack2.IsEmpty() {
		el:=  q.stack2.Pop()
		q.stack2.Push(el)
		return el
	} else {
		if !q.stack1.IsEmpty() {
			transfer(&q.stack1, &q.stack2)
			return q.Imprimir()
		} else {
			return ""
		}
	}
}
