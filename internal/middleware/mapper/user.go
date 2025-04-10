package mapper

import (
	"pvz/internal/model"
)

func RegisterRequestToUser(req model.RegisterRequest, hashedPassword []byte) model.User {
	return model.User{
		Email:    req.Email,
		Password: string(hashedPassword), // Пароль уже должен быть захэширован!
		Role:     req.Role,
	}
}

func UserToRegisterResponse(user model.User) model.RegisterResponse {
	return model.RegisterResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}

// func UserToModel(req *CreateUserRequest) model.User {
// 	id := uuid.New()

// 	return model.User{
// 		Id:   id,
// 		User: req.user,
// 	}
// }
