package main

import (
	"fmt"
//x     "io"
)

///////////////////////////////////////////////////////////////////////////////
//

func mooi_log(a ...interface{}) (n int, err error) {
 	return fmt.Println(a)
//x 	return 
}




func main() {

	args, err := NewCefArgs()
	error_check(err, "Invalid Command Line Args")
	args.dump()

	//x     header, err := ReadCef(&args)
	_, err = ReadCef(&args)
	error_check(err, "Error parsing header")

}
