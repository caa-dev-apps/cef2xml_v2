package main

import (
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type KVState int

const (
	B4_KEY KVState = iota
	KEY
	B4_EQ
	B4_VAL
	VAL_Q
	VAL
	B4_NEXT
	B4_NEXT_COMMENT
	B4_EOL
	EOL_COMMENT
	CONTINUE_NEXT_LINE
	COMMENT
    EOL
)

func state_str(s KVState) (str string) {

	switch s {
	case B4_KEY:
		return "B4_KEY              "
	case KEY:
		return "KEY                 "
	case B4_EQ:
		return "B4_EQ               "
	case B4_VAL:
		return "B4_VAL              "
	case VAL_Q:
		return "VAL_Q               "
	case VAL:
		return "VAL                 "
	case B4_NEXT:
		return "B4_NEXT             "
	case B4_NEXT_COMMENT:
		return "B4_NEXT_COMMENT     "
	case B4_EOL:
		return "B4_EOL              "
	case EOL_COMMENT:
		return "EOL_COMMENT         "
	case CONTINUE_NEXT_LINE:
		return "CONTINUE_NEXT_LINE  "
	case COMMENT:
		return "COMMENT             "
    case EOL: 
		return "EOL                 "
	default:
		return "????????            "
	}

	return
}

func state_diag(ch rune, s1, s2 KVState) {

	fmt.Println(string(ch), state_str(s1), state_str(s2))
}

func EachLine(i_path string) chan string {
	output := make(chan string)

	go func() {
		defer close(output)

		fi, err := os.Open(i_path)
		if err != nil {
			return
		}
		defer fi.Close()

		var reader *bufio.Reader
		if strings.HasSuffix(i_path, "gz") == false {
			reader = bufio.NewReader(fi)
		} else {
			fz, err := gzip.NewReader(fi)
			if err != nil {
				return
			}
			defer fz.Close()

			reader = bufio.NewReader(fz)
		}

		for {
			line, err := reader.ReadString('\n')
			output <- line
			if err == io.EOF {
				break
			}
		}
	}()

	return output
}

type KeyVal struct {
	key string
	val []string
}

func eachKeyVal(lines chan string) chan KeyVal {
	output := make(chan KeyVal)

	state := B4_KEY
	key := ""
	val := []string{}

	l_key := ""
	l_val := ""

	initVars := func() {
		state = B4_KEY

		key = ""
		val = []string{}

		l_key = ""
		l_val = ""
	}

	parse_line := func(i_line string) (in_err error) {

	skip_comment:
		for _, ch := range i_line {
            //debug state_0 := state
        
			switch state {
			case B4_KEY:
				switch {
				case ch == '!':
					//x state = COMMENT
					break skip_comment
				case unicode.IsSpace(ch):
					// No Change
				case unicode.IsLetter(ch):
					l_key += string(ch)
					state = KEY
				default:
					in_err = errors.New(`INVALID START OF LINE char -> ` + string(ch))
					return
				}
			case KEY:
				switch {
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-':
					l_key += string(ch)
				case unicode.IsSpace(ch):
					state = B4_EQ
				case ch == '=':
					key = l_key
					state = B4_VAL
				default:
					in_err = errors.New("INVALID KEY char -> " + string(ch))
					return
				}
			case B4_EQ:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '=':
					key = l_key
					state = B4_VAL
				default:
					in_err = errors.New("INVALID B4_EQ char -> " + string(ch))
					return
				}
			case B4_VAL:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '"':
					l_val = string(ch)
					state = VAL_Q
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-' || ch == '+' || ch == '$':
					l_val = string(ch)
					state = VAL
				default:
					in_err = errors.New("INVALID B4_VAL char -> " + string(ch))
					return
				}
			case VAL_Q:
				switch ch {
				case '"':
					l_val += string(ch)
					val = append(val, l_val)
					state = B4_EOL
				default:
					l_val += string(ch)

				}
			case VAL:
				switch {
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-' || ch == '+' || ch == ':' || ch == '/':
					l_val += string(ch)
				case unicode.IsSpace(ch):
					val = append(val, l_val)
					state = B4_EOL
				case ch == ',':
					val = append(val, l_val)
					state = B4_NEXT
				case ch == '!':
					val = append(val, l_val)
					state = EOL_COMMENT
				default:
					in_err = errors.New("INVALID VAL char -> " + string(ch))
					return
				}
			case B4_NEXT:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '\\':
					// No Change .. multiline
				case ch == '"':
					l_val = string(ch)
					state = VAL_Q
				case unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '.' || ch == '-' || ch == '+':
					l_val = string(ch)
					state = VAL
				case ch == '!':
					state = B4_NEXT_COMMENT
				default:
					in_err = errors.New("INVALID B4_NEXT char -> " + string(ch))
					return
				}

			case B4_NEXT_COMMENT:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == '\\':
					// No Change .. multiline
					state = B4_NEXT
				}

			case B4_EOL:
				switch {
				case unicode.IsSpace(ch):
					// No Change
				case ch == ',':
					state = B4_NEXT
				case ch == '!':
					break skip_comment
				default:
					in_err = errors.New("INVALID B4_EOL char -> " + string(ch))
					return
				}
			}
            
            //debug state_diag(ch , state_0, state)
		}

        if(state == VAL) {
            val = append(val, l_val)
            state = EOL
        }
        
		return
	}

	go func() {
		defer close(output)

		initVars()

		for l_line := range lines {
            //debug fmt.Println("< ", l_line)
            
			err := parse_line(l_line)
			if err != nil {
				println(err)
				break
			}

			if state != B4_NEXT {
                //debug fmt.Println("< ", KeyVal{key, val})
                  key = strings.ToUpper(key)

                if len(key) > 0 {
                    output <- KeyVal{key, val}
                }
                
				initVars()
			}
		}
	}()

	return output
}

