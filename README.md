# validate-registration-number
주민등록번호 / 외국인등록번호 유효성 검사하는 모듈

# Installation
```
go get -u github.com/beautiful-store/validate-registration-number
```

# Description

* 국적타입과 등록번호 입력하면 외국인등록번호와 주민등록번호 유효성 검사(regNum 숫자 13자리 입력)
```
func RegistrationNumber(nationalityType string, regNum string) bool {
    return bool
}
```
nationalityType 종류 : NATIVE, FOREIGN


# 사용법
```
	validate_registration_number.RegistrationNumber(nationalityType,regNum)
```

# 버전
2022-02-09 v0.0.0
