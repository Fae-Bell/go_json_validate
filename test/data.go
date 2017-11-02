package test

type (
	Test struct {
		One   int    `json:"one" json.required:"true"`
		Two   string `json:"two" json.pattern:"ABC"`
		Three bool   `json:"three"`
		Sub   TestTwo
	}

	TestTwo struct {
		SubField string `json:"sub_field" json.required:"true"`
	}
)
