package yolov5

type YoloV5 struct {
}

// DetectResult 表示对象检测的结果，包括位置、标签和置信度。
type DetectResult struct {
	X       int     `json:"X"`
	Y       int     `json:"Y"`
	Width   int     `json:"宽"`
	Height  int     `json:"高"`
	Label   string  `json:"标签"`
	Score   float64 `json:"精度"`
	CenterX int     `json:"-"` //中心坐标X
	CenterY int     `json:"-"` //中心坐标Y
}

// New 创建一个新的 YOLOv5 实例，并加载模型和标签。
// 参数:
//
//	cpuThreadNum: 用于模型推理的 CPU 线程数。
//	paramPath: 模型参数文件路径。
//	binPath: 模型二进制文件路径。
//	labels: 标签文件路径。
//
// 返回:
//
//	*YoloV5: 新创建的 YOLOv5 实例，如果加载失败则返回 nil。
func New(cpuThreadNum int, paramPath, binPath, labels string) *YoloV5 {
	return nil
}

// Detect 在指定的屏幕区域执行目标检测。
// 参数:
//
//	x1, y1: 检测区域的左上角坐标。
//	x2, y2: 检测区域的右下角坐标。如果 x2 或 y2 为 0，则表示使用设备的最大宽度或高度。
//
// 返回:
//
//	[]DetectResult: 检测结果列表。如果检测失败或没有检测到任何结果，则返回 nil。
func (y *YoloV5) Detect(x1, y1, x2, y2 int) []DetectResult {
	return nil
}

// Close 关闭 YOLOv5 模型实例，释放相关资源。
func (y *YoloV5) Close() {

}
