package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)
	}

	values,err:=readValues(*infile)

	if err==nil{
		fmt.Println("Read values:",values)
		writeValues(values,*outfile)
	}else{
		fmt.Println(err)
	}
}

func readValues(infile string) (values[] int,err error) {
	file,err:=os.Open(infile)
	if err!=nil{
		fmt.Println("Failed to open the input file ", infile)
		return
	}
	defer file.Close()
	br:=bufio.NewReader(file)
	values=make([] int,0)

	for{
		line,isPrefix,readErr:=br.ReadLine();
		if readErr!=nil{
			if readErr!=io.EOF{
				err=readErr
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}
		str := string(line) // 转换字符数组为字符串
		value, errConv := strconv.Atoi(str)
		if errConv != nil {
			err = errConv
			return
		}
		values = append(values, value)
	}
	return
}

func writeValues(values [] int,outfile string) (err error) {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}
	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}