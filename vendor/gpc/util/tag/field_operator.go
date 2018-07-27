package tag

import (
	"errors"
	"strings"

	plib "gpc/publib"
)

//-----------------------------------------------------------------------------------------------------------//

// 判断 操作符
func (self *FieldCheck) Operator(v interface{}, exp string) (ret *RetStatus) {
	ret = NewRetStatus()
	op, val, err := self.operatorSplit(exp)
	if err != nil {
		return ret.SetErr(err.Error())
	}
	if num, ok := v.(int64); ok {
		if strings.HasPrefix(op, "in") {
			return self.operatorInInt(num, op, val)
		} else {
			if n, err := plib.StrToInt64(val); err != nil {
				return ret.SetErr(err.Error())
			} else {
				return self.operatorInt(num, op, n)
			}
		}
	} else if str, ok := v.(string); ok {
		if strings.HasPrefix(op, "in") {
			return self.operatorInString(str, op, val)
		} else {
			return self.operatorString(str, op, val)
		}
	} else {
		return ret.SetErr("operator only support int64 and string type.")
	}
}

// 分割 操作符
func (self *FieldCheck) operatorSplit(opVal string) (op, val string, err error) {
	if strings.HasPrefix(opVal, "len") {
		opVal = strings.TrimPrefix(opVal, "len")
	}
	if len(opVal) >= 2 {
		isMatch := false
		if len(opVal) >= 3 {
			op = opVal[:2]
			switch op {
			case ">=", "<=", "==", "!=":
				val = opVal[2:]
				isMatch = true
			default:
				if strings.HasPrefix(opVal, "in=") {
					op = opVal[:3]
					val = opVal[3:]
					isMatch = true
				} else if strings.HasPrefix(opVal, "in!=") {
					op = opVal[:4]
					val = opVal[4:]
					isMatch = true
				}
			}
		}
		if !isMatch {
			op = opVal[:1]
			switch op {
			case ">", "<", "=":
				val = opVal[1:]
				isMatch = true
			}
		}
		if isMatch && val != "" {
			return
		}
	}
	return "", "", errors.New("invalid:" + opVal)
}

func (self *FieldCheck) operatorInt(v1 int64, op string, v2 int64) (ret *RetStatus) {
	ret = NewRetStatus()
	exp := plib.Int64ToStr(v1) + op + plib.Int64ToStr(v2)
	ok := true
	switch op {
	case ">":
		ok = (v1 > v2)
	case "<":
		ok = (v1 < v2)
	case "=", "==":
		ok = (v1 == v2)
	case ">=":
		ok = (v1 >= v2)
	case "<=":
		ok = (v1 <= v2)
	case "!=":
		ok = (v1 != v2)
	default:
		return ret.SetErr("invalid:" + exp)
	}
	if !ok {
		return ret.SetFail("fail:" + exp)
	}
	return
}

func (self *FieldCheck) operatorString(v1, op, v2 string) (ret *RetStatus) {
	ret = NewRetStatus()
	exp := v1 + " " + op + v2
	ok := true
	switch op {
	case ">":
		ok = (v1 > v2)
	case "<":
		ok = (v1 < v2)
	case "=", "==":
		ok = (v1 == v2)
	case ">=":
		ok = (v1 >= v2)
	case "<=":
		ok = (v1 <= v2)
	case "!=":
		ok = (v1 != v2)
	default:
		return ret.SetErr("invalid:" + exp)
	}
	if !ok {
		return ret.SetFail("fail:" + exp)
	}
	return
}

func (self *FieldCheck) operatorInInt(v1 int64, op, v2 string) (ret *RetStatus) {
	ret = NewRetStatus()
	exp := plib.Int64ToStr(v1) + " " + op + v2
	strArr := strings.Split(v2, ",")
	plib.StrArrayDeleteAllSpaceTrim(&strArr)
	var int64Arr plib.Int64Array
	for _, s := range strArr {
		if i, err := plib.StrToInt64(s); err != nil {
			return ret.SetErr(err.Error())
		} else {
			int64Arr.Add(i)
		}
	}
	if int64Arr.Len() == 0 {
		goto Invalid
	}
	switch op {
	case "in=":
		if int64Arr.Contains(v1) {
			return
		}
		goto Fail
	case "in!=":
		if int64Arr.Contains(v1) {
			goto Fail
		}
		return
	}
Invalid:
	return ret.SetErr("invalid:" + exp)
Fail:
	return ret.SetFail("fail:" + exp)
}

func (self *FieldCheck) operatorInString(v1, op, v2 string) (ret *RetStatus) {
	ret = NewRetStatus()
	exp := v1 + " " + op + v2
	strArr := strings.Split(v2, ",")
	plib.StrArrayDeleteAllSpaceTrim(&strArr)
	if len(strArr) == 0 {
		goto Invalid
	}
	switch op {
	case "in=":
		if plib.StrArrayContains(&strArr, v1) {
			return
		}
		goto Fail
	case "in!=":
		if plib.StrArrayContains(&strArr, v1) {
			goto Fail
		}
		return
	}
Invalid:
	return ret.SetErr("invalid:" + exp)
Fail:
	return ret.SetFail("fail:" + exp)
}

//-----------------------------------------------------------------------------------------------------------//
