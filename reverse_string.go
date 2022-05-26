package reverse_string

import (
	"github.com/rivo/uniseg"
)

func ReverseString(input string) (output string) {
	var lBuffer string

	gr := uniseg.NewGraphemes(input)
	for gr.Next() {
		lBuffer = string(gr.Runes()) + lBuffer
	}
	return lBuffer
}
