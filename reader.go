package main

import (
	"bufio"
	"compress/gzip"
	"strings"
	"fmt"
	"os"
	"regexp"
    "errors"
)

func check_data(line string) {
}

///////////////////////////////////////////////////////////////////////////////
//

func ReadHeader(args *CefArgs) (r_header CefHeaderData, r_err error) {
    
//x included := []string{}
    includedMap := map[string]bool{}
    ix := 0
    nestedLevel := 0
	re1 := regexp.MustCompile(`^\s*!.*$`)
	re2 := regexp.MustCompile(`^\s*(\S+)\s*=\s*(.+)\s*$`)
    r_header = CefHeaderData{ m_state: ATTR, m_data: CAA_MetaData{} }
    l_path := *args.m_cefpath

    // forward decl
    //x var doProcess func(i_path string)
    var doProcess func(i_path string) (err error) 
        
    ///////////////////////////////////////////////////////////////////////////////
    //
        
    doHeader := func(i_filename string) (err error) {
        done := false

        if(nestedLevel >= 8) {
            return errors.New("Include nested files limit reached (8 deep)")
        }
        
        for i := 0; i < len(args.m_includes) && done == false; i++ {
        
            l_path := args.m_includes[i] + `/` + i_filename
            
            _, included := includedMap[l_path]
            if included == true {
                return errors.New("Attempt to include duplcate ceh")
            } 
                
            fmt.Println("PATH: ", i, l_path)
               
            if exists, _ := fileExists(l_path); exists == true {
                includedMap[l_path] = true
                nestedLevel++
                
                if err = doProcess(l_path); err != nil {
                    return err
                }
                nestedLevel--
                done = true
            }
        }
        
        if(done == false) {
            return errors.New("Include file not found: " + i_filename)
        }
        
        return err
    }    
    
    
    doProcess = func(i_path string) (err error) {
    
        var fz *gzip.Reader    
        var r *bufio.Reader

        fi, err := os.Open(i_path)
        error_check(err, "Error opening "+i_path)
        defer fi.Close()

        if strings.HasSuffix(i_path, "gz") == false {
            r = bufio.NewReader(fi)
        } else {
            fz, err = gzip.NewReader(fi)
            error_check(err, "Error opening "+i_path)
            defer fz.Close()

            r = bufio.NewReader(fz)
        }

        l_data_until := false

        for {
            line, err := r.ReadString('\n')
            if err != nil {
                fmt.Println("Done reading file!!!")
                break;
            }
            ix++

            // check if scanning data
            if l_data_until == true {
                check_data(line)
            } else {
                
                it := re1.FindStringSubmatch(line)
                // check is not a comment
                if len(it) != 1  {
                    // check if key value pair (it[0]=full-match, it[1]=key, it[2]=value .. length = 3)
                    it := re2.FindStringSubmatch(line)
                    if len(it) == 3  {
                        k := strings.Trim(it[1], ` "'`)
                        v := strings.Trim(it[2], ` "'`)
                        v = strings.Trim(v, ` "'`)
                        
                        if strings.EqualFold("include", k) == true {
                          //x l_headerXml.add_kv("include-start", val)
                          //x include_ceh(val) 
                          
                            println("####", v)
                            err = doHeader(v)
                            if err != nil {
                                return err
                            }
                            
                          //x l_headerXml.add_kv("include-end", val)
                        } else {
                            l_data_until = strings.EqualFold("DATA_UNTIL", k)
                            r_header.add_kv(&k, &v)
                        }
                    }
                    // else either empty ine
                    // or malformed
                }
            }
        }
        
        return
    }
    
    r_err = doProcess(l_path)
    if r_err  != nil {
        return 
    }
        
//x     fmt.Printf("Lines read -> %d\n", ix)
//x     fmt.Println("----------XXXX------------")
    fmt.Printf("Lines read -> %d\n", ix)
//x     fmt.Println("----------XXXX------------")
    
    r_header.m_data.dump()
    
    fmt.Println("----------XXXX------------")
    
    return 
    
}
