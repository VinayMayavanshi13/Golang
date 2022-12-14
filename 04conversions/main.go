package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Program")
	fmt.Println("Please rate us between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// this input is a string you can't increment it by simply adding 1 as num can't be added to a string.
	// so you need a type conversion.

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}

// Error faced in this program using : numRating, err := strconv.ParseFloat(input, 64) at Line 23
//strconv.ParseFloat: parsing "3\r\n": invalid syntax
// here on hitting an Enter along with 3 \r i.e carriage return and \n i.e new line character is also get inserted.
//This can be solved by using strings.TrimSpace(input)

// All about ParseFloat

/*
ParseFloat(s string, bitSize int) (float64, error)
ParseFloat converts the string s to a floating-point number
with the precision specified by bitSize: 32 for float32, or 64 for float64.
When bitSize=32, the result still has type float64, but it will be
convertible to float32 without changing its value.

ParseFloat accepts decimal and hexadecimal floating-point number syntax.
If s is well-formed and near a valid floating-point number,
ParseFloat returns the nearest floating-point number rounded
using IEEE754 unbiased rounding.
(Parsing a hexadecimal floating-point value only rounds when
there are more bits in the hexadecimal representation than
will fit in the mantissa.)

The errors that ParseFloat returns have concrete type *NumError
and include err.Num = s.

If s is not syntactically well-formed, ParseFloat returns err.Err = ErrSyntax.

If s is syntactically well-formed but is more than 1/2 ULP
away from the largest floating point number of the given size,
ParseFloat returns f = ±Inf, err.Err = ErrRange.

ParseFloat recognizes the strings "NaN", and the (possibly signed) strings "Inf" and "Infinity"
as their respective special floating point values. It ignores case when matching.


*/
