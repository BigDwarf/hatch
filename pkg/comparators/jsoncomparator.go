package comparators

import (
	"github.com/sirupsen/logrus"
	"hatch/pkg/util"
)

type jsonComparator struct {
}

func NewJsonComparator() *jsonComparator {
	return &jsonComparator{}
}

func (comp *jsonComparator) Compare(inputFile1, inputFile2 string) (error, bool) {

	err, firstFileContents := util.ReadFromFile(inputFile1)
	if err != nil {
		return err, false
	}

	err, secondFileContents := util.ReadFromFile(inputFile2)
	if err != nil {
		return err, false
	}

	firstArraySize := len(firstFileContents)
	secondArraySize := len(secondFileContents)

	if firstArraySize != secondArraySize {
		logrus.Info("JSON array sizes are not identical.")
		return nil, false
	}

	hashedElemCount := make(map[string]int)

	for _, value := range firstFileContents {
		hashedElemCount[util.AsSha256(castToArrayItemType(value))] += 1
	}
	for _, value := range secondFileContents {
		hashedElemCount[util.AsSha256(castToArrayItemType(value))] -= 1
	}
	identical := true
	for _, value := range hashedElemCount {
		if value != 0 {
			identical = false
			break
		}
	}
	return nil, identical
}

// Definition of a json array item
type jsonRecord struct {
	Name string
	Id string
}

func castToArrayItemType(o map[string]interface{}) interface{} {
	return &jsonRecord{
		Name: o["name"].(string),
		Id: o["id"].(string),
	}
}
