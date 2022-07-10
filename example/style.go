package main

import "github.com/RinLovesYou/imgui-go"

func setStyle() {
	style := imgui.CurrentStyle()

	style.SetColor(imgui.StyleColorText, getColor(1, 1, 1, 1))
	style.SetColor(imgui.StyleColorTextDisabled, getColor(0.5, 0.5, 0.5, 1))
	style.SetColor(imgui.StyleColorWindowBg, getColor(0.06, 0.06, 0.06, 0.94))
	style.SetColor(imgui.StyleColorChildBg, getColor(1, 1, 1, 0))
	style.SetColor(imgui.StyleColorPopupBg, getColor(0.08, 0.08, 0.08, 0.94))
	style.SetColor(imgui.StyleColorBorder, getColor(0.43, 0.43, 0.5, 0.5))
	style.SetColor(imgui.StyleColorBorderShadow, getColor(0, 0, 0, 0))
	style.SetColor(imgui.StyleColorFrameBg, getColor(0.20, 0.21, 0.22, 0.54))
	style.SetColor(imgui.StyleColorFrameBgHovered, getColor(0.4, 0.4, 0.4, 0.4))
	style.SetColor(imgui.StyleColorFrameBgActive, getColor(0.18, 0.18, 0.18, 67))
	style.SetColor(imgui.StyleColorTitleBg, getColor(0.4, 0.4, 0.4, 1))
	style.SetColor(imgui.StyleColorTitleBgActive, getColor(0.29, 0.29, 0.29, 1.00))
	style.SetColor(imgui.StyleColorTitleBgCollapsed, getColor(0.00, 0.00, 0.00, 0.51))
	style.SetColor(imgui.StyleColorMenuBarBg, getColor(0.14, 0.14, 0.14, 1.00))
	style.SetColor(imgui.StyleColorScrollbarBg, getColor(0.02, 0.02, 0.02, 0.53))
	style.SetColor(imgui.StyleColorScrollbarGrab, getColor(0.31, 0.31, 0.31, 1.00))
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, getColor(0.41, 0.41, 0.41, 1.00))
	style.SetColor(imgui.StyleColorScrollbarGrabActive, getColor(0.51, 0.51, 0.51, 1.00))
	style.SetColor(imgui.StyleColorCheckMark, getColor(0.94, 0.94, 0.94, 1.00))
	style.SetColor(imgui.StyleColorSliderGrab, getColor(0.51, 0.51, 0.51, 1.00))
	style.SetColor(imgui.StyleColorSliderGrabActive, getColor(0.86, 0.86, 0.86, 1.00))
	style.SetColor(imgui.StyleColorButton, getColor(0.44, 0.44, 0.44, 0.40))
	style.SetColor(imgui.StyleColorButtonHovered, getColor(0.46, 0.47, 0.48, 1.00))
	style.SetColor(imgui.StyleColorButtonActive, getColor(0.42, 0.42, 0.42, 1.00))
	style.SetColor(imgui.StyleColorHeader, getColor(0.70, 0.70, 0.70, 0.31))
	style.SetColor(imgui.StyleColorHeaderHovered, getColor(0.70, 0.70, 0.70, 0.80))
	style.SetColor(imgui.StyleColorHeaderActive, getColor(0.48, 0.50, 0.52, 1.00))
	style.SetColor(imgui.StyleColorSeparator, getColor(0.43, 0.43, 0.50, 0.50))
	style.SetColor(imgui.StyleColorSeparatorHovered, getColor(0.72, 0.72, 0.72, 0.78))
	style.SetColor(imgui.StyleColorSeparatorActive, getColor(0.51, 0.51, 0.51, 1.00))
	style.SetColor(imgui.StyleColorResizeGrip, getColor(0.91, 0.91, 0.91, 0.25))
	style.SetColor(imgui.StyleColorResizeGripHovered, getColor(0.81, 0.81, 0.81, 0.67))
	style.SetColor(imgui.StyleColorResizeGripActive, getColor(0.46, 0.46, 0.46, 0.95))
	style.SetColor(imgui.StyleColorPlotLines, getColor(0.61, 0.61, 0.61, 1.00))
	style.SetColor(imgui.StyleColorPlotLinesHovered, getColor(1.00, 0.43, 0.35, 1.00))
	style.SetColor(imgui.StyleColorPlotHistogram, getColor(0.73, 0.60, 0.15, 1.00))
	style.SetColor(imgui.StyleColorPlotHistogramHovered, getColor(1.00, 0.60, 0.00, 1.00))
	style.SetColor(imgui.StyleColorTextSelectedBg, getColor(0.87, 0.87, 0.87, 0.35))
	style.SetColor(imgui.StyleColorModalWindowDarkening, getColor(0.80, 0.80, 0.80, 0.35))
	style.SetColor(imgui.StyleColorDragDropTarget, getColor(1.00, 1.00, 0.00, 0.90))
	style.SetColor(imgui.StyleColorNavHighlight, getColor(0.60, 0.60, 0.60, 1.00))
	style.SetColor(imgui.StyleColorNavWindowingHighlight, getColor(1.00, 1.00, 1.00, 0.70))

	style.SetWindowPadding(imgui.Vec2{X: 8, Y: 6})
	style.SetWindowRounding(3)
	style.SetFramePadding(imgui.Vec2{X: 5, Y: 7})
	style.SetItemSpacing(imgui.Vec2{X: 5, Y: 5})
}

func getColor(r, g, b, a float32) imgui.Vec4 {
	return imgui.Vec4{X: r, Y: g, Z: b, W: a}
}
