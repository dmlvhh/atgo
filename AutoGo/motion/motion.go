package motion

// TouchDown 模拟触摸屏按下操作。
// 参数:
//
//	x: 触摸点的 X 坐标。
//	y: 触摸点的 Y 坐标。
//	fingerID: 触摸点的指针 ID（1-10）。
func TouchDown(x, y, fingerID int) {

}

// TouchMove 模拟触摸屏移动操作。
// 参数:
//
//	x: 移动到的 X 坐标。
//	y: 移动到的 Y 坐标。
//	fingerID: 触摸点的指针 ID（1-10）。
func TouchMove(x, y, fingerID int) {

}

// TouchUp 模拟触摸屏抬起操作。
// 参数:
//
//	x: 抬起点的 X 坐标。
//	y: 抬起点的 Y 坐标。
//	fingerID: 触摸点的指针 ID（1-10）。
func TouchUp(x, y, fingerID int) {

}

// Click 模拟单击操作。
// 参数:
//
//	x: 单击点的 X 坐标。
//	y: 单击点的 Y 坐标。
//	fingerID: 触摸点的指针 ID（1-10）。
func Click(x, y, fingerID int) {

}

// LongClick 模拟长按操作。
// 参数:
//
//	x: 长按点的 X 坐标。
//	y: 长按点的 Y 坐标。
//	duration: 长按持续时间（毫秒）。
func LongClick(x, y, duration int) {

}

// Swipe 模拟滑动操作。
// 参数:
//
//	x1: 起始点的 X 坐标。
//	y1: 起始点的 Y 坐标。
//	x2: 结束点的 X 坐标。
//	y2: 结束点的 Y 坐标。
//	duration: 滑动持续时间（毫秒）。
func Swipe(x1, y1, x2, y2, duration int) {

}

// Swipe2 使用贝塞尔曲线方式进行滑动
// 参数:
//
//	x1: 起始点的 X 坐标。
//	y1: 起始点的 Y 坐标。
//	x2: 结束点的 X 坐标。
//	y2: 结束点的 Y 坐标。
//	duration: 滑动持续时间（毫秒）。
func Swipe2(x1, y1, x2, y2, duration int) {

}

// Home 模拟按下 Home 键。
func Home() {

}

// Back 模拟按下返回键。
func Back() {

}

// Recents 显示最近任务。
func Recents() {

}

// PowerDialog 弹出电源键菜单。
func PowerDialog() {

}

// Notifications 拉出通知栏。
func Notifications() {

}

// QuickSettings 显示快速设置(下拉通知栏到底)。
func QuickSettings() {

}

// VolumeUp 按下音量上键。
func VolumeUp() {

}

// VolumeDown 按下音量下键。
func VolumeDown() {

}

// Camera 模拟按下照相键。
func Camera() {

}

// KeyAction 模拟按键 code值参考KEYCODE_开头常量。
func KeyAction(code int) {

}
