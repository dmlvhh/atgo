package main

import (
	"github.com/Dasongzi1366/AutoGo/app"
	"github.com/Dasongzi1366/AutoGo/uiacc"
	"log"
)

func main() {
	a := uiacc.New()
	success := app.Launch("com.taobao.idlefish")
	if !success {
		log.Println("app不存在")
		return
	}
	for i := 0; i < 200; i++ {
		n := a.DescMatches("^.*,搜索,点击跳转到搜索激活页").FindOnce()
		if n != nil {
			n.Click()
			break
		}
	}

	for i := 0; i < 200; i++ {
		once := a.ClassName("android.widget.EditText").FindOnce()
		if once != nil {
			once.Click()
			once.SetText("手机")
			a.Desc("搜索").FindOnce().Click()
			break
		}
	}
	for i := 0; i < 200; i++ {
		av := a.Text("包邮").FindOnce()
		if av != nil {
			av.Click()
		}
	}

	//for i := 0; i < 10; i++ {
	//	once2 := a.ClassName("androidx.viewpager.widget.ViewPager").FindOnce()
	//	if once2 != nil {
	//		fmt.Println(once2.ToString())
	//		for _, object := range once2.GetChildren() {
	//			//打印节点信息
	//			fmt.Println(object.ToString())
	//			if object.GetClickable() {
	//				object.Click()
	//				return
	//			}
	//		}
	//	}
	//}

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
