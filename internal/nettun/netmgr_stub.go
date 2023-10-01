//go:build !darwin
// +build !darwin

package nettun

func provideSystemNetworkManager(_ zerolog.Logger) NetworkManager {
	// Stub for compile-time error
	return THIS_OPERATING_SYSTEM_IS_NOT_SUPPORTED
}
