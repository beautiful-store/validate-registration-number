package validate_registration_number

import (
	"strconv"
	"strings"
)


func RegistrationNumber(nationalityType string, regNum string) bool {
	regNum = strings.Replace(regNum, "-", "", 1)

	var arrCheckNum = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3, 4, 5}

	sum := 0

	for i := 0; i < 12; i++ {
		r, _ := strconv.Atoi(string(regNum[i]))
		sum += r * arrCheckNum[i]
	}
	if len(regNum) == 13 && nationalityType == "NATIVE" {
		lastValue, _ := strconv.Atoi(string(regNum[12]))
		verificationCode := 11 - (sum % 11)

		if verificationCode >= 10 && verificationCode <= 11 {
			verificationCode = verificationCode - 10
		}
		if lastValue == verificationCode {
			return true
		} else {
			return false
		}
	} else if  len(regNum) == 13 && nationalityType == "FOREIGN" {

		lastValue, _ := strconv.Atoi(string(regNum[12]))
		verificationCode := 13 - (sum % 11)

		if verificationCode >= 10 && verificationCode <= 13 {
			verificationCode = verificationCode - 10
		}
		if lastValue == verificationCode{
			return true
		} else {
			return false
		}
	}else {
		return false
	}
}