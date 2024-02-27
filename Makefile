# Custom configuration | 自定义配置
# 项目名称
SERVICE = core
# 项目格式化之后的名称
SERVICE = core
# 项目名称全小写格式
SERVICE_LOWER = core
# 项目名称下划线格式
SERVICE_SNAKE = core
# 项目名称短杠格式
SERVICE_DASH = core

# git仓库当前版本号
VERSION = $(shell git describe --tags --always)

# 项目文件命名风格
PROJECT_STYLE = go_zero

# Swagger文件类型，支持yml,json
SWAGGER_TYPE = json

# Ent 启用的官方特性(提供SQL查询执行，提供拦截器)
ENT_FEATURE = sql/execquery,intercept

# 构建框架
GOARCH = amd64

# rpc服务的proto文件的引入的包的目录
RPC_PROTO_IMPORT = /home/potatomine/code/pkg/mod

GO ?= go

# 使用双引号可以确保-s作为一个整体传递给gofmt命令
GOFMT ?= gofmt "-s"
GOFILES := $(shell find . -name "*.go")
# 链接器的标记和选项
LDFLAGS := -s -w

.PHONY: test
test: # 执行项目测试,但是要编写测试文件
	go test -v --cover ./api/internal/..
	go test -v --cover ./rpc/internal/..

.PHONY: fmt
fmt: # 格式化代码
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint: # 运行代码错误分析
	golangci-lint run -D staticcheck

.PHONY: tools
tools: # 安装必要工具
	$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	$(GO) install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: docker
docker: #构建docker镜像
	docker build -f Dockerfile-api -t ${DOCKER_USERNAME}/$(SERVICE_DASH)-api:${VERSION}
	docker build -f Dockerfile-rpc -t ${DOCKER_USERNAME}/$(SERVICE_DASH)-rpc:${VERSION}

.PHONY: publish-docker
publish-docker: # 发布docker镜像
	echo "${DOCKER_PASSWORD}" | docker login --username ${DOCKER_USERNAME} --password-stdin https://${REPO}
	docker push ${DOCKER_USERNAME}/${SERVICE_DASH}-rpc:${VERSION}
	docker push ${DOCKER_USERNAME}/${SERVICE_DASH}-api:${VERSION}
	@echo "Publish docker successfully"

.PHONY: gen-api
gen-api: # 生成api代码
	goctl api go --api ./api/desc/all.api --dir ./api --style=$(PROJECT_STYLE)
	swagger generate spec --output=./$(SERVICE_STYLE).$(SWAGGER_TYPE) --scan-models
	@echo "Generate API files successfully"

# 这里goctl原生的功能受限了,import的文件需要再proto文件中使用option go_package = ...导入 请注意！
# 鉴于在多个proto文件中添加option非常麻烦,所以使用了suyuan改版过的goctls工具更加方便的生成rpc代码
.PHONY: gen-rpc
gen-rpc: # 生成rpc代码
	goctls rpc protoc -I=. -I=$(RPC_PROTO_IMPORT) ./rpc/$(SERVICE_STYLE).proto --style=$(PROJECT_STYLE) --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc --styles=$(PROJECT_STYLE)
	@echo "Generate PRC files successfully"

# 模板：分页相关模板 设置可选字段模板
.PHONY: gen-ent
gen-ent: # 生成ent代码
	go run -mod=mod entgo.io/ent/cmd/ent generate --template glob=".rpc/ent/template/*.tmpl" ./rpc/ent/schema --feature $(ENT_FEATURE)
	@echo "Generate Ent files successfully"

# 这里使用goctls,可以根据schema中的go文件一步到位,就不需要再生成proto文件然后make gen-rpc了
.PHONY: gen-rpc-ent-logic
gen-rpc-ent-logic:
	goctls rpc ent --schema=./rpc/ent/schema --style=$(PROJECT_STYLE) --import_prefix=/rpc --service_name=$(SERVICE) --project_name=$(PROJECT_STYLE) -o=./rpc --model=$(model) -group=$(group) --proto_out=./rpc/desc/$(shell echo $(model) | tr A-Z a-z).proto --overwrite=true
	@echo "Generate logic codes from ent successfully"

.PHONY: api
api: # 运行api模块
	go run ./api/core.go -f ./api/etc/core.yaml

.PHONY: rpc
rpc: # 运行rpc模块
	go run ./rpc/core.go -f ./rpc/etc/core.yaml
