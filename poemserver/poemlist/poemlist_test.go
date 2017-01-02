package poemlist_test

import (
	"fmt"
	. "poemXML/poemserver/poemlist"
	"poemXML/utils"
	"testing"
)

func TestNewPoemList(t *testing.T) {
	aPoemList := New()
	if aPoemList == nil {
		t.Error("New PoemList initialized to zero")
	}
}

//TODO: Convert to table-driven test
func TestAddPoemIds(t *testing.T) {
	//Three ids case
	poemListManyIds := New()
	idsToBeAdded := []string{"ID_1", "ID_2", "ID_3"}
	poemListManyIds.AddPoemIds(idsToBeAdded)
	idsAdded := poemListManyIds.PoemList()
	if !utils.SliceOfStringEqual(idsToBeAdded, idsAdded) {
		t.Errorf("Ids added to poem list incorrectly.")
	}

	//Double addition of ids
	moreIds := []string{"ID_4"}
	allExpectedIds := append(idsToBeAdded, moreIds...)
	poemListManyIds.AddPoemIds(moreIds)
	idsAdded = poemListManyIds.PoemList()
	if !utils.SliceOfStringEqual(allExpectedIds, idsAdded) {
		t.Errorf("On second addition ids added to poem list incorrectly.")
	}

	//No ids case
	poemListNoIds := New()
	idsToBeAdded = []string{}
	poemListNoIds.AddPoemIds(idsToBeAdded)
	idsAdded = poemListNoIds.PoemList()
	if !utils.SliceOfStringEqual(idsToBeAdded, idsAdded) {
		t.Errorf("Ids added to poem list incorrectly. Expected no ids to be added")
	}
}

func TestPoemListString(t *testing.T) {
	//Empty list case
	aPoemList := New()
	descriptionString := aPoemList.String()
	expectedString := fmt.Sprintf(DESCRIPTION_TEMPLATE, 0) + "\n"
	if descriptionString != expectedString {
		t.Errorf("Expected poem list string expected: %s, got %s", expectedString, descriptionString)
	}

	//List with 3 elements
	poemListManyIds := New()
	idsToBeAdded := []string{"ID_1", "ID_2", "ID_3"}
	poemListManyIds.AddPoemIds(idsToBeAdded)
	descriptionString = poemListManyIds.String()
	expectedString = fmt.Sprintf(DESCRIPTION_TEMPLATE, len(idsToBeAdded)) + "\n"
	if descriptionString != expectedString {
		t.Errorf("Expected poem list string expected: %s, got %s", expectedString, descriptionString)
	}
}
