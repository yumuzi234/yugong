package smlrpc

func lowerIdent(in string) string {
	if len(in) == 0 {
		return in
	}
	start := in[0]
	if !isCap(start) && !isLow(start) {
		return in
	}
	var out []byte
	i := 0
	foundLower := false
	for ;i<len(in) && !foundLower;i++ {
		cur := in[i]
		cur, foundLower = toLow(cur)
		out = append(out, cur)
	}
	if i>2 && i!= len(in) {
		out[i-2] = out[i-2] - 32
	}
	for ;i<len(in);i++ {
		cur := in[i]
		if foundLower {
			out = append(out, cur)
			foundLower = isLow(cur)
		} else {
			cur, foundLower = toLow(cur)
			out = append(out, cur)
		}
	}
	return string(out)
}

func isCap(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isLow(b byte) bool {
	return b >='a' && b <= 'z'
}

func toLow(b byte) (byte, bool) {
	if isCap(b) {
		return b + 32, false
	}
	return b, true
}
