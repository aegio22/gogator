package main

import (
	"fmt"

	"github.com/aegio22/gogator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cfg.SetUser("dane")
	if err != nil {
		fmt.Println(err)
		return
	}

	updatedCfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	updatedCfg.Repr()

}
