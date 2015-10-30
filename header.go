package main

import (
	"bufio"
	"compress/gzip"
	"strings"
	"fmt"
	"os"
	"regexp"
)

///////////////////////////////////////////////////////////////////////////////
//


//x //x func add_kv(k, v string) {
//x func add_kv(k, v *string) {
//x     //
//x     fmt.Println(*k, *v)
//x }

func check_data(line string) {
}

///////////////////////////////////////////////////////////////////////////////
//

func ReadHeader(args *CefArgs) (CefHeaderData, error) {
    
    l_header := CefHeaderData{ debug: "This is a test 2"}
    
    l_path := *args.m_cefpath

	var fz *gzip.Reader    
	var r *bufio.Reader

	fi, err := os.Open(l_path)
	error_check(err, "Error opening "+l_path)
	defer fi.Close()

	if strings.HasSuffix(l_path, "gz") == false {
		r = bufio.NewReader(fi)
	} else {
		fz, err = gzip.NewReader(fi)
		error_check(err, "Error opening "+l_path)
		defer fz.Close()

		r = bufio.NewReader(fz)
	}

	re1 := regexp.MustCompile(`^\s*!.*$`)
	re2 := regexp.MustCompile(`^\s*(\S+)\s*=\s*(.+)\s*$`)
    
    l_data_until := false

	ix := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
            break;
		}
		ix++

        //i fmt.Print(line)

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
                    if strings.EqualFold("include", it[1]) == true {
                      //x l_headerXml.add_kv("include-start", val)
                      //x include_ceh(val) 
                      //x l_headerXml.add_kv("include-end", val)
                    } else {
                        l_data_until = strings.EqualFold("DATA_UNTIL", it[1])
                        l_header.add_kv(&it[1], &it[2])
                    }
                }
                // else either empty ine
                // or malformed
            }
        }
	}
    
    fmt.Printf("Lines read -> %d\n", ix)
    
    return l_header, nil
    
}
