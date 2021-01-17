package comparators

import (
	"github.com/Flaque/filet"
	"testing"
)

func TestEqualFilesAreEqual(t *testing.T) {
	jsonComparator := NewJsonComparator()

	defer filet.CleanUp(t)

	file1 := filet.TmpFile(t, "", "[\n  {\n  \"name\": \"azaliia\",\n  \"id\": \"11111\",\n  \"asdasd\": [3, 1, 2]\n  },\n  {\n  \"name\": \"arslan\",\n  \"id\": \"sadasdads\",\n  \"asdasd\": [4, 2, 3]\n  }\n]")
	file2 := filet.TmpFile(t, "", "[\n  {\n  \"name\": \"azaliia\",\n  \"id\": \"11111\",\n  \"asdasd\": [3, 1, 2]\n  },\n  {\n  \"name\": \"arslan\",\n  \"id\": \"sadasdads\",\n  \"asdasd\": [4, 2, 3]\n  }\n]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res != true {
		t.Errorf("Equal files are not threated as equal")
	}
}

func TestTopLevelArraysAreSorted(t *testing.T) {
	jsonComparator := NewJsonComparator()

	defer filet.CleanUp(t)

	file1 := filet.TmpFile(t, "", "[\n  {\n  \"name\": \"azaliia\",\n  \"id\": \"11111\",\n  \"asdasd\": [3, 1, 2]\n  },\n  {\n  \"name\": \"arslan\",\n  \"id\": \"sadasdads\",\n  \"asdasd\": [4, 2, 3]\n  }\n]")
	file2 := filet.TmpFile(t, "", "[\n  {\n  \"name\": \"azaliia\",\n  \"id\": \"11111\",\n  \"asdasd\": [3, 2, 1]\n  },\n  {\n  \"name\": \"arslan\",\n  \"id\": \"sadasdads\",\n  \"asdasd\": [2, 4, 3]\n  }\n]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res != true {
		t.Errorf("Top level arrays with pod types are not sorted properly")
	}
}

func TestTopLevelObjectOrderDoesntMatter(t *testing.T) {
	jsonComparator := NewJsonComparator()
	defer filet.CleanUp(t)
	file1 := filet.TmpFile(t, "","[\n  {\n  \"name\": \"azaliia\",\n  \"id\": \"11111\",\n  \"asdasd\": [3, 1, 2]\n  },\n  {\n  \"name\": \"arslan\",\n  \"id\": \"sadasdads\",\n  \"asdasd\": [4, 2, 3]\n  }\n]")
	file2 := filet.TmpFile(t, "","[\n  {\n    \"name\": \"arslan\",\n    \"id\": \"sadasdads\",\n    \"asdasd\": [4, 2, 3]\n  },\n  {\n    \"name\": \"azaliia\",\n    \"id\": \"11111\",\n    \"asdasd\": [3, 1, 2]\n  }\n]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res != true {
		t.Errorf("Top level object order matters")
	}
}

func TestInnerArraysAreSorted(t *testing.T) {
	jsonComparator := NewJsonComparator()

	defer filet.CleanUp(t)
	file1 := filet.TmpFile(t, "","[\n  {\n  \"a\": \"a\",\n  \"b\": \"a\",\n  \"c\": [3, 1, 2],\n  \"d\": {\n      \"b\": [5, 6, 7],\n      \"c\": [9, 8, 7],\n      \"d\": {\n        \"a\": [1, 3, 2]\n      }\n    }\n  }\n]")
	file2 := filet.TmpFile(t, "","[\n  {\n    \"a\": \"a\",\n    \"b\": \"a\",\n    \"c\": [3, 1, 2],\n    \"d\": {\n      \"b\": [7, 6, 5],\n      \"c\": [7, 8, 9],\n      \"d\": {\n        \"a\": [1, 2, 3]\n      }\n    }\n  }\n]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res != true {
		t.Errorf("Top level arrays with pod types are not sorted properly")
	}
}

func TestInnerArraysNotEqual(t *testing.T) {
	jsonComparator := NewJsonComparator()

	defer filet.CleanUp(t)
	file1 := filet.TmpFile(t, "","[\n  {\n  \"a\": \"a\",\n  \"b\": \"a\",\n  \"c\": [3, 1, 2],\n  \"d\": {\n      \"b\": [5, 6, 7],\n      \"c\": [9, 8, 7],\n      \"d\": {\n        \"a\": [1, 3, 2]\n      }\n    }\n  }\n]")
	file2 := filet.TmpFile(t, "","[\n  {\n    \"a\": \"a\",\n    \"b\": \"a\",\n    \"c\": [3, 1, 2],\n    \"d\": {\n      \"b\": [7, 6, 5],\n      \"c\": [7, 8, 9],\n      \"d\": {\n        \"a\": [1, 3]\n      }\n    }\n  }\n]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res == true {
		t.Errorf("Json's are not equal, but program says opposite :)")
	}
}

func TestEmptyJsonArray(t *testing.T) {
	jsonComparator := NewJsonComparator()
	defer filet.CleanUp(t)
	file1 := filet.TmpFile(t, "","[]")
	file2 := filet.TmpFile(t, "","[]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res != true {
		t.Errorf("Empty json arrays are not equal")
	}
}

func TestInnerObjectOrderDoesntMatter(t *testing.T) {
	jsonComparator := NewJsonComparator()
	defer filet.CleanUp(t)
	file1 := filet.TmpFile(t, "","[\n  {\n  \"a\": \"a\",\n  \"b\": \"a\",\n  \"c\": [3, 1, 2],\n  \"d\": {\n      \"c\": [9, 8, 7],\n      \"b\": [5, 6, 7],\n      \"d\": {\n        \"a\": [1, 3, 2]\n      }\n    }\n  }\n]")
	file2 := filet.TmpFile(t, "","[\n  {\n    \"a\": \"a\",\n    \"b\": \"a\",\n    \"c\": [3, 1, 2],\n    \"d\": {\n      \"b\": [7, 6, 5],\n      \"c\": [7, 8, 9],\n      \"d\": {\n        \"a\": [1, 2, 3]\n      }\n    }\n  }\n]")

	if _, res := jsonComparator.Compare(file1.Name(), file2.Name()); res != true {
		t.Errorf("Inner object order matters, but it shouldn't")
	}
}