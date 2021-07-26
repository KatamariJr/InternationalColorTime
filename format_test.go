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
			name:  "Short Name",
			field: ColorTime(Blue, 15, 0, 0),
			args: args{
				layout: ICTStandardShortName,
			},
			want: "BLU:15:00",
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
