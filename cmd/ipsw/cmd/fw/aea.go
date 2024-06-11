/*
Copyright © 2024 blacktop

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package fw

import (
	"os"

	"github.com/apex/log"
	"github.com/blacktop/ipsw/pkg/aea"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NOTES: key-val is 32 bytes or 64 char hex string

func init() {
	FwCmd.AddCommand(aeaCmd)

	aeaCmd.Flags().StringP("output", "o", "", "Folder to extract files to")
	aeaCmd.MarkFlagDirname("output")
	viper.BindPFlag("fw.aea.output", aeaCmd.Flags().Lookup("output"))
}

// aeaCmd represents the ane command
var aeaCmd = &cobra.Command{
	Use:    "aea",
	Short:  "🚧 Parse ANE1 DMGs",
	Args:   cobra.ExactArgs(1),
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		if viper.GetBool("verbose") {
			log.SetLevel(log.DebugLevel)
		}

		// flags
		output := viper.GetString("fw.aea.output")

		if err := os.MkdirAll(output, 0o750); err != nil {
			return err
		}

		return aea.Parse(args[0], output)
	},
}
