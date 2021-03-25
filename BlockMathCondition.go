package netroutine

import (
	"encoding/json"
)

const idBlockMathCondition = "BlockMathCondition"

type BlockMathCondition struct {
	FromKey    string
	CompareTo  float64
	Operation  string
	StatusPass string
	StatusFail string
}

func (b *BlockMathCondition) toBytes() ([]byte, error) {
	return json.Marshal(b)
}

func (b *BlockMathCondition) fromBytes(bytes []byte) error {
	return json.Unmarshal(bytes, b)
}

func (b *BlockMathCondition) kind() string {
	return idBlockMathCondition
}

func (b *BlockMathCondition) Run(wce *Environment) (string, error) {
	var statusPass error
	var statusFail error
	var err error
	var metCondition bool

	statusPass, err = statusFromString(b.StatusPass)
	if err != nil {
		return log(b, "invalid status string", Error)
	}

	statusFail, err = statusFromString(b.StatusFail)
	if err != nil {
		return log(b, "invalid status string", Error)
	}

	s1, ok := wce.getData(b.FromKey)
	if !ok {
		return log(b, "couldn't find the first source variable", Error)
	}

	s1f, err := toFloat64(s1)
	if err != nil {
		return log(b, "first source variable wasn't a float", Error)
	}

	switch b.Operation {
	case "==":
		if s1f == b.CompareTo {
			metCondition = true
		} else {
			metCondition = false
		}
	case "!=":
		if s1f != b.CompareTo {
			metCondition = true
		} else {
			metCondition = false
		}
	case "<":
		if s1f < b.CompareTo {
			metCondition = true
		} else {
			metCondition = false
		}
	case "<=":
		if s1f <= b.CompareTo {
			metCondition = true
		} else {
			metCondition = false
		}
	case ">":
		if s1f > b.CompareTo {
			metCondition = true
		} else {
			metCondition = false
		}
	case ">=":
		if s1f >= b.CompareTo {
			metCondition = true
		} else {
			metCondition = false
		}
	default:
		return log(b, "invalid operation", Error)
	}

	if !metCondition {
		return log(b, "failed to meet condition", statusFail)
	}

	return log(b, "met condition", statusPass)
}
