// +build gofuzz

package twkb

func Fuzz(data []byte) int {
	if _, err := Unmarshal(data); err != nil {
		return 0
	}
	return 1
}
