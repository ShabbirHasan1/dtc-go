package message

func CloneString(s string) string {
	to := make([]byte, len(s))
	to = append(to, s...)
	return string(to)
}
