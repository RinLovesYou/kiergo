package kiergo

// #cgo CPPFLAGS: -I./directx -DIMGUI_DISABLE_WIN32_DEFAULT_IME_FUNCTIONS
// #cgo CXXFLAGS: -std=c++11
// #cgo CXXFLAGS: -Wno-subobject-linkage
// #cgo CFLAGS: -I${SRCDIR}/lib -I.
// #cgo LDFLAGS: -L${SRCDIR}/lib/d3d11.a -L${SRCDIR}/lib/d3dcompiler.a -L${SRCDIR}/lib/dxgi.a -ld3dcompiler -ld3d11 -ldxgi -ldwmapi -lgdi32
import "C"
