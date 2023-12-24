package main

import (
	"fmt"
)

// The belowe code is define the main Command Authentication
type Authentication interface {
	Execute()
}

// This part to deinfde the the Receiver Authenticator
type Authenticator struct{}

// The main functions prefomed my the accessor to approch authentication
func (a *Authenticator) Login() {
	password := ""
	username := ""
	if username == "farazdaq" && password == "501755" {
		fmt.Print("Login successfully")
	} else {
		fmt.Print("worng usrname or password")
	}

}
func (a *Authenticator) Logout() {
	// will go to the main web page and end the session
}

func (a *Authenticator) CreateAccount() {

}

func (a *Authenticator) RemoveAccount() {

}

func (a *Authenticator) Recovder() {

}

func (a *Authenticator) Validate() {
	uersname := "farar"
	password := "50"
	if uersname == "" && password == "" {
		fmt.Print("you have to add the inputs")
	} else {
		fmt.Print("correct inputs")
	}
}

// this part is the creat the commands for the all functions
type LoginCommand struct {
	Authenticator *Authenticator
}

func (c *LoginCommand) Execute() {
	c.Authenticator.Login()
}

// ------------------------------------------s--------------
type LogoutCommand struct {
	Authenticator *Authenticator
}

func (c *LogoutCommand) Execute() {
	c.Authenticator.Logout()
}

// ---------------------------------------------------------
type CreateAccountCommand struct {
	Authenticator *Authenticator
}

func (c *CreateAccountCommand) Execute() {
	c.Authenticator.CreateAccount()
}

// ---------------------------------------------------------
type RemoveAccountCommand struct {
	Authenticator *Authenticator
}

func (c *RemoveAccountCommand) Execute() {
	c.Authenticator.RemoveAccount()
}

// -----------------------------------------------------------
type ReceiverCommand struct {
	Authenticator *Authenticator
}

func (c *ReceiverCommand) Execute() {
	c.Authenticator.Recovder()
}

// -------------------------------------------------------------
type ValidateCommand struct {
	Authenticator *Authenticator
}

func (c *ValidateCommand) Execute() {
	c.Authenticator.Validate()
}

// Tis the part is deine the invoker Accessor
type Accessor struct {
	authentication Authentication
}

func (r *Accessor) SetAuthentication(authentication Authentication) {
	r.authentication = authentication
}

func (r *Accessor) SubmitLogin() {
	r.authentication.Execute()
}

func (r *Accessor) SubmitLogout() {
	r.authentication.Execute()
}

func (r *Accessor) SubmitCreateAccount() {
	r.authentication.Execute()
}

func (r *Accessor) SubmaitRemoveAccount() {
	r.authentication.Execute()
}

func (r *Accessor) SubmitRecover() {
	r.authentication.Execute()
}

func (r *Accessor) SubmitValidate() {
	r.authentication.Execute()
}
