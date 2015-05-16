package main

// All states part of mainSm are expected to implement this.
type mainState interface {
	Update(timestamp float64)
	Draw(timestamp float64)
}
