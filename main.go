package main

import (
	"fmt"
	"github.com/Dasongzi1366/AutoGo/app"
	"github.com/Dasongzi1366/AutoGo/ime"
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
			once.SetText("脚本定制")
			a.Desc("搜索").FindOnce().Click()
			break
		}
	}

	for i := 0; i < 100; i++ {
		once := a.ClassName("androidx.recyclerview.widget.RecyclerView").FindOnce()
		if once != nil {
			fmt.Println("once", once.ToString())
			for _, object := range once.GetChildren() {
				//fmt.Println(object.ToString())
				if object.ToString() != "" {
					children := object.GetChildren()
					//fmt.Println("ccccc", children)
					for _, child := range children {
						fmt.Println("child", child.ToString())
						fmt.Println(child.ClickCenter())

						for i := 0; i < 10; i++ {
							aaa := a.DescStartsWith("在售").FindOnce()
							if aaa != nil {
								fmt.Println("在售", aaa.ToString())
								aaa.ClickCenter()
							}
						}

						for i := 0; i < 100; i++ {
							ss := a.DescStartsWith("商品价格").FindOnce()
							if ss != nil {
								ss.ClickCenter()
								fmt.Println("2222", ss.ToString())

								for i := 0; i < 100; i++ {
									bb := a.DescStartsWith("聊一聊").FindOnce()
									if bb != nil {
										bb.ClickCenter()
										fmt.Println("3333", bb.ToString())

									}
								}
								for i := 0; i < 100; i++ {
									bb := a.DescStartsWith("我想要").FindOnce()
									if bb != nil {
										bb.ClickCenter()
										fmt.Println("3333", bb.ToString())
									}
								}
								for i := 0; i < 5; i++ {
									dd := a.DescStartsWith("想跟TA说点什么...").FindOnce()
									if dd != nil {
										fmt.Println("想跟TA说点什么", dd.ToString())
										dd.Click()
										ime.InputText("你好")
									}
								}
								for i := 0; i < 5; i++ {
									fmt.Println("key")
									ime.KeyAction(0)
								}
							}
						}

					}
				}
			}
		}
	}
	for i := 0; i < 5; i++ {
		gg := a.DescStartsWith("发送").FindOnce()
		if gg != nil {
			fmt.Println("gg", gg.ToString())
			gg.Click()
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

}
