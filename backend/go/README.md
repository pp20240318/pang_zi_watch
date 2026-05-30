# 胖子腕表 · Go API

## 环境

MySQL（见 `说明.txt`）：

- 库名：`watch`
- 用户：`watch`
- 密码：`Aa123456`

可通过环境变量覆盖，前缀 `WATCH_`，参见 `env.example`。

## 启动

```bash
cd backend/go
go run ./cmd/server
```

默认监听 `:8080`，首次启动会：

1. 自动建表（GORM AutoMigrate）
2. 创建默认管理员 `admin` / `admin123`
3. 若商品表为空，导入与 C 端 `catalog.js` 一致的种子数据

## 接口

- 公开：`/api/public/*`（C 端商城）
- 管理：`/api/*`（JWT）
- 上传：`POST /api/upload`，静态文件 `/uploads/*`
