# How to compile the proto files

1. Install the `protoc-gen-go` plugin if it is not already installed:

   ``` shell
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   ```

2. Ensure that the directory where the Go binaries are installed (usually `~/go/bin`) is included in your PATH. You can add it to your PATH by adding the following line to your shell's configuration file (e.g., `~/.zshrc` for zsh):

   ``` shell
   export PATH=$PATH:$(go env GOPATH)/bin
    ```

If you run these commands, the `protoc-gen-go` executable should be available, and the script should run without encountering the same error.
