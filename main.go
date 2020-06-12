package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for i := 0; i >= 0; i++{
		fileUrl := fmt.Sprintf("https://compilers.iecc.com/crenshaw/tutor%d.txt", i)
		filePath := path.Base(fileUrl)
		err := DownloadFile(filePath, fileUrl)
		if err != nil {
			logrus.Fatalln(err)
		}
		logrus.Println("Downloaded: ", fileUrl)
	}

}


func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Fatalln(err)
	}

	out, err := os.Create(filepath)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}