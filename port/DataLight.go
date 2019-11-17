package port

// LightData  着色器灯光底层数据结构
type LightData struct {
	//? 灯光列表
	Light   [10]*Light
	IsLight [10]bool //? 灯光是否有效标记
}

// Init 初始化
func (L *LightData) Init() error {
	return nil
}

// Update 主更新
func (L *LightData) Update() {

}
