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
