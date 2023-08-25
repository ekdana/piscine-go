package piscine

func StrRev(s string) string {
	aString := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		asd := aString[i]
		aString[i] = aString[len(aString)-1-i]
		aString[len(aString)-1-i] = asd
	}
	return string(aString)
}
