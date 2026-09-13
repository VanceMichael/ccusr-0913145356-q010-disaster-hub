# 灾后安置点转移对账中枢

服务对账各安置点的人员进出与转移事件，保证同一时刻只有一个有效位置。

## 技术约定

服务采用 Go 1.25 与 SQLite。HTTP 健康检查固定为 `GET /healthz`，业务错误使用结构化 JSON 返回。领域词汇和数据边界记录在 `docs/domain.md`，数据库结构位于 `db/`。

## 本地运行

```bash
docker build -t q010-mudslide-hub .
docker run --rm -p 8080:8080 q010-mudslide-hub
```

镜像构建阶段会执行现有自动化测试。也可以运行 `scripts/verify.sh`，该脚本会先校验容器配置，再完成一次干净构建。

## 目录

- `src/`、`app/` 或 `cmd/`：服务入口与领域代码。
- `db/`：数据库迁移和约束。
- `fixtures/`：可公开的领域样例。
- `tests/`：可执行验证代码。
