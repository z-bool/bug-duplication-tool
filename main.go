package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var globalMap map[string]bool
var totalMap []string

func readFile() *bufio.Reader {
	fi, err := os.Open("domains.txt")
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(fi)
	return r

}

func writeFile(list []string) {
	file, err := os.OpenFile("result.txt", os.O_CREATE|os.O_APPEND, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for _, val := range list {
		file.WriteString(val + "\n")
	}
}

func init() {
	globalMap = make(map[string]bool)
	totalMap = make([]string, 0)
}

func main() {
	r := readFile()
	for {
		lineBytes, err := r.ReadBytes('\n')
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		if globalMap[line] == false {
			globalMap[line] = true
			totalMap = append(totalMap, line)
		}
	}
	writeFile(totalMap)
}
