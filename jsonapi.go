package nrdb

import (
	"fmt"
	"net/url"
)

type Links struct {
	Self     *string `json:"self,omitempty"`
	Related  *string `json:"related,omitempty"`
	First    *string `json:"first,omitempty"`
	Last     *string `json:"last,omitempty"`
	Previous *string `json:"prev,omitempty"`
	Next     *string `json:"next,omitempty"`
}

type Relationship struct {
	Links *Links `json:"links"`
}

type Response[T any] struct {
	Data  T      `json:"data"`
	Links *Links `json:"links"`
}

type Document[A any, R any] struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Attributes    *A     `json:"attributes"`
	Links         *Links `json:"links"`
	Relationships *R     `json:"relationships"`
}

type Params struct {
	PageSize   *uint64
	PageNumber *uint64
}

func (p Params) SetPageInfo(query url.Values) url.Values {
	if p.PageSize != nil {
		query.Set("page[size]", fmt.Sprintf("%d", *p.PageSize))
	}

	if p.PageNumber != nil {
		query.Set("page[number]", fmt.Sprintf("%d", *p.PageNumber))
	}

	return query
}

func (p Params) Query() (url.Values, error) {
	query := url.Values{}

	p.SetPageInfo(query)

	return query, nil
}
