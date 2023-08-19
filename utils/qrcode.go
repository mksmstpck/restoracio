package utils

import (
	"github.com/mksmstpck/restoracio/internal/config"
	"github.com/skip2/go-qrcode"
)

func QrGenerate(route string) ([]byte, error) {
	url := config.NewConfig().GinUrl
	code, err := qrcode.New(url+route, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	png, err := code.PNG(256)
	if err != nil {
		return nil, err
	}
	return png, nil
}
