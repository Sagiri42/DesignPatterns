package main

// 简单工厂模式
// 是一种 "创建型设计模式", 它提供了一种方法来创建一个结构体, 而不需要暴露结构体的创建逻辑

// Shape 一个形状的接口, 这个接口需要实现一个 "画" 的方法
type Shape interface {
	Draw() string
}

// Circle 这是一个圆型的结构体, 这个结构体需要实现Shape接口
type Circle struct {
}

// Draw 实现Shape接口, 画一个圆
func (c *Circle) Draw() string {
	return "Draw a circle"
}

// Rectangle 这是一个矩形结构体, 这个结构体需要实现Shape接口
type Rectangle struct {
}

// Draw 实现Shape接口, 画一个矩形
func (c *Rectangle) Draw() string {
	return "Draw a rectangle"
}

// NewShape 简单工厂模式, 生成一个形状
func NewShape(shapeType string) Shape {
	switch shapeType {
	case "circle":
		return new(Circle)
	case "rectangle":
		return new(Rectangle)
	}
	return nil
}
