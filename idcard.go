package idcard

import (
	"fmt"
	"strconv"
	"strings"
)

// Upgrade15To18 15位身份证升位为18位
// 入参必须是15位字符串，否者提示错误
func Upgrade15To18(oldID string) (string, error) {
	newID := strings.TrimSpace(oldID)

	if len(newID) != 15 {
		return "", fmt.Errorf("len(%s) is not 15", newID)
	}

	// 拆成字符数组["1", ....]
	newIDArr := strings.Split(newID, "")

	// step1: 15转17，插入年份“19”
	newIDArr = append(newIDArr[:6], append([]string{"1", "9"}, newIDArr[6:]...)...)

	// step2: 计算校验码
	newIDVer := calcVerificationCode(newIDArr)

	// step3: 追加校验码
	newIDArr = append(newIDArr[:17], newIDVer[:1])

	// step4: 重新装配成单个字符串
	newID = strings.Join(newIDArr, "")

	return newID, nil
}

// 输入17位身份证，计算校验位数值
func calcVerificationCode(idArr []string) string {
	sum := 0
	multi := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checkCodes := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

	for idx, idStr := range idArr {
		id, _ := strconv.Atoi(idStr)
		sum += id * multi[idx]
	}

	idx := sum % 11
	return checkCodes[idx]
}
