package env

import "syscall"

func GetEnv(name string) (string, bool) {
	return syscall.Getenv(name)
}

func GetDefault(name string, d string) string {
	if v, has := GetEnv(name); has {
		return v
	}
	return d
}
