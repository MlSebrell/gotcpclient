package main

import "fmt"

func main() {

	// unlike arrays slices are typed only by
	// the elements they contain 
	// not the number of elements.
	// to create an empty slice with non-zero
	// length, use the builtin `make`
	// here we make a slice of `strings`s of length
	// `3` (intially zero-values)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get just like the arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as 
	// expected.
	fmt.Println("len:", len(s))

	/* In addition to these basic operations,
	slices support several more that make them
	richer than arrays. One is a builtin `append`,
	which returns a slice containing one or more
	new values

	Note that we need to accept a return value
	from append as we may get a new slice value. */ 

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	
	// slices can also be `copy`d. here we create an
	// empty slice `c` of the same length as `s` and 
	// copy it into `c` from `s` . 
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	
	//slices support a "slice" operator with the syntax
	// `slice[low:high]`. for example, this gets a slice
	// of the elements `s[2]`. `s[3]`, and `s[4]`.
	l := s[2:5]
	fmt.Println("sl1:", l)
	
	// This slices up to (but excluding) `s[5]`.
	l = s[:5]
	fmt.Println("sl2:", l)
	
	// and this slices up from (and including) `s[2]`.
	l = s[2:]
	fmt.Println
	
