#pragma once

#include <stdbool.h>
#include <stdint.h>

#define WIN32_LEAN_AND_MEAN
#include <windows.h>

typedef long (*IDXGISwapChainPresent)(intptr_t pSwapChain, intptr_t SyncInterval, intptr_t Flags);

long InvokePresent(IDXGISwapChainPresent callback, intptr_t pSwapChain, intptr_t SyncInterval, intptr_t Flags);

extern long onPresent(void* pSwapChain, unsigned int SyncInterval, unsigned int Flags);
extern LRESULT wndProc(const HWND hWnd, UINT uMsg, WPARAM wParam, LPARAM lParam);