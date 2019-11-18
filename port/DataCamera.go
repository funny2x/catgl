package port

import (
	"gitee.com/RuiCat/catgl/lib/gl"
	"gitee.com/RuiCat/catgl/lib/mgl32"
)

///
// 摄像机数据对象
// ? 日志
//! 2019年11月1日 重写
//! 2019年11月18日
//		完成着色器底层绑定
///

// CameraLookAtVData 摄像机朝向结构体
type CameraLookAtVData struct {
	Up      mgl32.Vec3 //? 朝向
	Eye     mgl32.Vec3 //? 位置
	Center  mgl32.Vec3 //? 视点
	LookAtV mgl32.Mat4 //? 摄像机朝向
}

// Update 主更新
func (C *CameraLookAtVData) Update() {
	//? 计算朝向
	C.LookAtV = mgl32.LookAtV(C.Eye, C.Center, C.Up)
}

// CameraProjectionData 摄像机投影矩阵结构体
type CameraProjectionData struct {
	Fovy       float32    // 视角
	Aspect     float32    // 屏幕宽高比
	ZNear      float32    // 近平面距离
	ZFar       float32    // 远平面距离
	Projection mgl32.Mat4 // 投影矩阵
}

// Update 主更新
func (C *CameraProjectionData) Update() {
	// 计算投影矩阵
	C.Projection = mgl32.Perspective(mgl32.DegToRad(C.Fovy), C.Aspect, C.ZNear, C.ZFar)
}

// CameraData 摄像机数据结构体
type CameraData struct {
	CameraLookAtVData
	CameraProjectionData
	// 着色器参数
	GlobalMatrix uint32 // 着色器矩阵全局变量
}

// Init 初始化
func (C *CameraData) Init() error {
	gl.GenBuffers(1, &C.GlobalMatrix)                          // 创建缓冲器
	gl.BindBuffer(gl.UNIFORM_BUFFER, C.GlobalMatrix)           // 绑定缓冲区
	gl.BufferData(gl.UNIFORM_BUFFER, 128, nil, gl.STATIC_DRAW) // 分配内存空间
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)                        // 解除引用
	// 绑定着色器变量定义的绑定点,矩阵定义绑定在0号绑定点
	gl.BindBufferRange(gl.UNIFORM_BUFFER, 0, C.GlobalMatrix, 0, 128)
	// 更新朝向
	C.Update()
	// 更新投影矩阵
	C.UpdateProjection()
	return nil
}

// UpdateProjection 投影矩阵更新
func (C *CameraData) UpdateProjection() {
	// 更新投影矩阵
	C.CameraProjectionData.Update()
	// 更新着色器值
	gl.BindBuffer(gl.UNIFORM_BUFFER, C.GlobalMatrix)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 64, 64, gl.Ptr(&C.Projection[0]))
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}

// Update 主更新(视图矩阵)
func (C *CameraData) Update() {
	// 更新朝向
	C.CameraLookAtVData.Update()
	// 更新着色器值
	gl.BindBuffer(gl.UNIFORM_BUFFER, C.GlobalMatrix)
	gl.BufferSubData(gl.UNIFORM_BUFFER, 0, 64, gl.Ptr(&C.LookAtV[0]))
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)
}
