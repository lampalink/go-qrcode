package main

import (
	"fmt"

	qrcode "github.com/lampalink/go-qrcode"
)

func main() {
	qrc, err := qrcode.New("https://github.com/lampalink/go-qrcode")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
	}

	// save file
	if err := qrc.Save("../testdata/repo-qrcode.jpeg"); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
