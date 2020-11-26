package api

func (a *apiCondition) convert() *Condition {
	if a == nil {
		return nil
	}

	return &Condition{
		Name:        a.Name,
		Description: a.Desc,
	}
}
