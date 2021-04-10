package utils

import (
	"fmt"
	"testing"
)

func TestGetIPByDomain(t *testing.T) {
	fmt.Println(GetIPByDomain("www.baidu.com"))
}
