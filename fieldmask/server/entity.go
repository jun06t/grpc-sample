package main

type User struct {
	ID      string  `bson:"_id"`
	Name    string  `bson:"name"`
	Email   string  `bson:"email"`
	Age     int     `bson:"age"`
	Address Address `bson:"address"`
}

type Address struct {
	Country string `bson:"country"`
	State   string `bson:"state"`
	City    string `bson:"city"`
	Zipcode string `bson:"zipcode"`
}
