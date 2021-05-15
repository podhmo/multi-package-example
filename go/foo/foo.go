package foo

import "fmt"

const VERSION = "0.1.1"

func Hello() {

	fmt.Printf("hello from %s (%s)\n", "foo", VERSION)
}
