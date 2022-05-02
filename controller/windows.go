package controller

//______________________________________________________________________________

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"fyne.io/fyne/v2"

	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/layout"
)

// https://github.com/fpabl0/fyne/tree/feature/aligned-expanded-boxlayout
//E:\gowork\pkg\mod\fyne.io\fyne\v2@v2.1.4\cmd\fyne_demo
//https://github.com/fyne-io/fyne/pull/989
//https://github.com/andydotxyz/fyne/blob/22748166c14bd938dbe068d8880f6a7b3cb986b9/cmd/fyne_demo/screens/window.go

var tree *widget.Tree
var data map[string][]string
var input1 *widget.Entry
var input2 *widget.Entry
var label1 *widget.Label
var label2 *widget.Label
var DetailsContainer *fyne.Container
var selectedObjects []fyne.CanvasObject
var list *widget.List

//______________________________________________________________________________
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//  fmt.Println(path)
		// //楷体:simkai.ttf //
		//黑体:simhei.ttf //
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

//https://github.com/fyne-io/fyne/issues/2348
// https://github.com/namezis/fyne-tree-with-details-form/blob/main/main.go
//https://github.com/fyne-io/fyne/issues/2590自定义容器
func Start(mainW fyne.Window) {
	// os.Setenv("FYNE_FONT", "./font/Ubuntu Mono derivative Powerline Bold.ttf")

	data = makeData()

	// myApp := app.New()
	// myWindow := myApp.NewWindow("指纹浏览器管理窗口")
	// myApp := app.New()

	myWindow := fyne.CurrentApp().NewWindow("指纹浏览器管理窗口")

	tree = makeTree()
	input1 = widget.NewEntry()
	input2 = widget.NewEntry()
	bottom := container.NewVBox(input1, input2)

	content := container.NewBorder(
		widget.NewLabel("菜单列表"),
		bottom,
		nil, nil,
		tree)

	// widget := newTestWidget()
	DetailsContainer = container.NewMax()
	cc1:=container.NewCenter(container.NewHBox(
		widget.NewLabel("欢迎进入指纹系统"),
		widget.NewLabel("欢迎进入指纹系统"),
	))
	DetailsContainer.Objects = []fyne.CanvasObject{cc1}
	DetailsContainer.Refresh()
	cc := container.NewHSplit(content, DetailsContainer)

	cc.SetOffset(0.2)

	myWindow.SetContent(cc)
	myWindow.Resize(fyne.NewSize(800, 800))
	myWindow.SetMaster()
	myWindow.Show()
}

func updateList() {
	tree.Select(input2.Text)
}
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

//______________________________________________________________________________
func updatform(id string) {
	cc1:=container.NewCenter(container.NewHBox(
		widget.NewLabel("系统"),
		widget.NewLabel("不一样的统"),
	))
	DetailsContainer.Objects = []fyne.CanvasObject{cc1}
	DetailsContainer.Refresh()
}

// func makeUI() *fyne.CanvasObject {
// 	var grid *fyne.Container
// 	var button1, button2 *widget.Button

// 	const buttonsLength = 50

// 	buttons := make([]fyne.CanvasObject, buttonsLength)
// 	for i := range buttons {
// 		buttons[i] = widget.NewButton(strconv.Itoa(i), nil)
// 	}

// 	button1 = widget.NewButton("Add", func() {
// 		for i := range buttons {
// 			grid.Add(buttons[i])
// 		}
// 		button1.Disable()
// 		button2.Enable()
// 	})

// 	button2 = widget.NewButton("Remove", func() {
// 		for i := buttonsLength - 1; i >= 0; i-- {
// 			grid.Remove(buttons[i])
// 		}
// 		button1.Enable()
// 		button2.Disable()
// 	})

// 	button2.Disable()
// 	grid = container.NewGridWithColumns(20)
// 	return *container.NewVBox(button1, button2, grid)
// }
func makeTree() *widget.Tree {
	return &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return data[uid]
		},
		IsBranch: func(uid string) bool {
			children := data[uid]
			return len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("?")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			//list切割的数组
			list := strings.Split(uid, ".")
			// fmt.Println(list)
			// fmt.Println(uid)
			if branch {
				obj.(*widget.Label).SetText(fmt.Sprintf("%s", list[len(list)-1]))
			} else {
				obj.(*widget.Label).SetText(fmt.Sprintf("%s", list[len(list)-1]))
			}
		},
		OnSelected: func(uid string) {
			// input2.SetText(input1.Text)
			if uid == "UserA.新建" {
				str := "新建"
				input1.SetText(str)
				updatform(str)
			} else if uid == "UserB.配置" {
				str := "配置"
				input1.SetText(str)
				updatform(str)
			}

			// label1.SetText(uid)
			// label2.SetText(uid)
		},
	}
}

func makeData() map[string][]string {
	return map[string][]string{
		"": {
			"新建浏览器配置文件",
			"浏览器配置文件",
		},
		"新建浏览器配置文件": {
			"UserA.新建",
			"UserA.菜单二",
		},
		"UserA.Hints": {
			"UserA.Hints.Mail A",
			"UserA.Hints.Zoom A",
		},
		"浏览器配置文件": {
			"UserB.配置",
			"UserB.菜单二",
		},
		"UserB.Hints": {
			"UserB.Hints.Mail B",
			"UserB.Hints.Zoom B",
		},
	}
}
