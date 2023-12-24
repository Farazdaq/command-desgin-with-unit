package main

import (
	"testing"
)

func TestLogin(t *testing.T) {

	authenticatormain := &Authenticator{}
	loginmain := &LoginCommand{Authenticator: authenticatormain}

	remote := &Accessor{}

	remote.SetAuthentication(loginmain)
	remote.SubmitLogin()
	//Test case1 valid login
	//if !Login("farazdaq", "501755") {
	//
	// Test case 2 invalid login
	//	t.Error("Expected login to fail for farardaq, but it is sucessful.")
	//}
	// Test case 3 invalid login
	//if Login("farazdaq", "") {
	//t.Error("Expected login to fail for farardaq, but it is sucessful.")
	//}
	// Test case 4 invalid login
	//if Login("", "501755") {
	//t.Error("Expected login to fail for , but it is sucessful.")
	//}
	// Test case 5 invalid login
	//if Login("", "") {
	//t.Error("Expected login to fail for , but it is sucessful.")
	//}
	//if Login("$%##$%^&*", "%$#$%^^&") {
	//t.Error("Expected login to fail for %$#$%^^&, but it is sucessful.")
	//}

	// Note or the test case are ok just fill the uesname and password varble

}
