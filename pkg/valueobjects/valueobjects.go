package valueobjects

import (
	"fmt"
	"regexp"
)

type Date string
type Email string
type Flag bool
type Id string
type Money float64
type Name string
type Numeric int
type Text string
type Url string

func (d Date) IsValid() error {
	regex := `^\d{4}-\d{2}-\d{2}$`
	match, _ := regexp.MatchString(regex, string(d))
	if !match {
		return fmt.Errorf("invalid date")
	}
	return nil
}

func (e Email) IsValid() error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(regex, string(e))
	if !match {
		return fmt.Errorf("invalid email")
	}
	return nil
}

func (m Money) Format() string {
	return fmt.Sprintf("%.2f", m)
}

func (m Money) IsValid() error {
	if m < 0 {
		return fmt.Errorf("invalid money")
	}
	return nil
}

func (m Money) Sum(m2 Money) Money {
	return Money(float64(m) + float64(m2))
}

func (n Name) IsValid() error {
	regex := `^[a-zA-Z0-9]+$`
	match, _ := regexp.MatchString(regex, string(n))
	if !match {
		return fmt.Errorf("invalid name")
	}
	return nil
}
