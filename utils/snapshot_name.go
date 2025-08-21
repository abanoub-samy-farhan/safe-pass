package utils

import (
	"time"
)
func GetSnapshotName() string {
	t := time.Now()
	name := "safe-pass-" + t.Format("2006-01-02:15:04:05") + ".bin"
	return name
}
