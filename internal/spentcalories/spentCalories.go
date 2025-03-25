package spentcalories

import (
	"time"
	"strconv"
    "strings"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep = 0.65 // средняя длина шага.
	mInKm   = 1000 // количество метров в километре.
	minInH  = 60   // количество минут в часе.
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",") 
	if len(parts) != 3 {
	 return 0, "", 0, fmt.Errorf("invalid data format") 
	}
  
  
	steps, err := strconv.Atoi(parts[0]) 
	if err != nil {
	 return 0, "", 0, fmt.Errorf("invalid steps format: %w", err)
	}
  
  
	activityType := parts[1] 
  
  
	duration, err := time.ParseDuration(parts[2]) 
	if err != nil {
	 return 0, "", 0, fmt.Errorf("invalid duration format: %w", err)
	}
  
  
	return steps, activityType, duration, nil 
}

// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
func distance(steps int) float64 {
	return float64(steps) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
func meanSpeed(steps int, duration time.Duration) float64 {
	if duration == 0 {
        return 0
    }
    distance := distance(steps)
    return distance / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// data string - строка с данными.
// weight, height float64 — вес и рост пользователя.
func TrainingInfo(data string, weight, height float64) string {
	switch{
	case trainingType == "Бег":
		distance := distance(steps)
		speed := meanSpeed(steps,duration)
		calories :=  RunningSpentCalories(steps, weight, duration)
	case trainingType == "Ходьба":
		distance := distance(steps)
		speed := meanSpeed(steps,duration)
		calories :=  WalkingSpentCalories(steps, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)	
	}
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	speed := meanSpeed(steps,duration)
    RunningCalories := (runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift * weight / mInKm * duration * minInH)
    return RunningCalories
}



// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// duration time.Duration — длительность тренировки.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	speed := meanSpeed(action,duration)
	SpeedInMs := speed * minInH
	WalkingCalorie := (walkingCaloriesWeightMultiplier*weight + (math.Pow(SpeedInMs, 2)/(height/cmInM))*walkingSpeedHeightMultiplier*weight) * duration * minInH
    return WalkingCalorie


}
