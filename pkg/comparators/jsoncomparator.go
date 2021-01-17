package comparators

import (
	"github.com/sirupsen/logrus"
	"hatch/pkg/util"
	"reflect"
	"sort"
)

const (
	delta = 0.001
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

	comp.sortInnerPODArrays(firstFileContents)
	comp.sortInnerPODArrays(secondFileContents)
	comp.sortObjectArrays(firstFileContents)
	comp.sortObjectArrays(secondFileContents)

	for _, elem := range firstFileContents {
		hashedElemCount[util.AsSha256(elem)] += 1
	}
	for _, elem := range secondFileContents {
		hashedElemCount[util.AsSha256(elem)] -= 1
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

func (comp *jsonComparator) sortInnerPODArrays(input []map[string]interface{}) {
	for _, item := range input {
		comp.sortInnerPODArraysRecurse(item)
	}
}

func (comp *jsonComparator) sortInnerPODArraysRecurse(input map[string]interface{}) {
	for _, v := range input {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			comp.sortPODJsonArray(v.([]interface{}))
		case reflect.Map:
			comp.sortInnerPODArraysRecurse(v.(map[string]interface{}))
		default:
			logrus.Trace("Pod type, not interesting")
		}
	}
}

func (comp *jsonComparator) sortPODJsonArray(arr []interface{}) {
	if len(arr) == 0 {
		return
	}
	switch reflect.TypeOf(arr[0]).Kind() {
	case reflect.Float64:
		sort.Slice(arr, func(i, j int) bool {
			intValueI := arr[i].(float64)
			intValueJ := arr[j].(float64)
			return (intValueI - intValueJ) < delta
		})
	case reflect.String:
		sort.Slice(arr, func(i, j int) bool {
			stringI := arr[i].(string)
			stringJ := arr[j].(string)
			return stringI > stringJ
		})

	case reflect.Bool:
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].(bool) && !arr[j].(bool)
		})
	case reflect.Map:
		for _, item := range arr {
			comp.sortInnerPODArraysRecurse(item.(map[string]interface{}))
		}
	case reflect.Slice:
		for _, item := range arr {
			comp.sortPODJsonArray(item.([]interface{}))
		}
	default:
		logrus.Info("Not interesting...")
	}
}

func (comp *jsonComparator) sortObjectArrays(input []map[string]interface{}) {
	for _, item := range input {
		comp.sortObjectArraysRecurse(item)
	}
}

func (comp *jsonComparator) sortObjectArraysRecurse(input map[string]interface{}) {
	for _, v := range input {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			comp.sortObjectJsonArray(v.([]interface{}))
		case reflect.Map:
			comp.sortObjectArraysRecurse(v.(map[string]interface{}))
		default:
			logrus.Trace("Pod type, not interesting")
		}
	}
}
func (comp *jsonComparator) sortObjectJsonArray(arr []interface{}) {
	if len(arr) == 0 {
		return
	}
	switch reflect.TypeOf(arr[0]).Kind() {
	case reflect.Map:
		for _, item := range arr {
			comp.sortObjectArraysRecurse(item.(map[string]interface{}))
		}
	case reflect.Slice:
		for _, item := range arr {
			comp.sortObjectJsonArray(item.([]interface{}))
		}
		sort.Slice(arr, func(i, j int) bool {
			return util.AsSha256(arr[i]) > util.AsSha256(arr[j])
		})
	}
}