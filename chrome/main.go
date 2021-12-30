package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var port string = os.Getenv("PORT")

func doSnapshot(source string, out_file string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app := "/usr/bin/google-chrome-stable"
	args := []string{
		"--headless",
		"--disable-software-rasterizer",
		"--window-size=800,600",
		"--no-sandbox",
		"--disable-gpu",
		"--screenshot=" + out_file,
		source,
	}

	cmd := exec.CommandContext(ctx, app, args...)
	err := cmd.Run()
	return err
}

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
	out_file := "/tmp/snapshots/" + hex.EncodeToString(source_sha1[:]) + ".jpeg"

	_, err := os.Stat(out_file)

	if errors.Is(err, os.ErrNotExist) {
		output_err := doSnapshot(source, out_file)

		if output_err != nil {
			log.Println(output_err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("file exists: " + out_file)
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
