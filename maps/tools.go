// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package maps

import "strconv"

// GetBoolean defined GetBoolean
func GetBoolean(mapData map[string]string, key string, defaultValue bool) bool {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseBool(data)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}

// GetFloat64 defined GetFloat64
func GetFloat64(mapData map[string]string, key string, defaultValue float64) float64 {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseFloat(data, 64)
		if err == nil {
			return parsed
		}
	}

	return defaultValue
}

// GetFloat32 defined GetFloat32
func GetFloat32(mapData map[string]string, key string, defaultValue float32) float32 {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseFloat(data, 32)
		if err == nil {
			return float32(parsed)
		}
	}

	return defaultValue
}

// GetString defined GetString
func GetString(mapData map[string]string, key string, defaultValue string) string {
	data, present := mapData[key]
	if present {
		return data
	}

	return defaultValue
}

// GetInt64 defined GetInt64
func GetInt64(mapData map[string]string, key string, defaultValue int64) int64 {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseInt(data, 10, 64)
		if err == nil {
			return parsed
		}
	}
	return defaultValue
}

// GetInt defined GetInt
func GetInt(mapData map[string]string, key string, defaultValue int) int {
	data, present := mapData[key]
	if present {
		parsed, err := strconv.ParseInt(data, 10, 32)
		if err == nil {
			return int(parsed)
		}
	}
	return defaultValue
}

// Contains defined Contains
func Contains(mapData map[string]string, key string) bool {
	_, present := mapData[key]
	return present
}
