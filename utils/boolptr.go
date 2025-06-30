package utils

func boolPtrToVal(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}
