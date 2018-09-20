package main

import (
	"fmt"
)

func copy_and_append_string(slice []string, elem string) []string {
	// wrong: return append(slice, elem)
	return append(append([]string(nil), slice...), elem)
}

func PowerSet(s []string) [][]string {
	if s == nil {
		return nil
	}
	r := [][]string{[]string{}}
	for _, es := range s {
		var u [][]string
		for _, er := range r {
			u = append(u, copy_and_append_string(er, es))
		}
		r = append(r, u...)
	}
	return r
}

func main() {

	p := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		p = append(p, string(i))
	}

	fmt.Println(PowerSet(p))
}
