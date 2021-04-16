package internationalcolortime

import (
	"reflect"
	"testing"
	"time"
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
