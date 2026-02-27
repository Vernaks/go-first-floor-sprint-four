package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid data format: expected 'steps duration', got %q", data)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid steps value %q: %w", parts[0], err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("steps must be positive, got %d", steps)
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid duration value %q: %w", parts[1], err)
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("duration must be positive, got %v", duration)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Printf("Ошибка парсинга")
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanse := float64(steps) * stepLength
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	distanceKm := distanse / mInKm

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %0.2f км.\nВы сожгли %0.2f ккал.\n", steps, distanceKm, calories)
	return result

}
