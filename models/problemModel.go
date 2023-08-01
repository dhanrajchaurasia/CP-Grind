package models

import (
	"gorm.io/gorm"
	"time"
)

type Problem struct {
	gorm.Model
	Name      string `validate:"required, min = 3, max = 32"`
	URL       string `validate:"required, min = 3, max = 32"`
	Verdict   Verdict
	Note      string
	HintTaken float32
	TimeTaken time.Duration
	Platform  string `validate:"required, min = 3, max = 32"`
}

type ProblemGrind struct {
	gorm.Model
	User     string
	Problems []Problem
}

type Verdict string

const (
	AC       Verdict = "AC"
	WA       Verdict = "WA"
	TLE      Verdict = "TLE"
	RTE      Verdict = "RTE"
	CE       Verdict = "CE"
	NotTried Verdict = "NOT TRIED"
)
