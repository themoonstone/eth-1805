package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"math"
)

// 求面积
func GetCircleWithError(r float64) (float64, error) {
	if r < 0 {
		err := errors.New("半径非负")
		return -1, err
	}
	return math.Pi * r * r, nil
}

// 调用
func main() {
	area, err := GetCircleWithError(12.5)
	if nil != err {
		log.Errorf("error : %v\n", err)
		return
	}
	fmt.Println("area of circle is : ", area)
}