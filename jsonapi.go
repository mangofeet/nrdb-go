package nrdb

type Links struct {
	Self    string `json:"self,omitempty"`
	Related string `json:"related,omitempty"`
}

type Relationship struct {
	Links *Links `json:"links"`
}

type Response[T any] struct {
	Data []*T `json:"data"`
}

type Document[A any, R any] struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Attributes    *A     `json:"attributes"`
	Links         *Links `json:"links"`
	Relationships *R     `json:"relationships"`
}

type Params struct {
	PageLimit  *uint64
	PageOffset *uint64
}

func (p Params) SetPageInfo(query url.Values) url.Values {
	if p.PageLimit != nil {
		query.Set("filter[page]", fmt.Sprintf("%d", *p.PageLimit))
	}

	if p.PageOffset != nil {
		query.Set("filter[offset]", fmt.Sprintf("%d", *p.PageOffset))
	}

	return query
}

func (p Params) Query() (url.Values, error) {
	query := url.Values{}

	p.SetPageInfo(query)

	return query, nil
}
