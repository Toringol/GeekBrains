package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	sourceFile = flag.String("source", "", "sourceFile")
	findStr    = flag.String("strFind", "", "strFind")
)

type Opts struct {
	SourceFile string
	FindStr    string
}

func grep(source, strFind string) {
	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	fileScanner := bufio.NewScanner(sourceFile)

	for fileScanner.Scan() {
		textToCheck := fileScanner.Text()
		if strings.Contains(textToCheck, strFind) {
			fmt.Println(textToCheck)
		}
	}

	if fileScanner.Err() != nil {
		log.Fatal(err)
	}

}

func main() {
	flag.Parse()

	opts := Opts{
		SourceFile: *sourceFile,
		FindStr:    *findStr,
	}

	if len(flag.Args()) > 2 {
		log.Fatalf("Too many arguments")
	}

	if opts.SourceFile == "" || opts.FindStr == "" {
		log.Fatalf("You should write sourcefile and what you want to find like this\n" +
			"--source=name --strFind=name")
	}

	grep(opts.SourceFile, opts.FindStr)
}
