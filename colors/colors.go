package colors

// Generated code; DO NOT EDIT

// go generate: go run cmd/colorsgenerate/main.go --baseFile color.yml -o

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

var (
	Aqua = ColorData{
		ColorRGB: ColorRGB{
			B: int32(255),
			G: int32(255),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(255),
			G: int32(255),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "00ffff",
		HexRGBW: "00ffff00",
	}
	Indigo = ColorData{
		ColorRGB: ColorRGB{
			B: int32(210),
			G: int32(35),
			R: int32(56),
		},
		ColorRGBW: ColorRGBW{
			B: int32(210),
			G: int32(35),
			R: int32(56),
			W: int32(0),
		},
		HexRGB:  "3823d2",
		HexRGBW: "3823d200",
	}
	Brick = ColorData{
		ColorRGB: ColorRGB{
			B: int32(31),
			G: int32(31),
			R: int32(189),
		},
		ColorRGBW: ColorRGBW{
			B: int32(31),
			G: int32(31),
			R: int32(189),
			W: int32(0),
		},
		HexRGB:  "bd1f1f",
		HexRGBW: "bd1f1f00",
	}
	Sage = ColorData{
		ColorRGB: ColorRGB{
			B: int32(61),
			G: int32(177),
			R: int32(162),
		},
		ColorRGBW: ColorRGBW{
			B: int32(61),
			G: int32(177),
			R: int32(162),
			W: int32(0),
		},
		HexRGB:  "a2b13d",
		HexRGBW: "a2b13d00",
	}
	Lime = ColorData{
		ColorRGB: ColorRGB{
			B: int32(0),
			G: int32(255),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(0),
			G: int32(255),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "00ff00",
		HexRGBW: "00ff0000",
	}
	Lavender = ColorData{
		ColorRGB: ColorRGB{
			B: int32(213),
			G: int32(111),
			R: int32(172),
		},
		ColorRGBW: ColorRGBW{
			B: int32(213),
			G: int32(111),
			R: int32(172),
			W: int32(0),
		},
		HexRGB:  "ac6fd5",
		HexRGBW: "ac6fd500",
	}
	Maroon = ColorData{
		ColorRGB: ColorRGB{
			B: int32(85),
			G: int32(3),
			R: int32(176),
		},
		ColorRGBW: ColorRGBW{
			B: int32(85),
			G: int32(3),
			R: int32(176),
			W: int32(0),
		},
		HexRGB:  "b00355",
		HexRGBW: "b0035500",
	}
	Mustard = ColorData{
		ColorRGB: ColorRGB{
			B: int32(0),
			G: int32(197),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(0),
			G: int32(197),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ffc500",
		HexRGBW: "ffc50000",
	}
	Green = ColorData{
		ColorRGB: ColorRGB{
			B: int32(0),
			G: int32(192),
			R: int32(2),
		},
		ColorRGBW: ColorRGBW{
			B: int32(0),
			G: int32(192),
			R: int32(2),
			W: int32(0),
		},
		HexRGB:  "02c000",
		HexRGBW: "02c00000",
	}
	Purple = ColorData{
		ColorRGB: ColorRGB{
			B: int32(166),
			G: int32(7),
			R: int32(112),
		},
		ColorRGBW: ColorRGBW{
			B: int32(166),
			G: int32(7),
			R: int32(112),
			W: int32(0),
		},
		HexRGB:  "7007a6",
		HexRGBW: "7007a600",
	}
	Denim = ColorData{
		ColorRGB: ColorRGB{
			B: int32(187),
			G: int32(102),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(187),
			G: int32(102),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "0066bb",
		HexRGBW: "0066bb00",
	}
	Navy = ColorData{
		ColorRGB: ColorRGB{
			B: int32(128),
			G: int32(0),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(128),
			G: int32(0),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "000080",
		HexRGBW: "00008000",
	}
	Red = ColorData{
		ColorRGB: ColorRGB{
			B: int32(2),
			G: int32(0),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(2),
			G: int32(0),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ff0002",
		HexRGBW: "ff000200",
	}
	Pear = ColorData{
		ColorRGB: ColorRGB{
			B: int32(30),
			G: int32(216),
			R: int32(208),
		},
		ColorRGBW: ColorRGBW{
			B: int32(30),
			G: int32(216),
			R: int32(208),
			W: int32(0),
		},
		HexRGB:  "d0d81e",
		HexRGBW: "d0d81e00",
	}
	Mint = ColorData{
		ColorRGB: ColorRGB{
			B: int32(122),
			G: int32(255),
			R: int32(128),
		},
		ColorRGBW: ColorRGBW{
			B: int32(122),
			G: int32(255),
			R: int32(128),
			W: int32(0),
		},
		HexRGB:  "80ff7a",
		HexRGBW: "80ff7a00",
	}
	Pine = ColorData{
		ColorRGB: ColorRGB{
			B: int32(14),
			G: int32(114),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(14),
			G: int32(114),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "00720e",
		HexRGBW: "00720e00",
	}
	Grey = ColorData{
		ColorRGB: ColorRGB{
			B: int32(136),
			G: int32(136),
			R: int32(136),
		},
		ColorRGBW: ColorRGBW{
			B: int32(136),
			G: int32(136),
			R: int32(136),
			W: int32(0),
		},
		HexRGB:  "888888",
		HexRGBW: "88888800",
	}
	Teal = ColorData{
		ColorRGB: ColorRGB{
			B: int32(128),
			G: int32(128),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(128),
			G: int32(128),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "008080",
		HexRGBW: "00808000",
	}
	Blue = ColorData{
		ColorRGB: ColorRGB{
			B: int32(255),
			G: int32(0),
			R: int32(0),
		},
		ColorRGBW: ColorRGBW{
			B: int32(255),
			G: int32(0),
			R: int32(0),
			W: int32(0),
		},
		HexRGB:  "0000ff",
		HexRGBW: "0000ff00",
	}
	Pink = ColorData{
		ColorRGB: ColorRGB{
			B: int32(255),
			G: int32(0),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(255),
			G: int32(0),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ff00ff",
		HexRGBW: "ff00ff00",
	}
	Orange = ColorData{
		ColorRGB: ColorRGB{
			B: int32(0),
			G: int32(80),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(0),
			G: int32(80),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ff5000",
		HexRGBW: "ff500000",
	}
	Tangerine = ColorData{
		ColorRGB: ColorRGB{
			B: int32(0),
			G: int32(139),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(0),
			G: int32(139),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ff8b00",
		HexRGBW: "ff8b0000",
	}
	Yellow = ColorData{
		ColorRGB: ColorRGB{
			B: int32(0),
			G: int32(255),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(0),
			G: int32(255),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ffff00",
		HexRGBW: "ffff0000",
	}
	Rose = ColorData{
		ColorRGB: ColorRGB{
			B: int32(127),
			G: int32(54),
			R: int32(255),
		},
		ColorRGBW: ColorRGBW{
			B: int32(127),
			G: int32(54),
			R: int32(255),
			W: int32(0),
		},
		HexRGB:  "ff367f",
		HexRGBW: "ff367f00",
	}
)
