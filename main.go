package main

import (
	"github.com/sirupsen/logrus"
	"grim-dawn-warhouse/src/grimdawn/warehouse"
	"os"
)

func main() {
	logrus.SetReportCaller(true)
	wh := warehouse.Warehouse{}
	path := os.Args[1]
	err := wh.Load(path)
	if err != nil {
		logrus.Fatalf("Load warehosue from %s failed, cause: %s", path, err.Error())
		return
	}

}
