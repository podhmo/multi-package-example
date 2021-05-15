package bar

import "fmt"

const VERSION = "0.1.1"

func Hello() {

	fmt.Printf("hello from %s (%s)\n", "bar", VERSION)
}
