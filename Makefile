## build: Билд проекта
build: 
	go build -o .bin/ cmd/telegram-bot/main.go

## run: Запуск экзешника
run:
	./.bin/main
help: Makefile
	@echo
	@echo "Доступные команды "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo