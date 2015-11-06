package main

import (
    "errors"
	"fmt"
 	"flag"
)

///////////////////////////////////////////////////////////////////////////////
// required to handle cmd args - list of strings e.g. include folders

type strslice []string
 
func (ss *strslice) String() string {
    //x return *s
    return fmt.Sprintf("%s", *ss)
}
 
func (ss *strslice) Set(s string) error {
    *ss = append(*ss, s)
    return nil
}

///////////////////////////////////////////////////////////////////////////////
//

type CefArgs struct
{
    m_includes strslice
    m_cefpath *string
}

func (a1s *CefArgs) init() error {
    err := error(nil)
    
    flag.Var(&a1s.m_includes,   "i", "Include Folders")
    a1s.m_cefpath = flag.String("f", "", "Cef file path (.cef/.cef.gz)")
    
    flag.Parse()
    if flag.NFlag() == 0 {
        flag.PrintDefaults()
        err = errors.New("Error: Invalid number of cmdline args")
    } 
    
    return err
}

func (a1s *CefArgs) dump() {
    fmt.Printf("cef path: %s\n", *a1s.m_cefpath)
    fmt.Println("include folders", a1s.m_includes)
}

func NewCefArgs() (args CefArgs, err error) {
    args = CefArgs{}
    err = args.init()
    
    return args, err 
}

