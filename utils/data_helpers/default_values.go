package data_helpers

// Default values and constants used across the test framework

const (
	// Default Test Execution Variables
	SiteID      = "site1Id" // From system config
	DifID       = "DIF"     // From system config
	TcId        = "0000"    // 4-digit default value for testcase id variable
	VariantNN   = "00"      // 2-digit default value for variant variable
	TRUE_VALUE  = true      // Default true value for test cases
	FALSE_VALUE = false     // Default false value for test cases
)

// Default values for test data and constants
var (
	TestCode                  = "9999" // Default value for TestCode, to be replaced by actual calculation
	DefaultLongDescription    = "An example long description. This is too long to be a short description"
	DefaultImageUrl           = "http://www.example.com/stracciatella.png"
	DefaultProductClass       = "CONSUMABLE"
	DefaultProductHierarchyID = "TESTPH1"
	DefaultTemperatureClass   = "FROZEN"
	DefaultStorageArea        = "DMS"
	DefaultSellable           = false
	DefaultSecure             = true
	DefaultCatchWeight        = true
)

// Default values for reporting
var (
	TotalSteps       int = 0
	PassedSteps      int = 0
	FailedSteps      int = 0
	SkippedSteps     int = 0
	TotalScenarios   int = 0
	PassedScenarios  int = 0
	FailedScenarios  int = 0
	SkippedScenarios int = 0
)
