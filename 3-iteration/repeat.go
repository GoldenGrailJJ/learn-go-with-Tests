package iteration

func Repeat(character string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += character
	}
	return repeated
}
