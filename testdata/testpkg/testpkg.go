package testpkg

type TestStruct struct {
	A int
	B string
}

func TestMe() TestStruct {
	return TestStruct{
		A: 1,
		B: "2",
	}
}
