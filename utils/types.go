package utils

import (
	"github.com/google/uuid"
)

func StringToPointerWithNil(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func StringPointerToString(value *string) string {
	if value != nil {
		return *value
	}
	return ""
}

func StringPointerToStringWithDefault(value *string, defaultValue string) string {
	if value != nil {
		return *value
	}
	return defaultValue
}

func UUIDToNullUUID(value uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  value,
		Valid: value != uuid.Nil,
	}
}

func NullUUIDToUUIDPointer(value uuid.NullUUID) *uuid.UUID {
	if value.Valid {
		return &value.UUID
	}
	return nil
}
