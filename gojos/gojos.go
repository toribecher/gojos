package gojos

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run(addr string) {
	log.Println("RUNNING!")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) Initialize() {
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		log.Fatal()
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/getPNG/{png}", getPNG).Methods("GET")
}

func getPNG(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	png := params["png"]

	pngPath := fmt.Sprintf("./uploads/%s.png", png)
	pngImage, err := getImageFromFilePath(pngPath)
	if err != nil {
		http.Error(w, "The png extension cannot be found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "image/png")
	writeImage(w, &pngImage)
}

func getImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f)
	return image, err
}

func writeImage(w http.ResponseWriter, img *image.Image) {
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
