package server

import "testing"

func TestServer_Run(t *testing.T) {
	s := New()
	s.SetPort(8080)
	s.Run()
}
