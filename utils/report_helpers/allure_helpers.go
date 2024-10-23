package report_helpers

import (
	"fmt"
	"test-in-go/utils/db_helpers"

	"github.com/ozontech/allure-go/pkg/allure"
)

// RunAllureStep handles the repetitive logic for Allure steps.
func RunAllureStep(stepName string, stepFunc func() error) error {
	step := allure.NewSimpleStep(stepName).Begin()
	err := stepFunc()
	if err != nil {
		step.Failed().Finish()
		return err
	}
	step.Passed().Finish()
	return nil
}

// QueryDatabase abstracts the logic of querying the database.
func QueryDatabase(query, id string) (string, error) {
	var dbValue string
	err := db_helpers.DB.QueryRow(query, id).Scan(&dbValue)
	if err != nil {
		return "", fmt.Errorf("database query failed: %v", err)
	}
	return dbValue, nil
}
