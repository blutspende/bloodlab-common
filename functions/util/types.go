package util

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

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

func NullUUIDToUUIDPointer(value uuid.NullUUID) *uuid.UUID {
	if value.Valid {
		return &value.UUID
	}
	return nil
}

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
