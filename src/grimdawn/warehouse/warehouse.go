package warehouse

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"grim-dawn-warhouse/src/grimdawn/warehouse/archive"
	"os"
)

type Item struct {
}

type ExtendedItem struct {
	item Item
}

type Warehouse struct {
}

func (w Warehouse) Load(path string) error {
	file, err := os.Open(path)
	if err != nil {
		logrus.Errorf("Read %s failed, cause: %s", path, err.Error())
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Errorf("Close %s failed, cause: %s", path, err.Error())
		}
	}(file)
	arc := archive.Archive{}
	err = arc.Load(bufio.NewReader(file))
	if err != nil {
		return err
	}

	return nil
}
