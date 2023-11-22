package main

import (
	"errors"
	"fmt"

	"github.com/stretchr/testify/mock"
)

func main() {
	c := CustomerRepositoryMock{}
	c.On("GetCustomer", 1).Return("No", 20, nil)
	c.On("GetCustomer", 2).Return("", 0, errors.New("not found"))

	name, age, err := c.GetCustomer(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(name, age)
}

type CustomerRepositoryMock struct {
	mock.Mock
}

func (m *CustomerRepositoryMock) GetCustomer(id int) (name string, age int, err error) {
	args := m.Called(id)
	return args.String(0), args.Int(1), args.Error(2)
}
