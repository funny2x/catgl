//? 全局灯光
layout (std140) uniform GlobalLight //* 全局变量
{
    Light [10]Light;
    bool  [10]IsLight;
};