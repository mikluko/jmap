package sievescript

import "github.com/mikluko/jmap"

type Filter interface {
	implementsFilter()
}

type FilterOperator struct {
	Operator   jmap.Operator `json:"operator,omitempty"`
	Conditions []Filter      `json:"conditions,omitempty"`
}

func (fo *FilterOperator) implementsFilter() {}

type FilterCondition struct {
	Name     string `json:"name,omitempty"`
	IsActive bool   `json:"isActive,omitempty"`
}

func (fc *FilterCondition) implementsFilter() {}
