package domain

import (
	"context"
	"time"
)

type ExerciseSet struct {
	ID         string
	ExerciseID string
	WorkoutID  string
	Sets       int
	Reps       int
	Weight     int
}

type Workout struct {
	ID       string
	UserID   string
	Date     time.Time
	Notes    string
	Exercise []ExerciseSet
}

type WorkoutRepository interface {
	GetWorkouts(ctx context.Context, userID string) ([]Workout, error)
	CreateWorkout(ctx context.Context, workout *Workout) (*Workout, error)
}
