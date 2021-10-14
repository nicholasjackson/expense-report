REPO = nicholasjackson/expense-report:report-latest

build_docker:
	docker build -t ${REPO} .

build_and_push_docker: build_docker
	docker push ${REPO}