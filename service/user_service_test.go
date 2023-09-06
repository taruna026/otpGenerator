package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"sinarmas/models"
	"sinarmas/repo/mocks"
	"testing"
	"time"
)

type UserServiceTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	service IUserService
	repo    *mocks.MockIUserRepo
}

func (suite *UserServiceTestSuite) BeforeTest(suiteName, testName string) {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockIUserRepo(suite.ctrl)
	suite.service = NewUserService(suite.repo)
}

func (suite *UserServiceTestSuite) SetupTest(suiteName, testName string) {
}

func (suite *UserServiceTestSuite) AfterTest(suiteName, testName string) {
	suite.ctrl.Finish()
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

func (suite *UserServiceTestSuite) TestGenerateOtp_OtpAlreadyGenerated() {

	otpDetails := &models.User{
		Id:                  1,
		UserId:              "taruna",
		Otp:                 "12345",
		RequestId:           "123",
		OtpGeneratedAtEpoch: time.Now().Unix(),
		IsValidated:         false,
	}

	request := &models.OtpGenerationRequest{
		UserId:    "taruna",
		RequestId: "123",
	}

	suite.repo.EXPECT().GetByUserAndRequestId(otpDetails.UserId, otpDetails.RequestId).Return(otpDetails, nil)
	expectedResponse := &models.OtpGenerationResponse{
		UserId: otpDetails.UserId,
		Otp:    otpDetails.Otp,
	}
	res, err := suite.service.GenerateOtp(request)
	suite.Equal(expectedResponse, res)
	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestValidateOtp_InvalidOtp() {

	otpDetails := &models.User{
		Id:                  1,
		UserId:              "taruna",
		Otp:                 "12345",
		RequestId:           "123",
		OtpGeneratedAtEpoch: time.Now().Unix(),
		IsValidated:         false,
	}

	request := &models.OtpValidationRequest{
		UserId:    "taruna",
		RequestId: "123",
		Otp:       "1234",
	}

	expectedResponse := &models.OtpValidationErrorResponse{
		Error:            "invalid_otp",
		ErrorDescription: "OTP is not valid",
	}

	suite.repo.EXPECT().GetByUserAndRequestId(request.UserId, request.RequestId).Return(otpDetails, nil)
	res, err := suite.service.ValidateOtp(request)

	suite.Equal(expectedResponse, res)
	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestValidateOtp_InvalidAlreadyValidated() {

	request := &models.OtpValidationRequest{
		UserId:    "taruna",
		RequestId: "123",
		Otp:       "1234",
	}

	expectedResponse := &models.OtpValidationErrorResponse{
		Error:            "otp_not_found",
		ErrorDescription: "OTP not found",
	}

	suite.repo.EXPECT().GetByUserAndRequestId(request.UserId, request.RequestId).Return(nil, errors.New("record not found"))
	res, err := suite.service.ValidateOtp(request)

	suite.Equal(expectedResponse, res)
	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestValidateOtp_Success() {
	otpDetails := &models.User{
		Id:                  1,
		UserId:              "taruna",
		Otp:                 "12345",
		RequestId:           "123",
		OtpGeneratedAtEpoch: time.Now().Unix(),
		IsValidated:         false,
	}

	request := &models.OtpValidationRequest{
		UserId:    "taruna",
		RequestId: "123",
		Otp:       "12345",
	}

	updatedOtpDetails := &models.User{
		Id:                  1,
		UserId:              "taruna",
		Otp:                 "12345",
		RequestId:           "123",
		OtpGeneratedAtEpoch: time.Now().Unix(),
		IsValidated:         true,
	}

	expectedRes := &models.OtpValidationSuccessResponse{
		UserId:  request.UserId,
		Message: "Otp Validated Successfully",
	}
	suite.repo.EXPECT().GetByUserAndRequestId(request.UserId, request.RequestId).Return(otpDetails, nil)
	suite.repo.EXPECT().Save(updatedOtpDetails).Return(nil)
	res, err := suite.service.ValidateOtp(request)

	suite.Equal(expectedRes, res)
	suite.NoError(err)
}

func (suite *UserServiceTestSuite) TestValidateOtp_OtpExpired() {
	otpDetails := &models.User{
		Id:                  1,
		UserId:              "taruna",
		Otp:                 "12345",
		RequestId:           "123",
		OtpGeneratedAtEpoch: 0,
		IsValidated:         false,
	}

	request := &models.OtpValidationRequest{
		UserId:    "taruna",
		RequestId: "123",
		Otp:       "12345",
	}

	expectedRes := &models.OtpValidationErrorResponse{
		Error:            "otp_expired",
		ErrorDescription: "Otp has expired, please generate new otp",
	}
	suite.repo.EXPECT().GetByUserAndRequestId(request.UserId, request.RequestId).Return(otpDetails, nil)
	res, err := suite.service.ValidateOtp(request)

	suite.Equal(expectedRes, res)
	suite.NoError(err)
}
