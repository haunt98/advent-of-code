package debug

import "os"

func IsDebug() bool {
	return os.Getenv("DEBUG") != ""
}
