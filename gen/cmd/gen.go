package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

func main() {
	var model, target, home string
	flag.StringVar(&model, "model", "", "生成数据表方法名")
	flag.StringVar(&target, "target", "../../core/dao", "生成文件目标文件夹")
	flag.StringVar(&home, "home", "../dao.tpl", "生成文件模版")
	flag.Parse()

	if model == "" {
		panic("model is null")
	}

	target = strings.TrimSuffix(target, "/")

	filePath := fmt.Sprintf("%s/%s.go", target, convertToUpperCamelCase(model))
	filePath, _ = filepath.Abs(filePath)
	if isExist, _ := pathExists(filePath); isExist {
		fmt.Printf("%s exists, ignored generation\n", filePath)
		os.Exit(0)
	}

	tplText, err := readTpl(home)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(fmt.Sprintf("%s.go", model)).
		Funcs(template.FuncMap{"firstUpper": firstUpper}).
		Parse(tplText)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err = tmpl.Execute(f, map[string]interface{}{"Model": model}); err != nil {
		panic(err)
	}
	fmt.Printf("generate file: %s\n", filePath)
}

func firstUpper(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func readTpl(tplPath string) (string, error) {
	f, err := os.Open(tplPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	fd, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(fd), nil
}

// pathExists 判断所给路径文件/文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func convertToUpperCamelCase(input string) string {
	runes := []rune(input)
	var result []string

	for i := 0; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			if i > 0 {
				result = append(result, "_")
			}
			result = append(result, string(unicode.ToLower(runes[i])))
		} else {
			result = append(result, string(runes[i]))
		}
	}

	return strings.Join(result, "")
}
