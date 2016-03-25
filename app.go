package main

import (
//x 	"fmt"
)

func mooi_log(a ...interface{}) (n int, err error) {
	//x fmt.Println(a)

	return
	//x     return println(a)
}

//x func mooi_log(a ...interface{}) {
//x //x  	return fmt.Println(a)
//x     println(a)
//x }

func main() {

	args, err := NewCefArgs()
	error_check(err, "Invalid Command Line Args")
	args.dump()

	//x println("Hello, World!")

	_, err = ReadCef(&args)
	error_check(err, "Error parsing header")

}
