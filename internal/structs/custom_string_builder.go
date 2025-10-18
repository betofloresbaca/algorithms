package structs

type CustomStringBuilder struct {
	parts [][]byte
	count int
}

func (csb *CustomStringBuilder) WriteString(value string) {
	byteArr := []byte(value)
	csb.parts = append(csb.parts, byteArr)
	csb.count += len(byteArr)
}

func (csb *CustomStringBuilder) String() string {
	result := make([]byte, csb.count)
	pos := 0
	for _, part := range csb.parts {
		copy(result[pos:], part)
		pos += len(part)
	}
	return string(result)
}
