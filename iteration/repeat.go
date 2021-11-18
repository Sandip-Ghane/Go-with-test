package iteration

// RepeatedCount ...
const RepeatedCount = 5

// Repeat ...
func Repeat(char string) string {
	var repeatedString string
	for a := 1; a <= RepeatedCount; a++ {
		repeatedString += char
	}
	return repeatedString
}

func main() {
	Repeat("a")
}
