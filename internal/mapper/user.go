package mapper

import (
	"log/slog"
	"pvz/internal/handler/dto"
	"pvz/internal/model"
	"pvz/internal/repository/entities"

	"github.com/google/uuid"
)

func RegisterRequestToUser(req dto.RegisterRequest, hashedPassword []byte) model.User {
	return model.User{
		Email:    req.Email,
		Password: string(hashedPassword), // Пароль уже должен быть захэширован!
		Role:     req.Role,
	}
}

func UserToEntity(user model.User) entities.User {
	var id uuid.UUID
	var err error

	if user.ID == "" {
		id = uuid.New()
	} else {
		id, err = uuid.Parse(user.ID)
		if err != nil {
			slog.Error("parse UUID", "err", err)
			id = uuid.New()
		}
	}

	return entities.User{
		ID:       id,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
	}
}

func EntityToUser(ent entities.User) model.User {
	return model.User{
		ID:       ent.ID.String(),
		Email:    ent.Email,
		Role:     ent.Role,
		Password: ent.Password,
	}
}

func UserToRegisterResponse(user model.User) dto.RegisterResponse {
	return dto.RegisterResponse{
		ID:    user.ID,
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
