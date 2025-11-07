package utils

import (
	"fmt"
	"strings"
)

const (
	CheckTheName = "check the name"
	ConnStrErr = "invalid connection string, load envs"
	ConnectionString = "CONNECTION_STRING"
	DbConnectionSuccess = "db connection succeeded"
	LoadEnvFromPath = ".env"
	Port            = ":3000"
	ServerStatus    = "Running"
	FilterByName = "name =?"
	FilterByDeleted = "deleted = ?"
	DevUpdatedSuccess = "developer updated successfully"
	DevDeletedSuccess = "developer deleted successfully"
	MigrationsSuccess = "migration executed successfully"
)

func LoadEnvErr(envFilePath string, err error) error{
	return	fmt.Errorf("cant load env from path: %s, err {%+v}", envFilePath, err)
}

func FailedTo(operation string, model string, name string) string{
	return fmt.Sprintf("failed to %s %s %s: check the name", operation, model, strings.ToUpper(name))
}