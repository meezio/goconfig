package goflags

import (
	"os"
	"testing"
)

type testStruct struct {
	A int    `flag:"A" flagDefault:"100"`
	B string `flag:"B" flagDefault:"200"`
	C string
	D bool `cfg:"D" cfgDefault:"true"`
	F float64
	G float64 `cfg:"G" cfgDefault:"3.05"`
	N string  `flag:"-"`
	M int
	p string
	S testSub `flag:"S"`
}

type testSub struct {
	A int        `flag:"A" flagDefault:"300"`
	B string     `flag:"C" flagDefault:"400"`
	S testSubSub `flag:"S"`
}

type testSubSub struct {
	A int    `flag:"A" flagDefault:"500"`
	B string `flag:"S" flagDefault:"600"`
}

func TestParse(t *testing.T) {

	//os.Args = []string{"noop", "-flag1=val1", "arg1", "arg2"}
	//os.Args = []string{"program", "-h"}

	Setup("flag", "flagDefault", "flagDescription")

	os.Args = []string{
		"program",
		"-a=900",
		"-b=TEST",
		"-d=true",
		"-s_s_a=99999",
		"-f=23.6",
	}

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}

	Preserve = false
	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	if s.A != 900 {
		t.Fatal("s.A != 900, s.A:", s.A)
	}

	if s.B != "TEST" {
		t.Fatal("s.B != \"TEST\", s.B:", s.B)
	}

	if !s.D {
		t.Fatal("s.D == true, s.D:", s.D)
	}

	if s.F != 23.6 {
		t.Fatal("s.F != 23.6, s.F:", s.F)
	}

	if s.S.S.A != 99999 {
		t.Fatal("s.S.S.A != 99999, s.S.S.A :", s.S.S.A)
	}

}

func TestPreserve(t *testing.T) {

	//os.Args = []string{"noop", "-flag1=val1", "arg1", "arg2"}
	//os.Args = []string{"program", "-h"}

	os.Args = []string{
		"program",
		"-a=8888",
		"-b=TEST",
		"-s_s_a=99999",
	}

	s := &testStruct{A: 1, S: testSub{A: 1, B: "2"}}

	Reset()
	Preserve = true
	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	if s.S.A != 1 {
		t.Fatal("s.S.A != 1, s.S.A:", s.S.A)
	}

	if s.A != 8888 {
		t.Fatal("s.A != 8888, s.A:", s.A)
	}

	if s.B != "TEST" {
		t.Fatal("s.B != \"TEST\", s.B:", s.B)
	}

	if s.S.S.A != 99999 {
		t.Fatal("s.S.S.A != 99999, s.S.S.A:", s.S.S.A)
	}
}
