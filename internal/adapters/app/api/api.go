package api

import "github.com/MdSadiqMd/GoHexCore/internal/ports"

type Adapter struct {
	arith ports.APIPort
}

func NewAdapter(arith ports.APIPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.GetAddition(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (aria Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := aria.arith.GetSubtraction(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (aria Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := aria.arith.GetMultiplication(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (aria Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := aria.arith.GetDivision(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
