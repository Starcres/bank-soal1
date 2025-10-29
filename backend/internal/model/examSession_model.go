package model

import "time"

type ExamSession struct {
	Id         int           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId     int           `json:"user_id" gorm:not null"`
	ExamId     int           `json:"exam_id" gorm:"not null"`
	StartedAt  time.Time     `json:"started_at" gorm:"not null"`
	FinishedAt *time.Time    `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'in_progress'"`
	CurrentNo  int           `json:"current_no" gorm:"not null"`
	Score      *float64      `json:"score"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserAnswers []UserAnswer `gorm:"foreignKey:ExamSessionId"`
}

type UpdateExamSession struct {
	UserId     int           `json:"user_id" gorm:not null"`
	ExamId     int           `json:"exam_id" gorm:"not null"`
	StartedAt  time.Time     `json:"started_at" gorm:"not null"`
	FinishedAt *time.Time    `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'in_progress'"`
	CurrentNo  int           `json:"current_no" gorm:"not null"`
	Score      *float64      `json:"score"`
}

type UpdateCurrNo struct {
	CurrentNo int `json:"current_no" gorm:"not null"`
}

type UpdateStatus struct {
	Status SessionStatus `json:"status" gorm:"not null"`
}

type FinishExam struct {
	FinishedAt time.Time     `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'finished'"`
	Score      *float64      `json:"score"`
}

type ESessionResponse struct {
	Id         int           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId     int           `json:"user_id" gorm:not null"`
	ExamId     int           `json:"exam_id" gorm:"not null"`
	StartedAt  time.Time     `json:"started_at" gorm:"not null"`
	FinishedAt *time.Time    `json:"finished_at"`
	Status     SessionStatus `json:"status" gorm:"default:'in_progress'"`
	CurrentNo  int           `json:"current_no" gorm:"not null"`
	Score      *float64      `json:"score"`

	UserAnswers []UserAnswer `gorm:"foreignKey:ExamSessionId"`
}
