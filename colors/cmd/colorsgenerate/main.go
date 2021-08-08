package main

import (
	"fmt"
	"log"
	"os"

	"github.com/katamarijr/internationalcolortime/colors"

	"github.com/go-yaml/yaml"
	"github.com/spf13/pflag"

	. "github.com/dave/jennifer/jen"
)

type ColorGenerate struct {
	Colors map[string]struct {
		Hex  int32 `yaml:"hex"`
		HexW int64 `yaml:"hexW"`
	} `yaml:"colors"`
}

const (
	hexRMask = 0xff0000
	hexGMask = 0x00ff00
	hexBMask = 0x0000ff

	hexRwMask int64 = 0xff000000
	hexGwMask int64 = 0x00ff0000
	hexBwMask int64 = 0x0000ff00
	hexWwMask int64 = 0x000000ff
)

func main() {

	baseFile := pflag.String("baseFile", "", "filename of the yml to use for generating code")
	outputFilename := pflag.StringP("outputFile", "o", "", "output location of the output file")

	pflag.Parse()

	ymlFile, err := os.Open(*baseFile)
	if err != nil {
		log.Fatal(err)
	}

	colorsParse := &ColorGenerate{}

	err = yaml.NewDecoder(ymlFile).Decode(colorsParse)
	if err != nil {
		log.Fatal(err)
	}

	//sortedColorNames, sortedColorHex, sortedColorHexW := sortColors(*colorsParse)

	f := NewFile("colors")
	f.Comment("Generated code; DO NOT EDIT")
	f.Line()
	f.Comment("go generate: go run cmd/colorsgenerate/main.go --baseFile color.yml -o ")
	f.Line()

	f.Type().Id("ColorData").Struct(
		Id("HexRGB").String(),
		Id("HexRGBW").String(),
		Id("ColorRGB"),
		Id("ColorRGBW"),
	)

	f.Type().Id("ColorRGB").Struct(
		Id("R").Int32(),
		Id("G").Int32(),
		Id("B").Int32(),
	)

	f.Type().Id("ColorRGBW").Struct(
		Id("R").Int32(),
		Id("G").Int32(),
		Id("B").Int32(),
		Id("W").Int32(),
	)

	rgbVars := []Code{}

	//create colorRGB values
	for k, v := range colorsParse.Colors {
		c := colors.ColorRGB{
			R: (v.Hex & hexRMask) >> 16,
			G: (v.Hex & hexGMask) >> 8,
			B: v.Hex & hexBMask,
		}

		cw := colors.ColorRGBW{}
		hexW := ""

		if v.HexW != 0 {
			cw = colors.ColorRGBW{
				R: int32((v.HexW & hexRwMask) >> 24),
				G: int32((v.HexW & hexGwMask) >> 16),
				B: int32((v.HexW & hexBwMask) >> 8),
				W: int32(v.HexW & hexWwMask),
			}
			hexW = fmt.Sprintf("%08x", v.HexW)
		} else {
			cw = colors.ColorRGBW{
				R: c.R,
				G: c.G,
				B: c.B,
				W: 0,
			}
			hexW = fmt.Sprintf("%06x00", v.Hex)
		}

		rgbVars = append(rgbVars, Id(k).Op("=").Id("ColorData").Values(Dict{
			Id("HexRGB"):  Lit(fmt.Sprintf("%06x", v.Hex)),
			Id("HexRGBW"): Lit(hexW),
			Id("ColorRGB"): Id("ColorRGB").Values(Dict{
				Id("R"): Lit(c.R),
				Id("G"): Lit(c.G),
				Id("B"): Lit(c.B),
			}),
			Id("ColorRGBW"): Id("ColorRGBW").Values(Dict{
				Id("R"): Lit(cw.R),
				Id("G"): Lit(cw.G),
				Id("B"): Lit(cw.B),
				Id("W"): Lit(cw.W),
			}),
		}))

	}

	f.Var().Defs(rgbVars...)

	err = f.Save(*outputFilename)
	if err != nil {
		log.Fatal(err)
	}

}

func sortColors(g ColorGenerate) (names []string, hex []int32, hexW []int64) {
	//loop through the map and turn it into a determinate sorted slice so the outputs will be in order. sort by member name
	// Convert map to slice of keys.
	keys := []string{}
	for key, _ := range g.Colors {
		keys = append(keys, key)
	}

	// Convert map to slice of values.
	hexValues := []int32{}
	for _, value := range g.Colors {
		hexValues = append(hexValues, value.Hex)
	}
	hexWValues := []int64{}
	for _, value := range g.Colors {
		hexWValues = append(hexWValues, value.HexW)
	}

	return nil, nil, nil
}
