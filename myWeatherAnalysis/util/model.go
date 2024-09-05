package util

type S1 struct {
	Name        string `json:"s1 name"`
	CounterPart S2     `json:"counter part"`
}

type S2 struct {
	Name string `json:"s2 name"`
}
