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

const (
	ExtGolang = "go"
	ExtApi    = "api"
)

func main() {
	var model, dao, api, logic, home string
	flag.StringVar(&model, "model", "", "生成数据表方法名")
	flag.StringVar(&dao, "dao", "../../core/dao", "生成dao文件目标文件夹")
	flag.StringVar(&api, "api", "../../core/apis", "生成api文件目标文件夹")
	flag.StringVar(&logic, "logic", "../../core/internal/logic", "生成logic文件目标文件夹")
	flag.StringVar(&home, "home", "../tpl", "生成文件模版")
	flag.Parse()

	if model == "" {
		panic("model is null")
	}

	home = strings.TrimSuffix(home, "/")

	// render dao
	render(dao, ExtGolang, model, home+"/dao.tpl", convertToUpperCamelCase(model))
	// render api
	render(api, ExtApi, model, home+"/api.tpl", convertToUpperCamelCase(model))
	// render logic crud
	logic = strings.TrimSuffix(logic, "/")
	// - create
	render(
		fmt.Sprintf("%s/%s", logic, strings.ToLower(model)),
		ExtGolang,
		model,
		home+"/create.tpl",
		fmt.Sprintf("add_%s_logic", convertToUpperCamelCase(model)),
	)
	// - read
	render(
		fmt.Sprintf("%s/%s", logic, strings.ToLower(model)),
		ExtGolang,
		model,
		home+"/read.tpl",
		fmt.Sprintf("get_%s_logic", convertToUpperCamelCase(model)),
	)
	// - update
	render(
		fmt.Sprintf("%s/%s", logic, strings.ToLower(model)),
		ExtGolang,
		model,
		home+"/update.tpl",
		fmt.Sprintf("update_%s_logic", convertToUpperCamelCase(model)),
	)
	// -delete
	render(
		fmt.Sprintf("%s/%s", logic, strings.ToLower(model)),
		ExtGolang,
		model,
		home+"/delete.tpl",
		fmt.Sprintf("delete_%s_logic", convertToUpperCamelCase(model)),
	)
	// -utils
	render(
		fmt.Sprintf("%s/%s", logic, strings.ToLower(model)),
		ExtGolang,
		model,
		home+"/logic_utils.tpl",
		"utils",
	)
}

func render(target, ext, model, tpl, filename string) {
	target = strings.TrimSuffix(target, "/")
	filePath := fmt.Sprintf("%s/%s.%s", target, filename, ext)
	filePath, _ = filepath.Abs(filePath)
	if isExist, _ := pathExists(filePath); isExist {
		fmt.Printf("%s exists, ignored generation\n", filePath)
		return
	}

	tplText, err := readTpl(tpl)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").
		Funcs(template.FuncMap{
			"firstUpper":              firstUpper,
			"convertToUpperCamelCase": convertToUpperCamelCase,
			"toLower":                 strings.ToLower,
		}).
		Parse(tplText)
	if err != nil {
		panic(err)
	}

	if err = os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
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
