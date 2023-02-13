package ui

import (
	"errors"
	"fmt"
	"os"
	"text/scanner"
	"unicode"
)

var (
	sc     scanner.Scanner
	Prompt = "> "
	Quit   = errors.New("quit")
)

func init() {
	sc.Init(os.Stdin)
	sc.Filename = "stdin"
	sc.IsIdentRune = isIdent
	sc.Error = errorHandler
	sc.Mode = scanner.ScanIdents | scanner.ScanChars | scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanRawStrings
}

type Input []string

func GetInput() (Input, bool, error) {
	fmt.Print(Prompt)

	var input Input

	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		switch tok {
		case scanner.String, scanner.RawString, scanner.Char:
			str := sc.TokenText()
			if len(str) < 3 {
				input = append(input, "")
			} else {
				input = append(input, str[1:len(str)-1])
			}
		default:
			input = append(input, sc.TokenText())
		}
		if sc.Peek() == '\n' {
			break
		}
	}

	if input == nil || len(input) == 0 {
		return nil, true, nil
	}
	return input, false, nil
}

func isIdent(r rune, i int) bool {
	if unicode.IsSpace(r) || unicode.IsControl(r) {
		return false
	}

	if i == 0 {
		if unicode.IsPunct(r) {
			return false
		}
		switch r {
		case '"', '`', '\'':
			return false
		}
	}

	return true
}

func errorHandler(s *scanner.Scanner, msg string) {
	if msg == "invalid char literal" {
		return
	}
	fmt.Println("Error:", msg)
}

func GetString(prompt string) (string, error) {
	fmt.Printf("%s ", prompt)
	var s string
	_, err := fmt.Scanln(&s)
	return s, err
}
