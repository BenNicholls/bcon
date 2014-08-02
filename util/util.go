package util

type BconError struct {
	What string
}

func (e BconError) Error() string {
	return e.What
}
