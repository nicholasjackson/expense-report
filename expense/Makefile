REPO = nicholasjackson/expense-report:expense-latest

build_docker:
	docker build -t ${REPO} .

build_and_push_docker: build_docker
	docker push ${REPO}