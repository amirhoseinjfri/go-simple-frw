package go_simple_frw

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FileData struct {
	filedata []string
}

//just the file path as string and return data as array
func (f *FileData) ReadFile(path *os.File) FileData {
	bio_scanner := bufio.NewScanner(path)
	for bio_scanner.Scan() {
		data := bio_scanner.Text()
		f.filedata = append(f.filedata, data)
	}
	return *f
}

//Data as String To write to file , f as Typed filedata
//W : if true it will write to the end of file
//if false it will change whole file to the data you typed
//path as string "Just The file name and format like : hi.txt"
func WriteToFile(data string, f FileData, w bool, path *os.File) {
	if w {
		x := f.ReadFile(path)
		d := x.filedata
		d = append(d, data)
		for i, v := range d {
			if (i + 1) == len(d) {
				_, err := fmt.Fprintln(path, v)
				CheckError(err)
			}
		}
	} else {
		err := os.WriteFile(path.Name(), []byte(data), 0644)
		CheckError(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
