package main

import (
	"api/image"
	"golib/server"
	"golib/server/http"
	"log"
	"time"
)

func main() {
	image.Initialize(0x92, &image.Models{
		Upscaler: "models/upscaler.bin",
		HeadSeg:  "models/headseg.bin",
		Cartoon:  "models/cartoon.bin",
		Line:     "models/line.bin",
		Enhancer: "models/enhancer.bin",
		Colorizers: map[string]string{
			"eccv16":     "models/eccv16.bin",
			"siggraph17": "models/siggraph17.bin",
		},
		Styles: map[string]string{
			"candy":      "models/candy.bin",
			"mosaic":     "models/mosaic.bin",
			"pointilism": "models/pointilism.bin",
			"princess":   "models/princess.bin",
			"udnie":      "models/udnie.bin",
		},
	}, 4, "output.jpg")
	http.MaxLength = 20 << 20
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	http.ReturnErr = true
	log.Fatal(server.ServeHTTP(":3000", http.NewServer(), 5*time.Second))
}
