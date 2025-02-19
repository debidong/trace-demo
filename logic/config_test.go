package logic

import (
	"fmt"
	"testing"
)

func Test_MustLoadConfig(t *testing.T) {
	config := MustLoadConfig("../config.yaml")
	fmt.Println(config)
	fmt.Println(config.Server["china"])
}
