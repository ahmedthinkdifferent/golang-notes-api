package util

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"noteapp/core/env"
	"testing"
)

var token string

// We'll be able to store suite-wide
// variables and add methods to this
// test suite struct
type JwtTestSuite struct {
	suite.Suite
}

// This is an example test that will always succeed
func (suite *JwtTestSuite) TestExample() {
	suite.Equal(true, true)
}

// We need this function to kick off the test suite, otherwise
// "go test" won't know about our tests
func TestJwtTestSuite(t *testing.T) {
	suite.Run(t, new(JwtTestSuite))
}

// This will run right before the test starts
// and receives the suite and test names as input
func (suite *JwtTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Println("BeforeTest")
}

// This will run after test finishes
// and receives the suite and test names as input
func (suite *JwtTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("AfterTest")
}

// This will run before the tests in the suite are run
func (suite *JwtTestSuite) SetupSuite() {
	fmt.Println("SetupSuite")
	env.AppEnvironment = &env.AppEnv{App: env.App{Port: 3306, JwtToken: "token"}}
}

// This will run before each test in the suite
func (suite *JwtTestSuite) SetupTest() {
	fmt.Println("SetupTest")
}
func (suite *JwtTestSuite) TestGenerateJwtTokenReturnToken() {
	token = GenerateJwtToken(map[string]any{"user_id": 1})
	suite.NotEmpty(token)
}

func (suite *JwtTestSuite) TestVerifyJwtTokenReturnUserData() {
	claims, err := VerifyJwtToken(token)
	suite.Nil(err)
	suite.Equal(claims["user_id"], float64(1))
}
