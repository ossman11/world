package nural

// Operation of a nuron to calculate any input to any output
type operation func([]interface{}) interface{}

type nur struct {
	op   operation
	valC int
	val  []interface{}
	res  chan interface{}
	in   []*nur
	out  []*nur
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
	n.out = append(n.in, c)
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
	for i := 0; i < len(n.out); i++ {
		n.res <- v
	}
	close(n.res)
}

// Sets the val of the nur
func (n *nur) set(i int, v interface{}) {
	n.valC++
	n.val[i] = v
	if n.valC >= len(n.in) {
		n.exec()
	}
}

// Executes the op of the nur
func (n *nur) exec() {
	n.res <- n.op(n.val)
}

// Makes the nur wait for all its pars to be fin
func (n *nur) wait() {
	n.valC = 0
	n.val = make([]interface{}, len(n.in))
	for i, c := range n.in {
		n.set(i, <-c.res)
	}
}

type net struct {
	in  []nur
	out []nur
}
