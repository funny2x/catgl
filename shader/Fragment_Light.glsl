//! 静态灯光
struct Light {
	vec3 Position; //* 位置
	//* 颜色
	vec3 Ambient;		//? 环境色
	vec3 Diffuse;  		//? 漫反射
	vec3 Specular;		//? 高光反射
	//? 物理参数
	float Constant;    //* 常量
	float Linear;      //* 线性
	float Quadratic;   //* 二次方程式
};