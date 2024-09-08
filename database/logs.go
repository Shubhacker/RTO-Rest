package database

import (
	"log"
)

// For now added all logs in DB later will use proper logger
func Logs(logType, message string) error {
	var inputArgs []interface{}
	var reportId string
	query := `insert into public.logs (message, "Type") VALUES ($1, $2);`
	inputArgs = append(inputArgs, logType)
	inputArgs = append(inputArgs, message)

	err2 := DB.QueryRow(query, inputArgs...).Scan(&reportId)
	if err2 != nil {
		log.Println("Error from DB", err2.Error())
		return err
	}
	return nil
}
