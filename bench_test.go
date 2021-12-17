package translator

import "testing"

// go test -bench=. -benchmem ./...
func BenchmarkReturnSelf(b *testing.B) {
	type args struct {
		str string
	}

	benchmarks := []struct {
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

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ReturnSelf(bm.args.str)
			}
		})
	}
}
