package kiergo

// #include "directx/dxhook.h"
import "C"
import "unsafe"

func findPresent() (uintptr, error) {
	presentPtr := C.FindPresent()
	if presentPtr == nil {
		return 0, ErrFindPresent
	}

	return uintptr(presentPtr), nil
}

func SetupValuesGetDevice(swapChain unsafe.Pointer) (unsafe.Pointer, error) {
	devicePtr := C.SetupValuesGetDevice(swapChain)
	if devicePtr == nil {
		return nil, ErrSetupValuesGetDevice
	}

	return devicePtr, nil
}

func GetGameWindow() (unsafe.Pointer, error) {
	windowPtr := C.GetGameWindow()
	if windowPtr == nil {
		return nil, ErrGetGameWindow
	}

	return windowPtr, nil
}
func GetContext() (unsafe.Pointer, error) {
	contextPtr := C.GetContext()
	if contextPtr == nil {
		return nil, ErrGetContext
	}

	return contextPtr, nil
}
func SetRenderTargets() {
	C.SetRenderTargets()
}
