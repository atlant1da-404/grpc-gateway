package dto

import "encoding/json"

type Note struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Count int32  `json:"count"`
}

func (n Note) Marshal() []byte {

	bNote, err := json.Marshal(n)
	if err != nil {
		return nil
	}

	return bNote
}

func (n *Note) Unmarshal(bNote []byte) {
	_ = json.Unmarshal(bNote, &n)
}
