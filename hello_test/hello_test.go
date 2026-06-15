package hello_test_package

// testing is a standard lib that provides T type and helpers like t.Errorf.
import (
	hellopackage "hellomodule/hellodirectory"
	"testing"
)

// this takes a pointer to testing Type, and they do something.
func TestSayHello(t *testing.T) {
	want := "Hello, test!"
	got := hellopackage.Say("test")
	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}

func TestSayHello2(t *testing.T) {
	want := "Hello, Steven, Kim, Alex!"
	got := hellopackage.Say2([]string{"Steven", "Kim", "Alex"})
	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
