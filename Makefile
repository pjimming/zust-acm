api:
	goctl api go -api core/core.api -dir core --home template --style go_zero
	goctl api format --dir core/apis -declare

model:
	goctl model mysql ddl --dir core/model --home template --src sql/zust_acm.sql --style go_zero --strict true -i 'gmt_created,gmt_updated'