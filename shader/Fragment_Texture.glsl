//? 贴图列表
uniform sampler2D Texture[31];
uniform bool IsTexture[31];

//? 内部函数
vec4 GetTexture(id int,uv vec2) {
	if (IsTexture[id]) {
		return texture(Texture[id], uv);
	}
}