package spentcalories

import (
    "fmt"
    "math"
    "strconv"
    "strings"
    "time"
)

// Основные константы, необходимые для расчетов.
const (
    lenStep = 0.65  // средняя длина шага.
    mInKm   = 1000  // количество метров в километре.
    minInH  = 60    // количество минут в часе.
    cmInM   = 100   // количество сантиметров в метре.
)

// parseTraining парсит строку с данными тренировки и возвращает количество шагов, тип активности и продолжительность.
func parseTraining(data string) (int, string, time.Duration, error) {
    parts := strings.Split(data, ",")
    if len(parts) != 3 {
        return 0, "", 0, fmt.Errorf("invalid data format")
    }

    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, "", 0, fmt.Errorf("invalid steps format: %w", err)
    }

    activityType := strings.TrimSpace(parts[1]) // Удаляем лишние пробелы
    duration, err := time.ParseDuration(strings.TrimSpace(parts[2])) // Удаляем лишние пробелы
    if err != nil {
        return 0, "", 0, fmt.Errorf("invalid duration format: %w", err)
    }

    return steps, activityType, duration, nil
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
func distance(steps int) float64 {
    return float64(steps) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
    if duration == 0 {
        return 0
    }
    distance := distance(steps)
    return distance / duration.Hours()
}

// TrainingInfo возвращает строку с информацией о тренировке.
func TrainingInfo(data string, weight, height float64) (string, error) {
    steps, trainingType, duration, err := parseTraining(data)
    if err != nil {
        return "", err
    }
    distance := distance(steps)
    speed := meanSpeed(steps, duration)
    var calories float64

    switch trainingType {
    case "Бег":
        calories = RunningSpentCalories(steps, weight, duration)
        return fmt.Sprintf("Тип тренировки: %s\n Длительность: %.2f ч.\n Дистанция: %.2f км.\n Скорость: %.2f км/ч\n Сожгли калорий: %.2f\n", trainingType, duration.Hours(), distance, speed, calories), nil
    case "Ходьба":
        calories = WalkingSpentCalories(steps, duration, weight, height)
        return fmt.Sprintf("Тип тренировки: %s\n Длительность: %.2f ч.\n Дистанция: %.2f км.\n Скорость: %.2f км/ч\n Сожгли калорий: %.2f\n", trainingType, duration.Hours(), distance, speed, calories), nil
    default:
        return "", fmt.Errorf("unknown training type: %s", trainingType)
    }
}

// Константы для расчета калорий, расходуемых при беге.
const (
    runningCaloriesMeanSpeedMultiplier = 18.0
    runningCaloriesMeanSpeedShift      = 20.0
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
    speed := meanSpeed(steps, duration)
    calories := (runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift * weight / mInKm * duration.Hours())
    return calories
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
    walkingCaloriesWeightMultiplier = 0.035
    walkingSpeedHeightMultiplier    = 0.029
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(steps int, duration time.Duration, weight, height float64) float64 {
    speed := meanSpeed(steps, duration)
    SpeedInMs := speed * 1000 / 3600 // Перевод км/ч в м/с
    WalkingCalorie := (walkingCaloriesWeightMultiplier*weight + (math.Pow(SpeedInMs, 2)/(height*cmInM))*walkingSpeedHeightMultiplier*weight) * duration.Minutes()
    return WalkingCalorie
}
