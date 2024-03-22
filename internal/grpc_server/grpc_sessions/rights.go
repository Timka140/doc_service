package grpc_sessions

import (
	"fmt"

	"github.com/lib/pq"
)

const (
	CRenderDocx = 1 // Генерация docx
	CRenderPDF  = 2 // Конвертация в PDF
)

type TRights struct {
	Rights []int `json:"lvl"`
}

func NewRights() *TRights {
	t := &TRights{}

	return t
}

func (t *TRights) Set(rights map[string]interface{}) (*TRights, error) {
	if rights == nil {
		return nil, fmt.Errorf("SetRights(): список пуст")
	}

	out := &TRights{}

	administrator, ok := rights["docx"].(bool)
	if ok {
		out.update(administrator, CRenderDocx)
	}

	guest, ok := rights["pdf"].(bool)
	if ok {
		out.update(guest, CRenderPDF)
	}

	return out, nil
}
func (t *TRights) SetDB(val pq.Int64Array) {
	for _, v := range val {
		t.Rights = append(t.Rights, int(v))
	}
}

func (t *TRights) Vue() map[string]interface{} {
	list := make(map[string]interface{})

	if t.contains(CRenderDocx) {
		list["docx"] = true
	}

	if t.contains(CRenderPDF) {
		list["pdf"] = true
	}

	return list
}

func (t *TRights) Get() pq.Int64Array {
	var val pq.Int64Array

	for _, v := range t.Rights {
		val = append(val, int64(v))
	}

	if val == nil {
		val = pq.Int64Array{0}
	}

	return val
}

func (t *TRights) contains(val int) bool {
	res := false
	for _, v := range t.Rights {
		if v == val {
			res = true
		}
	}
	return res
}

// val - если true добавляет значение, false удаляет из списка
func (t *TRights) update(val bool, key int) {
	presence := false
	pos := 0

	for i, k := range t.Rights {
		if k == key {
			pos = i
			presence = true
			break
		}
	}

	//Удаляю права
	if !val && presence {
		t.Rights = append(t.Rights[:pos], t.Rights[pos+1:]...)
	}

	//Добавляю права
	if val && !presence {
		t.Rights = append(t.Rights, key)
	}

}

// Check проверка наличия прав
func (t *TRights) Check(rights []int) bool {
	val := false
	for _, r := range rights {
		for _, v := range t.Rights {
			if r == v {
				val = true
				break
			}
		}
		if val {
			break
		}
	}
	return val
}
