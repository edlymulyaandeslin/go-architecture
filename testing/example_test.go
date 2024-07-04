package testing

import "testing"

func TestSayHello_Success(t *testing.T) {
	expected := "hello Desi"
	actual, err := SayHello("Desi")

	if err != nil {
		t.Fatal("SayHello() failed")
	}
	if actual != expected {
		t.Fatalf(`SayHello() failed, actual %v, expected %v, nil`, actual, expected)
	}
}

func TestSayHello_Failed(t *testing.T) {
	expected := ""
	actual, err := SayHello("")

	if err == nil {
		t.Fatal("SayHello() failed")
	}
	if actual != expected {
		t.Fatalf(`SayHello() failed, actual %v, expected %v, nil`, actual, expected)
	}
}
