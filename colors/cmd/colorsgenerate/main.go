package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/katamarijr/internationalcolortime/colors"

	"github.com/go-yaml/yaml"
	"github.com/spf13/pflag"

	. "github.com/dave/jennifer/jen"
)

type ColorDefinitionFile struct {
	Colors map[string]colorHexValues `yaml:"colors"`
}

type colorHexValues struct {
	Hex  int32 `yaml:"hex"`
	HexW int64 `yaml:"hexW"`
}

type mappedColorData struct {
	Name string
	colorHexValues
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

// generator struct names
const (
	topLevelColorStructName = "ColorData"
	rgbColorStructName      = "ColorRGB"
	rgbwColorStructName     = "ColorRGBW"
)

func main() {

	baseFile := pflag.String("baseFile", "", "filename of the yml to use for generating code")
	outputFilename := pflag.StringP("outputFile", "o", "", "output location of the output file")

	pflag.Parse()

	ymlFile, err := os.Open(*baseFile)
	if err != nil {
		log.Fatal(err)
	}

	colorsParse := &ColorDefinitionFile{}

	err = yaml.NewDecoder(ymlFile).Decode(colorsParse)
	if err != nil {
		log.Fatal(err)
	}

	sortedColorData := sortColors(*colorsParse)

	f := NewFile("colors")
	f.Comment("Generated code; DO NOT EDIT")
	f.Line()
	f.Comment("//go:generate go run cmd/colorsgenerate/main.go --baseFile colors.yml -o colors.go")
	f.Line()

	f.Type().Id(topLevelColorStructName).Struct(
		Id("HexRGB").String(),
		Id("HexRGBW").String(),
		Id(rgbColorStructName),
		Id(rgbwColorStructName),
	)

	f.Type().Id(rgbColorStructName).Struct(
		Id("R").Int32(),
		Id("G").Int32(),
		Id("B").Int32(),
	)

	f.Type().Id(rgbwColorStructName).Struct(
		Id("R").Int32(),
		Id("G").Int32(),
		Id("B").Int32(),
		Id("W").Int32(),
	)

	rgbVars := []Code{}

	//create colorRGB values
	for _, colorData := range sortedColorData {
		hex := colorData.Hex
		c := colors.ColorRGB{
			R: (hex & hexRMask) >> 16,
			G: (hex & hexGMask) >> 8,
			B: hex & hexBMask,
		}

		cw := colors.ColorRGBW{}
		hexW := colorData.HexW
		hexWString := ""

		if hexW != 0 {
			cw = colors.ColorRGBW{
				R: int32((hexW & hexRwMask) >> 24),
				G: int32((hexW & hexGwMask) >> 16),
				B: int32((hexW & hexBwMask) >> 8),
				W: int32(hexW & hexWwMask),
			}
			hexWString = fmt.Sprintf("%08x", hexW)
		} else {
			cw = colors.ColorRGBW{
				R: c.R,
				G: c.G,
				B: c.B,
				W: 0,
			}
			hexWString = fmt.Sprintf("%06x00", hex)
		}

		rgbVars = append(rgbVars, Id(colorData.Name).Op("=").Id(topLevelColorStructName).Values(Dict{
			Id("HexRGB"):  Lit(fmt.Sprintf("%06x", hex)),
			Id("HexRGBW"): Lit(hexWString),
			Id(rgbColorStructName): Id(rgbColorStructName).Values(Dict{
				Id("R"): Lit(c.R),
				Id("G"): Lit(c.G),
				Id("B"): Lit(c.B),
			}),
			Id(rgbwColorStructName): Id(rgbwColorStructName).Values(Dict{
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

func sortColors(g ColorDefinitionFile) (theColorData []mappedColorData) {
	//loop through the map and turn it into a determinate sorted slice so the outputs will be in order. sort by member name
	// Convert map to slice of keys.
	theColorData = []mappedColorData{}
	for key, colors := range g.Colors {
		c := mappedColorData{
			Name: key,
			colorHexValues: colorHexValues{
				Hex:  colors.Hex,
				HexW: colors.HexW,
			},
		}
		theColorData = append(theColorData, c)
	}

	//sort the array
	sort.Slice(theColorData, func(i, j int) bool {
		return theColorData[i].Name < theColorData[j].Name
	})

	return theColorData
}
