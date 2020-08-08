package demo

import "testing"

func TestAddTest(t *testing.T){

	res := addUpper(4)
	if res != 10{
		t.Fatalf("add error")
	}
	t.Logf("test addUpper succ")

}
