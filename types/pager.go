package types

import (
	"github.com/go-playground/validator/v10"
)

type Pager struct {
	Offset int `query:"offset" validate:"min=0"`
	Limit  int `query:"limit" validate:"max=10"`
}

func (p *Pager) Validate() error {
	if p.Limit <= 0 {
		p.Limit = 10
	}

	var structValidator = validator.New()
	structError := structValidator.Struct(p)

	if structError != nil {
		return structError
	}

	return nil
}

type PagerResult struct {
	Next     int         `json:"next"`
	Previous int         `json:"previous"`
	Total    int         `json:"total"`
	Data     interface{} `json:"data"`
	Pager    Pager       `json:"-"`
}

func (p *PagerResult) Pages() *PagerResult {
	var current int

	if p.Pager.Offset == 0 {
		current = 1 * p.Pager.Limit
	} else {
		current = p.Pager.Offset * p.Pager.Limit
	}

	if current < p.Total {
		p.Next = p.Pager.Offset + 1
	} else {
		p.Next = 0
	}

	if current > p.Pager.Limit {
		p.Previous = p.Pager.Offset - 1
	} else {
		p.Previous = 0
	}

	return p
}
