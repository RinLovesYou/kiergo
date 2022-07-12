#include "dxhook.h"
#include <d3d11.h>
#include <dxgi.h>
#include <stdint.h>
#include <stdio.h>

#ifdef _UNICODE
#define KIERO_TEXT(text) L##text
#else
#define KIERO_TEXT(text) text
#endif

typedef HRESULT(__stdcall *D3D11PresentHook)(IDXGISwapChain *pSwapChain, UINT SyncInterval, UINT Flags);
typedef void(__stdcall *D3D11DrawIndexedHook)(ID3D11DeviceContext *pContext, UINT IndexCount, UINT StartIndexLocation, INT BaseVertexLocation);
typedef void(__stdcall *D3D11CreateQueryHook)(ID3D11Device *pDevice, const D3D11_QUERY_DESC *pQueryDesc, ID3D11Query **ppQuery);
typedef void(__stdcall *D3D11PSSetShaderResourcesHook)(ID3D11DeviceContext *pContext, UINT StartSlot, UINT NumViews, ID3D11ShaderResourceView *const *ppShaderResourceViews);
typedef void(__stdcall *D3D11ClearRenderTargetViewHook)(ID3D11DeviceContext *pContext, ID3D11RenderTargetView *pRenderTargetView, const FLOAT ColorRGBA[4]);

static HWND window = NULL;
static HMODULE g_hModule = NULL;
static ID3D11Device *device = NULL;
static ID3D11DeviceContext *pContext = NULL;
static IDXGISwapChain *swapChain = NULL;
static bool g_isInitialized;
WNDPROC oWndProc;

D3D11PresentHook phookD3D11Present = NULL;
D3D11DrawIndexedHook phookD3D11DrawIndexed = NULL;
D3D11CreateQueryHook phookD3D11CreateQuery = NULL;
D3D11PSSetShaderResourcesHook phookD3D11PSSetShaderResources = NULL;
D3D11ClearRenderTargetViewHook phookD3D11ClearRenderTargetViewHook = NULL;
ID3D11RenderTargetView *mainRenderTargetView;

DWORD_PTR *pSwapChainVTable = NULL;
DWORD_PTR *pDeviceVTable = NULL;
DWORD_PTR *pDeviceContextVTable = NULL;

WNDCLASSEX windowClass;

void *FindPresent()
{
    printf("FindPresent\n");
   windowClass = CreateDummyWindowClass();

   HWND window = ::CreateWindow(windowClass.lpszClassName, KIERO_TEXT("Kiergo DirectX Window"), WS_OVERLAPPEDWINDOW, 0, 0, 100, 100, NULL, NULL, windowClass.hInstance, NULL);

   HMODULE libD3D11;
   if ((libD3D11 =::LoadLibrary(KIERO_TEXT("d3d11.dll"))) == NULL)
   {
      ::DestroyWindow(window);
      ::UnregisterClass(windowClass.lpszClassName, windowClass.hInstance);
      printf("Failed to load d3d11.dll\n");
      return nullptr;
   }

   void *D3D11CreateDeviceAndSwapChain;
   if ((D3D11CreateDeviceAndSwapChain = reinterpret_cast<void*>(GetProcAddress(libD3D11, "D3D11CreateDeviceAndSwapChain"))) == NULL)
   {
      ::DestroyWindow(window);
      ::UnregisterClass(windowClass.lpszClassName, windowClass.hInstance);
      printf("Failed to get address of D3D11CreateDeviceAndSwapChain\n");
      return nullptr;
   }

   D3D_FEATURE_LEVEL levels[] = { D3D_FEATURE_LEVEL_11_0, D3D_FEATURE_LEVEL_10_1 };
   D3D_FEATURE_LEVEL obtainedLevel;

   DXGI_RATIONAL refreshRate;
   refreshRate.Numerator = 60;
   refreshRate.Denominator = 1;

   DXGI_MODE_DESC bufferDesc;
   bufferDesc.Width = 100;
   bufferDesc.Height = 100;
   bufferDesc.RefreshRate = refreshRate;
   bufferDesc.Format = DXGI_FORMAT_R8G8B8A8_UNORM;
   bufferDesc.ScanlineOrdering = DXGI_MODE_SCANLINE_ORDER_UNSPECIFIED;
   bufferDesc.Scaling = DXGI_MODE_SCALING_UNSPECIFIED;

   DXGI_SAMPLE_DESC sampleDesc;
   sampleDesc.Count = 1;
   sampleDesc.Quality = 0;

   DXGI_SWAP_CHAIN_DESC swapChainDesc;
   swapChainDesc.BufferDesc = bufferDesc;
   swapChainDesc.SampleDesc = sampleDesc;
   swapChainDesc.BufferUsage = DXGI_USAGE_RENDER_TARGET_OUTPUT;
   swapChainDesc.BufferCount = 1;
   swapChainDesc.OutputWindow = window;
   swapChainDesc.Windowed = 1;
   swapChainDesc.SwapEffect = DXGI_SWAP_EFFECT_DISCARD;
   swapChainDesc.Flags = DXGI_SWAP_CHAIN_FLAG_ALLOW_MODE_SWITCH;

   if (((long(__stdcall *)(
          IDXGIAdapter *,
          D3D_DRIVER_TYPE,
          HMODULE,
          UINT,
          const D3D_FEATURE_LEVEL *,
          UINT,
          UINT,
          const DXGI_SWAP_CHAIN_DESC *,
          IDXGISwapChain **,
          ID3D11Device **,
          D3D_FEATURE_LEVEL *,
          ID3D11DeviceContext **))(D3D11CreateDeviceAndSwapChain))(NULL, D3D_DRIVER_TYPE_HARDWARE, NULL, 0, levels, 1, D3D11_SDK_VERSION, &swapChainDesc, &swapChain, &device, &obtainedLevel, &pContext)
      < 0)
   {
      ::DestroyWindow(window);
      ::UnregisterClass(windowClass.lpszClassName, windowClass.hInstance);
      printf("Failed to create device and swap chain\n");
      return nullptr;
   }

   pSwapChainVTable = (DWORD_PTR *)(swapChain);
   pSwapChainVTable = (DWORD_PTR *)(pSwapChainVTable[0]);

   pDeviceVTable = (DWORD_PTR *)(device);
   pDeviceVTable = (DWORD_PTR *)pDeviceVTable[0];

   pDeviceContextVTable = (DWORD_PTR *)(pContext);
   pDeviceContextVTable = (DWORD_PTR *)(pDeviceContextVTable[0]);

   swapChain->Release();
   swapChain = NULL;

   device->Release();
   device = NULL;

   pContext->Release();
   pContext = NULL;

   DestroyWindow(window);
   UnregisterClass(windowClass.lpszClassName, windowClass.hInstance);

   return reinterpret_cast<void*>(pSwapChainVTable[8]);
}

void* SetupValuesGetDevice(void* swapChain) {
    if (!swapChain) {
        return nullptr;
    }

    IDXGISwapChain* pSwapChain = reinterpret_cast<IDXGISwapChain*>(swapChain);
    if (FAILED(pSwapChain->GetDevice(__uuidof(ID3D11Device), (void**)&device))) {
        return nullptr;
    }

    device->GetImmediateContext(&pContext);
    if (!pContext) {
        return nullptr;
    }

    DXGI_SWAP_CHAIN_DESC sd;
    if (FAILED(pSwapChain->GetDesc(&sd))) {
        return nullptr;
    }

    window = sd.OutputWindow;

    ID3D11Texture2D *pBackBuffer;
    if (FAILED(pSwapChain->GetBuffer(0, __uuidof(ID3D11Texture2D), (void**)&pBackBuffer))) {
        return nullptr;
    }

    if (FAILED(device->CreateRenderTargetView(reinterpret_cast<ID3D11Resource*>(pBackBuffer), NULL, &mainRenderTargetView))) {
        return nullptr;
    }

    pBackBuffer->Release();

    return reinterpret_cast<void*>(device);
}

void* GetGameWindow() {
    if (!window) {
        return nullptr;
    }
    return reinterpret_cast<void*>(window);
}

void* GetContext() {
    if (!pContext) {
        return nullptr;
    }
    return reinterpret_cast<void*>(pContext);
}

void SetRenderTargets() {
    pContext->OMSetRenderTargets(1, &mainRenderTargetView, NULL);
}

WNDCLASSEX CreateDummyWindowClass()
{
   windowClass.cbSize = sizeof(WNDCLASSEX);
   windowClass.style = CS_HREDRAW | CS_VREDRAW;
   windowClass.lpfnWndProc = DefWindowProc;
   windowClass.cbClsExtra = 0;
   windowClass.cbWndExtra = 0;
   windowClass.hInstance = GetModuleHandle(NULL);
   windowClass.hIcon = NULL;
   windowClass.hCursor = NULL;
   windowClass.hbrBackground = NULL;
   windowClass.lpszMenuName = NULL;
   windowClass.lpszClassName = KIERO_TEXT("Kiergo");
   windowClass.hIconSm = NULL;

   ::RegisterClassEx(&windowClass);

   return windowClass;
}