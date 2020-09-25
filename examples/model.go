package main

type ExampleOrder struct {
	Id       string
	User     ExampleUser      `dynamo:"-"`
	Products []ExampleProduct `dynamo:"-"`
}

type ExampleUser struct {
	Id        string
	FirstName string
	LastName  string
	Address   string
}

type ExampleProduct struct {
	Id          string
	Name        string
	Description string
}
