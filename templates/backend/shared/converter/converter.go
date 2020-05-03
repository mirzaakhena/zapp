package utils

import (
	"encoding/json"
	"errors"
	"fmt"

	"bitbucket.org/kfc/mthexperimental/shared/constant"
)

// GetMSI is
func GetMSI(input map[string]interface{}, key string) map[string]interface{} {
	objIntf, ok := input[key]
	Check(ok, fmt.Errorf("key %s not found", key))

	objMapintf, ok := objIntf.(map[string]interface{})
	Check(ok, fmt.Errorf("key %s cannot casting to map string interface", key))
	return objMapintf
}

// GetARI is
func GetARI(input map[string]interface{}, key string) []interface{} {
	objIntf, ok := input[key]
	Check(ok, fmt.Errorf("key %s not found", key))

	if objIntf == nil {
		return []interface{}{}
	}

	objArrIntf, ok := objIntf.([]interface{})
	Check(ok, fmt.Errorf("key %s cannot casting to array of interface", key))
	return objArrIntf
}

// ConMSI is
func ConMSI(input interface{}) map[string]interface{} {
	obj, ok := input.(map[string]interface{})
	Check(ok, fmt.Errorf("cannot casting to map string interface"))
	return obj
}

// GetMSIFromJSONString is
func GetMSIFromJSONString(bytes []byte) map[string]interface{} {
	var obj interface{}
	json.Unmarshal(bytes, &obj)
	return ConMSI(obj)
}

// GetString is
func GetString(input map[string]interface{}, key string) string {
	str, ok := input[key].(string)
	Check(ok, errors.New(constant.FailCastingToString))
	return str
}

// GetInt is
func GetInt(input map[string]interface{}, key string) int {
	float, ok := input[key].(float64)
	Check(ok, errors.New(constant.FailCastingToFloat))
	return int(float)
}

// GetFloat is
func GetFloat(input map[string]interface{}, key string) float64 {
	float, ok := input[key].(float64)
	Check(ok, errors.New(constant.FailCastingToFloat))
	return float
}

// GetBool is
func GetBool(input map[string]interface{}, key string) bool {
	boolean, ok := input[key].(bool)
	Check(ok, errors.New(constant.FailCastingToBoolean))
	return boolean
}

// Check is
func Check(ok bool, e error) {
	if ok {
		return
	}
	if e != nil {
		panic(e)
	}
}
