package util

const (
	CREATE_TOKEN = "CREATE_TOKEN"
	GRANT_ROLE   = "GRANT_ROLE"
	REVOKE_ROLE  = "REVOKE_ROLE"
	MINT_TOKEN   = "MINT_TOKEN"
)

// IsSupportedTxContext returns true if tthe tx context is supported
func IsSupportedTxContext(context string) bool {
	switch context {
	case CREATE_TOKEN, GRANT_ROLE, REVOKE_ROLE, MINT_TOKEN:
		return true
	}
	return false
}
