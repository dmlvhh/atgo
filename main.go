package main

import (
	"github.com/Dasongzi1366/AutoGo/uiacc"
)

// com.taobao.idlefish
func main() {
	//测试环境为雷电模拟器安卓9 分辨率为手机版1080x1920
	//创建一个节点操作的对象
	a := uiacc.New()
	////查找节点
	//a.Id("com.taobao.idlefish:id/search_bar_layout").FindOnce().Click()
	//ss := a.Index(3).FindOnce()
	//fmt.Println(ss.Click())

	//n := a.DescMatches("^.*,搜索,点击跳转到搜索激活页").FindOnce()
	//n.Click()
	//a.ClassName("android.widget.EditText").FindOnce().Click()
	for i := 0; i < 40; i++ {
		ss := a.ClassName("android.widget.EditText").FindOnce()
		if ss != nil {
			ss.SetText("纯流量")
		}
	}
	a.Desc("搜索").FindOnce().Click()

	//fmt.Println(a.ClassName("android.widget.TextView").Find())
	//f := a.ClassName("android.widget.LinearLayout").FindOnce().Click()
	//if f {
	//	fmt.Println(f)
	//}
	//t := a.ClassName("android.widget.EditText").FindOnce().Click()
	//fmt.Println(t)
	//obj := a.ClassNameContains("android.widget.EditText").FindOnce()
	//if obj != nil {
	//	fmt.Println("obj", obj)
	//	obj.SetText("12314")
	//}

	//obj := a.Text("在设置中搜索").FindOnce()

	//if obj != nil {
	//
	//	//打印节点信息
	//	fmt.Println(obj.GetClickable())
	//	obj.GetClickable()
	//	//判断节点如果可以点击的话直接点击节点
	//	if obj.GetClickable() {
	//		obj.Click()
	//		return
	//	}
	//
	//	//节点不能点击的情况下循环3次判断父节点是否可以点击,如果可以点击就点击后跳出
	//	for i := 0; i < 3; i++ {
	//		obj = obj.GetParent()
	//
	//		//打印节点信息
	//		fmt.Println(obj.ToString())
	//
	//		if obj.GetClickable() {
	//			obj.Click()
	//			return
	//		}
	//	}
	//
	//} else {
	//	fmt.Println("没有找到")
	//}

	//setText()
}
