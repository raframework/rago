package validation

type messages interface {
	List() (messages map[string]string)
}
