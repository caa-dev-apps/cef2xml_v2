package main

import (
//x 	"fmt"
)

///////////////////////////////////////////////////////////////////////////////
//

func main() {
    
    args, err := NewCefArgs()
    error_check(err, "Invalid Command Line Args")
    args.dump()
    
//x     header, err := ReadHeader(&args)
    _, err = ReadHeader(&args)
    error_check(err, "Error parsing header")
    
    
//x     fmt.Println("___________________")
//x     fmt.Println(header)
//x     fmt.Println("___________________")
    
}
