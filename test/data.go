package test

type (
	Test struct {
		One   int    `json:"one" json.required:"true"`
		Two   string `json:"two"`
		Three bool   `json:"three"`
	}
)
