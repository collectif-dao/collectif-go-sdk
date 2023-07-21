all: filecoin-ffi build-sdk build-client

SDK_NAME=collectif-go-sdk
CLIENT_NAME=collectif-client

build-sdk:
	go build -o $(SDK_NAME) -v ./cmd/sdk/

build-client:
	go build -o $(CLIENT_NAME) -v ./cmd/client/

run-sdk:
	go build -o $(SDK_NAME) -v ./cmd/sdk/
	./${SDK_NAME}

run-client:
	go build -o $(CLIENT_NAME) -v ./cmd/client/
	./${CLIENT_NAME}
 
clean:
	go clean
	$(MAKE) -C $(FFI_PATH) clean
	rm -f $(SDK_NAME)
	rm -f $(CLIENT_NAME)

## FFI

FFI_PATH:=extern/filecoin-ffi/
FFI_DEPS:=.install-filcrypto
FFI_DEPS:=$(addprefix $(FFI_PATH),$(FFI_DEPS))

$(FFI_DEPS): build/.filecoin-install ;

filecoin-ffi: $(FFI_PATH)
	$(MAKE) -C $(FFI_PATH) $(FFI_DEPS:$(FFI_PATH)%=%)
	@touch $@

MODULES+=$(FFI_PATH)
BUILD_DEPS+=filecoin-ffi
CLEAN+=filecoin-ffi

ffi-version-check:
	@[[ "$$(awk '/const Version/{print $$5}' extern/filecoin-ffi/version.go)" -eq 3 ]] || (echo "FFI version mismatch, update submodules"; exit 1)
BUILD_DEPS+=ffi-version-check

.PHONY: ffi-version-check
