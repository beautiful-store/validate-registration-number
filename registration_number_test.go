package validate_registration_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateBusinessNumber_외국인Ok(t *testing.T) {
	//given
	nationalityType := "FOREIGN"
	regNum := "9901015020063"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_내국인Ok(t *testing.T) {
	//given
	nationalityType := "NATIVE"
	regNum := "테스트하는사람주민등록적기"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_국적이유효하지않는경우(t *testing.T) {
	//given
	nationalityType := "KOREA"
	regNum := "9901015020063"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
}

func TestValidateBusinessNumber_등록번호가13자리가아닌경우(t *testing.T) {
	//given
	nationalityType := "FOREIGN"
	regNum := "990101502001"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(regNum), 12)
}

func TestValidateBusinessNumber_외국인인경우유효하지않는경우(t *testing.T) {
	//given
	nationalityType := "FOREIGN"
	regNum := "9901012000000"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_내국인인경우유효하지않는경우(t *testing.T) {
	//given
	nationalityType := "NATIVE"
	regNum := "9901015000000"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_내국인_2020년10월이후_Ok(t *testing.T) {
	//given
	nationalityType := "NATIVE"
	regNum := "2010013000000"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_내국인_2020년10월이후_No(t *testing.T) {
	//given
	nationalityType := "FOREIGN"
	regNum := "2010012000000"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_외국인_2020년10월이후_Ok(t *testing.T) {
	//given
	nationalityType := "FOREIGN"
	regNum := "2010017000000"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, true, actual)
	assert.Equal(t, len(regNum), 13)
}

func TestValidateBusinessNumber_내국인_한글있는_경우(t *testing.T) {
	//given
	nationalityType := "NATIVE"
	regNum := "6110073tu021q"

	//when
	actual := RegistrationNumber(nationalityType, regNum)
	//then
	assert.Equal(t, false, actual)
	assert.Equal(t, len(regNum), 13)
}
