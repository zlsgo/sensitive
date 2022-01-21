package sensitive

func ctyphe(r rune) rune {
	if r >= 65 && r <= 90 {
		r += 32
	} else if r >= 65313 && r <= 65338 {
		r = 97 + (r - 65313)
	} else if r >= 65345 && r <= 65370 {
		r = 97 + (r - 65345)
	} else if r >= 65280 && r <= 65375 {
		r = 32 + (r - 65280)
	}
	return r
}
