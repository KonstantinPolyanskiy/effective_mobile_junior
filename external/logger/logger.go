package logger

import "go.uber.org/zap"

func New(level string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}

	// Создаем конфиг и устанавливаем ему уровень логгирования
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl

	// Собираем логгер на основе конфига и возвращаем
	return cfg.Build()
}
