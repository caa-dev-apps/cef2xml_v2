package main

import (
//x 	"fmt"
)

func mooi_log(a ...interface{}) (n int, err error) {
	return
}

func main() {

	args, err := NewCefArgs()
	error_check(err, "Invalid Command Line Args")
	args.dump()

	//x println("Hello, World!")

	_, err = ReadCef(&args)
	error_check(err, "Error parsing header")

}
