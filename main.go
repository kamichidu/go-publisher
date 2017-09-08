package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
)

var appVersion string

func run(in io.Reader, out io.Writer, errOut io.Writer, args []string) int {
	log.SetOutput(errOut)
	log.SetPrefix(fmt.Sprintf("[%s] ", filepath.Base(args[0])))

	var (
		versionFlag bool
		pkgName     string
		typName     string
		ofilename   string
		noGofmtFlag bool
		buildTags   string
	)
	flg := flag.NewFlagSet(filepath.Base(args[0]), flag.ExitOnError)
	flg.BoolVar(&versionFlag, "v", false, "Show version")
	flg.StringVar(&pkgName, "p", "main", "Package name")
	flg.StringVar(&typName, "t", "", "Publisher type name")
	flg.StringVar(&ofilename, "o", "-", "Output filename")
	flg.BoolVar(&noGofmtFlag, "no-gofmt", false, "Don't use gofmt for generated go code")
	flg.StringVar(&buildTags, "tags", "", "go's build tags placed at top of generated code")
	if err := flg.Parse(args[1:]); err != nil {
		log.Printf("Argument error: %s", err)
		return 128
	}

	if versionFlag {
		fmt.Println(appVersion)
		return 0
	}

	buffer := new(bytes.Buffer)
	if err := generate(buffer, pkgName, typName, buildTags, flg.Args()); err != nil {
		log.Printf("Can't generate code: %s", err)
		return 1
	}

	var w io.Writer
	if ofilename == "-" {
		w = out
	} else {
		if f, err := os.Create(ofilename); err != nil {
			log.Printf("Can't create %v: %s", ofilename, err)
			return 128
		} else {
			defer f.Close()
			w = f
		}
	}

	var generatedContent []byte
	if !noGofmtFlag {
		var err error
		generatedContent, err = format.Source(buffer.Bytes())
		if err != nil {
			log.Panic("Can't gofmt output: %s", err)
		}
	} else {
		generatedContent = buffer.Bytes()
	}
	if _, err := w.Write(generatedContent); err != nil {
		log.Printf("Can't write output: %s", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Stdin, os.Stdout, os.Stderr, os.Args))
}
