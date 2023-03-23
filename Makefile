all: filecoin-ffi build

BINARY_NAME=collectif-go-sdk
 
build: 
	go build -o ${BINARY_NAME} -v cmd/sdk.go
 
run:
	go build -o ${BINARY_NAME} cmd/sdk.go
	./${BINARY_NAME}
 
clean:
	go clean
	$(MAKE) -C $(FFI_PATH) clean
	rm ${BINARY_NAME}
	

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
