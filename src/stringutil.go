package main

func startsWith(target string, source string) bool {
	sourceLen := len(source)
	if len(target) < sourceLen {
		return false
	}
	for i := 0; i < sourceLen; i++ {
		if target[i] != source[i] {
			return false
		}
	}
	return true
}

