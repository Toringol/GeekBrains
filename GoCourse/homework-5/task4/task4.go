package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	sourceFile      = flag.String("source", "", "sourceFile")
	destinationFile = flag.String("destination", "", "destinationFile")
)

type Opts struct {
	SourcePath      string
	DestinationPath string
}

func copy(source, destination string) {
	sourceFile, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	dataToCopy, err := ioutil.ReadAll(sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	destinationFile, err := os.Create(destination)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	_, err = destinationFile.Write(dataToCopy)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	opts := Opts{
		SourcePath:      *sourceFile,
		DestinationPath: *destinationFile,
	}

	if len(flag.Args()) > 2 {
		log.Fatalf("Too many arguments")
	}

	if opts.SourcePath == "" || opts.DestinationPath == "" {
		log.Fatalf("You should write sourcefile and destination file like this\n" +
			"--source=name --destination=name")
	}

	copy(opts.SourcePath, opts.DestinationPath)
}
