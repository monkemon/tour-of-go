// very punk solution to a stringer excercise from the tour of go
// https://go.dev/tour/methods/18

/*
Make the IPAddr type implement fmt.Stringer to print the address as a dotted quad.

For instance, IPAddr{1, 2, 3, 4} should print as "1.2.3.4".
*/




func (ip IPAddr) String() string {
	out:= fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
	return out
}
