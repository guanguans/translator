package translator

import (
	"testing"
)

// go test -v -cover -coverprofile=cover.out
// go tool cover -func=cover.out
// go tool cover -html=cover.out
func TestReturnSelf(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "apple",
			args: args{str: "apple"},
			want: "apple",
		},
		{
			name: "banana",
			args: args{str: "banana"},
			want: "banana",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReturnSelf(tt.args.str); got != tt.want {
				t.Errorf("ReturnSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
