package utils

import (
	"github.com/skip2/go-qrcode"
)

func QrGenerate(url string) ([]byte, error) {
	code, err := qrcode.New(url, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	png, err := code.PNG(256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
