// Package providers provides all built-in implementations for cran.
//
// Implementations are automatically registered when importing this module.
package providers

import "cran/domain"

func init() {
	domain.Register(&assemblee_nationale{})
}
