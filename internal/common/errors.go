package common

import "errors"

var (
	ErrNotFound         = errors.New("recurso não encontrado")
	ErrConflict         = errors.New("conflito de dados")
	ErrBadRequest       = errors.New("requisição inválida")
	ErrInternalServer   = errors.New("erro interno do servidor")
	ErrCannotDelete     = errors.New("não é possível deletar, existem itens vinculados")
)