package main

import (
	"bufio"
	"compress/gzip"
	"strings"
	"fmt"
	"os"
	"regexp"
)

func check_data(line string) {
}

///////////////////////////////////////////////////////////////////////////////
//

func ReadHeader(args *CefArgs) (CefHeaderData, error) {
    
//x     included := []string{}
    included := map[string]bool{}
    ix := 0
    nestedLevel := 0
	re1 := regexp.MustCompile(`^\s*!.*$`)
	re2 := regexp.MustCompile(`^\s*(\S+)\s*=\s*(.+)\s*$`)
    
    l_header := CefHeaderData{ m_state: ATTR, m_data: CAA_MetaData{} }
    
    l_path := *args.m_cefpath

    // forward decl
    var doProcess func(i_path string)
    
    doHeader := func(i_filename string) (err error) {
        //x err := error(nil)        
        
        fmt.Println("<---------------------------------------------------------------------- Include ->", i_filename)
        
        for _, l_folder := range args.m_includes {
            l_path := l_folder + `/` + i_filename
            
            _, duplicate := included[l_path]
            if duplicate == true {
                fmt.Println("Error: include duplicate X X X X X X X X X X X X X X X X X X X X X X X X X X X ")
            } else if l_exists, _ := fileExists(l_path); l_exists == true {
                included[l_path] = true
                
                nestedLevel++
                if(nestedLevel < 8 ) {
                    doProcess(l_path)
                    nestedLevel--
                } else {
                    fmt.Println("Error: max nested levels X X X X X X X X X X X X X X X X X X X X X X X X X X X ")
                }
            }
        }

        fmt.Println("----------------------------------------------------------------------- Include ->", i_filename)
        
        return err
    }    
    
    doProcess = func(i_path string) {
    
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
                        
//x                         if strings.EqualFold("include", it[1]) == true {
                        if strings.EqualFold("include", k) == true {
                          //x l_headerXml.add_kv("include-start", val)
                          //x include_ceh(val) 
                            err = doHeader(v)
                          //x l_headerXml.add_kv("include-end", val)
                        } else {
    //x                         l_data_until = strings.EqualFold("DATA_UNTIL", it[1])
    //x                         l_header.add_kv(&it[1], &it[2])
//x                             k := strings.Trim(it[1], ` "'`)
//x                             v := strings.Trim(it[2], ` "'`)
                            
    //x                         l_data_until = strings.EqualFold("DATA_UNTIL", it[1])
                            l_data_until = strings.EqualFold("DATA_UNTIL", k)
                            l_header.add_kv(&k, &v)
                        }
                    }
                    // else either empty ine
                    // or malformed
                }
            }
        }
    }
    
    
    doProcess(l_path)
    
        
//x     fmt.Printf("Lines read -> %d\n", ix)
//x     fmt.Println("----------XXXX------------")
    fmt.Printf("Lines read -> %d\n", ix)
//x     fmt.Println("----------XXXX------------")
    
    l_header.m_data.dump()
    
    fmt.Println("----------XXXX------------")
    
    return l_header, nil
    
}
