package gosca

import (
	"fmt"
	"github.com/spf13/cobra"
	"gosca/pkg/gosca"
	"gosca/utils"
	"os"
)

const (
	Python string = "python"
	Golang string = "Golang"
	Json   string = "json"
	Yaml   string = "yaml"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "initialize a scaffolding config file",
	Aliases: []string{"filepath"},
	Run: func(cmd *cobra.Command, args []string) {
		n := utils.Input("enter the project name: ")
		t := utils.Input("enter the scaffolding config filetype [%s, %s] (default: %s): ", Json, Yaml, Json)
		p := utils.Input("enter a predefined scaffolding [%s, %s] (optional): ", Python, Golang)

		var s *gosca.Scaffolding
		switch p {
		case Python:
			s = gosca.CreatePythonScaffolding(n)
		case Golang:
			s = gosca.CreateGolangScaffolding(n)
		default:
			s = gosca.CreateScaffolding(n)
		}
		var b *[]byte

		if t == Yaml {
			if y, yerr := s.ToYaml(); yerr != nil {
				panic(utils.WrapError(yerr, "scaffolding conversion to yaml fail at init command"))
			} else {
				b = &y
			}
		} else {
			if j, jerr := s.ToJson(); jerr != nil {
				panic(utils.WrapError(jerr, "scaffolding conversion to json fail at init command"))
			} else {
				b = &j
			}
		}

		if werr := os.WriteFile(args[0], *b, 0777); werr != nil {
			panic(utils.WrapError(werr, "cannot create the scaffolding file: %s at init command", args[0]))
		}

		fmt.Printf("scaffolding was succesfuly initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
