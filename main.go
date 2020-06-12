package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
)

func main() {
	fileUrl := "https://compilers.iecc.com/crenshaw/tutor2.txt"
	err := DownloadFile("tutor2.txt", fileUrl)
	if err != nil {
		logrus.Fatalln(err)
	}
	logrus.Println("Downloaded: ", fileUrl)
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