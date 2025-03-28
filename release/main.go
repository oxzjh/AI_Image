package main

import (
	"api/image"
	"api/sd"
	"context"
	"embed"
	"fmt"
	"golib/chrome"
	"golib/server"
	"golib/server/http"
	"log"
)

//go:embed dist
var dist embed.FS

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
	sd.Initialize("models", "models/taesd", "models/lora", "output.jpg")
	http.MaxLength = 20 << 20
	http.Domains = []string{"*"}
	http.AllowHeaders = "*"
	l1, port1, err := server.ServeLocal(http.NewServer())
	if err != nil {
		log.Fatal(err)
	}
	defer l1.Close()
	apiURL := fmt.Sprintf("http://127.0.0.1:%d", port1)

	c, err := chrome.New(context.Background(), "", "", 800, 600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	c.Bind("fetchURL", func() string {
		return apiURL
	})

	l2, port2, err := server.ServeLocalFile(dist)
	if err != nil {
		log.Fatal(err)
	}
	defer l2.Close()
	c.Load(fmt.Sprintf("http://127.0.0.1:%d/dist", port2))
	c.DisableContextMenu()
	c.DisableDevTools()
	<-c.Done()
}
