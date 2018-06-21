package nural

import (
	"fmt"
	"testing"
)

func testFunction(a []interface{}) interface{} {
	r := 1
	for _, v := range a {
		r += v.(int)
	}
	fmt.Println(a)
	fmt.Println(r)
	return r
}

func TestNur(t *testing.T) {
	n := []nur{
		*NewNur(testFunction),
		*NewNur(testFunction),
		*NewNur(testFunction),
		*NewNur(testFunction),
	}
	n[0].parent(&n[3])
	n[0].parent(&n[2])
	n[0].parent(&n[1])

	n[1].parent(&n[3])
	n[1].parent(&n[2])

	n[2].parent(&n[3])

	n[3].fin(0)

	<-n[0].don
	fmt.Println("Last finished")
	n[0].read()
	n[1].read()
	n[2].read()
}
