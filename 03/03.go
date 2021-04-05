package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	// fmt.Println(runtime.GOOS)
	_, b, _, _ := runtime.Caller(0)
	a := func() string {
		if runtime.GOOS != "windows" {
			return "file.txt"
		} else {
			return path.Dir(b) + "/file.txt"
		}
	}()

	file, err := os.Open(a) // in windows should be specified root,

	if err != nil {
		log.Fatal(err)
	}

	counts := map[string]int{}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		counts[word]++
	}

	fmt.Println("aku:", counts["aku"], ", ingin:", counts["ingin"], ", dapat:", counts["dapat"])
}
