PATH="$PATH:${GOPATH}/bin:${HOME}/go/bin" protoc --go_out=../dfs/ ./*.proto
# PATH="$PATH:${GOPATH}/bin:${HOME}/go/bin" protoc --go_out=../chat/ ./*.proto