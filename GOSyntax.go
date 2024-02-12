/* Go Syntax
   A Go file consists of the following parts:
   Package declaration
   Import packages
   Functions
   Statements and expressions */

package main /* In Go, every program is part of a package.
                We define this using the package keyword.
				In this example, the program belongs to the main package. */

import (
	"fmt" //fmt package is used to implement formatted Input/Output with functions.
)

func main() { //func: It is used to create a function in Go language.
	fmt.Println("Hello World!") // Println() is used to print the output.
}
