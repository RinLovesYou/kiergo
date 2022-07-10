package kiergo

import "errors"

var (
	ErrFindPresent          = errors.New("failed to find d3d11 present")
	ErrSetupValuesGetDevice = errors.New("failed to get d3d11 device")
	ErrGetGameWindow        = errors.New("failed to get game window")
	ErrGetContext           = errors.New("failed to get context")
)
