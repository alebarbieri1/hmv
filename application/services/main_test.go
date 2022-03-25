package services

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	suite.Run(t, new(AnalystServiceTestSuite))
	suite.Run(t, new(AuthenticationServiceTestSuite))
	suite.Run(t, new(EmergencyServiceTestSuite))
	suite.Run(t, new(PacientServiceTestSuite))
	suite.Run(t, new(RescuerServiceTestSuite))
	suite.Run(t, new(UserServiceTestSuite))
}
