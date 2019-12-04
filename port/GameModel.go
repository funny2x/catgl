package port

import (
	"gitee.com/RuiCat/catgl/lib/gl"
	"gitee.com/RuiCat/catgl/lib/mgl32"
)

// model 着色器变量名
var model = gl.Str("Model\x00")

// GameModelPositions 顶点数组对象
type GameModelPositions struct {
	mgl32.Mat4
}

// Init 初始化
func (M *GameModelPositions) Init() {
	M.Mat4 = mgl32.Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// Update 模型位置更新
func (M *GameModelPositions) Update(Program uint32) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(Program, model), 1, false, &(M.Mat4[0]))
}
