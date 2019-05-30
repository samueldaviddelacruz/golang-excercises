package cobra

import (
	"fmt"

	secret "github.com/samueldaviddelacruz/golang-exercises/secret_api_cli"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key, value := args[0], args[1]
		err := v.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
