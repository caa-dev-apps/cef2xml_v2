package main

import (
	"fmt"
)


func mooi_log(a ...interface{}) (n int, err error) {
// 	return fmt.Println(a)
	if (err!=nil){
	return fmt.Println("%s",err)}
return
}

func main() {

	args, err := NewCefArgs()
	error_check(err, "Invalid Command Line Args")
	args.dump()

	_, err = ReadCef(&args)
	error_check(err, "Error parsing header")

}
