package helper

import (
	"bnt/bnt-box-service/internal/core/domain/entity"
	"bnt/bnt-box-service/internal/core/domain/mapper"
	helper "bnt/bnt-box-service/internal/core/helper/error-helper"
	"fmt"
	"regexp"
	"strconv"


	"github.com/google/uuid"
)

func Generate(prefixStart string, prefixEnd string, prefixtype string, indent entity.Indent) (interface{}, error) {
	var boxes []entity.Box
	var _runtype = getRunType(indent.Denomination.Value)
	var prefix []string
	_prefixStart, err := validatePrefix(prefixStart, prefixtype, indent.Denomination.Value)
	if err != nil {
		return nil, err
	}
	_prefixEnd, err := validatePrefix(prefixEnd, prefixtype, indent.Denomination.Value)
	if err != nil {
		return nil, err
	}

	if _prefixEnd.(int) < _prefixStart.(int) {
		return nil, helper.ErrorMessage(helper.ValidationError, "Sorry, end prefix must be greater than start prefix.")
	}

	prefixStartIndices := prefixToArray(_prefixStart.(int))
	prefixEndIndices := prefixToArray(_prefixEnd.(int))

	count := len(prefixStartIndices)
	max := 90
	_prefix := ""
	switch count {
	case 1:
		a := prefixStartIndices[0]
		b := prefixEndIndices[0]
		for i := a; i <= b; i++ {
			if i == 73 || i == 79 {
				goto a
			}
			prefix = append(prefix, string(rune(i)))
		a:
		}
	case 2:
		a := prefixStartIndices[0]
		b := prefixEndIndices[0]
		c := prefixStartIndices[1]
		d := prefixEndIndices[1]
		for i := a; i <= b; i++ {
			if (c > 64 && c < 91) && (d > 64 && d < 91) {
				if i > a {
					c = 65
				}
				if i == b {
					max = d
				}
				if i == 73 || i == 79 {
					goto b
				}
				for j := c; j <= max; j++ {
					if j == 73 || j == 79 {
						goto c
					}
					_prefix = string(rune(i)) + string(rune(j))
					prefix = append(prefix, _prefix)
				c:
				}
			} else {
				for j := c; j <= d; j++ {

					_prefix = string(rune(i)) + fmt.Sprintf("%02d", j)
					prefix = append(prefix, _prefix)

				}
			}
		b:
		}
	case 3:
		a := prefixStartIndices[0]
		b := prefixEndIndices[0]
		c := prefixStartIndices[1]
		d := prefixEndIndices[1]
		e := prefixStartIndices[2]
		f := prefixEndIndices[2]
		for i := a; i <= b; i++ {
			if (c > 64 && c < 91) && (d > 64 && d < 91) {
				if i > a {
					c = 65
				}
				if i == b {
					max = d
				}
				if i == 73 || i == 79 {
					goto d
				}
				for j := c; j <= max; j++ {
					if j == 73 || j == 79 {
						goto e
					}
					
					for k := e; k <= f; k++ {
						_prefix = string(rune(i)) + string(rune(j))
						_prefix = _prefix + fmt.Sprintf("%02d", k)
						prefix = append(prefix, _prefix)

					}
				e:
				}
			} else {
				for j := c; j <= d; j++ {

					_prefix = string(rune(i)) + fmt.Sprintf("%02d", j)
					prefix = append(prefix, _prefix)

				}
			}
		d:
		}
	}

	for i := 0; i < len(prefix); i++ {
		_box := mapper.MapBoxConfigDto(prefix[i], indent)

		for j := 1; j <= _runtype; j++ {
			_box.Reference = uuid.New().String()
			var prefixFormat = regexp.MustCompile(`\d`).MatchString
			if !prefixFormat(prefix[i]){
				_box.BoxNo = prefix[i] + fmt.Sprintf("%03d", j)
			} else{
				last2Letters  := prefix[i][len(prefix[i])-2:]
				firstLetters:= prefix[i][0:len(prefix[i])-len(last2Letters)]
				suffix, _ := strconv.Atoi(last2Letters)
				boxno:=0
					
					if j<_runtype {
						suffix -=1
						boxno,_= strconv.Atoi(fmt.Sprint(suffix) + fmt.Sprintf("%02d", j))
						_box.BoxNo = firstLetters + fmt.Sprintf("%04d", boxno)
					}else{
						boxno = suffix * j
					}
					
				_box.BoxNo = firstLetters + fmt.Sprintf("%04d", boxno)
				
			}
			
			boxes = append(boxes, _box)
		}
	}

	return boxes, nil
}

func getRunType(denominationValue int) int {
	if denominationValue >= 100 {
		return 100
	}
	return 1000
}

func isHigher(denominationValue int) bool {
	return denominationValue >= 200

}

func validatePrefix(prefix string, prefixtype string, denomination int) (interface{}, error) {
	var isSingleHighDenomination = regexp.MustCompile(`^[A-Z]{1}[0-9]{2}$`).MatchString
	var isSingleLowDenomination = regexp.MustCompile(`^[A-Z]{1}$`).MatchString
	var isDoubleHighDenomination = regexp.MustCompile(`^[A-Z]{2}[0-9]{2}$`).MatchString
	var isDoubleLowDenomination = regexp.MustCompile(`^[A-Z]{2}$`).MatchString
	var ascii = getPrefixCode(prefix)
	switch prefixtype {
	case "Single":
		if isHigher(denomination) {
			if isSingleHighDenomination(prefix) {
				return ascii, nil
			}
		} else {
			if isSingleLowDenomination(prefix) {
				return ascii, nil
			}
		}
	case "Double":
		if isHigher(denomination) {
			if isDoubleHighDenomination(prefix) {
				return ascii, nil
			}
		} else {
			if isDoubleLowDenomination(prefix) {
				return ascii, nil
			}
		}
	}
	return nil, helper.ErrorMessage(helper.ValidationError, "The supplied prefix does not match the specified denomination value. Please stand adviced.")
}

func getPrefixCode(str string) int {
	_returnvalue := 0
	for _, char := range str {
		ascii := int(char)
		if ascii < 65 || ascii > 91 {

			_returnvalue = _returnvalue*10 + (int(char) - 48)
		} else {
			_returnvalue = _returnvalue*100 + int(char)
		}

	}
	return _returnvalue
}

func prefixToArray(prefix int) []int {
	array := []int{}

	for prefix != 0 {
		array = append(array, prefix%100)
		prefix /= 100
	}

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return array
}
