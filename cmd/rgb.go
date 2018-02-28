package cmd

import (
	"github.com/spf13/cobra"
	"github.com/lucasb-eyer/go-colorful"
	"time"
	"github.com/ottojo/lights"
)

var rgbInterval int

var rgbCmd = &cobra.Command{
	Use:   "rgb",
	Short: "RGB ALL THE THINGS!!1elfcos(0)-e^(iπ)!",
	Long:  `This command will RGB-cycle all the LEDs in the room as long as it is running.`,
	Run: func(cmd *cobra.Command, args []string) {

		for {
			for h := 0; h < 360; h++ {
				c := colorful.Hsv(float64(h), 1, 1)
				go lights.SetAll(lights.Color(c))
				time.Sleep(time.Duration(rgbInterval) * time.Millisecond)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(rgbCmd)

	rgbCmd.Flags().IntVarP(&rgbInterval, "rgbInterval", "i", 10, "time in ms for one color (360 colors in total)")
}
