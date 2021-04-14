package internationalcolortime

import "time"

type Color int

// Color iotas are arranged by UTC.
const (
	Red Color = iota
	Brick
	Orange
	Tangerine
	Mustard
	Yellow
	Pear
	Sage
	Mint
	Lime
	Green
	Pine
	Grey
	Aqua
	Teal
	Denim
	Blue
	Navy
	Indigo
	Purple
	Lavender
	Maroon
	Pink
	Rose
)

var colorToString = map[Color]string{
	Red:       "Red",
	Brick:     "Brick",
	Orange:    "Orange",
	Tangerine: "Tangerine",
	Mustard:   "Mustard",
	Yellow:    "Yellow",
	Pear:      "Pear",
	Sage:      "Sage",
	Mint:      "Mint",
	Lime:      "Lime",
	Green:     "Green",
	Pine:      "Pine",
	Grey:      "Grey",
	Aqua:      "Aqua",
	Teal:      "Teal",
	Denim:     "Denim",
	Blue:      "Blue",
	Navy:      "Navy",
	Indigo:    "Indigo",
	Purple:    "Purple",
	Lavender:  "Lavender",
	Maroon:    "Maroon",
	Pink:      "Pink",
	Rose:      "Rose",
}

// InternationalColorTime represents a time on the International Color Time clock. An ICT value does not contain any
// date information, only time, and as such there are no equivalent After() or Before() functions like there are in the
// standard time package, since they cannot be calculated with certainty.
type InternationalColorTime struct {
	hour    Color
	minutes int
}

func TimeToICT(t time.Time) InternationalColorTime {

}

func ColorTime(hour Color, min, sec, nsec int) InternationalColorTime {

}

// Now returns the current color time.
func Now() InternationalColorTime {

}

func Parse(layout, value string) (InternationalColorTime, error) {

}

// String implements the fmt.Stringer interface.
func (c Color) String() string {
	str, ok := colorToString[c]
	if !ok {
		return ""
	}
	return str
}

func (i InternationalColorTime) IsZero() bool {

}

func (i InternationalColorTime) In(location time.Location) time.Time {

}

func (i InternationalColorTime) Equal(i2 *InternationalColorTime) bool {

}

func (i InternationalColorTime) Add(duration time.Duration) InternationalColorTime {

}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time). If d <= 0, Truncate returns
// t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the zero time; it does not operate on the presentation
// form of the time. Thus, Truncate(Hour) may return a time with a non-zero minute, depending on the time's Location.
func (i InternationalColorTime) Truncate(d time.Duration) InternationalColorTime {

}

// String returns the time formatted using the format string
//  "PNK:04:05.999999999"
// which is the ICT equivalent of the default time format found in the standard "time" package, 15:04:05.999999 -0700 MST.
func (i InternationalColorTime) String() string {

}

func (i InternationalColorTime) Hour() Color {

}

func (i InternationalColorTime) Minute() int {

}

func (i InternationalColorTime) Second() int {

}

func (i InternationalColorTime) Nanosecond() int {

}

func (i InternationalColorTime) Unix() int64 {

}

func (i InternationalColorTime) UnixNano() int64 {

}

func (i InternationalColorTime) Round(d time.Duration) InternationalColorTime {

}

func (i InternationalColorTime) Format(format string) string {

}

func (i InternationalColorTime) Clock() (hour Color, min, sec int) {

}

func (i InternationalColorTime) GobEncode() ([]byte, error) {

}

func (i InternationalColorTime) GobDecode(data []byte) error {

}

func (i InternationalColorTime) Local() time.Time {

}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (i InternationalColorTime) MarshalBinary() ([]byte, error) {

}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
func (i *InternationalColorTime) UnmarshalBinary(data []byte) error {

}

// MarshalJSON implements the json.Marshaler interface. The time is a quoted string in RFC 3339 format, with sub-second
// precision added if present.
func (i InternationalColorTime) MarshalJSON() ([]byte, error) {

}

// UnmarshalJSON implements the json.Unmarshaler interface. The time is expected to be a quoted string in RFC 3339
// format.
func (i *InternationalColorTime) UnmarshalJSON(data []byte) error {

}

// MarshalText implements the encoding.TextMarshaler interface. The time is formatted in RFC 3339 format, with
// sub-second precision added if present.
func (i InternationalColorTime) MarshalText() ([]byte, error) {

}

// UnmarshalText implements the encoding.TextUnmarshaler interface. The time is expected to be in RFC 3339 format.
func (i *InternationalColorTime) UnmarshalText(data []byte) error {

}
