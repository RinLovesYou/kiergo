#pragma once

#include <Windows.h>

#ifdef __cplusplus
extern "C" {
#endif

extern void* FindPresent();
extern void* SetupValuesGetDevice(void* swapChain);
extern void* GetGameWindow();
extern void* GetContext();
extern void SetRenderTargets();


WNDCLASSEX CreateDummyWindowClass();

#ifdef __cplusplus
}
#endif