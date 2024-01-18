package logger

import "go.uber.org/zap"

func New(level string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	// Если ошибка в получении уровня, возвращаем стандартный продакшен логгер
	if err != nil {
		return zap.NewProduction()
	}

	// Создаем конфиг и устанавливаем ему уровень логгирования
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	// Собираем логгер на основе конфига и возвращаем
	return cfg.Build()
}
