package stringbuilder

import "strconv"

type StringBuilder struct {
	data []byte
}

// NewStringBuilder returns a new StringBuilder
func NewStringBuilder() *StringBuilder {
	return &StringBuilder{
		data: make([]byte, 0),
	}
}

// Append appends a string to the StringBuilder
func (sb *StringBuilder) Append(s string) {
	sb.data = append(sb.data, s...)
}

// AppendByte appends a byte to the StringBuilder
func (sb *StringBuilder) AppendByte(b byte) {
	sb.data = append(sb.data, b)
}

// AppendString appends a string to the StringBuilder
func (sb *StringBuilder) AppendString(s string) {
	sb.data = append(sb.data, s...)
}

// AppendInt appends an int to the StringBuilder
func (sb *StringBuilder) AppendInt(i int) {
	sb.data = append(sb.data, strconv.Itoa(i)...)
}

// to string
func (sb *StringBuilder) String() string {
	return string(sb.data)
}
