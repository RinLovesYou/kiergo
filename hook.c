#include "hook.h"

long InvokePresent(IDXGISwapChainPresent callback, intptr_t pSwapChain, intptr_t SyncInterval, intptr_t Flags)
{
    return callback(pSwapChain, SyncInterval, Flags);
}