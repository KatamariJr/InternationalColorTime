package structs

type ColorData struct {
	HexRGB  string
	HexRGBW string
	ColorRGB
	ColorRGBW
}
type ColorRGB struct {
	R int32
	G int32
	B int32
}
type ColorRGBW struct {
	R int32
	G int32
	B int32
	W int32
}
