package internal

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

// CustomerMostSpent is the struct that represents a customer that spent the most.
type CustomerMostSpent struct {
	// FirstName is the first name of the customer.
	FirstName string `json:"first_name"`
	// LastName is the last name of the customer.
	LastName string `json:"last_name"`
	// Ammount is the total spent by the customer.
	Ammount float32 `json:"ammount"`
}
