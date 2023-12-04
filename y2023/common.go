package y2023

import "os"

func Must1[T any](v T, err error) T {
	return v
}

func fatalErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFromFile(filename string) string {
	b, err := os.ReadFile(filename)
	fatalErr(err)

	return string(b)
}
