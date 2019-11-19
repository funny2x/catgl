package port

///
// 灯光底层绑定
//? 日志
//!  2019年11月19日
//     绑定底层
///
/*
// LightObjectData 静态灯光
type LightObjectData struct {
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

// LightData 灯光
type LightData struct {
	// 着色器变量指针
	GlobalLight uint32 // 着色器灯光全局变量地址
	//? 灯光数据
	Light   LightObjectData
	IsLight bool
	Number  int
}

// Init 初始化
func (L *LightData) Init() error {
	gl.GenBuffers(1, &L.GlobalLight)                           // 创建缓冲器
	gl.BindBuffer(gl.UNIFORM_BUFFER, L.GlobalLight)            // 绑定缓冲区
	gl.BufferData(gl.UNIFORM_BUFFER, 116, nil, gl.STATIC_DRAW) // 分配内存空间
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0)                        // 解除引用
	// 绑定着色器变量定义的绑定点,灯光定义绑定在1号绑定点
	gl.BindBufferRange(gl.UNIFORM_BUFFER, 1, L.GlobalLight, 0, 116)
	return nil
}

// Update 主更新
func (L *LightData) Update() {
	Offset := L.Number * 116
	gl.BindBuffer(gl.UNIFORM_BUFFER, L.GlobalLight) // 绑定缓冲区
	{
		// 同一写
		bsub := func(data unsafe.Pointer, size int) {
			gl.BufferSubData(gl.UNIFORM_BUFFER, Offset, size, data)
			Offset += size
		}
		bsub(gl.Ptr(&(L.Light.Position)[0]), 12) // 位置
		bsub(gl.Ptr(&(L.Light.Ambient)[0]), 12)  // 环境色
		bsub(gl.Ptr(&(L.Light.Diffuse)[0]), 12)  // 漫反射
		bsub(gl.Ptr(&(L.Light.Specular)[0]), 12) // 高光反射
		bsub(gl.Ptr(&L.Light.Linear), 4)         // 线性
		bsub(gl.Ptr(&L.Light.Constant), 4)       // 常量
		bsub(gl.Ptr(&L.Light.Quadratic), 4)      // 二次方程式
	}
	//var B int8
	//if L.IsLight {
	//	B = 1
	//} else {
	//	B = 0
	//}
	//gl.BufferSubData(gl.UNIFORM_BUFFER, Offset+72, 4, gl.Ptr(&B)) // IsLight
	gl.BindBuffer(gl.UNIFORM_BUFFER, 0) // 解除引用
}
*/
