package envy

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

const (
	NIL = "<nil>"
)

var parselnTests = []struct {
	// input
	in string
	// output
	key string
	val string
	err string
}{
	{"PORT=9090", "PORT", "9090", NIL},
	{"PORT9090", "", "", "missing delimiter '='"},
	{"PORT =9090", "PORT", "9090", NIL},
	{`PORT="9090"`, "PORT", "9090", NIL},
	{`PORT='9090'`, "PORT", "9090", NIL},
	{"PORT= 9090", "PORT", "9090", NIL},
}

func Test_Simple_Parseln(t *testing.T) {
	for _, tt := range parselnTests {
		key, val, err := Parseln(tt.in)
		expect(t, key, tt.key)
		expect(t, val, tt.val)
		expect(t, fmt.Sprint(err), tt.err)
	}
}

func Test_Load(t *testing.T) {
	buf := bytes.NewBufferString("PORT=9090\nMARTINI_ENV=dev\nHELLO='world'")

	err := Load(buf)
	expect(t, fmt.Sprint(err), NIL)
	expect(t, os.Getenv("PORT"), "9090")
	expect(t, os.Getenv("MARTINI_ENV"), "dev")
	expect(t, os.Getenv("HELLO"), "world")
}

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
