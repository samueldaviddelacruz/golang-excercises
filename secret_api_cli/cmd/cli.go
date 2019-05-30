package main

import (
	"github.com/samueldaviddelacruz/golang-exercises/secret_api_cli/cmd/cobra"
)

func main() {
	cobra.RootCmd.Execute()
	/*
		v := secret.File("my-fake-keys", "secrets")

		err := v.Set("demo_key", "Just a value")
		if err != nil {
			panic(err)
		}

		plain, err := v.Get("demo_key")
		if err != nil {
			panic(err)
		}
		fmt.Println(plain)
	*/
}
