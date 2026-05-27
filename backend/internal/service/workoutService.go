package service

import (
	"context"

	"github.com/Oliverstalsy/egolifter/internal/domain"
	"github.com/Oliverstalsy/egolifter/internal/random"
)

type WorkoutService struct {
	repo domain.WorkoutRepository
}

func NewWorkoutService(repo domain.WorkoutRepository) *WorkoutService {
	return &WorkoutService{
		repo: repo,
	}
}

func ChooseExercise() []domain.Exercise {
	return random.Exercises
}

func (w *WorkoutService) FindWorkouts(ctx context.Context, userID string) ([]domain.Workout, error) {
	return w.repo.GetWorkouts(ctx, userID)
}
