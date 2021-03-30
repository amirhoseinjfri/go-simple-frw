package go_simple_frw

import (
	"bufio"
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
