PROGRAM_NAME=alfred-new

GO_COMPILER=go
GO_FLAGS=GOOS=darwin GOARCH=amd64

BUILD_PATH=./dist

build:
	go run . > $(BUILD_PATH)/workflow.json

clean:
	rm -rf $(BUILD_PATH)/*