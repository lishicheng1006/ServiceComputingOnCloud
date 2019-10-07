package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

//Define variables corresponding to command line parameters
//IntP is like Int, but accepts a shorthand letter that can be used after a single dash.
var startPage = pflag.IntP("startPage", "s", -1, "Specify the start page of the page range to extract")
var endPage = pflag.IntP("endPage", "e", -1, "Specify the end page of the page range to extract")
var lineNum = pflag.IntP("lineNum", "l", 72, "Fixed number of page rows")
var fToSkip = pflag.BoolP("fToSkip", "f", false, "Pages of this type of text are delimited by ASCII page-feed characters")
var destPrinter = pflag.StringP("destPrinter", "d", "", "Acceptable print destination name")

func findRead() (*bufio.Reader, error) {
	//NArg is the number of arguments remaining after flags have been processed.
	if pflag.NArg() == 0 {
		return bufio.NewReader(os.Stdin), nil
	} else if pflag.NArg() == 1 {
		fp, err := os.Open(pflag.Arg(0))
		if err != nil {
			return nil, err
		}
		return bufio.NewReader(fp), nil
	} else {
		return nil, errors.New("Excess parameters detected")
	}
}

func findWrite() (*bufio.Writer, *exec.Cmd, io.WriteCloser, error) {
	if *destPrinter == "" {
		return bufio.NewWriter(os.Stdout), nil, nil, nil
	} else {
		cmd := exec.Command("findstr", "1")
		input, err := cmd.StdinPipe()
		if err != nil {
			return nil, nil, nil, err
		}
		cmd.Stdout = os.Stdout
		return bufio.NewWriter(input), cmd, input, nil
	}
}

func check() error {
	if *startPage == -1 || *endPage == -1 {
		return errors.New("\"-sNumber\" and \"-eNumber\" are required and\r\nmust be the first two arguments on the command line after the command name selpg")
	}
	if *startPage <= 0 {
		return errors.New("The start page is out of range")
	}
	if *endPage <= 0 {
		return errors.New("The end page is out of range")
	}
	if *startPage > *endPage {
		return errors.New("The end page must be no smaller than the start page")
	}
	if *lineNum != 72 && *fToSkip {
		return errors.New("The \"-lNumber\" and \"-f\" options are mutually exclusive")
	}
	return nil
}

func printRes() error {
	err := check()
	if err != nil {
		return err
	}
	reader, err := findRead()
	if err != nil {
		return err
	}
	writer, cmd, input, err := findWrite()
	if err != nil {
		return err
	}
	page := 1
	line := 0
	var delim byte
	if *fToSkip {
		delim = '\f'
	} else {
		delim = '\n'
	}
	for {
		content, err := reader.ReadBytes(delim)
		if err != nil {
			if err != io.EOF {
				return err
			} else {
				break
			}
		}
		if page >= *startPage && page <= *endPage {
			_, err := writer.WriteString(string(content))
			writer.Flush()
			if err != nil {
				return err
			}
		}
		if *fToSkip {
			page++
		} else {
			if line++; line == *lineNum {
				line = 0
				page++
			}
		}
	}
	if cmd != nil {
		input.Close()
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	if page < *startPage || page < *endPage {
		return errors.New("The start page is out of the range of text pages")
	}
	return nil
}

func main() {
	pflag.Parse()
	err := printRes()
	if err != nil {
		log.Fatal(err)
	}
}
