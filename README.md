# 🚀 Go Fiber Template

Một project template production-ready cho REST API viết bằng **Go** + **Fiber**, tích hợp sẵn MySQL, Redis, JWT Auth, Queue Worker, và Scheduler — giúp bạn khởi động dự án mới nhanh chóng mà không cần setup từ đầu.

---

## ✨ Tính năng

- **[Fiber v2](https://gofiber.io/)** — HTTP framework hiệu năng cao, cú pháp tương tự Express.js
- **[GORM](https://gorm.io/)** — ORM mạnh mẽ, hỗ trợ soft delete, migration
- **MySQL** — Database chính với connection pooling
- **[Redis](https://redis.io/)** — Cache layer và blacklist JWT token khi logout
- **JWT Authentication** — Đăng nhập / đăng xuất với token blacklist
- **[Asynq](https://github.com/hibiken/asynq)** — Background job queue chạy qua Redis
- **[Gocron](https://github.com/go-co-op/gocron)** — Scheduled tasks (cron jobs)
- **[golang-migrate](https://github.com/golang-migrate/migrate)** — Database migration có version
- **Dependency Injection** — Container pattern, dễ test và mở rộng
- **Layered Architecture** — Handler → Service → Repository rõ ràng
- **Validation** — Request validation với `go-playground/validator`
- **Viper** — Quản lý config từ file `.env`

---

## 🏗️ Kiến trúc

```
myapp/
├── main.go                     # Entry point — khởi động HTTP server
├── .env / .env.example         # Biến môi trường
├── Makefile                    # Các lệnh phổ biến
│
├── config/                     # Khởi tạo DB, Redis, load config
│   ├── config.go
│   ├── database.go
│   └── redis.go
│
├── app/
│   ├── bootstrap/              # Wiring toàn bộ dependencies (DI)
│   ├── container/              # Container struct chứa tất cả layers
│   ├── http/
│   │   ├── handlers/           # HTTP handlers (Controller layer)
│   │   ├── requests/           # Request structs + validation rules
│   │   └── responses/          # Response helpers (success, error, pagination...)
│   ├── services/               # Business logic
│   ├── repositories/           # Database queries (GORM)
│   └── middleware/             # Auth, CORS, Error handler
│
├── models/                     # GORM models
├── routes/                     # Đăng ký routes, gắn middleware
│
├── pkg/
│   └── cache/                  # Redis cache wrapper
│
├── database/
│   └── migrations/             # SQL migration files (up/down)
│
├── cmd/
│   ├── migrate/                # CLI tool chạy migration
│   ├── worker/                 # Background job worker (Asynq)
│   └── schedule/               # Cron job scheduler (Gocron)
│
├── jobs/                       # Job definitions và handlers
└── schedule/                   # Scheduler registration
```

### Luồng xử lý request

```
HTTP Request
    └─► Middleware (CORS, Auth, Logger)
            └─► Handler  (parse & validate request)
                    └─► Service  (business logic)
                            └─► Repository  (database / cache)
                                    └─► Response
```

---

## ⚙️ Yêu cầu

| Công cụ | Phiên bản |
|---------|-----------|
| Go      | >= 1.21   |
| MySQL   | >= 8.0    |
| Redis   | >= 6.0    |

---

## 🚀 Bắt đầu nhanh

### 1. Clone và cài đặt dependencies

```bash
git clone <your-repo-url> myapp
cd myapp
go mod tidy
```

### 2. Cấu hình môi trường

```bash
cp .env.example .env
```

Chỉnh sửa file `.env`:

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

### 3. Chạy migration

```bash
# Tạo database trước (nếu chưa có)
mysql -u root -p -e "CREATE DATABASE myapp_db CHARACTER SET utf8mb4;"

# Chạy toàn bộ migration
make migrate-up
```

### 4. Khởi động server

```bash
make run
# hoặc
go run main.go
```

Server sẽ chạy tại `http://localhost:3000`.

---

## 📡 API Endpoints

Tất cả routes đều có prefix `/api`.

### Auth

| Method | Endpoint         | Mô tả                       | Auth required |
|--------|------------------|-----------------------------|---------------|
| POST   | `/api/auth/login`  | Đăng nhập, trả về JWT token | ❌            |
| GET    | `/api/auth/me`     | Lấy thông tin user hiện tại | ✅            |
| POST   | `/api/auth/logout` | Đăng xuất, blacklist token  | ✅            |

### Users (CRUD)

| Method | Endpoint          | Mô tả              | Auth required |
|--------|-------------------|--------------------|---------------|
| GET    | `/api/users`        | Danh sách users    | ✅            |
| POST   | `/api/users`        | Tạo user mới       | ✅            |
| GET    | `/api/users/:id`    | Chi tiết user      | ✅            |
| PUT    | `/api/users/:id`    | Cập nhật user      | ✅            |
| DELETE | `/api/users/:id`    | Xóa user (soft)    | ✅            |

### Transactions

| Method | Endpoint              | Mô tả                  | Auth required |
|--------|-----------------------|------------------------|---------------|
| GET    | `/api/transactions`     | Danh sách transactions | ✅            |
| POST   | `/api/transactions`     | Tạo transaction mới    | ✅            |

### Health Check

| Method | Endpoint      | Mô tả          |
|--------|---------------|----------------|
| GET    | `/api/health`   | Kiểm tra server |

### Xác thực (Authentication)

Các endpoint có Auth required cần gửi header:

```
Authorization: Bearer <jwt_token>
```

---

## 🛠️ Makefile Commands

```bash
make run              # Chạy development server
make build            # Build binary ra bin/myapp
make tidy             # Dọn dẹp go.mod / go.sum
make test             # Chạy toàn bộ test

make migrate-up       # Chạy tất cả migration
make migrate-down     # Rollback 1 bước
make migrate-reset    # Rollback toàn bộ về ban đầu
make migrate-version  # Xem version migration hiện tại
make migrate-create name=tên_migration  # Tạo file migration mới
```

---

## 🗄️ Database Migration

Template sử dụng `golang-migrate` với file SQL thuần, giúp việc quản lý schema rõ ràng và có thể rollback.

### Tạo migration mới

```bash
make migrate-create name=create_products_table
```

Lệnh này tạo ra 2 file trong `database/migrations/`:
- `000004_create_products_table.up.sql` — SQL để apply
- `000004_create_products_table.down.sql` — SQL để rollback

### Ví dụ nội dung migration

```sql
-- up.sql
CREATE TABLE products (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- down.sql
DROP TABLE IF EXISTS products;
```

---

## ⚡ Background Jobs (Queue)

Template tích hợp **Asynq** để xử lý các tác vụ nặng bất đồng bộ (gửi email, xử lý ảnh, v.v.) thông qua Redis queue.

### Khởi động worker

```bash
go run cmd/worker/main.go
```

Worker xử lý 3 queue với độ ưu tiên khác nhau: `critical (60%)` > `default (30%)` > `low (10%)`.

### Tạo và dispatch job mới

**1. Định nghĩa job trong `jobs/`:**

```go
// jobs/send_notification.go
const TypeSendNotification = "notification:send"

type SendNotificationPayload struct {
    UserID  uint   `json:"user_id"`
    Message string `json:"message"`
}

func NewSendNotificationTask(payload SendNotificationPayload) (*asynq.Task, error) {
    data, _ := json.Marshal(payload)
    return asynq.NewTask(TypeSendNotification, data), nil
}

type SendNotificationHandler struct{}

func (h *SendNotificationHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
    var p SendNotificationPayload
    json.Unmarshal(t.Payload(), &p)
    // logic gửi notification...
    return nil
}
```

**2. Đăng ký handler trong `cmd/worker/main.go`:**

```go
mux.Handle(jobs.TypeSendNotification, jobs.NewSendNotificationHandler())
```

**3. Dispatch job từ service:**

```go
task, _ := jobs.NewSendNotificationTask(jobs.SendNotificationPayload{
    UserID:  user.ID,
    Message: "Welcome!",
})
client.Enqueue(task, asynq.Queue("default"))
```

---

## 📅 Scheduled Tasks (Cron)

Template tích hợp **Gocron** để chạy các tác vụ định kỳ (cleanup, report, sync...).

### Khởi động scheduler

```bash
go run cmd/schedule/main.go
```

### Thêm scheduled task

Chỉnh sửa `schedule/scheduler.go`:

```go
func Register(s gocron.Scheduler) {
    // Chạy mỗi ngày lúc 2:00 AM
    s.NewJob(
        gocron.CronJob("0 2 * * *", false),
        gocron.NewTask(cleanupExpiredTokens),
    )

    // Chạy mỗi 5 phút
    s.NewJob(
        gocron.DurationJob(5 * time.Minute),
        gocron.NewTask(syncExternalData),
    )
}
```

---

## 🔧 Mở rộng project

### Thêm một resource mới (ví dụ: Product)

**1. Tạo model** `models/product.go`

**2. Tạo migration** `make migrate-create name=create_products_table`

**3. Tạo repository** `app/repositories/product_repository.go` (implement interface)

**4. Đăng ký trong** `app/repositories/repositories.go`

**5. Tạo service** `app/services/product_service.go` (implement interface)

**6. Đăng ký trong** `app/services/services.go`

**7. Tạo handler** `app/http/handlers/product_handler.go`

**8. Đăng ký trong** `app/http/handlers/handlers.go`

**9. Wire trong** `app/bootstrap/app.go`

**10. Thêm routes** `routes/product.go` và gọi trong `routes/routes.go`

---

## 🧪 Test

```bash
make test
# hoặc
go test ./... -v
```

Mỗi layer (service, repository) đều có interface riêng — dễ dàng mock để unit test.

---

## 📦 Dependencies chính

| Package | Mục đích |
|---------|----------|
| `gofiber/fiber/v2` | HTTP framework |
| `gorm.io/gorm` | ORM |
| `golang-jwt/jwt/v5` | JWT authentication |
| `redis/go-redis/v9` | Redis client |
| `hibiken/asynq` | Background job queue |
| `go-co-op/gocron/v2` | Cron scheduler |
| `golang-migrate/migrate/v4` | Database migration |
| `go-playground/validator/v10` | Request validation |
| `spf13/viper` | Config management |
| `golang.org/x/crypto` | Password hashing (bcrypt) |

---

## 📄 License

MIT
