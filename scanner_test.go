package fingerprint

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func Test_ScannerIdentify(t *testing.T) {
	type args struct {
		strObjects []fmt.Stringer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should reproduce a hash the mock device object",
			args: args{strObjects: []fmt.Stringer{mockDevice{"foobar"}}},
			want: "216dfbe4-a2ed-acf2-f363-dafeac2fbd21",
		},
		{
			name: "should reproduce a hash the from a mock object and string",
			args: args{strObjects: []fmt.Stringer{mockDevice{"foobar"}, Stringer("hello world")}},
			want: "40ae8a2a-7b65-6441-0f22-cbc22925f517",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc := NewScanner()
			sc.SetHashKeyWithSeed(1)

			got, err := sc.Identify(tt.args.strObjects...)
			if err != nil {
				t.Fatalf("Scanner.Identify() error = %v", err)
			}

			if got != tt.want {
				t.Fatalf("Scanner.Identify() = %v, want %v", got, tt.want)
			}

			_, err = uuid.Parse(got)
			if err != nil {
				t.Fatalf("Scanner.Identify() did not generate a valid uuid in form err = %v", err)
			}
		})
	}
}

//BenchmarkScannerIdentifyNothing-12            	 6482433	       178 ns/op	      96 B/op	       3 allocs/op
//BenchmarkScannerIdentifyNothing-12            	 6619848	       181 ns/op	      96 B/op	       3 allocs/op
//BenchmarkScannerIdentifyNothing-12            	 6306696	       182 ns/op	      96 B/op	       3 allocs/op
func BenchmarkScannerIdentifyNothing(b *testing.B) {
	sc := NewScanner()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, _ := sc.Identify()
		TestString = x
	}
}

//BenchmarkScannerIdentifyWithStringer-12       	 4960764	       264 ns/op	     112 B/op	       4 allocs/op
//BenchmarkScannerIdentifyWithStringer-12       	 4376788	       267 ns/op	     112 B/op	       4 allocs/op
//BenchmarkScannerIdentifyWithStringer-12       	 4555501	       269 ns/op	     112 B/op	       4 allocs/op
func BenchmarkScannerIdentifyWithStringer(b *testing.B) {
	sc := NewScanner()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, _ := sc.Identify(mockDevice{"foobar"})
		TestString = x
	}
}

//BenchmarkScannerIdentifyWithStringerTwo-12    	 3737829	       274 ns/op	     112 B/op	       4 allocs/op
//BenchmarkScannerIdentifyWithStringerTwo-12    	 4413387	       266 ns/op	     112 B/op	       4 allocs/op
//BenchmarkScannerIdentifyWithStringerTwo-12    	 3742868	       294 ns/op	     112 B/op	       4 allocs/op
func BenchmarkScannerIdentifyWithStringerTwo(b *testing.B) {
	sc := NewScanner()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, _ := sc.Identify(mockDevice{"foobar"}, Stringer("hello world"))
		TestString = x
	}
}
