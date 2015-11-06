package main

import (
	"fmt"
	"os"
)

///////////////////////////////////////////////////////////////////////////////
//

func error_check(err error, i_s string) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(i_s)
		os.Exit(1)
	}
}


