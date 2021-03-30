package go_simple_frw

import (
	"bufio"
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
		*&f.filedata = append(*&f.filedata, data)
	}
	return *f
}

func WriteToFile(data string, f FileData, w bool, path *os.File) {
	if w {

	} else {
		b := bufio.NewWriter(path)
		_, err := b.WriteString(data)
		CheckError(err)
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
