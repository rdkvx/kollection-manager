gorun:
	go run cmd/main.go

composeup:
	docker compose -f deployment/postgree-dockercompose.yml up -d

composerestart:
	docker compose -f "deployment/postgree-dockercompose.yml" restart