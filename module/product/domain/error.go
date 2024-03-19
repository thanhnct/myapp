package productdomain

import "errors"

var (
	ErrorProductNameCannotBeBlank = errors.New("product name can not be blank")
)
