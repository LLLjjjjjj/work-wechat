package work

import (
	"fmt"
	"github.com/gogf/gf/os/gfile"
	"testing"
)

func TestNewImgUpload(t *testing.T) {
	token := "4FbJHIe9VF-z49EBjC8NUtrwYxFTiJQ9mZtGchePPY-Bp72LyJXLDDSMc3UXNnNLMbbY4XXALpoQUxNM6ZVTkn_W9oYyQWYZgpoHE_TQNIz_w_FJ1Nb1E89vPJlW76IMcsb2KLLZqJEEU9wppkp6Y5JKG3kQ55yFiLuQtsTf90sqYrzsDGjEawXhJ5qgb-tVumD6BDn6omDK8PiN8PSdrQ"
	fileStream := gfile.GetBytes("logo.png")
	result := NewImgUpload(token, fileStream)
	resultByte, err := result.GetRequestBody()
	fmt.Println(string(resultByte), err)
}

func TestNewMediaGet(t *testing.T) {

}
