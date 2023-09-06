package service

import (
	"sinarmas/models"
	"sinarmas/repo"
	"time"
)

type userService struct {
	repo repo.IUserRepo
}

func NewUserService(repo repo.IUserRepo) IUserService {
	return &userService{repo: repo}
}

//go:generate mockgen -source=./user_service.go -package=mocks -destination=./mocks/user_service.go
type IUserService interface {
	GenerateOtp(request *models.OtpGenerationRequest) (*models.OtpGenerationResponse, error)
	ValidateOtp(request *models.OtpValidationRequest) (interface{}, error)
}

func (u userService) GenerateOtp(request *models.OtpGenerationRequest) (*models.OtpGenerationResponse, error) {

	data, err := u.repo.GetByUserAndRequestId(request.UserId, request.RequestId)

	if data != nil {
		expiryTime := time.Unix(data.OtpGeneratedAtEpoch, 0).Add(2 * time.Minute).Unix()
		if time.Now().Unix() < expiryTime {
			return &models.OtpGenerationResponse{
				UserId: data.UserId,
				Otp:    data.Otp,
			}, nil
		}
	}

	otp := GenerateOtp()
	res := &models.User{
		UserId:              request.UserId,
		Otp:                 otp,
		RequestId:           request.RequestId,
		OtpGeneratedAtEpoch: time.Now().Unix(),
		IsValidated:         false,
	}
	err = u.repo.Create(res)
	if err != nil {
		return nil, err
	}
	return &models.OtpGenerationResponse{
		UserId: request.UserId,
		Otp:    otp,
	}, nil
}

func (u userService) ValidateOtp(request *models.OtpValidationRequest) (interface{}, error) {

	data, err := u.repo.GetByUserAndRequestId(request.UserId, request.RequestId)
	if err != nil && err.Error() == "record not found" {
		return &models.OtpValidationErrorResponse{
			Error:            "otp_not_found",
			ErrorDescription: "OTP not found",
		}, nil
	}

	if data.Otp != request.Otp {
		return &models.OtpValidationErrorResponse{
			Error:            "invalid_otp",
			ErrorDescription: "OTP is not valid",
		}, nil
	}

	expiryTime := time.Unix(data.OtpGeneratedAtEpoch, 0).Add(2 * time.Minute).Unix()
	if time.Now().Unix() < expiryTime {
		//validate otp in db
		res := &models.User{
			Id:                  data.Id,
			UserId:              data.UserId,
			Otp:                 data.Otp,
			RequestId:           data.RequestId,
			OtpGeneratedAtEpoch: data.OtpGeneratedAtEpoch,
			IsValidated:         true,
		}

		err := u.repo.Save(res)
		if err != nil {
			return nil, err
		}

		return &models.OtpValidationSuccessResponse{
			UserId:  request.UserId,
			Message: "Otp Validated Successfully",
		}, nil
	}

	return &models.OtpValidationErrorResponse{
		Error:            "otp_expired",
		ErrorDescription: "Otp has expired, please generate new otp",
	}, nil
}
