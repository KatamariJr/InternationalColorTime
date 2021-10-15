
// Generated code; DO NOT EDIT

//go:generate go run ../../cmd/colorsgenerate/main.go --baseFile ../../cmd/colorsgenerate/colors.yml -o .

typedef struct {
	int R;
	int G;
	int B;
} ColorRGB;

typedef struct {
	int R;
	int G;
	int B;
	int W;
} ColorRGBW;

typedef struct {
	char* HexRGB;
	char* HexRGBW;
	ColorRGB rgb;
	ColorRGBW rgbw;
} ColorData;


ColorData Aqua = {
	.HexRGB = "00ffff",
	.HexRGBW = "00ffff00",
	.rgb = {0, 255, 255},
	.rgbw = {0, 255, 255, 0}
};

ColorData Blue = {
	.HexRGB = "0000ff",
	.HexRGBW = "0000ff00",
	.rgb = {0, 0, 255},
	.rgbw = {0, 0, 255, 0}
};

ColorData Brick = {
	.HexRGB = "bd1f1f",
	.HexRGBW = "bd1f1f00",
	.rgb = {189, 31, 31},
	.rgbw = {189, 31, 31, 0}
};

ColorData Denim = {
	.HexRGB = "0066bb",
	.HexRGBW = "0066bb00",
	.rgb = {0, 102, 187},
	.rgbw = {0, 102, 187, 0}
};

ColorData Green = {
	.HexRGB = "02c000",
	.HexRGBW = "02c00000",
	.rgb = {2, 192, 0},
	.rgbw = {2, 192, 0, 0}
};

ColorData Grey = {
	.HexRGB = "888888",
	.HexRGBW = "88888800",
	.rgb = {136, 136, 136},
	.rgbw = {136, 136, 136, 0}
};

ColorData Indigo = {
	.HexRGB = "3823d2",
	.HexRGBW = "3823d200",
	.rgb = {56, 35, 210},
	.rgbw = {56, 35, 210, 0}
};

ColorData Lavender = {
	.HexRGB = "ac6fd5",
	.HexRGBW = "ac6fd500",
	.rgb = {172, 111, 213},
	.rgbw = {172, 111, 213, 0}
};

ColorData Lime = {
	.HexRGB = "00ff00",
	.HexRGBW = "00ff0000",
	.rgb = {0, 255, 0},
	.rgbw = {0, 255, 0, 0}
};

ColorData Maroon = {
	.HexRGB = "b00355",
	.HexRGBW = "b0035500",
	.rgb = {176, 3, 85},
	.rgbw = {176, 3, 85, 0}
};

ColorData Mint = {
	.HexRGB = "80ff7a",
	.HexRGBW = "80ff7a00",
	.rgb = {128, 255, 122},
	.rgbw = {128, 255, 122, 0}
};

ColorData Mustard = {
	.HexRGB = "ffc500",
	.HexRGBW = "ffc50000",
	.rgb = {255, 197, 0},
	.rgbw = {255, 197, 0, 0}
};

ColorData Navy = {
	.HexRGB = "000080",
	.HexRGBW = "00008000",
	.rgb = {0, 0, 128},
	.rgbw = {0, 0, 128, 0}
};

ColorData Orange = {
	.HexRGB = "ff5000",
	.HexRGBW = "ff500000",
	.rgb = {255, 80, 0},
	.rgbw = {255, 80, 0, 0}
};

ColorData Pear = {
	.HexRGB = "d0d81e",
	.HexRGBW = "d0d81e00",
	.rgb = {208, 216, 30},
	.rgbw = {208, 216, 30, 0}
};

ColorData Pine = {
	.HexRGB = "00720e",
	.HexRGBW = "00720e00",
	.rgb = {0, 114, 14},
	.rgbw = {0, 114, 14, 0}
};

ColorData Pink = {
	.HexRGB = "ff00ff",
	.HexRGBW = "ff00ff00",
	.rgb = {255, 0, 255},
	.rgbw = {255, 0, 255, 0}
};

ColorData Purple = {
	.HexRGB = "7007a6",
	.HexRGBW = "7007a600",
	.rgb = {112, 7, 166},
	.rgbw = {112, 7, 166, 0}
};

ColorData Red = {
	.HexRGB = "ff0000",
	.HexRGBW = "ff000000",
	.rgb = {255, 0, 0},
	.rgbw = {255, 0, 0, 0}
};

ColorData Rose = {
	.HexRGB = "ff367f",
	.HexRGBW = "ff367f00",
	.rgb = {255, 54, 127},
	.rgbw = {255, 54, 127, 0}
};

ColorData Sage = {
	.HexRGB = "a2b13d",
	.HexRGBW = "a2b13d00",
	.rgb = {162, 177, 61},
	.rgbw = {162, 177, 61, 0}
};

ColorData Tangerine = {
	.HexRGB = "ff8b00",
	.HexRGBW = "ff8b0000",
	.rgb = {255, 139, 0},
	.rgbw = {255, 139, 0, 0}
};

ColorData Teal = {
	.HexRGB = "008080",
	.HexRGBW = "00808000",
	.rgb = {0, 128, 128},
	.rgbw = {0, 128, 128, 0}
};

ColorData Yellow = {
	.HexRGB = "ffff00",
	.HexRGBW = "ffff0000",
	.rgb = {255, 255, 0},
	.rgbw = {255, 255, 0, 0}
};
