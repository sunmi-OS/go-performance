package main

import (
	"encoding/json"
	"os"
	"runtime"
	"sort"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo"
	"github.com/sunmi-OS/gocore/api"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

var jsonString = `
{
    "l1": {
        "l1_1": [
            "l1_1_1",
            "l1_1_2",
        ],
        "l1_2": {
            "l1_2_1": 121
        }
    },
    "l2": {
        "l2_1": null,
        "l2_2": true,
        "l2_3": {}
    }
}
`

type Demo struct {
	L1 L1Demo `json:"l1"`
	L2 L2Demo `json:"l2"`
}

type L1Demo struct {
	L1_1 []string `json:"l1_1"`
	L1_2 L1_2Demo `json:"l1_2"`
}

type L1_2Demo struct {
	L1_2_1 int `json:"l1_2_1"`
}

type L2Demo struct {
	L2_1 interface{} `json:"l2_1"`
	L2_2 interface{} `json:"l2_2"`
	L2_3 interface{} `json:"l2_3"`
}

type EchoApi struct {
}

var eApi EchoApi

func (a *EchoApi) echoStart(c *cli.Context) error {
	// Echo instance
	e := echo.New()

	// Middleware
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())

	// Route => handler
	e.Any("/gjson", func(c echo.Context) error {

		response := api.NewResponse(c)

		gj := gjson.Parse(jsonString)

		gj.Get("l1.l1_2.l1_2_1").String()
		gj.Get("l1.l1_1").Array()
		gj.Get("l2").Map()

		return response.RetSuccess("sccess")
	})

	e.Any("/json", func(c echo.Context) error {

		response := api.NewResponse(c)

		demo := &Demo{}

		json.Unmarshal([]byte(jsonString), demo)

		return response.RetSuccess("sccess")
	})

	e.Any("/jsoniter", func(c echo.Context) error {

		response := api.NewResponse(c)

		demo := &Demo{}

		var json = jsoniter.ConfigCompatibleWithStandardLibrary

		json.Unmarshal([]byte(jsonString), demo)

		return response.RetSuccess("sccess")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func main() {

	runtime.GOMAXPROCS(1)

	app := cli.NewApp()
	app.Name = "IOT-seanbox"
	app.Usage = "IOT-seanbox"
	app.Email = "wenzhenxi@sunmi.com"
	app.Version = "1.0.0"
	app.Usage = "IOT-seanbox"
	app.Action = eApi.echoStart

	// 指定对于的命令
	app.Commands = []cli.Command{
		{
			Name:    "api",
			Aliases: []string{"a"},
			Usage:   "api",
			Subcommands: []cli.Command{
				{
					Name:   "start",
					Usage:  "开启API-DEMO",
					Action: eApi.echoStart,
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)

}
