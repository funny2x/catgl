package port

import "gitee.com/RuiCat/catgl/lib/gl"

// GameVertexlBuffersType 游戏顶点类型
type GameVertexlBuffersType int

const (
	// GameVertexSTATIC 表示该缓存区不会被修改
	GameVertexSTATIC GameVertexlBuffersType = gl.STATIC_DRAW
	// GameVertexDyNAMIC 表示该缓存区会被周期性更改
	GameVertexDyNAMIC GameVertexlBuffersType = gl.DYNAMIC_DRAW
	// GameVertexSTREAM 表示该缓存区会被频繁更改
	GameVertexSTREAM GameVertexlBuffersType = gl.STREAM_DRAW
)

// GameVertexData 顶点定义
type GameVertexData struct {
	Positions [3]float32 // 位置
	Normals   [3]float32 // 法线
	Color     [3]float32 // 颜色
	Uv        [2]float32 // Uv
	Tangent   [3]float32 // 切线
	Bitangent [3]float32 // 副切线
}

// GameVertexSync 顶点同步标记
type GameVertexSync int

const (
	// GameVertexSyncDefault 默认
	GameVertexSyncDefault GameVertexSync = 0
	// GameVertexSyncRetrieve 取回
	GameVertexSyncRetrieve GameVertexSync = 1
	// GameVertexSyncDelivery 发送
	GameVertexSyncDelivery GameVertexSync = 2
)

// GameVertexDataList 顶点数据列表
type GameVertexDataList struct {
	VBO    uint32                 // 顶点缓冲区指针
	Type   GameVertexlBuffersType // 顶点的类型: 静态,动态,实时. 用于 Data 跟GPU缓冲区的同步关系
	Data   []GameVertexData       // 存储顶点数据的列表(静态),需要设置更新到缓冲区或者从缓存区读取到最新的.
	Status GameVertexSync         // 顶点状态,用于取回跟更新顶点数据标记用.
}
