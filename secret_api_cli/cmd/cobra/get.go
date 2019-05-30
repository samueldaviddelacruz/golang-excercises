package cobra

import (
	"fmt"

	secret "github.com/samueldaviddelacruz/golang-exercises/secret_api_cli"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a secret from your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := secret.File(encodingKey, secretsPath())
		key := args[0]
		value, err := v.Get(args[0])
		if err != nil {
			fmt.Println("No value set")
			return
		}
		fmt.Printf("%s = %s\n", key, value)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
