api:
	goctl api go -api core/core.api -dir core --home template --style go_zero
	goctl api format --dir core/apis -declare

gentool:
	gentool -c gen/gen.yaml

crud:
	@gentool -c gen/gen.yaml
	@echo '📢Info: gentool exec complete!!!'
	@go run gen/cmd/gen.go -model=$(model) -dao=core/dao -api=core/apis -logic=core/internal/logic -home=gen/tpl
	@goctl api format --dir core/apis -declare
	@echo '📢Info: gen crud complete!!!'

test:
	go run gen/cmd/gen.go -model=$(model) -dao=core/dao -api=core/apis -logic=core/internal/logic -home=gen/tpl