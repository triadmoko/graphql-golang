package helpers

import "github.com/google/uuid"

func UUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func IsValidUUID(u string) bool {
    _, err := uuid.Parse(u)
    return err == nil
 }
