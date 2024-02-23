api:
	goctl api go -api core/core.api -dir core --home template --style go_zero
	goctl api format --dir core/apis -declare

gentool:
	gentool -c gen/gen.yaml

dao:
	@goctl model mysql datasource --url "root:123456@tcp(127.0.0.1:3306)/zust_acm" --table $(model) --dir core/model --home template --style go_zero

crud:
	@gentool -c gen/gen.yaml
	@echo '📢Info: gentool exec complete!!!'
	@make dao model=$(model)
	@go run gen/cmd/gen.go -model=$(model) -api=core/apis -logic=core/internal/logic -home=gen/tpl
	@goctl api format --dir core/apis -declare
	@echo '📢Info: gen crud complete!!!'

test:
	goctl model mysql datasource --url "root:123456@tcp(127.0.0.1:3306)/zust_acm" --table $(model) --dir core/model --home template --style go_zero