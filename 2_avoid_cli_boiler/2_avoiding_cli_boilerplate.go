package avoidboiler

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	//create a new app
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count Up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "Value to count up to",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("stop cannot be negative.")
				}
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Count Down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Start counting down from",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start > 0 {
					fmt.Println("Start cannot be negative.")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	// //First iteratiion of appliaion- just prints hello
	// //set up a global flac
	// app.Flags = []cli.Flag{
	// 	cli.StringFlag{
	// 		Name:  "name, n",
	// 		Value: "World",
	// 		Usage: "Who to say hello to",
	// 	},
	// }

	// //define the action to run
	// app.Action = func(c *cli.Context) error {
	// 	name := c.GlobalString("name")
	// 	fmt.Printf("Hello %s!\n", name)
	// 	return nil
	// }
	// app.Run(os.Args)

	app.Run(os.Args)
}
