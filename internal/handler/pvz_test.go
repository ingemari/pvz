package handler_test

// func TestHandleCreatePvz_Success(t *testing.T) {
// 	// Подготовка
// 	mockPvzService := new(mocks.MockPvzService)

// 	// Создаем логгер с текстовым обработчиком
// 	logger := logs.SetupLogger()
// 	slog.SetDefault(logger)

// 	handler := handler.NewPvzHandler(mockPvzService, logger)

// 	// Входные данные для запроса
// 	req := dto.PvzCreateRequest{
// 		City: "Москва",
// 	}

// 	// Генерация UUID и даты регистрации для теста
// 	pvzID := uuid.New()
// 	regDate := time.Now().Truncate(time.Second) // Округляем до секунд

// 	// Мокаем сервис
// 	expectedPvz := model.Pvz{
// 		Id:               pvzID,
// 		RegistrationDate: regDate,
// 		City:             req.City,
// 	}

// 	mockPvzService.On("CreatePvz", mock.Anything, mock.AnythingOfType("model.Pvz")).Return(expectedPvz, nil)

// 	// Запрос
// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request, _ = http.NewRequest(http.MethodPost, "/pvz", bytes.NewBuffer([]byte(`{
//         "city": "Москва"
//     }`)))

// 	// Обработка запроса
// 	handler.HandleCreatePvz(c)

// 	// Проверка
// 	assert.Equal(t, http.StatusCreated, w.Code)

// 	var resp dto.PvzCreateResponse
// 	err := json.Unmarshal(w.Body.Bytes(), &resp)
// 	assert.NoError(t, err)

// 	// Приводим ответное время к типу time.Time (форматируем строку времени в RFC3339)
// 	respTime, err := time.Parse(time.RFC3339, resp.RegDate)
// 	assert.NoError(t, err)

// 	// Сравниваем только по времени до секунд
// 	assert.True(t, respTime.Equal(expectedPvz.RegistrationDate), "Expected time: %v, got: %v", expectedPvz.RegistrationDate, respTime)

// 	// Сравниваем другие поля
// 	assert.Equal(t, expectedPvz.Id.String(), resp.Id) // Конвертируем UUID в строку
// 	assert.Equal(t, expectedPvz.City, resp.City)

// 	mockPvzService.AssertExpectations(t)
// }
