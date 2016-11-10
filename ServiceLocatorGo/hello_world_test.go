package ServiceLocatorGo

import "testing"

func TestHelloWorld(t *testing.T) {
	uut := ServiceLocatorGo{}

	cases := []struct {
		in, want string
	}{
		{"", "Hello, World"},
	}
	for _, c := range cases {
		got, _ := uut.HelloWorld()
		if got != c.want {
			t.Errorf("HelloWorld() == %q, want %q", got, c.want)
		}
	}
}
