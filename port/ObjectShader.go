package port

import (
	"gitee.com/RuiCat/catgl/lib/gl"
	"gitee.com/RuiCat/catgl/lib/mgl32"
)

///
// 着色器数据包
// ? 日志
//! 2019年11月1日 重写
///

// Color 颜色
type Color struct {
	Color    mgl32.Vec3 //* 颜色
	Texture  int32      //* 绑定贴图(索引)
	Strength float32    //* 强度
}

// UpdateShader 更新颜色
func (color *Color) UpdateShader(Program uint32, Name string) {
	gl.Uniform3fv(gl.GetUniformLocation(Program, gl.Str(Name+".Color")), 1, &color.Color[0])
	gl.Uniform1i(gl.GetUniformLocation(Program, gl.Str(Name+".Texture")), color.Texture)
	gl.Uniform1f(gl.GetUniformLocation(Program, gl.Str(Name+".Strength")), color.Strength)
}

// Light 静态灯光
type Light struct {
	Position mgl32.Vec3 //* 位置
	//? 颜色
	Ambient  mgl32.Vec3 //* 环境色
	Diffuse  mgl32.Vec3 //* 漫反射
	Specular mgl32.Vec3 //* 高光反射
	//? 物理参数
	Linear    float32 //* 线性
	Constant  float32 //* 常量
	Quadratic float32 //* 二次方程式
}

// UpdateShader 更新灯光
func (light *Light) UpdateShader(Program uint32, Name string) {
	gl.Uniform3fv(gl.GetUniformLocation(Program, gl.Str(Name+"Position\x00")), 1, &light.Position[0])
	gl.Uniform3fv(gl.GetUniformLocation(Program, gl.Str(Name+"Ambient\x00")), 1, &light.Ambient[0])
	gl.Uniform3fv(gl.GetUniformLocation(Program, gl.Str(Name+"Diffuse\x00")), 1, &light.Diffuse[0])
	gl.Uniform3fv(gl.GetUniformLocation(Program, gl.Str(Name+"Specular\x00")), 1, &light.Specular[0])
	gl.Uniform1f(gl.GetUniformLocation(Program, gl.Str(Name+"Linear\x00")), light.Linear)
	gl.Uniform1f(gl.GetUniformLocation(Program, gl.Str(Name+"Constant\x00")), light.Constant)
	gl.Uniform1f(gl.GetUniformLocation(Program, gl.Str(Name+"Quadratic\x00")), light.Quadratic)
}
