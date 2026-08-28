package model

import "testing"

func TestEntitiesValidation(t *testing.T) {
	if !NewRecord("1", "g", "p", "", 1).Valid() {
		t.Fatal()
	}
	if (Profile{ID: "p", DisplayName: "P"}).Valid() == false {
		t.Fatal()
	}
}
