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

func (f *FileData) ReadFile(path *os.File) FileData {
	bio_scanner := bufio.NewScanner(path)
	for bio_scanner.Scan() {
		data := bio_scanner.Text()
		f.filedata = append(f.filedata, data)
	}
	return *f
}

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
