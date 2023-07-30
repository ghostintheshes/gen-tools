package slice

import "fmt"

func ErrOutOfRange(length, out int) error {
	return fmt.Errorf("index invalid or out of range, length is %d, desire is %d", length, out)
}
