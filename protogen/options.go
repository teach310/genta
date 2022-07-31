package protogen

import "strings"

type Options struct {
	TemplatesPath string // .tmplファイルのあるフォルダのパス
	Model         string // .protoをなんのモデルに変換するのかを指定。 cs, txt
}

// paramter CodeGeneratorRequest.GetParameterで取得できるカンマ区切りのパラメータ
func parseOptions(parameter string) *Options {
	options := &Options{}
	for _, param := range strings.Split(parameter, ",") {
		i := strings.Index(param, "=")
		if i < 0 {
			continue // bool値でも必ずtrue, falseを明示的に指定することを必須とする。
		}

		key := param[0:i]
		value := param[i+1:]

		switch key {
		case "templates_path":
			options.TemplatesPath = value
		case "model":
			options.Model = value
		}
	}
	return options
}
