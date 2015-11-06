package main

import (
	"fmt"
)

///////////////////////////////////////////////////////////////////////////////
//

func main() {
    
    args, err := NewCefArgs()
    error_check(err, "Invalid Command Line Args")
    args.dump()
    
    header, err := ReadHeader(&args)
    error_check(err, "Error parsing header")
    
    
    fmt.Println("___________________")
    fmt.Println(header)
    fmt.Println("___________________")
    
}
