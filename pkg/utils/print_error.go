package utils

import (
	m "ascii/pkg/models"
	"fmt"
)

func PrintError(e m.Error) {
	err := fmt.Sprintf("------ Error -----\nError: %s\nDetail: %s\n----- Thank You -----", e.Error, e.Detail)
	fmt.Println(err)
}
