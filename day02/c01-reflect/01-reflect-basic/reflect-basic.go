package _1_reflect_basic

import (
	"encoding/json"
	"fmt"
)

// 开发者
type Developer struct {
	Name 	string		`json:"developerName"`
	Sex		string		`json:"developerSex"`
	Age		int			`json:"developerAge"`
	Exp		float64		`json:"developerExp"`
}

// 获取开发人员信息
func GetDeveloperInfo() {
	devops := Developer{
		Name: "谭祎",
		Age :28,
		Sex: "boy",
		Exp: 5.5,
	}
	data, _ := json.Marshal(devops)
	fmt.Printf("data : %v\n", string(data))
}
