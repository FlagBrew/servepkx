package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var pksm = []uint8{0x50, 0x4B, 0x53, 0x4D, 0x4F, 0x54, 0x41}

// readpkmn reads the pokemon from the file and passes it to be uploaded
func readpkmn(fileInfo os.FileInfo, file, ip string) {
	filesize := fileInfo.Size()
	if filesize == 232 || filesize == 264 || filesize == 784 {
		filesize = fileInfo.Size() + 7
		if filesize >= 273 {
			filesize = 273
		}
	}

	payload, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println("Error reading file: " + err.Error())
		return
	}

	pkx := make([]uint8, filesize)
	j := 0
	if fileInfo.Size() == 784 {
		j = 520
	}

	for i := 0; i < 7; i++ {
		pkx[i] = pksm[i]
	}

	for i := int64(7); i < filesize; i++ {
		pkx[i] = payload[j]
		// check to see if we haved a wc6full or wc7full, we have to do some
		// fixes for this because golang will go out of range if we don't
		if strings.HasSuffix(fileInfo.Name(), "wc7full") || strings.HasSuffix(fileInfo.Name(), "wc6full") {
			if int64(j) != fileInfo.Size()-1 {
				j++
			}
		} else {
			j++
		}
	}
	upload(ip, pkx)
}
