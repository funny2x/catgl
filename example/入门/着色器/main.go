package main

import (
	"fmt"

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
	// 游戏元素初始化
	Game(Scene)
	// 等待结束
	fmt.Println(<-err)
}

// Game 游戏对象
func Game(Scene *catgl.Scene) {
	// 默认着色器
	shader := &port.ShaderData{
		Vertex: `
		#version 330 core
		//? 默认数据顶点
		layout (location = 0) in vec3 Positions; //* 位置
		layout (location = 1) in vec3 Normals;	 //* 法线
		layout (location = 2) in vec3 Color;	 //* 颜色
		layout (location = 3) in vec2 Uv;		 //* Uv
		layout (location = 4) in vec3 Tangent;	 //* 切线
		layout (location = 5) in vec3 Bitangent; //* 副切线
		out vec4 vertexColor; 
		void main()
		{
			gl_Position = vec4(Positions, 1.0);
			vertexColor = vec4(Color, 1.0); 
		}
		`,
		Fragment: `
		#version 330 core
		out vec4 color;
		in vec4 vertexColor; 
		void main()
		{
			color = vertexColor;
		}
		`,
	}
	Scene.AddObject(shader)
	Scene.ShaderObjectList["默认着色器"] = shader
	// 默认游戏对象
	game := &GameObject{
		ShaderProgram: shader.Program,
		Vertices: []float32{
			-0.5, -0.5, 0.0, 1.0, 0.0, 0.0,
			0.5, -0.5, 0.0, 0.0, 1.0, 0.0,
			0.0, 0.5, 0.0, 0.0, 0.0, 1.0,
		},
	}
	Scene.AddObject(game)
	Scene.GameObjectList["三角形"] = game
}
