package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"strconv"
	_ "github.com/ottojo/lights"
	"github.com/ottojo/lights"
)

// allCmd represents the all command
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Set all lights in the room to one color",
	Long: `Set all lights in the room to one color`,
	Example:"leds all 1.0 0.4 0 sets all lights to 100% red, 40% green, 0% blue",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			log.Fatal("r g b")
		}

		r, err := strconv.ParseFloat(args[0], 64)
		g, err := strconv.ParseFloat(args[1], 64)
		b, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			log.Fatal(err)
		}

		c := lights.Color{R: r, G: g, B: b}
		lights.SetAll(c)
	},
}

func init() {
	rootCmd.AddCommand(allCmd)
}
