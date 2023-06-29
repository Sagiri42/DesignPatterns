package main

import (
	"log"
	"testing"
)

func TestNewShape(t *testing.T) {
	circle := NewShape("circle")       // 创建一个圆型
	rectangle := NewShape("rectangle") // 创建一个矩形

	log.Println(circle.Draw())    // 画一个圆形
	log.Println(rectangle.Draw()) // 画一个矩形
}
