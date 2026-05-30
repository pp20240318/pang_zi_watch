package handler

import (
	"encoding/json"
)

func parseJSONArray(raw string) []string {
	if raw == "" {
		return []string{}
	}
	var arr []string
	if err := json.Unmarshal([]byte(raw), &arr); err != nil {
		return []string{}
	}
	return arr
}

func mustJSONArray(items []string) string {
	if items == nil {
		items = []string{}
	}
	b, _ := json.Marshal(items)
	return string(b)
}

func parseJSONObjects(raw string) []map[string]interface{} {
	if raw == "" {
		return []map[string]interface{}{}
	}
	var arr []map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &arr); err != nil {
		return []map[string]interface{}{}
	}
	return arr
}

func mustJSONObjects(items []map[string]interface{}) string {
	if items == nil {
		items = []map[string]interface{}{}
	}
	b, _ := json.Marshal(items)
	return string(b)
}
