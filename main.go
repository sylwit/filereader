package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

type FileData struct {
	info    os.FileInfo
	content []byte
}

func getFileData(secretPath string) (*FileData, error) {
	fileInfo, err := os.Stat(secretPath)
	//fmt.Fprintf(w, "fileinfo = %#v\n", fileInfo)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadFile(secretPath)
	if err != nil {
		return nil, err
	}

	return &FileData{
		info:    fileInfo,
		content: content,
	}, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		secret := query.Get("file")
		secret = path.Clean(secret)

		chroot := "."
		if value, ok := os.LookupEnv("FILEREADER_CHROOT"); ok {
			chroot = value
		}

		secretPath := path.Join(chroot, secret)

		data, err := getFileData(secretPath)
		if err != nil {
			w.WriteHeader(400)
			if _, err := fmt.Fprintf(w, "%s", err); err != nil {
				log.Fatal(err)
			}
			return
		}

		if _, err := fmt.Fprintf(w, "Content of file %s\n"+
			"size = %d\n"+
			"last modification = %s\n"+
			"---\n%s",
			secretPath, data.info.Size(), data.info.ModTime(), data.content); err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
        	w.Header().Set("Content-Type", "text/html")

        	fmt.Fprintf(w, "<html><body>")
        	fmt.Fprintf(w, "<h1>Environment Variables</h1>")
        	fmt.Fprintf(w, "<ul>")

        	for _, env := range os.Environ() {
                fmt.Fprintf(w, "<li>%s</li>", env)
        	}

        	fmt.Fprintf(w, "</ul>")
        	fmt.Fprintf(w, "</body></html>")
    })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
