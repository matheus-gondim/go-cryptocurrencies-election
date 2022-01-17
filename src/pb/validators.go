package pb

import (
	validate "github.com/go-ozzo/ozzo-validation"
)

func (c *CreateCryptocurrency) IsValid() error {
	return validate.ValidateStruct(c,
		validate.Field(&c.Name, validate.Length(5, 255), validate.Required),
		validate.Field(&c.Symbol, validate.Length(2, 5), validate.Required),
	)
}

func (c *UpdateCryptocurrency) IsValid() error {
	return validate.ValidateStruct(c,
		validate.Field(&c.Id, validate.Required, validate.Min(1)),
		validate.Field(&c.Name, validate.Length(5, 255)),
		validate.Field(&c.Symbol, validate.Length(2, 5)),
	)
}
