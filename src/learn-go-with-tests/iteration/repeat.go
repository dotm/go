package iteration

//Repeat multiply a character and returns it as a string
func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated += character
	}
	return repeated
}

//OptimizedRepeat multiply a character and returns it as a string
func OptimizedRepeat(character rune) string {
	var repeated [5]rune
	for i := 0; i < 5; i++ {
		repeated[i] = character
	}

	return string(repeated[:])
}
