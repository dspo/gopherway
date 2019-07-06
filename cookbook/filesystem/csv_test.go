package filesystem

import "testing"

var (
	file1 = "testcsv1.csv"
	file2 = "testcsv2.csv"
)

func TestMergeFiles(t *testing.T) {
	MergeFiles(file1, file2)
}
