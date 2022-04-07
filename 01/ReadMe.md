### Chú ý buổi 1
Có 3 loại module:
1. Module chuẩn của Google
2. Public module chia sẻ cho mọi user
3. Private module chỉ cho phép những nhân viên hoặc thành viên giới hạn được xem.

Đã phát hành module là chia sẻ toàn bộ mã nguồn

Để debug thì cần phải tạo module bằng lệnh `go mod init`

Đây là lỗi khi debug mà chưa tạo module

```
Build Error: go build -o d:\golang\01\__debug_bin.exe -gcflags all=-N -l .
go: go.mod file not found in current directory or any parent directory; see 'go help modules' (exit status 1)
```

Tạo module
```
go mod init lab01
go mod tidy
```

`go mod tidy` dọn dẹp lại các modules

### 1.1 Khai báo biến:
Có các cách sau:

Khai báo + gán trong 1 lệnh. Recommended
```go
height := 1.76
```

Khai báo và Gán tách thành 2 lệnh. Not Recommended
```go
var height float64
height = 1.76
```

Tên hàm bắt đầu bằng chữ thường chỉ có phạm vi hoạt động trong nội bộ package (private function)
Tên hàm bắt đầu bằng chữ hoa có phạm vi hoạt động (scope) trong và ngoài package (public function)

Golang không có kế thừa, chỉ có interface
Golang không có private, public... nhưng có thể khai báo struct public bằng cách viết Hoa chữ cái đầu struct