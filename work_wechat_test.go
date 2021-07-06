/*
@Time : 2021/7/6 11:24 上午
@Author : 21
@File : we_work_test
@Software: GoLand
*/
package work

import (
	"fmt"
	"testing"
)

func TestNewWeWork(t *testing.T) {
	testClass := NewWeWork(SetProviderCorpID("1"))
	fmt.Println(testClass)
}
