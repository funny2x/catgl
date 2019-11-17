package main

import (
	"gitee.com/RuiCat/catgl"
)

func main() {
	// 创建场景
	Scene := &catgl.Scene{}
	// 场景默认对象
	Scene.Light = &Light{}
	Scene.Camera = &Camera{}
	// 渲染窗口
	R := catgl.Renderer{Width: 500, Height: 500, Title: "测试窗口", Scene: Scene}
	// 创建窗口
	err := R.New()
	if (<-err) != nil {
		return
	}
	<-err
}
