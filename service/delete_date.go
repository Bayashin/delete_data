package service

import (
	"time"

	"github.com/Bayashin/delete_data/lib"
)

type Customer struct {
	CustomerId    int       `json:"customerId"`
	Position      int       `json:"callNumber"`
	WaitingStatus string    `json:"waitingStatus"`
	Date          time.Time `json:"date"`
	FirebaseToken string    `json:"firebaseToken"`
	OwnerId       string    `json:"ownerId"`
}

func DeleteData() {
	db := lib.SqlConnect()
	db.Unscoped().Where("1=1").Delete(&Customer{})
}
