package main

import (
	"errors"
	//x 	"fmt"

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

	//x 	re1 := regexp.MustCompile(`^!.*$`)
	//x 	// need to ignore '!' if inside quotes (of value part) otherwise it's an end of line comment.
	//x 	re2a := regexp.MustCompile(`^(\S+)\s*=\s*(".+")\s*!*.*$`)
	//x 	re2b := regexp.MustCompile(`^(\S+)\s*=\s*(.+)\s*!*.*$`)

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

	//old_regex  	doProcess_regex = func(i_path string) (err error) {
	//old_regex
	//old_regex  		var fz *gzip.Reader
	//old_regex  		var r *bufio.Reader
	//old_regex
	//old_regex  		fi, err := os.Open(i_path)
	//old_regex  		error_check(err, "Error opening "+i_path)
	//old_regex  		defer fi.Close()
	//old_regex
	//old_regex  		if strings.HasSuffix(i_path, "gz") == false {
	//old_regex  			r = bufio.NewReader(fi)
	//old_regex  		} else {
	//old_regex  			fz, err = gzip.NewReader(fi)
	//old_regex  			error_check(err, "Error opening "+i_path)
	//old_regex  			defer fz.Close()
	//old_regex
	//old_regex  			r = bufio.NewReader(fz)
	//old_regex  		}
	//old_regex
	//old_regex  		l_data_until := false
	//old_regex  		scanner := bufio.NewScanner(r)
	//old_regex
	//old_regex  		for scanner.Scan() {
	//old_regex  			line := scanner.Text()
	//old_regex              line = strings.Trim(line, ` \t`)
	//old_regex
	//old_regex  			// check if now scanning data
	//old_regex  			if l_data_until == true {
	//old_regex                  break
	//old_regex              }
	//old_regex
	//old_regex              if len(line) == 0 {
	//old_regex                  continue
	//old_regex              }
	//old_regex
	//old_regex              // check is line commented out
	//old_regex              it := re1.FindStringSubmatch(line)
	//old_regex              if len(it) == 1 {
	//old_regex                  continue
	//old_regex              }
	//old_regex
	//old_regex              // checking (key)(=)("value") ! comment pattern
	//old_regex              it = re2a.FindStringSubmatch(line)
	//old_regex              if len(it) != 3 {
	//old_regex                  // checking (key)(=)(value) ! comment pattern
	//old_regex                  it = re2b.FindStringSubmatch(line)
	//old_regex                  if len(it) != 3 {
	//old_regex                      err = errors.New("reader: Malformerd Line -> " + line)
	//old_regex                      break
	//old_regex                  }
	//old_regex              }
	//old_regex
	//old_regex              k := strings.Trim(it[1], ` "'`)
	//old_regex              v := strings.Trim(it[2], ` "'`)
	//old_regex
	//old_regex              if strings.EqualFold("include", k) == true {
	//old_regex              	ceh_path, err := getIncludePath(v)
	//old_regex              	if err != nil {
	//old_regex              		return err
	//old_regex              	}
	//old_regex
	//old_regex              	includedMap[ceh_path] = true
	//old_regex              	nestedLevel++
	//old_regex
	//old_regex              	if err = doProcess_regex(ceh_path); err != nil {
	//old_regex              		return err
	//old_regex              	}
	//old_regex              	nestedLevel--
	//old_regex              } else {
	//old_regex              	l_data_until = strings.EqualFold("DATA_UNTIL", k)
	//old_regex              	r_header.add_kv(&k, &v)
	//old_regex              }
	//old_regex
	//old_regex  			ix++
	//old_regex  		}
	//old_regex
	//old_regex  		return
	//old_regex  	}

	doProcess = func(i_filepath string) (err error) {
		l_data_until := false
		l_lines := EachLine(i_filepath)

		for kv := range eachKeyVal(l_lines) {
			//x fmt.Println(cx, kv)
			//x cx++

			//x         if strings.EqualFold("include", k) == true {
			if strings.EqualFold("include", kv.key) == true {
				//x             ceh_path, err := getIncludePath(v)
				ceh_path, err := getIncludePath(kv.val[0])
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
				//x             l_data_until = strings.EqualFold("DATA_UNTIL", k)
				l_data_until = strings.EqualFold("DATA_UNTIL", kv.key)
				//x             	r_header.add_kv(&k, &v)
				r_header.add_kv(&kv)
			}

			ix++
		}
        
        if l_data_until == false {
            return errors.New("DATA_UNTIL not found")
        }

		return
	}

	//old_regex r_err = doProcess_regex
	r_err = doProcess(l_path)
	if r_err != nil {
		return
	}

	mooi_log("Lines read -> %d\n", ix)

	r_header.m_data.dump()

	return

}
