package _2_monster

import "testing"

func TestMonster_Store(t *testing.T) {
	// 创建实例对象并赋值
	monster := Monster{
		Name: "哥斯拉",
		Age: 3,
		Skill: "潜水",
	}

	if res := monster.Store(); res == false {
		t.Fatalf("monster.Store(). expected=%v, get=%v\n", true, res)
	}
	t.Logf("monster.Store() 测试成功..\n")
}

func TestMonster_ReadFromFile(t *testing.T) {
	m := Monster{}
	res := m.ReadFromFile()
	if !res {
		t.Fatalf("monster.ReadFromFile(). expected=%v, get=%v\n", true, res)
	}
	if m.Name != "哥斯拉" {
		t.Fatalf("monster.ReadFromFile(). expected=%v, get=%v\n", "哥斯拉", m.Name)
	}
	t.Logf("monster.ReadFromFile() 测试成功..\n")
}

/*
	订单相关功能
		TestOrder_xxx
	用户相关功能
		TestUser_xxx
	支持相关功能
		TestPay_xxx
*/