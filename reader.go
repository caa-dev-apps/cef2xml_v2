package main

import (
	"bufio"
	"compress/gzip"
	"errors"
//x 	"fmt"
	"os"
	"regexp"
	"strings"
)

//x func check_data(line string) {
//x }

///////////////////////////////////////////////////////////////////////////////
//

func ReadCef(args *CefArgs) (r_header CefHeaderData, r_err error) {

	includedMap := map[string]bool{}
	ix := 0
	nestedLevel := 0
//x 	re1 := regexp.MustCompile(`^\s*!.*$`)
//x 	re2 := regexp.MustCompile(`^\s*(\S+)\s*=\s*(.+)\s*$`)
    
	re1 := regexp.MustCompile(`^!.*$`)
    // need to ignore '!' if inside quotes (of value part) otherwise it's an end of line comment.
	re2a := regexp.MustCompile(`^(\S+)\s*=\s*(".+")\s*!*.*$`)
	re2b := regexp.MustCompile(`^(\S+)\s*=\s*(.+)\s*!*.*$`)
    
	r_header = CefHeaderData{m_state: ATTR, m_data: CAA_MetaData{}}
	l_path := *args.m_cefpath

	// forward decl
	var doProcess func(i_path string) (err error)

	///////////////////////////////////////////////////////////////////////////////
	//

	getIncludePath := func(i_filename string) (r_path string, err error) {
		done := false

		if nestedLevel >= 8 {
			return r_path, errors.New("Include nested files limit reached (8 deep)")
		}

		for i := 0; i < len(args.m_includes) && done == false; i++ {

			r_path = args.m_includes[i] + `/` + i_filename

			_, included := includedMap[r_path]
			if included == true {
				return r_path, errors.New("Attempt to include duplcate ceh")
			}

			mooi_log("PATH: ", i, r_path)

			done, _ = fileExists(r_path)
		}

		if done == false {
			err = errors.New("Include file not found: " + i_filename)
		}

		return
	}

	///////////////////////////////////////////////////////////////////////////////
	//

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
		scanner := bufio.NewScanner(r)

		for scanner.Scan() {
			line := scanner.Text()
            line = strings.Trim(line, ` \t`)

			// check if now scanning data
			if l_data_until == true {
                break
            }
            
            if len(line) == 0 {
                continue
            }
            
            // check is line commented out
            it := re1.FindStringSubmatch(line)
            if len(it) == 1 {
                continue
            }
            
            // checking (key)(=)("value") ! comment pattern
            it = re2a.FindStringSubmatch(line)
            if len(it) != 3 {
                // checking (key)(=)(value) ! comment pattern
                it = re2b.FindStringSubmatch(line)
                if len(it) != 3 {
                    err = errors.New("reader: Malformerd Line -> " + line)
                    break
                } 
            } 

            k := strings.Trim(it[1], ` "'`)
            v := strings.Trim(it[2], ` "'`)
                
            if strings.EqualFold("include", k) == true {    
            	ceh_path, err := getIncludePath(v)    
            	if err != nil {    
            		return err    
            	}    
                
            	includedMap[ceh_path] = true    
            	nestedLevel++    
                
            	if err = doProcess(ceh_path); err != nil {    
            		return err    
            	}    
            	nestedLevel--    
            } else {    
            	l_data_until = strings.EqualFold("DATA_UNTIL", k)    
            	r_header.add_kv(&k, &v)    
            }    
                
			ix++
		}

		return
	}

	r_err = doProcess(l_path)
	if r_err != nil {
		return
	}

	mooi_log("Lines read -> %d\n", ix)

	r_header.m_data.dump()

	return

}
