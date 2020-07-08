package fingerprint

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestIdentify(t *testing.T) {
	SetHashKeyWithSeed(1)
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
			got, err := Identify(tt.args.strObjects...)
			if err != nil {
				t.Fatalf("Identify() error = %v", err)
			}

			if got != tt.want {
				t.Fatalf("Identify() = %v, want %v", got, tt.want)
			}

			_, err = uuid.Parse(got)
			if err != nil {
				t.Fatalf("Identify() did not generate a valid uuid in form err = %v", err)
			}
		})
	}
}

var TestString string

//BenchmarkIdentifyNothing-12              7456538               155 ns/op             272 B/op          3 allocs/op
//BenchmarkIdentifyNothing-12              7818858               156 ns/op             272 B/op          3 allocs/op
//BenchmarkIdentifyNothing-12              7684363               157 ns/op             272 B/op          3 allocs/op
func BenchmarkIdentifyNothing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, _ := Identify()
		TestString = x
	}
}

//BenchmarkIdentifyWithStringer-12         5619817               216 ns/op             288 B/op          4 allocs/op
//BenchmarkIdentifyWithStringer-12         5560032               214 ns/op             288 B/op          4 allocs/op
//BenchmarkIdentifyWithStringer-12         5631970               214 ns/op             288 B/op          4 allocs/op
func BenchmarkIdentifyWithStringer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, _ := Identify(mockDevice{"foobar"})
		TestString = x
	}
}

//BenchmarkIdentifyWithStringerTwo-12      5182345               228 ns/op             288 B/op          4 allocs/op
//BenchmarkIdentifyWithStringerTwo-12      5208236               225 ns/op             288 B/op          4 allocs/op
//BenchmarkIdentifyWithStringerTwo-12      5305508               225 ns/op             288 B/op          4 allocs/op
func BenchmarkIdentifyWithStringerTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x, _ := Identify(mockDevice{"foobar"}, Stringer("hello world"))
		TestString = x
	}
}

type mockDevice struct {
	ua string
}

func (m mockDevice) String() string {
	return m.ua
}
