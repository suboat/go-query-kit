package gosql

import "encoding/json"

func (l *LimitRoot) ParseJSONString(str string) error {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		return err
	}
	return l.Parse(m)
}

func (l *LimitRoot) Parse(m map[string]interface{}) error {
	if m == nil || len(m) == 0 {
		return nil
	}
	for k, v := range m {
		if v == nil {
		} else if IsLimitKey(k) {
			if l.Values == nil {
				l.Values = make([]ILimit, 0, len(m))
			}
			value := &LimitValue{Key: k}
			if err := value.Parse(v); err == nil {
				l.Values = append(l.Values, value)
			}
		}
	}
	return nil
}

func (l *LimitValue) Parse(obj interface{}) error {
	switch i := obj.(type) {
	case float64:
		l.Value = int(i)
	case float32:
		l.Value = int(i)
	case int64:
		l.Value = int(i)
	case int:
		l.Value = i
	}
	if !l.IsLimited() {
		return ErrTypeValue
	}
	return nil
}
