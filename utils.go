package bt

import (
	"bytes"
	"math/rand"
)

const (
	predefineValue = 8
)

type attr struct {
	k string
	v string
}

type attributes struct {
	attrSlc []attr
	attrMap map[string]string
	isInMap bool
}

func (ab *attributes) copySlcToMap() {
	ab.attrMap = make(map[string]string)
	for _, val := range ab.attrSlc {
		ab.attrMap[val.k] = val.v
	}
	ab.isInMap = true
	ab.attrSlc = nil
}

func (ab *attributes) add(k, v string) {
	if !ab.isInMap {
		if len(ab.attrSlc) < predefineValue {
			ab.attrSlc = append(ab.attrSlc, attr{k: k, v: v})
			return
		}
		ab.copySlcToMap()
	}
	ab.attrMap[k] = v
}

func (ab *attributes) findValueByKey(key string) (string, bool) {
	if ab.isInMap {
		v, ok := ab.attrMap[key]
		return v, ok
	}

	for _, val := range ab.attrSlc {
		if val.k == key {
			return val.v, true
		}
	}
	return "", false
}

func (ab *attributes) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("[")
	if ab.isInMap {
		i := 0
		for k, v := range ab.attrMap {
			buf.WriteString(k + ":" + v)
			if i < len(ab.attrMap)-1 {
				buf.WriteString(",")
			}
			i++
		}
	} else {
		for i, val := range ab.attrSlc {
			buf.WriteString(val.k + ":" + val.v)
			if i < len(ab.attrSlc)-1 {
				buf.WriteString(",")
			}
		}
	}
	buf.WriteString("]")
	return buf.String()
}

// 乱序一个整数数组
func reArrangeIntSlice(intSlc []int) {
	for i := 1; i < len(intSlc); i++ {
		r := rand.Intn(i)
		intSlc[i], intSlc[r] = intSlc[r], intSlc[i]
	}
}
