// Code generated by sqlcgetters; DO NOT EDIT.

package querytest

func (x Foo) GetName() string {
	return x.Name
}

func (x Foo) GetBio() string {
	return x.Bio
}

func (x AtParamsParams) GetSlug() string {
	return x.Slug
}

func (x AtParamsParams) GetFilter() bool {
	return x.Filter
}

func (x FuncParamsParams) GetSlug() string {
	return x.Slug
}

func (x FuncParamsParams) GetFilter() bool {
	return x.Filter
}

func (x InsertAtParamsParams) GetName() string {
	return x.Name
}

func (x InsertAtParamsParams) GetBio() string {
	return x.Bio
}

func (x InsertFuncParamsParams) GetName() string {
	return x.Name
}

func (x InsertFuncParamsParams) GetBio() string {
	return x.Bio
}

func (x UpdateParams) GetSetName() bool {
	return x.SetName
}

func (x UpdateParams) GetName() string {
	return x.Name
}