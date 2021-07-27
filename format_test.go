package internationalcolortime

import "testing"

func TestInternationalColorTime_Format(t *testing.T) {
	type args struct {
		layout string
	}
	tests := []struct {
		name  string
		field InternationalColorTime
		args  args
		want  string
	}{
		{
			name:  "Long Name",
			field: ColorTime(Blue, 15, 0, 0),
			args: args{
				layout: ICTStandardLongName,
			},
			want: "Blue:15:00",
		},
		{
			name:  "Long Name 2",
			field: ColorTime(Pink, 59, 33, 0),
			args: args{
				layout: ICTStandardLongName,
			},
			want: "Pink:59:33",
		},
		{
			name:  "Long Name 3 w/ nanos",
			field: ColorTime(Denim, 22, 17, 999000000),
			args: args{
				layout: ICTStandardLongName,
			},
			want: "Denim:22:17.999",
		},
		{
			name:  "Short Name",
			field: ColorTime(Blue, 15, 0, 0),
			args: args{
				layout: ICTStandardShortName,
			},
			want: "BLU:15:00",
		},
		{
			name:  "Short Name 2",
			field: ColorTime(Green, 45, 22, 0),
			args: args{
				layout: ICTStandardShortName,
			},
			want: "GRN:45:22",
		},
		{
			name:  "Short Name 3 w/ nanos",
			field: ColorTime(Yellow, 5, 2, 987623443),
			args: args{
				layout: ICTStandardShortName,
			},
			want: "YLW:05:02.987623443",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.field
			if got := i.Format(tt.args.layout); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
