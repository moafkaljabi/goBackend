package models

// change account to devices.

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   float64
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        99,
		FirstName: firstName,
		LastName:  lastName,
		Number:    100,
		Balance:   1000,
	}
}
