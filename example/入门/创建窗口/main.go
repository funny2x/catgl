package main

import (
	"gitee.com/RuiCat/catgl"
	"gitee.com/RuiCat/catgl/port"
)

func main() {
	// 初始化
	catgl.NOScreenRender()
	// 创建场景
	Scene := &catgl.Scene{}
	// 场景默认对象
	Scene.Light = &port.LightData{}
	Scene.Camera = &port.CameraData{}
	// 渲染窗口
	R := catgl.Renderer{Width: 500, Height: 500, Title: "测试窗口", Scene: Scene}
	// 创建窗口
	err := R.New()
	if (<-err) != nil {
		return
	}
	<-err
}
