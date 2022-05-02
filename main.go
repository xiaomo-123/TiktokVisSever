package main

import (
	// "fmt"
	"TiktokVisSever/controller"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// https://github.com/lusingander/fyne-font-example
//https://www.vmlogin.com.cn/Instruction.pdf
func main() {
	a := app.New()
	w := a.NewWindow("登录窗口")
	// w:= fyne.CurrentApp().NewWindow("登录窗口")

	f := widget.NewForm()
	f.Append("username", widget.NewEntry())
	f.Append("password", widget.NewEntry())

	btn := widget.NewButton("Go", func() {
		w.Close()
		controller.Start(w)
		// fmt.Println("submit")
	})
	w.SetContent(
		container.NewVBox(f, btn))
	// container.NewBorder(f, btn, nil, nil, nil))

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(400, 300))
	w.SetFixedSize(true)
	w.ShowAndRun()
}
