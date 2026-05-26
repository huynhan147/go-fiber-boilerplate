# myapp — Go Fiber (MySQL + Redis)

## Cấu trúc project

```
myapp/
├── main.go                         # Entry point + DI + server start
├── config/
│   ├── config.go                   # Load .env (Viper)
│   ├── database.go                 # Kết nối MySQL + AutoMigrate
│   └── redis.go                    # Kết nối Redis
├── models/
│   └── user.go                     # GORM model
├── routes/
│   └── api.go                      # Định nghĩa routes
├── pkg/
│   └── cache/
│       └── cache.go                # Redis cache helper (Get/Set/Delete)
├── app/
│   ├── http/
│   │   ├── handlers/               # Controllers
│   │   ├── requests/               # Form Requests + Validation
│   │   └── responses/              # API Resources
│   ├── services/                   # Business logic (+ Redis cache)
│   ├── repositories/               # Data access (MySQL via GORM)
│   └── middleware/
│       ├── auth.go                 # JWT + Redis blacklist check
│       ├── cors.go
│       └── error_handler.go
```

## Redis dùng cho gì?

| Feature | Key pattern | TTL |
|---|---|---|
| Cache user by ID | `user:{id}` | 5 phút |
| Token blacklist (logout) | `blacklist:{token}` | = JWT expire |

## Setup

```bash
cp .env.example .env   # sửa DB + Redis config
go mod tidy
make run
```

## API Endpoints

| Method | URL                  | Auth | Mô tả              |
|--------|----------------------|------|--------------------|
| GET    | /api/health          | ❌   | Health check       |
| POST   | /api/auth/login      | ❌   | Đăng nhập → JWT    |
| POST   | /api/auth/logout     | ✅   | Đăng xuất (blacklist token) |
| GET    | /api/auth/me         | ✅   | Thông tin bản thân |
| GET    | /api/users           | ✅   | Danh sách (phân trang) |
| POST   | /api/users           | ✅   | Tạo user           |
| GET    | /api/users/:id       | ✅   | Chi tiết (có cache)|
| PUT    | /api/users/:id       | ✅   | Cập nhật (xóa cache) |
| DELETE | /api/users/:id       | ✅   | Xóa (xóa cache)    |

## Ví dụ request

```bash
# Login
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password123"}'

# Logout (blacklist token)
curl -X POST http://localhost:3000/api/auth/logout \
  -H "Authorization: Bearer <token>"

# Tạo user
curl -X POST http://localhost:3000/api/users \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{"name":"John","email":"john@example.com","password":"password123"}'
```
