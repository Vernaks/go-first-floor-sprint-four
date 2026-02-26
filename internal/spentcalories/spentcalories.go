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
	// TODO: реализовать функцию
	parts := strings.Split(data, " ")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid len data")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps value")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid time.Duration")
	}

	return steps, parts[1], duration, nil
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
	// TODO: реализовать функцию
	steps, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("invalid")
	}

	switch trainingType {
	case "Бег":
		runCall, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("invalid runCall")
		}
		strRun := fmt.Sprintf("Тип тренировки: Бег\n Длительность: %v \n Дистанция: %.2f км\n Cкорость: %.2f км/ч\n Сожгли калорий: %.2f",
			duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), runCall)

		return strRun, nil
	case "ходьба":
		walkCall, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", fmt.Errorf("invalid walkCall")
		}
		strWalk := fmt.Sprintf("Тип тренировки: Ходьба \n Длительность: %v \n Дистанция: %.2f км\n Cкорость: %.2f км/ч\n Сожгли калорий: %.2f",
			duration.Hours(), distance(steps, height), meanSpeed(steps, height, duration), walkCall)
		return strWalk, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("Invalid steps")
	}

	if weight == 0 || height == 0 {
		return 0, fmt.Errorf("invalid height or weight")
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

	if weight == 0 || height == 0 {
		return 0, fmt.Errorf("invalid height or weight")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("Invlaid duration")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * meanSpeed * durationInMinutes / minInH

	return calories * walkingCaloriesCoefficient, nil
}
