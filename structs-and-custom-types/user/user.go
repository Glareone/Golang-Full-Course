package user

import (
	"errors"
	"fmt"
	"time"
)

// making it upper case you make it available outside of this file
// making local types you can use lowerCase custom type names
type User struct {
	// should start uppercase because otherwise they will be unavailable
	// from the outside
	// TIP: Ideally to have them unexported because it forces the end user to use newUser constructor
	// instead of "new User{}" and have direct access to User's fields
	FirstName string    // Exported
	LastName  string    // Exported
	BirthDate string    // Exported
	CreatedAt time.Time // Exported

	// it's also possible to create variables with the same name which are lower case
	// ======
	// firstName string    // Unexported
	// lastName  string    // Unexported
	// birthDate string    // Unexported
	// createdAt time.Time // Unexported
}

// Creation Method
// newUser - is just convention (pattern), new+StructName,
// it's not a built-in feature in Go
func NewUser(firstName string, lastName string, birthDate string) (*User, error) {
	// We can validate the input before sending the instance back
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("wrong input")
	}

	// we return the pointer
	// it prevents of making extra copies of the value
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil

	// Returning User like this we after method invocation from another place we will return the copy of this instance
	// Therefore we say we return the pointer from the method and use & ampersand like I showed above
	//
	//return User{
	//	firstName: firstName,
	//	lastName:  lastName,
	//	birthDate: birthDate,
	//	createdAt: time.Now(),
	//}
}

// this function is attached to the original struct User
// In parentheses I also use "Receiver Argument" (or just Receiver) in the name to get access to the properties within the User struct
func (user User) OutputUserDetails() {
	// "user" is a copy of the struct we sent to the function calling it with "outputUserDetails(appUser)"
	// in order to use a pointer we need to change the declaration
	// it is a shallow copy of the original struct
	// here we use the copy of original user Struct
	fmt.Println("appUser: ", user.FirstName, user.LastName, user.BirthDate, user.CreatedAt)
}

// it's possible to use * asterisk in such methods as well
// to ensure that no extra memory occupied due to created copies
func (user *User) OutputUserDetailsAsterisk() {
	fmt.Println("appUser: ", user.FirstName, user.LastName, user.BirthDate, user.CreatedAt)
}

// MUTATOR
// you must use POINTER in Mutator methods because otherwise you mutate only the copy of "Receiver Argument", not original User Struct fields
func (user *User) ClearUserNameAsterisk() {
	user.FirstName = ""
	user.LastName = ""
	user.BirthDate = ""
}

// MUTATOR. WRONG USAGE. SHOULD BE USED WITH POINTER
// using regular value here would be a problem because we change the values inside the copy
// Go sends the copy to "Receiver Argument" as well if we dont use * asterisk
func (user User) ClearUserName() {
	user.FirstName = ""
	user.LastName = ""
	user.BirthDate = ""
}
