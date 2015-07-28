package main

import (
	"fmt"
	"github.com/ulule/deepcopier"
)

// Model
type User struct {
	Name string
}

func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
	// do whatever you want
	return ""
}

// Resource
type UserResource struct {
	DisplayName            string `deepcopier:"field:Name"`
	SkipMe                 string `deepcopier:"skip"`
	MethodThatTakesContext string `deepcopier:"context"`
}

func main() {
	user := &User{
		Name: "gilles",
	}

	resource := &UserResource{}

	deepcopier.Copy(user).To(resource)

	fmt.Println(resource.DisplayName)
}
