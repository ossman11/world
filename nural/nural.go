package nural

import (
	"fmt"
)

// Operation of a nuron to calculate any input to any output
type operation func([]interface{}) interface{}

type nur struct {
	op  operation
	run bool
	don chan bool
	res interface{}
	in  []*nur
	out []*nur
}

// Checks if the nur exists inside the slice of nurs
func contains(a []*nur, c *nur) bool {
	for _, v := range a {
		if v == c {
			return true
		}
	}
	return false
}

// Checks if the nur is a child
func (n *nur) isChild(c *nur) bool {
	return contains(n.out, c)
}

// Checks if the nur is a parent
func (n *nur) isParent(c *nur) bool {
	return contains(n.in, c)
}

// Adds the nur as a child
func (n *nur) child(c *nur) {
	if n.isChild(c) {
		return
	}
	n.out = append(n.out, c)
	c.parent(n)
}

// Adds the nur as a parent
func (n *nur) parent(c *nur) {
	if n.isParent(c) {
		return
	}
	n.in = append(n.in, c)
	c.child(n)
}

// Finishes a nur
// and propegates the result to all the children
func (n *nur) fin(v interface{}) {
	n.res = v
	for _, c := range n.out {
		go c.exec()
	}
	close(n.don)
}

// Executes the op of the nur
func (n *nur) exec() {
	if n.run {
		return
	}
	n.run = true
	val := make([]interface{}, len(n.in))
	for i, c := range n.in {
		<-c.don
		val[i] = c.res
	}
	n.fin(n.op(val))
}

func (n *nur) read() {
	<-n.don
	fmt.Println(n.res)
}

func (n *nur) reset() {
	n.don = make(chan bool)
	n.run = false
}

func NewNur(op operation) *nur {
	r := &nur{
		op,
		false,
		nil,
		nil,
		[]*nur{},
		[]*nur{},
	}
	r.reset()
	return r
}

type net struct {
	in  []nur
	out []nur
}
