run:
	templ generate
	npx @tailwindcss/cli -i ./pkg/views/css/input.css -o ./pkg/views/css/output.css 
	go run cmd/main.go

up:
	goose -dir pkg/database/migrations mysql "root:12345@tcp(172.30.128.1:3306)/seed" up 

down:
	goose -dir pkg/database/migrations mysql "root:12345@tcp(172.30.128.1:3306)/seed" down
