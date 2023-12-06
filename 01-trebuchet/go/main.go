package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func getCode(input string) int {
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
                break
            }
        }
        result += code1 * 10 + code2
    }
    return result
}

func main() {

    fmt.Println("Test result:", getCode("test"))
    fmt.Println("Output result:", getCode("input"))

}

