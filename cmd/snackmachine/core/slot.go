package core

type Slot struct {
	ID           int
	Snack        Snack
	Quantity     int
	Price        float64
	SnackMachine SnackMachine
	Position     int
}
