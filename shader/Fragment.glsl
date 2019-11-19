//! 静态灯光对象           
struct lightObject {
	vec3 Position; 		// 位置
	//* 颜色
	vec3 Ambient;		// 环境色
	vec3 Diffuse;  		// 漫反射
	vec3 Specular;		// 高光反射
	//? 物理参数
	float Linear;       // 线性
	float Constant;     // 常量
	float Quadratic;    // 二次方程式
};

//! 静态灯光数据            
struct lightData {
 	float [17]Light; 
	bool  IsLight; 
};


//? 全局灯光
layout (std140, binding = 1) uniform GlobalLight //* 全局变量
{
	lightData []Light; 
};