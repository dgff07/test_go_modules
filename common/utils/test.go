package utils

import (
	"fmt"

	"github.com/dgff07/test_go_modules/common/entities"
)

func DoSomething() {
	random := entities.Deck{"random"}
	fmt.Println(random)
}
