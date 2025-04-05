package daysteps

import (
    "fmt"
    "strconv"
    "strings"
    "time"
)

var (
    StepLength = 0.65 // длина шага в метрах
)

// parsePackage разбирает строку с данными о шагах и времени.
func parsePackage(data string) (int, time.Duration, error) {
    parts := strings.Split(data, ",")

    if len(parts) != 2 {
        return 0, 0, fmt.Errorf("invalid data format: %s", data)
    }

    steps, err := strconv.Atoi(parts[0])
    if err != nil {
        return 0, 0, fmt.Errorf("invalid steps format: %s", parts[0])
    }

    if steps <= 0 {
        return 0, 0, fmt.Errorf("steps must be positive: %d", steps)
    }

    duration, err := time.ParseDuration(parts[1])
    if err != nil {
        return 0, 0, fmt.Errorf("invalid duration format: %s", parts[1])
    }

    return steps, duration, nil
}

// DayActionInfo обрабатывает данные и возвращает информацию о дневной активности.
func DayActionInfo(data string, weight, height float64) string {
    steps, duration, err := parsePackage(data)
    if err != nil {
        fmt.Println(err)
        return ""
    }

    distanceMeters := float64(steps) * StepLength
    distanceKilometers := distanceMeters / 1000.0
    calories := WalkingSpentCalories(duration, weight)

    return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKilometers, calories)
}

// WalkingSpentCalories рассчитывает потраченные калории при ходьбе.
func WalkingSpentCalories(duration time.Duration, weight float64) float64 {
    const (
        caloriesPerHourPerKg = 60.0 // Примерное значение калорий в час на кг веса
    )
    durationHours := duration.Hours()
    calories := caloriesPerHourPerKg * weight * durationHours
    return calories
}
