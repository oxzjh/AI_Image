package main

import (
	"ai/sd"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var gSD *sd.SD

type progress struct {
	Step  int     `json:"step"`
	Steps int     `json:"steps"`
	Time  float32 `json:"time"`
}

func main() {
	var (
		cfgScale    float64
		inputImage  string
		outputImage string
	)
	mParams := sd.ModelParams{
		FreeParamsImmediately: true,
		GgmlType:              sd.SD_TYPE_COUNT,
		RngType:               sd.CUDA_RNG,
	}
	params := newParams()
	flag.StringVar(&mParams.ModelPath, "m", "", "Model path")
	flag.StringVar(&mParams.TaesdPath, "taesd", "", "Taesd model path")
	flag.StringVar(&mParams.LoraModelDir, "lora-dir", "", "Lora models dir")
	flag.IntVar(&mParams.NumThreads, "t", runtime.NumCPU(), "Number of threads")
	flag.StringVar(&params.Prompt, "p", "", "Prompt")
	flag.StringVar(&params.NegativePrompt, "n", "", "Negative prompt")
	flag.Float64Var(&cfgScale, "cfg-scale", 1, "Unconditional guidance scale")
	flag.IntVar(&params.Width, "width", 512, "Image width")
	flag.IntVar(&params.Height, "height", 512, "Image height")
	flag.IntVar(&params.SampleMethod, "sample-method", 0, "Sampling method: 0-euler_a, 1-euler, 2-heun, 3-dpm2, 4-dpm++2s_a, 5-dpm++2m, 6-dpm++2mv2, 7-lcm")
	flag.IntVar(&params.SampleSteps, "steps", 1, "Sample steps")
	flag.Int64Var(&params.Seed, "s", -1, "RNG seed, use random seed for -1")
	flag.IntVar(&params.BatchCount, "b", 1, "Batch count")
	flag.StringVar(&inputImage, "i", "", "Input image for image to image")
	flag.StringVar(&outputImage, "o", "output.jpg", "Output image path")
	flag.Parse()

	var err error
	if gSD, err = sd.New(&mParams); err != nil {
		log.Fatal(err)
	}

	params.CfgScale = float32(cfgScale)
	var m image.Image
	if inputImage != "" {
		if m, err = readImageFromFile(inputImage); err != nil {
			log.Fatal(err)
		}
	}
	generateImages(os.Stdout, m, params, outputImage)
}

func newParams() *sd.I2IParams {
	var params sd.I2IParams
	params.ClipSkip = -1
	params.ControlCond = sd.Image{}
	params.ControlStrength = 0.9
	params.StyleStrength = 20
	params.Strength = 0.75
	return &params
}

func generateImages(w io.Writer, m image.Image, p *sd.I2IParams, o string) {
	encoder := json.NewEncoder(w)
	sd.ProgressCB = func(step, steps int, time float32) {
		encoder.Encode(&progress{step, steps, time})
	}
	var ms sd.Images
	if m == nil { //T2I
		ms = gSD.T2I(&p.T2IParams)
	} else {
		p.InitImage = sd.NewImage(m)
		ms = gSD.I2I(p)
		p.InitImage.Close()
	}
	sd.ProgressCB = nil
	encoder.Encode(saveImages(ms, o))
}

func saveImages(ms sd.Images, output string) []string {
	var (
		outputPrefix string
		outputSuffix string
	)
	idx := strings.LastIndexByte(output, '.')
	if idx > 0 {
		outputPrefix = output[:idx]
		outputSuffix = output[idx:]
		if outputSuffix != ".png" {
			outputSuffix = ".jpg"
		}
	} else {
		outputPrefix = output
		outputSuffix = ".jpg"
	}
	var names []string
	for i, m := range ms.List() {
		var name string
		if i == 0 {
			name = outputPrefix + outputSuffix
		} else {
			name = fmt.Sprintf("%s_%d%s", outputPrefix, i, outputSuffix)
		}
		if outputSuffix == ".png" {
			m.WritePNG(name)
		} else {
			m.WriteJPG(name, nil)
		}
		m.Close()
		names = append(names, name)
	}
	ms.Close()
	return names
}

func readImageFromFile(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	m, _, err := image.Decode(f)
	f.Close()
	return m, err
}
