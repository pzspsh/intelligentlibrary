/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:15:50
*/
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// A basic set of examples showing that %v is the default format, in this
	// case decimal for integers, which can be explicitly requested with %d;
	// the output is just what Println generates.
	integer := 23
	// Each of these prints "23" (without the quotes).
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// The special verb %T shows the type of an item rather than its value.
	fmt.Printf("%T %T\n", integer, &integer)
	// Result: int *int

	// Println(x) is the same as Printf("%v\n", x) so we will use only Printf
	// in the following examples. Each one demonstrates how to format values of
	// a particular type, such as integers or strings. We start each format
	// string with %v to show the default output and follow that with one or
	// more custom formats.

	// Booleans print as "true" or "false" with %v or %t.
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// Result: true true

	// Integers print as decimals with %v and %d,
	// or in hex with %x, octal with %o, or binary with %b.
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// Result: 42 42 2a 52 101010

	// Floats have multiple formats: %v and %g print a compact representation,
	// while %f prints a decimal point and %e uses exponential notation. The
	// format %6.2f used here shows how to set the width and precision to
	// control the appearance of a floating-point value. In this instance, 6 is
	// the total width of the printed text for the value (note the extra spaces
	// in the output) and 2 is the number of decimal places to show.
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// Result: 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00

	// Complex numbers format as parenthesized pairs of floats, with an 'i'
	// after the imaginary part.
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// Result: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

	// Runes are integers but when printed with %c show the character with that
	// Unicode value. The %q verb shows them as quoted characters, %U as a
	// hex Unicode code point, and %#U as both a code point and a quoted
	// printable form if the rune is printable.
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// Result: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	// Strings are formatted with %v and %s as-is, with %q as quoted strings,
	// and %#q as backquoted strings.
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// Result: foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	// Maps formatted with %v show keys and values in their default formats.
	// The %#v form (the # is called a "flag" in this context) shows the map in
	// the Go source format. Maps are printed in a consistent order, sorted
	// by the values of the keys.
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// Result: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	// Structs formatted with %v show field values in their default formats.
	// The %+v form shows the fields by name, while %#v formats the struct in
	// Go source format.
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// Result: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	// The default format for a pointer shows the underlying value preceded by
	// an ampersand. The %p verb prints the pointer value in hex. We use a
	// typed nil for the argument to %p here because the value of any non-nil
	// pointer would change from run to run; run the commented-out Printf
	// call yourself to see.
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))
	// Result: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// Result: &{Kim 22} 0x010203 // See comment above.

	// Arrays and slices are formatted by applying the format to each element.
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// Result: [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// Result: [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}

	// Byte slices are special. Integer verbs like %d print the elements in
	// that format. The %s and %q forms treat the slice like a string. The %x
	// verb has a special form with the space flag that puts a space between
	// the bytes.
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// Result: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

	// Types that implement Stringer are printed the same as strings. Because
	// Stringers return a string, we can print them using a string-specific
	// verb such as %q.
	now := time.Unix(123456789, 0).UTC() // time.Time implements fmt.Stringer.
	fmt.Printf("%v %q\n", now, now)
	// Result: 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"

}
