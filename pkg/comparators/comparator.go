package comparators

type Comparator interface {
	compare(inputFile1, inputFile2 string) (error, bool)
}
