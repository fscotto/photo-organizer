package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fscotto/photo-organizer/fileutil"
	"github.com/fscotto/photo-organizer/util"
	"github.com/rwcarlsen/goexif/exif"
)

const (
	SRC  = "/Users/plague/Downloads/My Photo"
	DEST = "/Users/plague/Downloads/Foto"
)

func main() {
	start := time.Now()
	err := filepath.Walk(SRC, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if util.IsJpegFile(path) {

				handleWriteFile := func(sourceFilePath string, folderName string) {
					destinationPath := filepath.Join(DEST, string(os.PathSeparator), folderName)
					if ok := fileutil.CreateNewFolder(destinationPath); ok {
						fileutil.MoveFile(sourceFilePath, destinationPath)
					}
				}

				image, err := os.Open(path)
				if err != nil {
					log.Println(err)
					return err
				}

				var folderName = "1970"
				metadata, err := exif.Decode(image)
				if err != nil {
					handleWriteFile(path, folderName)
				} else {
					time, err := metadata.DateTime()
					if err != nil {
						handleWriteFile(path, folderName)
					} else {
						folderName = strconv.Itoa(time.Year())
						handleWriteFile(path, folderName)
					}
				}
			} else {
				fmt.Println(path, " format not valid")
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Parsing directory error %v\n", err)
	}
	fmt.Printf("Successful in %d seconds\n", time.Since(start)/time.Second)
}
