package catgl

import (
	"errors"
	"runtime"
	"time"

	"gitee.com/RuiCat/catgl/lib/gl"
	"gitee.com/RuiCat/catgl/lib/glfw"
	"gitee.com/RuiCat/catgl/port"
)

///
// 渲染器
//  创建渲染窗口与管理窗口事件,绑定场景等操作
// ? 日志
//! 2019年11月2日 重写
///

// Renderer 渲染器
type Renderer struct {
	window *glfw.Window
	err    chan error
	//? 场景对象
	Scene port.SceneObject
	//? 窗口参数
	Width  int
	Height int
	Title  string
	Color  [4]float32
	//? 结构体参数
	AspectRatio float32 //* 屏幕高宽比
}

// New 创建窗口
func (R *Renderer) New() chan error {
	R.err = make(chan error, 1)
	if R.Width > 0 && R.Height > 0 {
		//? 错误检测
		{
			if R.Scene == nil {
				R.err <- errors.New("窗口创建失败:未绑定场景")
				return R.err
			}
		}
		R.AspectRatio = float32(R.Width / R.Height)
		WindowManager.add(R)
		return R.err
	}
	R.err <- errors.New("窗口创建失败: 窗口大小为空")
	return R.err
}

// Update 更新函数
//? 在返回前住线程会被阻断
func (R *Renderer) Update() {
	gl.ClearColor(R.Color[0], R.Color[1], R.Color[2], R.Color[3])
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	// 更新场景
	R.Scene.Update()
}

// manager 窗口管理程序
type manager struct {
	Fps      *time.Ticker
	IsStatus bool
	//? 内部参数
	new  func()
	add  func(R *Renderer)
	list []*Renderer
}

// WindowManager 窗口管理程序
var WindowManager *manager

func init() {
	NewWindow := make(chan *Renderer)
	WindowManager = &manager{
		//? 默认窗口刷新频率
		Fps:      time.NewTicker(time.Millisecond * time.Duration(6)),
		IsStatus: true,
		//? 初始化
		new: func() {
			//? 绑定到主进程
			runtime.LockOSThread()
			defer runtime.UnlockOSThread()
			//? 初始化 glfw
			defer glfw.Terminate()
			if err := glfw.Init(); err != nil {
				panic(err)
			}
			//? 设置
			glfw.WindowHint(glfw.Samples, 8)                            //* MSAA
			glfw.WindowHint(glfw.Resizable, glfw.False)                 //* 固定窗口大小
			glfw.WindowHint(glfw.ContextVersionMajor, 4)                //* 主版本号
			glfw.WindowHint(glfw.ContextVersionMinor, 6)                //* 次版本号
			glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) //* 核心模式
			glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    //* OpenGL 兼容
			//? 主更新周期
			for WindowManager.IsStatus {
				select {
				case R := <-NewWindow: //? 窗口创建事件
					//? 创建窗口
					var err error
					R.window, err = glfw.CreateWindow(R.Width, R.Height, R.Title, nil, nil)
					if err != nil {
						R.err <- err
						continue
					}
					//? 设置当前窗口上下文
					R.window.MakeContextCurrent()
					//? 初始化Gl库
					if err := gl.Init(); err != nil {
						R.err <- err
						continue
					}
					//? 初始化场景
					err = R.Scene.Init()
					if err != nil {
						R.err <- err
						continue
					}
					//? 分离上下文
					glfw.DetachCurrentContext()
					//? 创建完成
					WindowManager.list = append(WindowManager.list, R)
					R.err <- nil
				case <-WindowManager.Fps.C: //? 刷新周期锁
					glfw.PollEvents() //? 轮查事件
					//? 循环处理窗口事件
					for i, R := range WindowManager.list {
						if !R.window.ShouldClose() {
							R.window.MakeContextCurrent() //? 设置当前窗口上下文
							//! 用户绘画处理
							R.Update()
							//? 交换缓冲区
							R.window.SwapBuffers()
							glfw.DetachCurrentContext() //? 分离上下文
						} else {
							R.window.Destroy()
							WindowManager.list = append(WindowManager.list[:i], WindowManager.list[i+1:]...)
							R.err <- nil
						}
					}
				}
			}
		},
		add: func(R *Renderer) {
			NewWindow <- R //? 创建窗口
		},
	}
	go WindowManager.new()
}
