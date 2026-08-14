# grpc

RankSystem 的 gRPC 服务端工程（Go 实现）。生成代码提交在本工程，本工程不负责生成。

构建：

```bash
go build ./...
go run ./cmd/server
```

服务监听 `:9100`，提供 `MathUtil.Add` / `MathUtil.Sub` 两个 RPC。

工程结构：

```text
grpc/
├── cmd/server/main.go           # 服务入口
├── internal/service/mathutil.go # 手写的业务实现（UtilServer）
├── pb/math/                     # 生成代码（math.pb.go / math_grpc.pb.go，提交入库）
└── util/math.go                 # 工具函数
```

协议唯一来源：`../grpc_proto/math.proto`。协议变更时在 grpc_proto 跑 `gen.sh` / `gen.bat`，
生成代码会输出到本工程的 `pb/math/`，重新生成后把 `math.proto` 和 `pb/math/` 一起提交。
