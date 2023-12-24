package main

func loginmain() {
	authenticatormain := &Authenticator{}
	loginmain := &LoginCommand{Authenticator: authenticatormain}

	remote := &Accessor{}

	remote.SetAuthentication(loginmain)
	remote.SubmitLogin()
}

func main() {
	loginmain()
}
