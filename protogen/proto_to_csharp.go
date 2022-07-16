package protogen

import (
	"errors"
	"fmt"

	"github.com/teach310/genta/generator"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (protoFile *ProtoFile) BuildCSharpFile() (*generator.CSharpFile, error) {
	typeConverter := &CSharpTypeConverter{}
	namespace := protoFile.Proto.GetPackage()
	var classList []*generator.CSharpClass
	for _, message := range protoFile.Proto.GetMessageType() {
		var fields []*generator.CSharpClassField
		for _, protoField := range message.GetField() {
			typeName, err := typeConverter.GetScalarTypeName(protoField.GetType())
			if err != nil {
				return nil, err
			}

			csharpField := &generator.CSharpClassField{
				Name:     protoField.GetName(),
				TypeName: typeName,
			}
			fields = append(fields, csharpField)
		}
		class := &generator.CSharpClass{
			Name:   message.GetName(),
			Fields: fields,
		}
		classList = append(classList, class)
	}
	return &generator.CSharpFile{
		Namespace: namespace,
		ClassList: classList,
	}, nil
}

type CSharpTypeConverter struct {
}

var protoToCSharpTypeMap = map[descriptorpb.FieldDescriptorProto_Type]string{
	descriptorpb.FieldDescriptorProto_TYPE_INT32:  "int",
	descriptorpb.FieldDescriptorProto_TYPE_INT64:  "long",
	descriptorpb.FieldDescriptorProto_TYPE_BOOL:   "bool",
	descriptorpb.FieldDescriptorProto_TYPE_STRING: "string",
}

func (c *CSharpTypeConverter) GetScalarTypeName(typeProto descriptorpb.FieldDescriptorProto_Type) (string, error) {
	if typeProto == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		return "", errors.New("not scalar")
	}
	typeName, ok := protoToCSharpTypeMap[typeProto]
	if !ok {
		return "", fmt.Errorf("[CSharpTypeConverter]unsupported type error %v", typeProto)
	}
	return typeName, nil
}
