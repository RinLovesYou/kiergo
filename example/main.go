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

		setStyle()

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
