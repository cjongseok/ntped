package ntped

import (
	"testing"
)

func TestSyncTimeout0(t *testing.T) {
	Sync(0, 0)
}
func TestSyncTimeout100(t *testing.T) {
	Sync(0, 100)
}
func TestSyncTimeout400(t *testing.T) {
	Sync(0, 400)
}
func TestSyncTimeout1000(t *testing.T) {
	Sync(0, 1000)
}
func TestSyncTimeout2000(t *testing.T) {
	Sync(0, 2000)
}
func TestSyncTimeout4000(t *testing.T) {
	Sync(0, 4000)
}

