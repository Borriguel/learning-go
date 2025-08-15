package iteration

func Repeat(caracter string) string {
	repetitions := ""
	for range 5 {
		repetitions += caracter
	}
	return repetitions
}
