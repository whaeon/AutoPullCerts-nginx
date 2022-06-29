package app

import (
	"archive/zip"
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func UnzipWithSave(data []byte, domain *string) {
	// read decoded data
	decodeReader := bytes.NewReader(data)
	// get the content of data
	content, err := ioutil.ReadAll(decodeReader)
	if err != nil {
		log.Println("get content error:", err)
		return
	}
	// create zip reader to read zip archive data
	zipReader, err := zip.NewReader(bytes.NewReader(content), int64(len(content)))
	if err != nil {
		log.Println("create zip reader error: ", err)
		return
	}

	fullPemName := *domain + ".pem"
	fullKeyName := *domain + ".key"
	fullPemPath := "/etc/nginx/certs/" + fullPemName
	fullKeyPath := "/etc/nginx/certs/" + fullKeyName

	// Read all the files from zip archive
	for _, zipFile := range zipReader.File {
		// the Open func return ReadCloser will be ignore by continue or return
		// so, we create a new func to resolve
		if fullPemName == zipFile.Name {
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			file, err := os.OpenFile(fullPemPath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Println("open file failed, ", err)
				return
			}
			fullPemWriter := bufio.NewWriter(file)
			fullPemWriter.WriteString(string(unzippedFileBytes))
			fullPemWriter.Flush()
			file.Close()
		} else if fullKeyName == zipFile.Name {
			unzippedFileBytes, err := readZipFile(zipFile)
			if err != nil {
				log.Println(err)
				continue
			}
			file, err := os.OpenFile(fullKeyPath, os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Println("open file failed, ", err)
				return
			}
			fullKeyWriter := bufio.NewWriter(file)
			fullKeyWriter.WriteString(string(unzippedFileBytes))
			fullKeyWriter.Flush()
			file.Close()
		} else {
			continue
		}
	}
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
