package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var port string = os.Getenv("PORT")

func indexHandler(w http.ResponseWriter, r *http.Request) {

	sources, source_err := r.URL.Query()["source"]

	if !source_err {
		print(source_err)
		log.Println("source is missing")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(sources[0]) < 1 {
		print(sources)
		log.Println("source is empty")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	source := sources[0]
	source_sha1 := sha1.Sum([]byte(source))

	app := "wkhtmltoimage"

	out_file := "/tmp/snapshots/" + hex.EncodeToString(source_sha1[:]) + ".jpeg"

	log.Println(app + " " + source + " " + out_file)

	// chrome --headless --disable-gpu --screenshot --window-size=1280,1696 https://www.chromestatus.com/
	cmd := exec.Command(app, source, out_file)
	_, output_err := cmd.Output()

	if output_err != nil {
		log.Println(output_err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	img, fileopen_err := os.Open(out_file)
	if fileopen_err != nil {
		log.Println(fileopen_err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, img)
}

func main() {
	if port == "" {
		port = "80"
	}

	fmt.Println("Server started on port: ", port)

	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}
