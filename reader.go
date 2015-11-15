package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check_data(line string) {
}

///////////////////////////////////////////////////////////////////////////////
//

func ReadCef(args *CefArgs) (r_header CefHeaderData, r_err error) {

	includedMap := map[string]bool{}
	ix := 0
	nestedLevel := 0
	re1 := regexp.MustCompile(`^\s*!.*$`)
	re2 := regexp.MustCompile(`^\s*(\S+)\s*=\s*(.+)\s*$`)
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

			fmt.Println("PATH: ", i, r_path)

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
			ix++

			// check if scanning data
			if l_data_until == true {
				check_data(line)
			} else if len(line) > 0 {

				it := re1.FindStringSubmatch(line)
				// check is not a comment
				if len(it) != 1 {
					// check if key value pair (it[0]=full-match, it[1]=key, it[2]=value .. length = 3)
					it := re2.FindStringSubmatch(line)
					if len(it) == 3 {
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
					} else {
						fmt.Println("Error: Malformed Line Pattern mismatch!!!!!!!!!!!!!!!!", "[", line, "]", len(line))
					}
					// else either empty ine
					// or malformed
				} else {
					//x fmt.Println("This is a comment", line)
				}
			}
		}

		return
	}

	r_err = doProcess(l_path)
	if r_err != nil {
		return
	}

	fmt.Printf("Lines read -> %d\n", ix)

	r_header.m_data.dump()

	return

}
