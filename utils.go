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


func fileExists(name string) (bool, error) {
    
  _, err := os.Stat(name)
  
  if os.IsNotExist(err) {
    return false, nil
  }
  return err != nil, err
}
