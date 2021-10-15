package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	colors "github.com/katamarijr/internationalcolortime/structs"

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

const goGenerateTag = "//go:generate go run ../../cmd/colorsgenerate/main.go --baseFile ../../cmd/colorsgenerate/colors.yml -o ."

func main() {

	baseFile := pflag.String("baseFile", "", "filename of the yml to use for generating code. Required")
	outputFileNameBase := pflag.StringP("outputDirectory", "o", "", "directory of the output files. Required")

	pflag.Parse()

	if outputFileNameBase == nil || baseFile == nil {
		pflag.PrintDefaults()
		os.Exit(1)
	}

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

	goOutputName := fmt.Sprintf("%s/colors.go", *outputFileNameBase)
	err = generateGoCode(sortedColorData, goOutputName)
	if err != nil {
		log.Fatal(err)
	}

	cOutputName := fmt.Sprintf("%s/colors.c", *outputFileNameBase)
	err = generateCCode(sortedColorData, cOutputName)
	if err != nil {
		log.Fatal(err)
	}

}

func generateGoCode(sortedColorData []mappedColorData, outputFilename string) error {
	f := NewFile("colors")
	f.Comment("Generated code; DO NOT EDIT")
	f.Line()
	f.Comment(goGenerateTag)
	f.Line()

	f.ImportName("github.com/katamarijr/internationalcolortime/structs", "structs")

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

		rgbVars = append(rgbVars, Id(colorData.Name).Op("=").Qual("github.com/katamarijr/internationalcolortime/structs", topLevelColorStructName).Values(Dict{
			Id("HexRGB"):  Lit(fmt.Sprintf("%06x", hex)),
			Id("HexRGBW"): Lit(hexWString),
			Id(rgbColorStructName): Qual("github.com/katamarijr/internationalcolortime/structs", rgbColorStructName).Values(Dict{
				Id("R"): Lit(c.R),
				Id("G"): Lit(c.G),
				Id("B"): Lit(c.B),
			}),
			Id(rgbwColorStructName): Qual("github.com/katamarijr/internationalcolortime/structs", rgbwColorStructName).Values(Dict{
				Id("R"): Lit(cw.R),
				Id("G"): Lit(cw.G),
				Id("B"): Lit(cw.B),
				Id("W"): Lit(cw.W),
			}),
		}))

	}

	f.Var().Defs(rgbVars...)

	return f.Save(outputFilename)
}

func generateCCode(sortedColorData []mappedColorData, outputFilename string) error {
	f, err := os.OpenFile(outputFilename, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}

	mainStructs := `
// Generated code; DO NOT EDIT

` + goGenerateTag + `

typedef struct {
	int R;
	int G;
	int B;
} ` + rgbColorStructName + `;

typedef struct {
	int R;
	int G;
	int B;
	int W;
} ` + rgbwColorStructName + `;

typedef struct {
	char* HexRGB;
	char* HexRGBW;
	ColorRGB rgb;
	ColorRGBW rgbw;
} ` + topLevelColorStructName + `;

`

	const theStructStuffFormat = `
` + topLevelColorStructName + ` %s = {
	.HexRGB = "%s",
	.HexRGBW = "%s",
	.rgb = {%d, %d, %d},
	.rgbw = {%d, %d, %d, %d}
};
`
	colorStructs := []string{}

	for _, colorData := range sortedColorData {
		hex := colorData.Hex
		c := colors.ColorRGB{
			R: (hex & hexRMask) >> 16,
			G: (hex & hexGMask) >> 8,
			B: hex & hexBMask,
		}

		hexW := colorData.HexW
		hexWString := ""

		cw := colors.ColorRGBW{}

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

		hexString := fmt.Sprintf("%06x", hex)

		s := fmt.Sprintf(theStructStuffFormat, colorData.Name,
			hexString, hexWString,
			cw.R, cw.G, cw.B,
			cw.R, cw.G, cw.B, cw.W)

		colorStructs = append(colorStructs, s)
	}

	//dump it all into the text file
	_, err = f.WriteString(mainStructs)
	if err != nil {
		return err
	}

	for _, v := range colorStructs {
		_, err = f.WriteString(v)
		if err != nil {
			return err
		}
	}

	return f.Close()
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
