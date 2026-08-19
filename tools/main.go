package main

import (
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	qrcode.WriteFile(os.Args[1], qrcode.Medium, 512, os.Args[2])
}
