package index

import (
	"fmt"
	"testing"
)

const FILEPATH_TO_INDEX_EXAMLE = "./test_indices/unit_test_index.xml"

func (ind *Index) GetFilepath() (filepath string) {
	return ind.indexPath
}

func (ind *Index) RemoveEveryElement() {
	ind.removeAllElements()
}

func GetSampleIndex() *Index {
	ind := New(FILEPATH_TO_INDEX_EXAMLE)
	ind.Elements = make([]*indexElement, 3, 3)

	ind.Elements[0] = &indexElement{PoemId: "ID_1", Path: "Path_1"}
	ind.Elements[1] = &indexElement{PoemId: "ID_2", Path: "Path_2"}
	ind.Elements[2] = &indexElement{PoemId: "ID_3", Path: "Path_3"}

	return ind
}

func (ind1 *Index) Equal(ind2 *Index) (equal bool, inequalityMessage string) {
	if ind1 == nil || ind2 == nil {
		if ind1 == nil && ind2 == nil {
			return true, ""
		}

		return false, "One of the indices is nil, the other one isn't."
	}

	if ind1.indexPath != ind2.indexPath {
		return false, "Index path I doesn't equal index path II."
	}

	if len(ind1.Elements) != len(ind2.Elements) {
		return false, "Number of element in index I unequal to number of elements in index II."
	}
	for i, elem := range ind1.Elements {
		if !elem.Equal(ind2.Elements[i]) {
			return false, fmt.Sprintf("Elements %d of indices unequal", i)
		}
	}

	return true, ""
}

func TestIndexEqual(t *testing.T) {
	//Setting test variables
	indexNoElements := GetSampleIndex()
	indexNoElements.Elements = nil

	indexOneDiffElement := GetSampleIndex()
	indexOneDiffElement.Elements[1].PoemId = "Different_value"

	indexDifferentPath := GetSampleIndex()
	indexDifferentPath.indexPath = "./different_path"

	indexDifferentElementNumber := GetSampleIndex()
	indexDifferentElementNumber.Elements = append(indexDifferentPath.Elements, &indexElement{PoemId: "_", Path: "_"})

	cases := []struct {
		ind1        *Index
		ind2        *Index
		expectEqual bool
	}{
		{nil, nil, true},
		{nil, New(""), false},
		{New(""), nil, false},
		{GetSampleIndex(), GetSampleIndex(), true},
		{GetSampleIndex(), indexNoElements, false},
		{indexOneDiffElement, GetSampleIndex(), false},
		{indexOneDiffElement, indexOneDiffElement, true},
		{GetSampleIndex(), indexDifferentPath, false},
		{indexDifferentElementNumber, GetSampleIndex(), false},
	}

	for i, aCase := range cases {
		equals, _ := aCase.ind1.Equal(aCase.ind2)

		if equals && !aCase.expectEqual {
			t.Errorf("Error in case %d. Expected not equal, got equal.", i)
		}
		if !equals && aCase.expectEqual {
			t.Errorf("Error in case %d. Expected equal, got not equal.", i)
		}
	}
}

func (elem1 *indexElement) Equal(elem2 *indexElement) bool {
	if elem1 == nil || elem2 == nil {
		if elem1 == nil && elem2 == nil {
			return true
		}
		return false
	}

	if elem1.Path != elem2.Path || elem1.PoemId != elem2.PoemId {
		return false
	}

	return true
}

func TestElementsEqual(t *testing.T) {
	cases := []struct {
		elem1       *indexElement
		elem2       *indexElement
		expectEqual bool
	}{
		{nil, nil, true},
		{&indexElement{}, nil, false},
		{nil, &indexElement{}, false},
		{&indexElement{PoemId: "2", Path: "./aPath"}, &indexElement{PoemId: "2", Path: "./aPath"}, true},
		{&indexElement{PoemId: "9", Path: "./aPath"}, &indexElement{PoemId: "2", Path: "./aPath"}, false},
		{&indexElement{PoemId: "2", Path: "./aPath"}, &indexElement{PoemId: "2", Path: "./otherPath"}, false},
	}

	for i, aCase := range cases {
		equals := aCase.elem1.Equal(aCase.elem2)
		if equals && !aCase.expectEqual {
			t.Errorf("Error in case %d. Expected not equal, got equal", i)
		}
		if !equals && aCase.expectEqual {
			t.Errorf("Error in case %d. Expected equal, got not equal", i)
		}
	}
}
