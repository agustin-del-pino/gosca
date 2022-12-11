package gosca

import (
	"fmt"
	"github.com/spf13/cobra"
	"gosca/pkg/gosca"
	"gosca/utils"
)

var useYml bool

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "creates the project scaffolding",
	Aliases: []string{"config", "where"},
	Run: func(cmd *cobra.Command, args []string) {
		var s gosca.Scaffolding
		placeholders := map[string]string{}

		if b, berr := gosca.PreprocessConfig(args[0], func(s string) string {
			if v, ok := placeholders[s]; ok {
				return v
			}

			r := utils.Input("enter the %s: ", s)

			placeholders[s] = r

			return r

		}); berr != nil {
			panic(utils.WrapError(berr, "error at preprocessing the config file in create command"))
		} else {
			if !useYml {
				if jerr := s.FromJsonBytes(b); jerr != nil {
					panic(utils.WrapError(jerr, "error at parsing the json configuration file in create command"))
				}
			} else {
				if yerr := s.FromYamlBytes(b); yerr != nil {
					panic(utils.WrapError(yerr, "error at parsing the yaml configuration file in create command"))
				}
			}
		}

		if cerr := gosca.ConstructScaffolding(args[1], &s); cerr != nil {
			panic(utils.WrapError(cerr, "error at constructing the scaffolding"))
		}

		fmt.Printf("scaffolding was succesfuly created!")
	},
}

func init() {
	createCmd.Flags().BoolVarP(&useYml, "use-yml", "y", false,
		"read the configuration as yml file")
	rootCmd.AddCommand(createCmd)
}
