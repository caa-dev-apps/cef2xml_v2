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

//x func fileExists(name string) (bool, error) {
//x     
//x     info, err := os.Stat(name)
//x   
//x     if(err != nil) {
//x         return false, err
//x     }
//x   
//x     return info.Mode().IsRegular(), err
//x }
//x 

func fileExists(name string) (isReq bool, err error) {
    
    info, err := os.Stat(name)
  
    if(err == nil) {
        isReq = info.Mode().IsRegular()
    }
  
    return
}
