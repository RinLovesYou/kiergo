package kiergo

//#include "hook.h"
import "C"
import (
	"errors"
	"unsafe"

	"github.com/nanitefactory/gominhook"
)

var (
	presentFunc func(unsafe.Pointer, uint32, uint32) error
	wndProcFunc func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) error
	initialized bool

	originalPresent C.IDXGISwapChainPresent
	originalWndProc uintptr
	presentFunction uintptr
)

func Hook(present func(unsafe.Pointer, uint32, uint32) error, wndProc func(unsafe.Pointer, uint32, unsafe.Pointer, unsafe.Pointer) error) error {
	if initialized {
		return errors.New("kiergo: already initialized")
	}

	presentFunc = present
	wndProcFunc = wndProc
	var err error

	presentFunction, err = findPresent()
	if err != nil {
		return err
	}

	err = gominhook.Initialize()
	if err != nil {
		return err
	}

	err = gominhook.CreateHook(presentFunction, uintptr(C.onPresent), uintptr(unsafe.Pointer(&originalPresent)))
	if err != nil {
		return err
	}

	return gominhook.EnableHook(presentFunction)
}

var inited bool

//export onPresent
func onPresent(pSwapChain unsafe.Pointer, SyncInterval, FlagsT uint32) C.long {
	var err error

	if !inited {
		_, err = SetupValuesGetDevice(pSwapChain)
		if err != nil {
			goto invoke
		}

		_, err = GetContext()
		if err != nil {
			goto invoke
		}

		window, err := GetGameWindow()
		if err != nil {
			goto invoke
		}

		originalWndProc = uintptr(C.SetWindowLongPtr((C.HWND)(window), C.GWLP_WNDPROC, C.LONG_PTR(uintptr(C.wndProc))))
		inited = true
	}

	err = presentFunc(pSwapChain, SyncInterval, FlagsT)
	if err != nil {
		gominhook.DisableHook(presentFunction)
	}

invoke:
	return C.InvokePresent(originalPresent, C.intptr_t(uintptr(pSwapChain)), C.intptr_t(SyncInterval), C.intptr_t(FlagsT))
}

//export wndProc
func wndProc(hwnd C.HWND, msg C.UINT, wparam C.WPARAM, lparam C.LPARAM) C.LRESULT {
	wndProcFunc(unsafe.Pointer(hwnd), uint32(msg), unsafe.Pointer(&wparam), unsafe.Pointer(&lparam))

	return C.CallWindowProc((*[0]byte)(unsafe.Pointer(originalWndProc)), hwnd, msg, wparam, lparam)
}
