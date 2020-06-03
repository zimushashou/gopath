package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var infile *string = flag.String("i", "infile", "file contains values for sorting")
var outfile *string = flag.String("o", "outfile", "file to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func readValue(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	defer file.Close()
	values = make([]int, 0)
	reader := bufio.NewReader(file)
	for  {
		line, prefix, err1 := reader.ReadLine()
		if err != nil {
			if err1 != io.EOF {
				 err = err1
			}
			break
		}
		if prefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		strLine := string(line)
		value, err1 := strconv.Atoi(strLine)
		fmt.Println("value: ", value)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}
	return values, nil
}
func main()  {
	// input args parse
	flag.Parse()
	if infile != nil{
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}
	// read file from infile
	values, err := readValue(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Read values:", values)
	}
}
