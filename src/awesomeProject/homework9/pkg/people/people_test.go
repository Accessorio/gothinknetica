package people

import (
	"bytes"
	"log"
	"reflect"
	"testing"
)

func TestPeople_MaxAge(t *testing.T) {
	f := &Client{1, "White", 18}
	s := &Employee{1, "Black", 21}
	th := &Client{2, "Green", 1}

	got, err := MaxAge(f, s, th)
	if err != "OK" {
		log.Fatalf("ERROR: %v", err)
	}
	want := s.AGE
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MaxAge() = %v, want %v", got, want)
	}
}

func TestPeople_MaxAgeAny(t *testing.T) {
	f := Client{1, "White", 18}
	s := Employee{1, "Black", 21}
	th := Client{2, "Green", 1}

	got := MaxAgeAny(f, s, th)
	want := s
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MaxAgeAny() = %v, want %v", got, want)
	}
}

func TestPeople_OnlyStrings(t *testing.T) {
	f := &Client{1, "White", 18}
	s1, s2, s3 := "1", "2", "3"
	s := bytes.NewBuffer([]byte{})
	got := OnlyStrings(s, s1, s2, s3, f)
	want := []string{"1", "2", "3"}
	if reflect.DeepEqual(got, want) != true {
		t.Fatalf("получили %v, ожидалось %v", got, want)
	}
}
