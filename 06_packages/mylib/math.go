package mylib

func Average(s []int) int {

	total := 0
	for _, v := range s {
		total += v
	}
	return int(total / len(s))
}
