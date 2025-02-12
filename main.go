package main

import (
	"github.com/Dasongzi1366/AutoGo/app"
	"github.com/Dasongzi1366/AutoGo/motion"
	"github.com/Dasongzi1366/AutoGo/uiacc"
	"github.com/Dasongzi1366/AutoGo/utils"
	"log"
)

func main() {
	a := uiacc.New()
	success := app.Launch("com.taobao.idlefish")
	if !success {
		log.Println("app不存在")
		return
	}
	utils.Sleep(4000)
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

	//for i := 0; i < 200; i++ {
	//	objectss := a.ClassName("androidx.recyclerview.widget.RecyclerView").Id("nested_recycler_view").FindOnce()
	//	if objectss != nil {
	//		fmt.Println(objectss.ToString())
	//		return
	//	}
	//}

	//var l1 *uiacc.UiObject
	//for i := 0; i < 200; i++ {
	//	sc := a.ClassName("androidx.recyclerview.widget.RecyclerView").
	//		Id("nested_recycler_view").
	//		PackageName("com.taobao.idlefish").FindOnce()
	//	if sc != nil {
	//		fmt.Println("sc", sc.ToString())
	//		l1 = sc
	//		break
	//	}
	//}
	//
	//var i = 0
	////var childrenNode []*uiacc.UiObject
	//fmt.Println(len(l1.GetChildren()))
	//for _, object := range l1.GetChildren() {
	//	i++
	//	if i <= 1 {
	//		continue
	//	}
	//	if object != nil {
	//		fmt.Println("l1", object.ToString())
	//		object.ClickCenter()
	//		motion.Back()
	//		utils.Sleep(1000)
	//	}
	//
	//}
	utils.Sleep(2000)
	motion.Swipe(635, 2203, 556, 344, 500)
	//motion.Swipe2(635, 2203, 556, 344, 500)
	//var title map[string]string

}
