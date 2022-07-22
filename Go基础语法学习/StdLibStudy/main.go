package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func timeTransfer() {
	layout := "20060102150405"
	timeStr := "1998-05-06 15:04:05"
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	fmt.Println(t.Format(layout))
}
func MergeTxtFile(path string) error {
	if fileinfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileinfo := range fileinfos {
			if fileinfo.IsDir() {
				MergeTxtFile(filepath.Join(path, fileinfo.Name()))
			} else {
				filename := fileinfo.Name()
				if strings.Contains(filename, ".txt") {
					filename = filepath.Join(path, filename)
					if fout, err := os.OpenFile(filename, os.O_RDONLY, 0666); err != nil {
						return err
					} else {
						reader := bufio.NewReader(fout)
						line, _, err := reader.ReadLine()
						var fin *os.File
						for err == nil {
							if fin, err = os.OpenFile(filepath.Join(path, "merge.txt"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
								writer := bufio.NewWriter(fin)
								n, err := writer.WriteString(string(line) + "\n")
								fmt.Println(err, n)
								writer.Flush()
							}
							line, _, err = reader.ReadLine()
						}
						fin.Close()
						fout.Close()
						if err == io.EOF {
							continue
						} else {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

func main() {
	//timeTransfer()
	err := MergeTxtFile(".")
	fmt.Println(err)
}
