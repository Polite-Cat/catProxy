package main_test

import (
	"fmt"
	"github.com/networm6/catProxy/confbox"
	"github.com/networm6/catProxy/data"
	"log"
)

func main() {
	scene80 := make(map[data.SrcUrl]data.DstUrl)
	scene80["local.simon-app.cn"] = "http://localhost:3001"
	scene80["local.polite.cat"] = "http://localhost:3001"
	list := data.PortList{
		PointLists: []data.PointList{
			{
				Port:   80,
				Points: scene80,
			}, {
				Port:   8080,
				Points: scene80,
			},
		},
	}
	fmt.Println(list)
	err := confbox.Save("I:\\ReverseProxy\\test.yaml", &list)
	if err != nil {
		log.Fatalln(err)
	}
}
