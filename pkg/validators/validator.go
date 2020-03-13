package validators

type Validator interface {
	Validate(source, target interface{}) (bool, error)
}
