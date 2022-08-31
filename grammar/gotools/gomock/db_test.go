package gomock

import (
	"testing"

	"github.com/golang/mock/gomock"
	"step/grammar/gotools/gomock/mocks"
)

func TestGetFromDB(t *testing.T) {
	// 创建gomock控制器，用来记录后续的操作信息
	ctrl := gomock.NewController(t)
	// 断言期望的方法都被执行
	// Go1.14+的单测中不再需要手动调用该方法
	defer ctrl.Finish()
	// 调用mockgen生成代码中的NewMockDB方法
	// 这里mocks是我们生成代码时指定的package名称
	m := mocks.NewMockDB(ctrl)
	// 打桩（stub）
	// 当传入Get函数的参数为liwenzhou.com时返回1和nil
	m.
		EXPECT().
		Get(gomock.Eq("liwenzhou.com")). // 参数
		Return(1, nil).                  // 返回值
		Times(1)                         // 调用次数

	// 调用GetFromDB函数时传入上面的mock对象m
	if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
		t.Fatal()
	}
}
