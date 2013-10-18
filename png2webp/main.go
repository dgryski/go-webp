package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"reflect"

	"github.com/dgryski/go-webp"
)

func main() {

	ifile := os.Args[1]
	ofile := os.Args[2]

	r, err := os.Open(ifile)
	if err != nil {
		log.Fatal("error during open: ", err)
	}

	img, _, err := image.Decode(r)

	if err != nil {
		log.Fatal("error decoding file: ", err)
	}

	var output []byte

	switch t := img.(type) {
	case *image.NRGBA:
		output = webp.WebPEncodeLosslessRGBA(t.Pix, t.Rect.Dx(), t.Rect.Dy(), t.Stride)
	case *image.RGBA:
		output = webp.WebPEncodeLosslessRGBA(t.Pix, t.Rect.Dx(), t.Rect.Dy(), t.Stride)
	default:
		log.Println("unknown type: ", reflect.TypeOf(img))
	}

	ioutil.WriteFile(ofile, output, 0666)

	webp.FreeWebP(output)
}
