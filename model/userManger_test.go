package model

import "testing"

func TestCreateUser(t *testing.T) {
	err := CreateUser("name2", "pass")
	if err != nil {
		panic(err)
	}
}
