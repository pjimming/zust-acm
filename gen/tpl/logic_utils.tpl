package {{toLower .Model}}

func modelToTypes(from *model.{{firstUpper .Model}}) *types.{{firstUpper .Model}} {
	tmp := &types.{{firstUpper .Model}}{}
	_ = copier.Copy(tmp, from)
	tmp.CreatedAt = from.CreatedAt.UnixMilli()
    tmp.UpdatedAt = from.UpdatedAt.UnixMilli()
	// todo: custom trans
	return tmp
}
