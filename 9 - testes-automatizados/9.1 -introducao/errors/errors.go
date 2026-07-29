package errors

import "errors"

var (
	ErroListaVazia     = errors.New("a lista não pode estar vazia")
	ErroNumeroNegativo = errors.New("números negativos não permitidos")
)
