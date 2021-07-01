package internationalcolortime

import (
	"time"
)

type Color int

const hourNanoseconds = 3600000000000
const secondNanoseconds = 1000000000

// Colors are arranged according to the order they run in UTC.
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
	Count
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

var colorToStringShort = map[Color]string{
	Red:       "RED",
	Brick:     "BRK",
	Orange:    "ORG",
	Tangerine: "TNG",
	Mustard:   "MRD",
	Yellow:    "YLW",
	Pear:      "PER",
	Sage:      "SAG",
	Mint:      "MNT",
	Lime:      "LIM",
	Green:     "GRN",
	Pine:      "PIN",
	Grey:      "GRY",
	Aqua:      "AQA",
	Teal:      "TEL",
	Denim:     "DNM",
	Blue:      "BLU",
	Navy:      "NVY",
	Indigo:    "ING",
	Purple:    "PPL",
	Lavender:  "LVR",
	Maroon:    "MRN",
	Pink:      "PNK",
	Rose:      "ROS",
}

var stringToColor = make(map[string]Color, len(colorToString))
var stringShortToColor = make(map[string]Color, len(colorToStringShort))

func init() {
	//create stringToColor map
	for k, v := range colorToString {
		stringToColor[v] = k
	}

	//create stringShortToColor map
	for k, v := range colorToStringShort {
		stringShortToColor[v] = k
	}
}

// InternationalColorTime represents a time on the International Color Time clock. An ICT value does not contain any
// date information, only time, and as such there are no equivalent After() or Before() functions like there are in the
// standard time package, since they cannot be calculated with certainty.
type InternationalColorTime struct {
	hour  Color
	nanos int64 //nanoseconds past the hour
}

// TimeToICT will convert a time t into an InternationalColorTime.
func TimeToICT(t time.Time) InternationalColorTime {
	return ColorTime(Color(t.In(time.UTC).Hour()), t.Minute(), t.Second(), t.Nanosecond())
}

// ColorTime will return the color time struct for a given hour, minute, sec, and nanosecond.
func ColorTime(hour Color, min, sec, nsec int) InternationalColorTime {
	return InternationalColorTime{}.add(int(hour), min, sec, nsec)
}

//// Now returns the current color time.
//func Now() InternationalColorTime {
//
//}
//
// Parse a InternationalColorTime from value using a given layout string. This follows the rules laid out by the time.Time
// library. If the specific reference time is
//  Mon Jan 2 15:04:05 MST 2006
// then the corresponding InternationColorTime reference time is
//  PNK:04:05.999999999
// where PNK is the 3-letter code for your given ColorHour.
//func Parse(layout, value string) (InternationalColorTime, error) {
//
//}

// String implements the fmt.Stringer interface.
func (c Color) String() string {
	str, ok := colorToString[c]
	if !ok {
		return ""
	}
	return str
}

// IsZero returns true if this is a default, uninitialized time.
//func (i InternationalColorTime) IsZero() bool {
//
//}
//

// In returns a time.Time with the equivalent timezone time as denoted by location. Since ICT has no date component,
// the local date value is used. For a more specific translocation, use InDate()
func (i InternationalColorTime) In(location *time.Location) time.Time {
	t := time.Now()
	return i.InDate(location, t.Year(), t.Month(), t.Day())
}

// InDate returns a time.Time with the equivalent timezone time as denoted by location, and the accompanying date as
// specified by the year/month/day params.
// todo maybe the date info should be the date as it refers to the UTC date, then translate that backwards/forwards as needed as well
func (i InternationalColorTime) InDate(location *time.Location, year int, month time.Month, day int) time.Time {

	_, offset := time.Date(0, 0, 0, 0, 0, 0, 0, location).Zone()

	//calculate how far from utc to move the time
	//we know that ICT is a UTC time, so we just tick the nanos in a direction

	_ = offset * secondNanoseconds

	t := time.Date(year, month, day, 0, 0, 0, 0, location)
	//todo finish

	return t

}

// add an amount of duration onto a color time
func (i InternationalColorTime) add(hour, min, sec, nsec int) InternationalColorTime {
	ns := int64(nsec+(sec*1000000000)+(min*60*1000000000)) + i.nanos
	nsHourIncrs := int(ns / hourNanoseconds)
	if nsHourIncrs >= 1 {
		ns -= int64(nsHourIncrs * hourNanoseconds)
		hour = (hour + nsHourIncrs) % (int(Count) - 1)
	}
	return InternationalColorTime{
		hour:  Color(hour),
		nanos: ns,
	}
}

// Equal will compare two InternationalColorTimes to see if they are the same.
//func (i InternationalColorTime) Equal(i2 *InternationalColorTime) bool {
//
//}
//

// Add will tick i forward by duration.
//func (i InternationalColorTime) Add(duration time.Duration) InternationalColorTime {
//	duration.Nanoseconds()
//	//todo finish
//}

// Truncate returns the result of rounding t down to a multiple of d (since the zero time). If d <= 0, Truncate returns
// t stripped of any monotonic clock reading but otherwise unchanged.
//
// Truncate operates on the time as an absolute duration since the zero time; it does not operate on the presentation
// form of the time. Thus, Truncate(Hour) may return a time with a non-zero minute, depending on the time's Location.
//func (i InternationalColorTime) Truncate(d time.Duration) InternationalColorTime {
//
//}

// String returns the time formatted using the format string
//  PNK:04:05.999999999
// which is the InternationalColorTime equivalent of the default time format found in the standard "time"
//  package, 15:04:05.999999 -0700 MST.
//func (i InternationalColorTime) String() string {
//
//}
//

// Hour returns the hour portion of an InternationalColorTime.
func (i InternationalColorTime) Hour() Color {
	return i.hour
}

// Minute returns the minute portion of an InternationalColorTime.
func (i InternationalColorTime) Minute() int {
	return int((i.nanos / 60000000000) % 60)
}

// Second returns the second portion of an InternationalColorTime.
func (i InternationalColorTime) Second() int {
	return int((i.nanos / 1000000000) % 60)
}

// Nanosecond returns teh nanosecond portion of an InternationalColorTime.
func (i InternationalColorTime) Nanosecond() int {
	return int(i.nanos % 1000000000)
}

// Round i to the nearest interval expressed by d.
//func (i InternationalColorTime) Round(d time.Duration) InternationalColorTime {
//
//}
//
// Format i in the given format string. The formatting directive should include placeholder values from the
// InternationalColorTime reference string
//  PNK:04:05.999999999
//func (i InternationalColorTime) Format(format string) string {
//
//}
//
// Clock returns the hour, minute, and second components of an InternationalColorTime.
//func (i InternationalColorTime) Clock() (hour Color, min, sec int) {
//
//}
//
// GobEncode implements the encoding.GobEncoder interface.
//func (i InternationalColorTime) GobEncode() ([]byte, error) {
//
//}
//
// GobDecode implements the encoding.GobDecoder interface.
//func (i InternationalColorTime) GobDecode(data []byte) error {
//
//}
//
// Local will return the time.Time object that corresponds to the InternationalColorTime equivalent for your local system
// time zone. Since InternationalColorTime structs do not contain date components, the local date will be used.
//func (i InternationalColorTime) Local() time.Time {
//
//}
//
//// MarshalBinary implements the encoding.BinaryMarshaler interface.
//func (i InternationalColorTime) MarshalBinary() ([]byte, error) {
//
//}
//
//// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//func (i *InternationalColorTime) UnmarshalBinary(data []byte) error {
//
//}
//
//// MarshalJSON implements the json.Marshaler interface. The time is a quoted string in RFC 3339 format, with sub-second
//// precision added if present.
//func (i InternationalColorTime) MarshalJSON() ([]byte, error) {
//
//}
//
//// UnmarshalJSON implements the json.Unmarshaler interface. The time is expected to be a quoted string in RFC 3339
//// format.
//func (i *InternationalColorTime) UnmarshalJSON(data []byte) error {
//
//}
//
//// MarshalText implements the encoding.TextMarshaler interface. The time is formatted in RFC 3339 format, with
//// sub-second precision added if present.
//func (i InternationalColorTime) MarshalText() ([]byte, error) {
//
//}
//
//// UnmarshalText implements the encoding.TextUnmarshaler interface. The time is expected to be in RFC 3339 format.
//func (i *InternationalColorTime) UnmarshalText(data []byte) error {
//
//}
