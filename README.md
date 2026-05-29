# 🚀 Go Fiber Boilerplate

> Production-ready REST API template — **Go** + **Fiber v2** + **MySQL** + **Redis** + **JWT** + **Queue** + **Scheduler**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Fiber](https://img.shields.io/badge/Fiber-v2.52-00ACD7?style=flat)](https://gofiber.io/)
[![License](https://img.shields.io/badge/License-MIT-blue?style=flat)](LICENSE)

Boilerplate giúp bạn khởi động dự án Go mới **trong vài phút**, không cần setup từ đầu. Tích hợp sẵn authentication, background jobs, cron scheduler, migration và logging.

---

## ✨ Tính năng

| Tính năng | Chi tiết |
|---|---|
| ⚡ **HTTP Framework** | [Fiber v2](https://gofiber.io/) — hiệu năng cao, cú pháp quen thuộc như Express.js |
| 🗄️ **ORM** | [GORM](https://gorm.io/) — hỗ trợ soft delete, hooks, eager loading |
| 🔐 **JWT Auth** | Đăng nhập / đăng xuất với token blacklist trên Redis |
| 🚦 **Background Jobs** | [Asynq](https://github.com/hibiken/asynq) — queue worker chạy qua Redis |
| ⏰ **Cron Scheduler** | [Gocron v2](https://github.com/go-co-op/gocron) — scheduled tasks linh hoạt |
| 📦 **Migration** | [golang-migrate](https://github.com/golang-migrate/migrate) — SQL migration có version, rollback được |
| 🪵 **Logging** | [Zap](https://github.com/uber-go/zap) + [Lumberjack](https://github.com/natefinch/lumberjack) — structured logging, log rotation |
| 🔧 **Config** | [Viper](https://github.com/spf13/viper) — load `.env`, hỗ trợ nhiều môi trường |
| ✅ **Validation** | [go-playground/validator v10](https://github.com/go-playground/validator) — validate request đầu vào |
| 🏗️ **Architecture** | Layered: Handler → Service → Repository + Dependency Injection |

---

## 📁 Cấu trúc project

```
myapp/
├── main.go                      # Entry point — khởi động HTTP server
├── .env / .env.example          # Biến môi trường
├── Makefile                     # Các lệnh tiện ích
│
├── config/                      # Khởi tạo kết nối DB, Redis, load config
│   ├── config.go
│   ├── database.go
│   └── redis.go
│
├── app/
│   ├── bootstrap/               # Wiring toàn bộ dependencies (DI container)
│   ├── container/               # Struct chứa tất cả layers
│   ├── http/
│   │   ├── handlers/            # Controller — parse request, gọi service
│   │   ├── requests/            # Request structs + validation rules
│   │   └── responses/           # Response helpers (success, error, pagination)
│   ├── services/                # Business logic
│   ├── repositories/            # Database queries (GORM)
│   └── middleware/              # Auth, CORS, error handler
│
├── models/                      # GORM models
├── routes/                      # Đăng ký routes, gắn middleware
│
├── pkg/
│   └── cache/                   # Redis cache wrapper
│
├── database/
│   └── migrations/              # SQL migration files (*.up.sql / *.down.sql)
│
├── cmd/
│   ├── migrate/                 # CLI tool chạy migration
│   ├── worker/                  # Background job worker (Asynq)
│   └── schedule/                # Cron scheduler (Gocron)
│
├── jobs/                        # Job definitions & handlers
└── schedule/                    # Đăng ký scheduled tasks
```

### Luồng xử lý request

```
HTTP Request
    └─► Middleware (CORS · Auth · Logger)
            └─► Handler  — parse & validate request
                    └─► Service  — business logic
                            └─► Repository  — database / cache
                                    └─► Response (JSON)
```

---

## ⚙️ Yêu cầu

| Công cụ | Phiên bản tối thiểu |
|---|---|
| Go | 1.21+ |
| MySQL | 8.0+ |
| Redis | 6.0+ |

---

## 🚀 Bắt đầu nhanh

### 1. Clone & cài dependencies

```bash
git clone https://github.com/huynhan147/go-fiber-boilerplate.git myapp
cd myapp
go mod tidy
```

### 2. Cấu hình môi trường

```bash
cp .env.example .env
```

Chỉnh sửa `.env` theo môi trường của bạn:

```env
APP_NAME=myapp
APP_ENV=local
APP_PORT=3000

DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=myapp_db
DB_USERNAME=root
DB_PASSWORD=secret
DB_CHARSET=utf8mb4

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

JWT_SECRET=your-super-secret-jwt-key
JWT_EXPIRE_HOURS=24
```

### 3. Tạo database & chạy migration

```bash
mysql -u root -p -e "CREATE DATABASE myapp_db CHARACTER SET utf8mb4;"
make migrate-up
```

### 4. Khởi động server

```bash
make run
# hoặc
go run main.go
```

Server chạy tại `http://localhost:3000` 🎉

---

## 📡 API Endpoints

Base path: `/api`

### Auth

| Method | Endpoint | Mô tả | Auth |
|---|---|---|---|
| `POST` | `/api/auth/login` | Đăng nhập, trả JWT token | ❌ |
| `GET` | `/api/auth/me` | Thông tin user hiện tại | ✅ |
| `POST` | `/api/auth/logout` | Đăng xuất, blacklist token | ✅ |

### Users

| Method | Endpoint | Mô tả | Auth |
|---|---|---|---|
| `GET` | `/api/users` | Danh sách users (phân trang) | ✅ |
| `POST` | `/api/users` | Tạo user mới | ✅ |
| `GET` | `/api/users/:id` | Chi tiết user | ✅ |
| `PUT` | `/api/users/:id` | Cập nhật user | ✅ |
| `DELETE` | `/api/users/:id` | Xóa user (soft delete) | ✅ |

### Transactions

| Method | Endpoint | Mô tả | Auth |
|---|---|---|---|
| `GET` | `/api/transactions` | Danh sách transactions | ✅ |
| `POST` | `/api/transactions` | Tạo transaction mới | ✅ |

### Health Check

| Method | Endpoint | Mô tả |
|---|---|---|
| `GET` | `/api/health` | Kiểm tra trạng thái server |

### Xác thực

Các endpoint có Auth ✅ yêu cầu gửi header:

```
Authorization: Bearer <jwt_token>
```

---

## 🛠️ Makefile Commands

```bash
# Phát triển
make run                              # Chạy development server
make build                            # Build binary → bin/myapp
make tidy                             # Dọn dẹp go.mod / go.sum
make test                             # Chạy tất cả test

# Migration
make migrate-up                       # Apply tất cả migration
make migrate-down                     # Rollback 1 bước
make migrate-reset                    # Rollback toàn bộ
make migrate-version                  # Xem version hiện tại
make migrate-create name=<tên>        # Tạo file migration mới
```

---

## 🗄️ Database Migration

Template dùng `golang-migrate` với SQL thuần — rõ ràng, dễ review, dễ rollback.

### Tạo migration mới

```bash
make migrate-create name=create_products_table
```

Tạo ra 2 file trong `database/migrations/`:
- `000004_create_products_table.up.sql`
- `000004_create_products_table.down.sql`

### Ví dụ

```sql
-- up.sql
CREATE TABLE products (
    id         BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name       VARCHAR(255)   NOT NULL,
    price      DECIMAL(10,2)  NOT NULL,
    created_at TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP      NULL
);

-- down.sql
DROP TABLE IF EXISTS products;
```

---

## ⚡ Background Jobs (Asynq)

Xử lý tác vụ nặng bất đồng bộ (gửi email, xử lý ảnh, webhook...) qua Redis queue.

### Khởi động worker

```bash
go run cmd/worker/main.go
```

Worker xử lý 3 queue theo độ ưu tiên: `critical (60%)` > `default (30%)` > `low (10%)`.

### Thêm job mới

**1. Định nghĩa job trong `jobs/`:**

```go
// jobs/send_notification.go
const TypeSendNotification = "notification:send"

type SendNotificationPayload struct {
    UserID  uint   `json:"user_id"`
    Message string `json:"message"`
}

func NewSendNotificationTask(p SendNotificationPayload) (*asynq.Task, error) {
    data, _ := json.Marshal(p)
    return asynq.NewTask(TypeSendNotification, data), nil
}

type SendNotificationHandler struct{}

func (h *SendNotificationHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
    var p SendNotificationPayload
    json.Unmarshal(t.Payload(), &p)
    // ... gửi notification
    return nil
}
```

**2. Đăng ký handler trong `cmd/worker/main.go`:**

```go
mux.Handle(jobs.TypeSendNotification, &jobs.SendNotificationHandler{})
```

**3. Enqueue từ service:**

```go
task, _ := jobs.NewSendNotificationTask(jobs.SendNotificationPayload{
    UserID:  user.ID,
    Message: "Welcome!",
})
client.Enqueue(task, asynq.Queue("default"))
```

---

## ⏰ Scheduled Tasks (Gocron)

Chạy các tác vụ định kỳ (cleanup, sync, report...).

### Khởi động scheduler

```bash
go run cmd/schedule/main.go
```

### Thêm task mới trong `schedule/scheduler.go`

```go
func Register(s gocron.Scheduler) {
    // Mỗi ngày lúc 2:00 AM
    s.NewJob(
        gocron.CronJob("0 2 * * *", false),
        gocron.NewTask(cleanupExpiredTokens),
    )

    // Mỗi 5 phút
    s.NewJob(
        gocron.DurationJob(5*time.Minute),
        gocron.NewTask(syncExternalData),
    )
}
```

---

## 🔧 Thêm resource mới (ví dụ: Product)

Quy trình chuẩn để thêm một CRUD resource:

1. **Model** — tạo `models/product.go`
2. **Migration** — `make migrate-create name=create_products_table`
3. **Repository** — tạo `app/repositories/product_repository.go` (implement interface)
4. **Đăng ký repository** — thêm vào `app/repositories/repositories.go`
5. **Service** — tạo `app/services/product_service.go` (implement interface)
6. **Đăng ký service** — thêm vào `app/services/services.go`
7. **Handler** — tạo `app/http/handlers/product_handler.go`
8. **Đăng ký handler** — thêm vào `app/http/handlers/handlers.go`
9. **Wiring DI** — cập nhật `app/bootstrap/app.go`
10. **Routes** — tạo `routes/product.go`, gọi trong `routes/routes.go`

---

## 🧪 Testing

```bash
make test
# hoặc
go test ./... -v
```

Mỗi layer đều có interface riêng → dễ mock để viết unit test độc lập.

---

## 📦 Dependencies chính

| Package | Version | Mục đích |
|---|---|---|
| `gofiber/fiber/v2` | v2.52 | HTTP framework |
| `gorm.io/gorm` | v1.25 | ORM |
| `golang-jwt/jwt/v5` | v5.2 | JWT authentication |
| `redis/go-redis/v9` | v9.4 | Redis client |
| `hibiken/asynq` | v0.24 | Background job queue |
| `go-co-op/gocron/v2` | v2.2 | Cron scheduler |
| `golang-migrate/migrate/v4` | v4.17 | Database migration |
| `go-playground/validator/v10` | v10.16 | Request validation |
| `spf13/viper` | v1.18 | Config management |
| `go.uber.org/zap` | v1.26 | Structured logging |
| `golang.org/x/crypto` | v0.18 | Password hashing (bcrypt) |

---

## 📄 License

[MIT](LICENSE) © huynhan147
