# Kiero-Golang-DirectX11-Hook

This is a very dirty Golang implementation of [Kiero Hook for Directx11](https://github.com/Rebzzel/kiero). <br>
For a QuickStart, you can check out the example folder.

## Information
This works due to a little modification of the [imgui-go wrapper by inkyblackness](https://github.com/inkyblackness/imgui-go).<br>
you can get it with `go get github.com/RinLovesYou/imgui-go`<br>

!! THIS MAKES A LOT OF ASSUMPTIONS ABOUT YOUR ENVIRONMENT. IT EXPECTS DEPENDENCIES LIKE MINHOOK/DX11/ETC TO BE INJECTED/PRESENT !!

![image](https://user-images.githubusercontent.com/29461788/178161976-84d73af1-1407-48c3-90e7-26b723a6cbf1.png)


## QuickStart

```go
package main

import (
	"fmt"
	"unsafe"

	"github.com/RinLovesYou/imgui-go"
	"github.com/RinLovesYou/kiergo"
)

func init() {
	err := kiergo.Hook(OnPresent, WndProc)
	if err != nil {
		fmt.Println(err)
	}
}

var gInitialized bool
var device unsafe.Pointer
var context unsafe.Pointer
var window unsafe.Pointer

var showDemo bool
var showGuide bool

func OnPresent(pSwapChain unsafe.Pointer, SyncInterval, FlagsT uint32) error {
	if !gInitialized {
		var err error

		device, err = kiergo.SetupValuesGetDevice(pSwapChain)
		if err != nil {
			return err
		}

		context, err = kiergo.GetContext()
		if err != nil {
			return err
		}

		window, err = kiergo.GetGameWindow()
		if err != nil {
			return err
		}

		imgui.CreateContext(nil)
		io := imgui.CurrentIO()
		io.SetConfigFlags(imgui.ConfigFlagsNavEnableKeyboard)

		imgui.Win32Init(window)
		imgui.Dx11Init(device, context)

		gInitialized = true
	}

	imgui.Dx11NewFrame()
	imgui.Win32NewFrame()
	imgui.NewFrame()

	imgui.Begin("hello, golang")

	imgui.Text("\"it just works\" - RinLovesYou 2022")
	if imgui.Button("Show Demo") {
		showDemo = !showDemo
	}

	if imgui.Button("Show Guide") {
		showGuide = !showGuide
	}

	imgui.End()

	if showDemo {
		imgui.ShowDemoWindow(&showDemo)
	}

	if showGuide {
		imgui.ShowUserGuide()
	}

	imgui.Render()
	kiergo.SetRenderTargets()
	imgui.Dx11RenderDrawData(imgui.RenderedDrawData())

	return nil
}

func WndProc(hwnd unsafe.Pointer, msg uint32, wparam, lparam unsafe.Pointer) error {
	imgui.Win32WndProcHandler(hwnd, msg, wparam, lparam)
	return nil
}

func main() {

}

```

compile with `go build --buildmode=c-shared -o YourDll.dll main.go`
