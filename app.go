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
    
//x     header, err := ReadCef(&args)
    _, err = ReadCef(&args)
    error_check(err, "Error parsing header")
    
}
