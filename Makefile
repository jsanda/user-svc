
PHONY: code-gen
code-gen:
	@protoc -I pkg/pb pkg/pb/user_service.proto --go_out=plugins=grpc:pkg/pb


