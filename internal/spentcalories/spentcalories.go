package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps value")
	}

	name := strings.TrimSpace(parts[1])
	if name == "" {
		return 0, "", 0, fmt.Errorf("invalid name value")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid time.Duration")
	}

	return steps, name, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient

	return (float64(steps) * stepLength) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	hours := duration.Hours()
	averageSpeed := dist / hours
	return averageSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// Проверка входных параметров
	if weight <= 0 {
		return "", fmt.Errorf("weight must be positive")
	}

	if height <= 0 {
		return "", fmt.Errorf("height must be positive")
	}

	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Printf("parse training error: %v", err)
		return "", fmt.Errorf("invalid training data: %w", err)
	}

	trainingType = strings.TrimSpace(trainingType)

	switch trainingType {
	case "Бег":
		runCall, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("failed to calculate running calories: %w", err)
		}

		hours := float64(duration.Hours())
		durationStr := fmt.Sprintf("%0.2f ч.", hours)

		strRun := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %s\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			durationStr, distance(steps, height), meanSpeed(steps, height, duration), runCall)

		return strRun, nil

	case "Ходьба":
		walkCall, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("failed to calculate walking calories: %w", err)
		}

		hours := float64(duration.Hours())
		durationStr := fmt.Sprintf("%0.2f ч.", hours)

		strWalk := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %s\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			durationStr, distance(steps, height), meanSpeed(steps, height, duration), walkCall)
		return strWalk, nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", trainingType)
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("Invalid steps")
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid height")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("Invlaid duration")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()

	return weight * meanSpeed * durationInMinutes / minInH, nil

}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("Invalid steps")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("invalid  weight")
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid height ")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("Invlaid duration")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * meanSpeed * durationInMinutes / minInH

	return calories * walkingCaloriesCoefficient, nil
}
