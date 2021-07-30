package internationalcolortime

import (
	"reflect"
	"testing"
	"time"
)

//timezones
var (
	BST = time.FixedZone("BST", 1*60*60)  //(+01:00)
	MST = time.FixedZone("MST", -7*60*60) //(-07:00)
	EST = time.FixedZone("EST", -5*60*60) //(-05:00)
	KST = time.FixedZone("KST", 9*60*60)  //(+09:00)
)

func TestColor_String(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want string
	}{
		{
			name: "A Good Color 1",
			c:    Yellow,
			want: "Yellow",
		},
		{
			name: "A Good Color 2",
			c:    Green,
			want: "Green",
		},
		{
			name: "A Bad Color 1",
			c:    99999,
			want: "",
		},
		{
			name: "A Bad Color 2",
			c:    -99999,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorTime(t *testing.T) {
	type args struct {
		hour Color
		min  int
		sec  int
		nsec int
	}
	tests := []struct {
		name string
		args args
		want InternationalColorTime
	}{
		{
			name: "test zero",
			args: args{
				hour: Red,
				min:  0,
				sec:  0,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  0,
				nanos: 0,
			},
		},
		{
			name: "test time min",
			args: args{
				hour: Red,
				min:  1,
				sec:  0,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 60000000000,
			},
		},
		{
			name: "test time sec",
			args: args{
				hour: Red,
				min:  0,
				sec:  1,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 1000000000,
			},
		},
		{
			name: "test time nsec",
			args: args{
				hour: Red,
				min:  0,
				sec:  0,
				nsec: 1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 1,
			},
		},
		{
			name: "test time sec and nsec",
			args: args{
				hour: Red,
				min:  0,
				sec:  1,
				nsec: 1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 1000000001,
			},
		},
		{
			name: "test time min and nsec",
			args: args{
				hour: Red,
				min:  1,
				sec:  0,
				nsec: 1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 60000000001,
			},
		},
		{
			name: "test time min and sec",
			args: args{
				hour: Red,
				min:  1,
				sec:  1,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 61000000000,
			},
		},
		{
			name: "test time min and sec and nsec",
			args: args{
				hour: Red,
				min:  1,
				sec:  1,
				nsec: 1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 61000000001,
			},
		},
		{
			name: "test time min and sec and nsec diff color",
			args: args{
				hour: Rose,
				min:  1,
				sec:  1,
				nsec: 1,
			},
			want: InternationalColorTime{
				hour:  Rose,
				nanos: 61000000001,
			},
		},
		{
			name: "Accept values outside the normal range for durations",
			args: args{
				hour: Red,
				min:  0,
				sec:  100,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 100000000000,
			},
		},
		{
			name: "Something that ticks over the hour mark should increment the hour value and reset nanos",
			args: args{
				hour: Red,
				min:  60,
				sec:  0,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  Brick,
				nanos: 0,
			},
		},
		{
			name: "Something that ticks over the hour mark should increment the hour value and reset nanos and keep them incrementing",
			args: args{
				hour: Red,
				min:  60,
				sec:  30,
				nsec: 0,
			},
			want: InternationalColorTime{
				hour:  Brick,
				nanos: 30000000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorTime(tt.args.hour, tt.args.min, tt.args.sec, tt.args.nsec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_String1(t *testing.T) {
	tests := []struct {
		name string
		c    Color
		want string
	}{
		{
			name: "Red",
			c:    Red,
			want: "Red",
		},
		{
			name: "Teal",
			c:    Teal,
			want: "Teal",
		},
		{
			name: "Lime",
			c:    Lime,
			want: "Lime",
		},
		{
			name: "Rose",
			c:    Rose,
			want: "Rose",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_Hour(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Color
	}{
		{
			name: "Yellow",
			fields: fields{
				hour:  Yellow,
				nanos: 99999999999999,
			},
			want: Yellow,
		},
		{
			name: "Yellow, nanos dont matter",
			fields: fields{
				hour:  Yellow,
				nanos: 2222,
			},
			want: Yellow,
		},
		{
			name: "Rose",
			fields: fields{
				hour:  Rose,
				nanos: 2222,
			},
			want: Rose,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.Hour(); got != tt.want {
				t.Errorf("Hour() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_Minute(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test 1",
			fields: fields{
				hour:  Green,
				nanos: 60000000000,
			},
			want: 1,
		},
		{
			name: "test 1, hours dont matter",
			fields: fields{
				hour:  Indigo,
				nanos: 60000000000,
			},
			want: 1,
		},
		{
			name: "Should tick over if mins is greater than or equal to 60",
			fields: fields{
				hour:  Indigo,
				nanos: 3600000000001,
			},
			want: 0,
		},
		{
			name: "Big numbah",
			fields: fields{
				hour:  Indigo,
				nanos: 3540000000000,
			},
			want: 59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.Minute(); got != tt.want {
				t.Errorf("Minute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeToICT(t *testing.T) {

	tests := []struct {
		name string
		t    time.Time
		want InternationalColorTime
	}{
		{
			name: "time 1",
			t:    time.Date(2000, 7, 11, 12, 0, 0, 0, time.UTC),
			want: ColorTime(Grey, 0, 0, 0),
		},
		{
			name: "time 2",
			t:    time.Date(1000, 9, 11, 16, 0, 0, 0, time.UTC),
			want: ColorTime(Blue, 0, 0, 0),
		},
		{
			name: "time 3",
			t:    time.Date(1000, 9, 11, 20, 0, 0, 0, time.UTC),
			want: ColorTime(Lavender, 0, 0, 0),
		},
		{
			name: "time 3 w/ mins+sec",
			t:    time.Date(1000, 9, 11, 20, 30, 30, 0, time.UTC),
			want: ColorTime(Lavender, 30, 30, 0),
		},
		{
			name: "Testing a time in a non-UTC timezone",
			t:    time.Date(1500, 9, 11, 12, 0, 0, 0, time.FixedZone("UTC-8", -8*60*60)),
			want: ColorTime(Lavender, 0, 0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeToICT(tt.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeToICT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_Nanosecond(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test 1",
			fields: fields{
				hour:  Tangerine,
				nanos: 1000,
			},
			want: 1000,
		},
		{
			name: "test 1, hour doesnt matter",
			fields: fields{
				hour:  Denim,
				nanos: 1000,
			},
			want: 1000,
		},
		{
			name: "should tick over if its past the amount of nanos for one second",
			fields: fields{
				hour:  Tangerine,
				nanos: 1000000001,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.Nanosecond(); got != tt.want {
				t.Errorf("Nanosecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_Second(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test 1",
			fields: fields{
				hour:  Tangerine,
				nanos: 1000000000,
			},
			want: 1,
		},
		{
			name: "test 1, hour doesnt matter",
			fields: fields{
				hour:  Denim,
				nanos: 1000000000,
			},
			want: 1,
		},
		{
			name: "should tick over if its past the amount of nanos for one minute",
			fields: fields{
				hour:  Tangerine,
				nanos: 61000000000,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.Second(); got != tt.want {
				t.Errorf("Second() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_InDate(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	type args struct {
		location *time.Location
		year     int
		month    time.Month
		day      int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		{
			name: "UTC timezone should always be the same",
			fields: fields{
				hour:  Grey,
				nanos: 1000000000,
			},
			args: args{
				location: time.FixedZone("UTC", 0),
				year:     2000,
				month:    4,
				day:      7,
			},
			want: time.Date(2000, 4, 7, 12, 0, 1, 0, time.UTC),
		},
		{
			name: "BST test",
			fields: fields{
				hour:  Grey,
				nanos: 1000000000,
			},
			args: args{
				location: BST,
				year:     2000,
				month:    4,
				day:      7,
			},
			want: time.Date(2000, 4, 7, 13, 0, 1, 0, BST),
		},
		{
			name: "MST test",
			fields: fields{
				hour:  Grey,
				nanos: 1000000000,
			},
			args: args{
				location: MST,
				year:     2000,
				month:    4,
				day:      7,
			},
			want: time.Date(2000, 4, 7, 5, 0, 1, 0, MST),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.InDate(tt.args.location, tt.args.year, tt.args.month, tt.args.day); !got.Equal(tt.want) {
				t.Errorf("InDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_add(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	type args struct {
		hour int
		min  int
		sec  int
		nsec int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   InternationalColorTime
	}{
		//nanoseconds
		{
			name: "add 1 nanosecond, no tickover",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				0, 0, 0, 1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 1,
			},
		},
		{
			name: "subtract 1 nanosecond, no tickover",
			fields: fields{
				hour:  Red,
				nanos: 500,
			},
			args: args{
				0, 0, 0, -1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 499,
			},
		},
		{
			name: "add 1 nanosecond, with tickover",
			fields: fields{
				hour:  Red,
				nanos: 3599999999999,
			},
			args: args{
				0, 0, 0, 1,
			},
			want: InternationalColorTime{
				hour:  Brick,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 nanosecond, with tickover",
			fields: fields{
				hour:  Brick,
				nanos: 0,
			},
			args: args{
				0, 0, 0, -1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 3599999999999,
			},
		},
		{
			name: "subtract 1 nanosecond, with tickover #2",
			fields: fields{
				hour:  Blue,
				nanos: 0,
			},
			args: args{
				0, 0, 0, -1,
			},
			want: InternationalColorTime{
				hour:  Denim,
				nanos: 3599999999999,
			},
		},
		{
			name: "add 1 nanosecond, with tickover and wraparound",
			fields: fields{
				hour:  Rose,
				nanos: 3599999999999,
			},
			args: args{
				0, 0, 0, 1,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 nanosecond, with tickover and wraparound",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				0, 0, 0, -1,
			},
			want: InternationalColorTime{
				hour:  Rose,
				nanos: 3599999999999,
			},
		},

		//seconds
		{
			name: "add 1 second, no tickover",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				0, 0, 1, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 1000000000,
			},
		},
		{
			name: "subtract 1 second, no tickover",
			fields: fields{
				hour:  Red,
				nanos: 1000000500,
			},
			args: args{
				0, 0, -1, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 500,
			},
		},
		{
			name: "add 1 second, with tickover",
			fields: fields{
				hour:  Red,
				nanos: 3599000000000,
			},
			args: args{
				0, 0, 1, 0,
			},
			want: InternationalColorTime{
				hour:  Brick,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 second, with tickover",
			fields: fields{
				hour:  Brick,
				nanos: 0,
			},
			args: args{
				0, 0, -1, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 3599000000000,
			},
		},
		{
			name: "add 1 second, with tickover and wraparound",
			fields: fields{
				hour:  Rose,
				nanos: 3599000000000,
			},
			args: args{
				0, 0, 1, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 second, with tickover and wraparound",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				0, 0, -1, 0,
			},
			want: InternationalColorTime{
				hour:  Rose,
				nanos: 3599000000000,
			},
		},

		//minutes
		{
			name: "add 1 minute, no tickover",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				0, 1, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 60000000000,
			},
		},
		{
			name: "subtract 1 minute, no tickover",
			fields: fields{
				hour:  Red,
				nanos: 60000000500,
			},
			args: args{
				0, -1, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 500,
			},
		},
		{
			name: "add 1 minute, with tickover",
			fields: fields{
				hour:  Red,
				nanos: 3540000000000,
			},
			args: args{
				0, 1, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Brick,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 minute, with tickover",
			fields: fields{
				hour:  Brick,
				nanos: 0,
			},
			args: args{
				0, -1, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 3540000000000,
			},
		},
		{
			name: "add 1 minute, with tickover and wraparound",
			fields: fields{
				hour:  Rose,
				nanos: 3540000000000,
			},
			args: args{
				0, 1, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 minute, with tickover and wraparound",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				0, -1, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Rose,
				nanos: 3540000000000,
			},
		},

		//hour
		{
			name: "add 1 hour, with tickover",
			fields: fields{
				hour:  Red,
				nanos: 500,
			},
			args: args{
				1, 0, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Brick,
				nanos: 500,
			},
		},
		{
			name: "subtract 1 hour, with tickover",
			fields: fields{
				hour:  Brick,
				nanos: 0,
			},
			args: args{
				-1, 0, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 0,
			},
		},
		{
			name: "add 1 hour, with tickover and wraparound",
			fields: fields{
				hour:  Rose,
				nanos: 0,
			},
			args: args{
				1, 0, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Red,
				nanos: 0,
			},
		},
		{
			name: "subtract 1 hour, with tickover and wraparound",
			fields: fields{
				hour:  Red,
				nanos: 0,
			},
			args: args{
				-1, 0, 0, 0,
			},
			want: InternationalColorTime{
				hour:  Rose,
				nanos: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.add(tt.args.hour, tt.args.min, tt.args.sec, tt.args.nsec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_Add(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   InternationalColorTime
	}{
		{
			name: "add 5 seconds",
			fields: fields{
				hour:  Green,
				nanos: 0,
			},
			args: args{
				time.Second * 5,
			},
			want: InternationalColorTime{
				hour:  Green,
				nanos: 5000000000,
			},
		},
		{
			name: "add 7 hours",
			fields: fields{
				hour:  Green,
				nanos: 0,
			},
			args: args{
				time.Hour * 7,
			},
			want: InternationalColorTime{
				hour:  Navy,
				nanos: 0,
			},
		},
		{
			name: "add 4 hours 20 minutes",
			fields: fields{
				hour:  Green,
				nanos: 0,
			},
			args: args{
				(time.Hour * 4) + (time.Minute * 20),
			},
			want: InternationalColorTime{
				hour:  Teal,
				nanos: 1200000000000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := InternationalColorTime{
				hour:  tt.fields.hour,
				nanos: tt.fields.nanos,
			}
			if got := i.Add(tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_String(t *testing.T) {
	type fields struct {
		hour  Color
		nanos int64
	}
	tests := []struct {
		name   string
		fields InternationalColorTime
		want   string
	}{
		{
			name:   "no nanos",
			fields: ColorTime(Yellow, 14, 50, 0),
			want:   "YLW:14:50",
		},
		{
			name:   "yes nanos",
			fields: ColorTime(Yellow, 14, 50, 999),
			want:   "YLW:14:50.000000999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.fields
			if got := i.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInternationalColorTime_Clock(t *testing.T) {
	tests := []struct {
		name     string
		time     InternationalColorTime
		wantHour Color
		wantMin  int
		wantSec  int
	}{
		{
			name:     "Test 1",
			time:     ColorTime(Indigo, 5, 33, 0),
			wantHour: Indigo,
			wantMin:  5,
			wantSec:  33,
		},
		{
			name:     "Test 2",
			time:     ColorTime(Indigo, 45, 01, 0),
			wantHour: Indigo,
			wantMin:  45,
			wantSec:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.time
			gotHour, gotMin, gotSec := i.Clock()
			if gotHour != tt.wantHour {
				t.Errorf("Clock() gotHour = %v, want %v", gotHour, tt.wantHour)
			}
			if gotMin != tt.wantMin {
				t.Errorf("Clock() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotSec != tt.wantSec {
				t.Errorf("Clock() gotSec = %v, want %v", gotSec, tt.wantSec)
			}
		})
	}
}

func TestInternationalColorTime_Truncate(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		time InternationalColorTime
		args args
		want InternationalColorTime
	}{
		{
			name: "any value of d that is greater than one hour is reduced to one hour",
			time: ColorTime(Green, 30, 17, 6),
			args: args{
				d: time.Hour + time.Second,
			},
			want: ColorTime(Green, 0, 0, 0),
		},
		{
			name: "Truncate to hour",
			time: ColorTime(Green, 30, 17, 6),
			args: args{
				d: time.Hour,
			},
			want: ColorTime(Green, 0, 0, 0),
		},
		{
			name: "Truncate to minute",
			time: ColorTime(Green, 30, 17, 6),
			args: args{
				d: time.Minute,
			},
			want: ColorTime(Green, 30, 0, 0),
		},
		{
			name: "Truncate to 8 minutes",
			time: ColorTime(Green, 30, 17, 6),
			args: args{
				d: time.Minute * 8,
			},
			want: ColorTime(Green, 24, 0, 0),
		},
		{
			name: "Truncate to second",
			time: ColorTime(Green, 30, 17, 6),
			args: args{
				d: time.Second,
			},
			want: ColorTime(Green, 30, 17, 0),
		},
		{
			name: "Truncate to 5 seconds",
			time: ColorTime(Green, 30, 17, 6),
			args: args{
				d: time.Second * 5,
			},
			want: ColorTime(Green, 30, 15, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.time
			if got := i.Truncate(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}
