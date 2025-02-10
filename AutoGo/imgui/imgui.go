package imgui

import (
	"image/color"
)

// TextItem 表示颜色和文本的组合。
type TextItem struct {
	TextColor color.Color //文字颜色 如果为 nil，则默认使用白色 (255,255,255)
	Text      string
}

// HudInit 初始化状态条。
// 参数:
//
//	x1, y1: 状态条的左上角坐标。
//	x2, y2: 状态条的右下角坐标。
//	bgColor: 状态条的背景颜色，如果为 nil，则默认使用灰色 (100, 100, 100)。
//	textSize: 状态条上的文字大小，如果小于等于 0，则默认使用 45。
//
// 返回:
//
//	无返回值。
func HudInit(x1, y1, x2, y2 int, bgColor color.Color, textSize int) {

}

// HudSetText 设置状态条的文本。
// 参数:
//
//	items: 一个包含多个 TextItem 的切片，每个 TextItem 包含文字颜色和文本内容。
//
// 返回:
//
//	无返回值。
func HudSetText(items []TextItem) {

}

// HudClose 销毁状态条。
// 返回:
//
//	无返回值。
func HudClose() {

}

// DrawRect 绘制矩形。
// 参数:
//
//	x1, y1: 矩形的左上角坐标。
//	x2, y2: 矩形的右下角坐标。
//	c: 矩形的颜色。
//
// 返回:
//
//	无返回值。
func DrawRect(x1, y1, x2, y2 int, c color.Color) {

}

// ClearRect 清除绘制的矩形。
// 返回:
//
//	无返回值。
func ClearRect() {

}

// Toast 显示 Toast 提示信息。
// 参数:
//
//	message: 要显示的提示信息。
//
// 返回:
//
//	无返回值。
func Toast(message string) {

}
