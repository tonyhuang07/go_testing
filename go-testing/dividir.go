package testing

import "errors"

func dividir(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("el dominador no puede ser 0")
	}
	return a / b, nil
}
