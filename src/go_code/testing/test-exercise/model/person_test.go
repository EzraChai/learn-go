package model

import (
	"testing"
)

func TestPerson_SerializePerson(t *testing.T) {
	person := &Person{
		Name:  "Chloe Gan",
		Age:   16,
		Hobby: "Read books",
	}
	err := person.SerializePerson()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPerson_DeserializePerson(t *testing.T) {
	person := &Person{}
	err := person.DeserializePerson()
	if err != nil {
		t.Fatal(err)
	}
	personCopied := &Person{
		Name:  "Chloe Gan",
		Age:   16,
		Hobby: "Read books",
	}
	if *person != *personCopied {
		t.Fatal("person not same")
	}
}
