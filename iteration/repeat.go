package iteration

func Repeat(caracter string, count int) string {
	repetitions := ""
	for range count {
		repetitions += caracter
	}
	return repetitions
}
