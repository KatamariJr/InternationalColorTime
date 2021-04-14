package internationalcolortime

import "testing"

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
