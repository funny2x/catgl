package port

import (
	"gitee.com/RuiCat/catgl/lib/mgl32"
)

///
// 摄像机数据对象
// ? 日志
//! 2019年11月1日 重写
///

// CameraLookAtVData 摄像机朝向结构体
type CameraLookAtVData struct {
	Up      mgl32.Vec3 //? 朝向
	Eye     mgl32.Vec3 //? 位置
	Center  mgl32.Vec3 //? 视点
	LookAtV mgl32.Mat4 //? 摄像机朝向
}

// Init 初始化
func (C *CameraLookAtVData) Init() error {
	//? 计算朝向
	C.LookAtV = mgl32.LookAtV(C.Eye, C.Center, C.Up)
	return nil
}

// Update 主更新
func (C *CameraLookAtVData) Update() {

}

// CameraProjectionData 摄像机投影矩阵结构体
type CameraProjectionData struct {
	Fovy       float32    // 视角
	Aspect     float32    // 屏幕宽高比
	ZNear      float32    // 近平面距离
	ZFar       float32    // 远平面距离
	Projection mgl32.Mat4 // 投影矩阵
}

// Init 初始化
func (C *CameraProjectionData) Init() error {
	// 计算投影矩阵
	C.Projection = mgl32.Perspective(mgl32.DegToRad(C.Fovy), C.Aspect, C.ZNear, C.ZFar)
	return nil
}

// Update 主更新
func (C *CameraProjectionData) Update() {

}
