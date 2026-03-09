package utils

import "fmt"

func GetPort(port string) string {
	return fmt.Sprintf(":%s", port)
}
