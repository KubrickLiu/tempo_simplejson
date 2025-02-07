package simplejson

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

type Json struct {
	data interface{}
	err  error
}

func New() *Json {
	return &Json{
		data: make(map[string]interface{}),
	}
}

func NewJson(b []byte) (*Json, error) {
	r := bytes.NewBuffer(b)
	return NewJsonFromReader(r)
}

func NewJsonFromReader(r io.Reader) (*Json, error) {
	j := new(Json)
	dec := json.NewDecoder(r)
	dec.UseNumber()
	if err := dec.Decode(&j.data); err != nil {
		return nil, err
	}
	return j, nil
}

func (j *Json) Data() interface{} {
	return j.data
}

func (j *Json) MarshalJSON() ([]byte, error) {
	return json.Marshal(&j.data)
}

func (j *Json) IsValid() bool {
	return j.err == nil
}

func (j *Json) Set(key string, val interface{}) error {
	m, err := j.Map()
	if err != nil {
		return err
	}
	m[key] = val
	return nil
}

func (j *Json) SetPath(branch []string, val interface{}) error {
	length := len(branch)
	if length == 0 {
		return nil
	}

	cur, err := j.Map()
	if err != nil {
		return err
	}

	var content string
	var tmp interface{}
	var ok bool
	for i := 0; i < length-1; i++ {
		content = branch[i]
		if tmp, ok = cur[content]; !ok {
			nCur := make(map[string]interface{})
			cur[content] = nCur
			cur = nCur
			continue
		}

		if _, ok = tmp.(map[string]interface{}); !ok {
			return errors.New("type assertion to map[string]interface{} failed")
		}

		cur = tmp.(map[string]interface{})
	}

	cur[branch[len(branch)-1]] = val
	return nil
}

func (j *Json) Del(key string) error {
	m, err := j.Map()
	if err != nil {
		return err
	}

	delete(m, key)
	return nil
}

func (j *Json) Get(key string) *Json {
	m, err := j.Map()
	if err == nil {
		if v, ok := m[key]; ok {
			return &Json{
				data: v,
			}
		}
	}

	return &Json{}
}

func (j *Json) GetPath(branch ...string) *Json {
	ret := j
	for _, content := range branch {
		ret = ret.Get(content)
	}
	return ret
}

func (j *Json) GetIndex(index int) *Json {
	arr, err := j.Array()
	if err != nil {
		return &Json{
			err: errors.New("invalid data"),
		}
	}

	val := arr[index]
	ret := &Json{
		data: val,
	}
	return ret
}

func (j *Json) Map() (map[string]interface{}, error) {
	if m, ok := j.Data().(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[string]interface{} failed")
}

func (j *Json) Array() ([]interface{}, error) {
	if arr, ok := j.Data().([]interface{}); ok {
		return arr, nil
	}
	return nil, errors.New("type assertion to []interface{} failed")
}

func (j *Json) Int() (int, error) {
	if val, ok := j.Data().(int); ok {
		return val, nil
	}
	return 0, errors.New("type assertion to int failed")
}

func (j *Json) Float64() (float64, error) {
	if val, ok := j.Data().(float64); ok {
		return val, nil
	}
	return 0, errors.New("type assertion to float64 failed")
}

func (j *Json) Bool() (bool, error) {
	if val, ok := j.Data().(bool); ok {
		return val, nil
	}
	return false, errors.New("type assertion to bool failed")
}

func (j *Json) String() (string, error) {
	if val, ok := j.Data().(string); ok {
		return val, nil
	}
	return "", errors.New("type assertion to string failed")
}

func (j *Json) Bytes() ([]byte, error) {
	if val, ok := j.Data().([]byte); ok {
		return val, nil
	}
	return nil, errors.New("type assertion to []byte failed")
}
