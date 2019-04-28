package models

import (
	"fmt"
	"database/sql/driver"
	"encoding/json"
)

type CustomerJson struct {
	JsonObject interface{}
	JsonBytes  []byte
}

func NewCustomerJson(jsonObject interface{}) (CustomerJson, error) {
	var err error
	customerJson := CustomerJson{
		JsonObject: jsonObject,
	}
	customerJson.JsonObject, err = json.Marshal(jsonObject)
	return customerJson, err
}

func (c CustomerJson) MarshalJSON() ([]byte, error) {
	if c.JsonBytes == nil && c.JsonObject == nil {
		return []byte("{}"), nil
	}

	if (c.JsonBytes == nil || len(c.JsonBytes) == 0) && c.JsonObject != nil {
		bytes, err := json.Marshal(c.JsonObject)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	}

	return c.JsonBytes, nil
}

func (c *CustomerJson) UnmarshalJSON(b []byte) error {

	c.JsonBytes = nil
	c.JsonBytes = append(c.JsonBytes, b...)
	err := json.Unmarshal(c.JsonBytes, &c.JsonObject)
	if err != nil {
		return err
	}

	return nil
}

func (c *CustomerJson) Scan(src interface{}) error {

	c.JsonBytes = nil
	switch d := src.(type) {
	case string:
		c.JsonBytes = []byte(d)
	case []byte:
		c.JsonBytes = append(c.JsonBytes, d...)
	default:
		return fmt.Errorf("know type: %#v", src)
	}
	if len(c.JsonBytes) == 0 {
		return nil
	}

	err := json.Unmarshal(c.JsonBytes, &c.JsonObject)
	if err != nil {
		return err
	}

	return nil
}

func (c CustomerJson) Value() (driver.Value, error) {
	return c.MarshalJSON()
}

func (c CustomerJson) GetJsonObject() interface{} {
	return c.JsonObject
}

func (c CustomerJson) GetJsonBytes() []byte {
	return c.JsonBytes
}

