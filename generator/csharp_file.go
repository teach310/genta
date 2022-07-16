package generator

type CSharpFile struct {
	Namespace string
	ClassList []*CSharpClass
}

type CSharpClass struct {
	Name   string
	Fields []*CSharpClassField
}

type CSharpClassField struct {
	Name     string
	TypeName string
}
