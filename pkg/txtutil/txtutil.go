package txtutil

// SplitSingleLongTxt does nothing.
// Deprecated: This is a no-op for backwards compatibility.
func SplitSingleLongTxt(records any) {
}

// ToChunks returns the string as chunks of 255-octet strings (the last string being the remainder).
func ToChunks(s string) []string {
	return splitChunks(s, 255)
}

// Segment returns the string as 255-octet segments, the last being the remainder.
func Segment(s string) []string {
	return splitChunks(s, 255)
}

func splitChunks(buf string, lim int) []string {
	var chunk string
	chunks := make([]string, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
