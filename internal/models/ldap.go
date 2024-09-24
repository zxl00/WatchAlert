package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSON []byte

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return string(j), nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

func (j *JSON) UnmarshalToJSON(i interface{}) error {
	err := json.Unmarshal(*j, i)
	return err
}

type Ldap struct {
	ID        uint   `gorm:"primarykey" json:"id"` // 主键ID
	Address   string `json:"address"`
	DN        string `json:"dn"`
	AdminUser string `json:"admin_user"`
	Password  string `json:"password"`
	OU        string `json:"ou"`
	Filter    string `json:"filter"`
	Mapping   JSON   `gorm:"type:json" json:"mapping"`
	SSL       uint   `json:"ssl"`
	Status    uint   `json:"status"`
}
