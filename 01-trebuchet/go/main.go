package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	"strings"
	"unicode"

)

func getCodePart1(line []rune) int {
    i := 0
    j := len(line) - 1
    code1 := -1
    code2 := -1

    for (i <= j) {
        if(code1 == -1) {
            c1 := line[i]
            if(unicode.IsDigit(c1)) {
                code1 = int(c1 - '0')
            } else {
                i++
            }
        }

        if(code2 == -1) {
            c2 := line[j]
            if(unicode.IsDigit(c2)) {
                code2 = int(c2 - '0')
            } else {
                j--
            }
        }

        if(code1 != -1 && code2 != -1) {
            // fmt.Println(string(line), code1, code2)
            return code1 *10 + code2
        }
    }

    return 0

}

func getCodePart2(line []rune) int {

    numStrs := [9]string{
        "one",
        "two",
        "three",
        "four",
        "five",
        "six",
        "seven",
        "eight",
        "nine",
    }

    fp := 0
    bp := len(line) - 1
    code1 := -1
    code2 := -1

    for(fp <= bp) {

        if(code1 == -1) {
            c1 := line[fp]
            if(unicode.IsDigit(c1)) {
                code1 = int(c1 - '0')
            } else {
                for numIndex, num := range numStrs {
                    if(strings.HasSuffix(string(line[0:fp+1]), num)) {
                        code1 = numIndex + 1
                    }
                }
                if(code1==-1) {fp++}
            }
        }

        if(code2 == -1) {
            c2 := line[bp]
            if(unicode.IsDigit(c2)) {
                code2 = int(c2 - '0')
            } else {
                for numIndex, num := range numStrs {
                    if(strings.HasPrefix(string(line[bp:]), num)) {
                        code2 = numIndex + 1
                    }
                }
                if(code2==-1) {bp--}
            }
        }

        if(code1 != -1 && code2 != -1) {
            // fmt.Println(string(line), code1, code2)
            return code1 *10 + code2
        }

    }

    return 0

}

func readCode(input string, algoPart uint8) int {
    fi, err := os.Open(input)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    var result int = 0

    scanner := bufio.NewScanner(fi)

    for scanner.Scan() {
        line := [] rune(scanner.Text())
        switch algoPart {
            case 1:
            result += getCodePart1(line)
            case 2:
            result += getCodePart2(line)
        }
    }
    return result
}

func main() {

    fmt.Println("Test result(1):", readCode("test", 1))
    fmt.Println("Output result(1):", readCode("input",1))

    fmt.Println("Test result(2):", readCode("test2", 2))
    fmt.Println("Output result(2):", readCode("input",2))
}

