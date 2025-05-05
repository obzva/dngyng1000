export GOOS=linux
export GOARCH=amd64

.PHONY: create update

bootstrap: ./cmd/lambda/main.go
	@echo "□ building go binary"
	@go build -tags lambda.norpc -o bootstrap ./cmd/lambda
	@echo "■ done"

blogFunction.zip: bootstrap
	@echo "□ creating zip file"
	@zip blogFunction.zip bootstrap
	@echo "■ done"
	@echo "□ removing a temporary file: bootstrap"
	@rm -f bootstrap
	@echo "■ done"
	@echo "□ setting file permissions"
	@chmod 644 blogFunction.zip
	@echo "■ done"

create: blogFunction.zip
	@echo "□ creating lambda function"
	@aws lambda create-function \
	--function-name blogFunction \
	--runtime provided.al2023 \
	--handler bootstrap \
	--role arn:aws:iam::283218086904:role/lambda-basic-ex \
	--zip-file fileb://blogFunction.zip
	@echo "■ done"
	@$(MAKE) clean_zip

update: blogFunction.zip
	@echo "□ updating lambda function"
	@aws lambda update-function-code \
	--function-name blogFunction \
	--zip-file fileb://blogFunction.zip
	@echo "■ done"
	@$(MAKE) clean_zip

clean_zip:
	@echo "□ removing a temporary file: blogFunction.zip"
	@rm -f blogFunction.zip
	@echo "■ done"