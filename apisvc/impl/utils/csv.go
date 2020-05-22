package utils

func CSVError(err error) [][]string {
	res := [][]string{
		{"error"},
		{err.Error()},
	}
	return res
}
