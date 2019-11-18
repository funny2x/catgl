package port

import (
	"gitee.com/RuiCat/catgl/lib/gl"
)

// GameVertexVAO 顶点数组对象
type GameVertexVAO struct {
	VAO   uint32 // 顶点指针
	First int32  // 起始位置
	Count int32  // 顶点数量
}

// Update 更新顶点
func (V *GameVertexVAO) Update() {
	gl.BindVertexArray(V.VAO)
	gl.DrawArrays(gl.TRIANGLES, V.First, V.Count)
	gl.BindVertexArray(0)
}

// GameVertexVBO 顶点缓冲区对象
type GameVertexVBO []float32

// New 创建
func (V *GameVertexVBO) New(T GameVertexlBuffersType) uint32 {
	var VBO uint32
	gl.GenBuffers(1, &(VBO))                                         //? 创建 VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)                              //? 激活缓冲区
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(*V), gl.Ptr(*V), uint32(T)) //? 添加缓冲区数据
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	return VBO
}
