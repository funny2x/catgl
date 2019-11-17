package port

import (
	"errors"
	"strings"

	"gitee.com/RuiCat/catgl/lib/gl"
)

// ID uuid用于标记用
type ID string

// Init 初始化接口
type Init interface {
	Init() error
}

// CallInit 回调初始化
type CallInit struct {
	Init
	Call chan error
}

// GameObject 游戏对象
type GameObject interface {
	UpdateBeforehand() //? 预更新
	UpdateRendering()  //? 渲染更新
}

// SceneObject 场景对象
type SceneObject interface {
	Init
	Update()
}

// CameraObject 摄像机对象
type CameraObject interface {
	Init
	Update() //? 摄像机更新

}

// LightObject 灯光对象
type LightObject interface {
	Init
	Update() //? 灯光更新
}

// ShaderObject 着色器对象
type ShaderObject interface {
	Init
	GetProgram() uint32 //? 得到着色器指针
}

// GetLog 得到错误日志
func GetLog(Program uint32) error {
	var logLength int32
	gl.GetProgramiv(Program, gl.INFO_LOG_LENGTH, &logLength)
	log := strings.Repeat("\x00", int(logLength+1))
	gl.GetProgramInfoLog(Program, logLength, nil, gl.Str(log))
	if log == "\x00" {
		return nil
	}
	return errors.New(log)
}
