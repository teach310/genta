// 公式のcompiler/protogenパッケージはgoの生成に特化している。
// go_packageの入力規則があるなど、縛りもきつい。
// https://pkg.go.dev/google.golang.org/protobuf/compiler/protogen
// ほしいのは汎用的なものなので自作する。
// インタフェースはOptionからRunするのが違和感あるので揃えない。
package protogen

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// 標準入力を変換して標準出力に流すっていう機能をもった構造体
type Plugin struct {
}

// オプションを渡す。
func NewPlugin() *Plugin {
	plugin := &Plugin{}
	return plugin
}

// Run executes a function as a protoc plugin.
//
// It reads a CodeGeneratorRequest message from os.Stdin, invokes the plugin
// function, and writes a CodeGeneratorResponse message to os.Stdout.
//
// If a failure occurs while reading or writing, Run prints an error to
// os.Stderr and calls os.Exit(1).
func (plugin *Plugin) Run() {
	if err := plugin.run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}

func (plugin *Plugin) run() error {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return err
	}
	resp := plugin.generate(req)
	out, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(out); err != nil {
		return err
	}
	return nil
}

func (plugin *Plugin) generate(req *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	protoFiles := make(map[string]*ProtoFile, len(req.ProtoFile))
	for _, fdesc := range req.ProtoFile {
		protoFile := &ProtoFile{Proto: fdesc}
		protoFiles[fdesc.GetName()] = protoFile
	}

	responseFiles := make([]*pluginpb.CodeGeneratorResponse_File, 0)
	for _, filename := range req.FileToGenerate {
		var sb strings.Builder
		protoFile := protoFiles[filename]
		sb.WriteString(plugin.getMessageInfoPrototype(protoFile.Proto.GetMessageType()))
		sb.WriteString("\n\nEND!\n")
		outputPath := strings.Replace(filename, ".proto", ".pb.txt", 1)
		content := sb.String()
		responseFiles = append(responseFiles, &pluginpb.CodeGeneratorResponse_File{
			Name:    proto.String(outputPath),
			Content: proto.String(content),
		})
	}
	resp := &pluginpb.CodeGeneratorResponse{
		File: responseFiles,
	}
	return resp
}

// 調査用。DescriptorProtoからの情報を文字列としてぬいて出力する
func (plugin *Plugin) getMessageInfoPrototype(messageTypes []*descriptorpb.DescriptorProto) string {
	var sb strings.Builder
	for _, messageType := range messageTypes {
		sb.WriteString("DescriptorProto.GetName(): ")
		sb.WriteString(messageType.GetName())
	}
	return sb.String()
}

type ProtoFile struct {
	Proto *descriptorpb.FileDescriptorProto
	// ToGenerate bool // true if we should generate code for this file TODO: 追記
}
