package labeler

import (
	drawer "analyzer/drawer"
	"errors"
	"fmt"
)

func GatherData(size uint, drawer *drawer.Drawer) error {

	contests, err := drawer.GetSome(10, 0)
	if err != nil {
		return err
	}

	fmt.Println(len(contests))

	return errors.New("continue")

}
