/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 14:10:26
*/
package main

import (
	"fmt"
	"math/big"
)

// Use the classic continued fraction for e
//
//	e = [1; 0, 1, 1, 2, 1, 1, ... 2n, 1, 1, ...]
//
// i.e., for the nth term, use
//
//	   1          if   n mod 3 != 1
//	(n-1)/3 * 2   if   n mod 3 == 1
func recur(n, lim int64) *big.Rat {
	term := new(big.Rat)
	if n%3 != 1 {
		term.SetInt64(1)
	} else {
		term.SetInt64((n - 1) / 3 * 2)
	}

	if n > lim {
		return term
	}

	// Directly initialize frac as the fractional
	// inverse of the result of recur.
	frac := new(big.Rat).Inv(recur(n+1, lim))

	return term.Add(term, frac)
}

// This example demonstrates how to use big.Rat to compute the
// first 15 terms in the sequence of rational convergents for
// the constant e (base of natural logarithm).
func main() {
	for i := 1; i <= 15; i++ {
		r := recur(0, int64(i))

		// Print r both as a fraction and as a floating-point number.
		// Since big.Rat implements fmt.Formatter, we can use %-13s to
		// get a left-aligned string representation of the fraction.
		fmt.Printf("%-13s = %s\n", r, r.FloatString(8))
	}

}
