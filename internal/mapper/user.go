package mapper

import (
	"pvz/internal/handler/dto"
	"pvz/internal/model"
)

func RegisterRequestToUser(req dto.RegisterRequest, hashedPassword []byte) model.User {
	return model.User{
		Email:    req.Email,
		Password: string(hashedPassword), // Пароль уже должен быть захэширован!
		Role:     req.Role,
	}
}

func UserToRegisterResponse(user model.User) dto.RegisterResponse {
	return dto.RegisterResponse{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
	}
}

func LoginRequestToUser(req dto.LoginRequest, hpass []byte) model.User {
	return model.User{
		Email:    req.Email,
		Password: string(hpass),
	}
}
