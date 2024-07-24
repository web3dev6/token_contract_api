package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/web3dev6/token_transaction/util"
)

var validTxContext validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if txContext, ok := fieldLevel.Field().Interface().(string); ok {
		// check if context is supported
		return util.IsSupportedTxContext(txContext)

	}
	return false
}
