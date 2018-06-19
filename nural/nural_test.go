package nural

import (
	"fmt"
	"testing"
)

func testFunction(a []interface{}) interface{} {
	return a[0]
}

func TestNur(t *testing.T) {
	n := []nur{
		nur{
			testFunction, 0,
			make([]interface{}, 0),
			make(chan interface{}, 2),
			[]*nur{},
			[]*nur{},
		},
		nur{
			testFunction, 0,
			make([]interface{}, 0),
			make(chan interface{}, 2),
			[]*nur{},
			[]*nur{},
		},
	}
	n[0].child(&n[1])
	go n[0].fin(0)
	go n[1].wait()
	go fmt.Println(<-n[1].res)
}
