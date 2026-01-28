package utils

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// String
func StringToPointer(value string) *string {
	return &value
}
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
func NullStringToString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}
func NullStringToStringPointer(value sql.NullString) *string {
	if value.Valid {
		return &value.String
	}
	return nil
}

// UUID
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

// Time
func NullTimeToTimePointer(value sql.NullTime) *time.Time {
	if value.Valid {
		return &value.Time
	}
	return nil
}
func TimePointerToNullTime(value *time.Time) sql.NullTime {
	if value != nil {
		return sql.NullTime{
			Time:  *value,
			Valid: true,
		}
	}
	return sql.NullTime{}
}

// Int
func IntToPointer(value int) *int {
	return &value
}
