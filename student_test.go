package coverage

import "testing"

func TestConcatenator(t *testing.T) {
	want := "GolangUnited"
	_, got := concatenator("Golang", "United")
	if got != want {
		t.Errorf("Error: want and got do not match\nwant: %v, got: %v\n", want, got)
	}
}

/*func TestSummator(t *testing.T) {
	want := 10
	got := summator(4, 6)
	if got != want {
		t.Errorf("Error: want and got do not match\nwant: %v, got: %v\n", want, got)
	}
}*/
