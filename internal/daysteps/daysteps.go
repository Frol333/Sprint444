package daysteps

import (
	"fmt"
	"time"
	"strconv"
    "strings"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	 // Разделяем строку на слайс строк по разделителю ","
	 parts := strings.Split(data, ",")


	 // Проверяем, что в слайсе ровно 2 элемента
	 if len(parts) != 2 {
	  return 0, 0, fmt.Errorf("invalid data format: %s", data)
	 }
   
	 // Преобразуем количество шагов в int
	 steps, err := strconv.Atoi(parts[0])
	 if err != nil {
	  return 0, 0, fmt.Errorf("invalid steps format: %s", parts[0])
	 }
   
	 // Проверяем, что количество шагов больше 0
	 if steps <= 0 {
	  return 0, 0, fmt.Errorf("steps must be positive: %d", steps)
	 }
   
	 // Преобразуем продолжительность прогулки в time.Duration
	 duration, err := time.ParseDuration(parts[1])
	 if err != nil {
	  return 0, 0, fmt.Errorf("invalid duration format: %s", parts[1])
	 }
   
	 //  возвращаем значения
	 return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {

		steps, duration, err := parsePackage(data) // Получаем данные из строки
		if err != nil {
			fmt.Println(err)
			return ""
		}
	
		if steps <= 0 {
			return ""
		}
	
		distanceMeters := float64(steps) * StepLength       // Вычисляем дистанцию в метрах
		distanceKilometers := distanceMeters / 1000.0          // Переводим в километры
		calories := WalkingSpentCalories(duration, weight) // Вычисляем калории
	
		return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distanceKilometers, calories)
	}
	

