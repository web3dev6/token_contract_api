package util

const (
	CREATE_TOKEN   = "CREATE_TOKEN"
	MINT_TOKEN     = "MINT_TOKEN"
	BURN_TOKEN     = "BURN_TOKEN"
	TRANSFER_TOKEN = "TRANSFER_TOKEN"
)

// IsSupportedTxContext returns true if tthe tx context is supported
func IsSupportedTxContext(context string) bool {
	switch context {
	case CREATE_TOKEN, MINT_TOKEN, BURN_TOKEN, TRANSFER_TOKEN:
		return true
	}
	return false
}
