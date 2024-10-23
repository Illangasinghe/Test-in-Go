package report_helpers

import (
	"fmt"
	"os"
	"test-in-go/utils/data_helpers"
)

var reportFile *os.File

// InitPrettyReport initializes a file to store the report
func InitPrettyReport() error {
	var err error
	reportFile, err = os.Create("./reports/pretty-report.txt")
	if err != nil {
		return fmt.Errorf("failed to create report file: %v", err)
	}
	return nil
}

// PrettyLogStep logs step results in a human-readable format
func PrettyLogStep(stepName, status, details string) error {
	log := fmt.Sprintf("STEP: %s | STATUS: %s | DETAILS: %s\n", stepName, status, details)
	fmt.Print(log)
	_, err := reportFile.WriteString(log)
	if err != nil {
		return fmt.Errorf("failed to write step log: %v", err)
	}
	return nil
}

// PrettyLogScenario logs scenario results in a human-readable format
func PrettyLogScenario(scenarioName, status string) error {
	log := fmt.Sprintf("SCENARIO: %s | STATUS: %s\n", scenarioName, status)
	fmt.Print(log)
	_, err := reportFile.WriteString(log)
	if err != nil {
		return fmt.Errorf("failed to write scenario log: %v", err)
	}
	return nil
}

// FinalizePrettyReport finalizes the report with a summary by directly accessing global variables
func FinalizePrettyReport() error {
	summary := fmt.Sprintf(
		"\n--- FINAL REPORT ---\nTOTAL STEPS: %d | PASSED: %d | FAILED: %d | SKIPPED: %d\nTOTAL SCENARIOS: %d | PASSED: %d | FAILED: %d | SKIPPED: %d\n",
		data_helpers.TotalSteps, data_helpers.PassedSteps, data_helpers.FailedSteps, data_helpers.SkippedSteps,
		data_helpers.TotalScenarios, data_helpers.PassedScenarios, data_helpers.FailedScenarios, data_helpers.SkippedScenarios,
	)
	fmt.Print(summary)
	_, err := reportFile.WriteString(summary)
	if err != nil {
		return fmt.Errorf("failed to write final report summary: %v", err)
	}
	return ClosePrettyReport()
}

// ClosePrettyReport closes the report file
func ClosePrettyReport() error {
	if err := reportFile.Close(); err != nil {
		return fmt.Errorf("failed to close report file: %v", err)
	}
	return nil
}

// PassedStep increments passed steps count and logs it
func PassedStep() error {
	data_helpers.PassedSteps++
	data_helpers.TotalSteps++
	return nil
}

// FailedStep increments failed steps count and logs it
func FailedStep() error {
	data_helpers.FailedSteps++
	data_helpers.TotalSteps++
	return nil
}

// SkippedStep increments skipped steps count
func SkippedStep() error {
	data_helpers.SkippedSteps++
	data_helpers.TotalSteps++
	return nil
}

// PassedScenario increments passed scenarios count and logs it
func PassedScenario() error {
	data_helpers.PassedScenarios++
	return nil
}

// FailedScenario increments failed scenarios count and logs it
func FailedScenario() error {
	data_helpers.FailedScenarios++
	return nil
}

// SkippedScenario increments skipped scenarios count
func SkippedScenario() error {
	data_helpers.SkippedScenarios++
	return nil
}
