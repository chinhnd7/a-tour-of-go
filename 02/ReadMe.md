# Lecture Notes:

## 1. package - module
package main chứa func main là nơi bắt đầu chạy chương trình go.

- Một ứng dụng Go (module go) có thể chứa nhiều package.
- Mỗi package nằm ở một thư mục khác nhau.
- Trong 1 package có thể chứa 1-nhiều file go. Mỗi file go hãy giữ dưới 200 dòng.
- Một ứng dụng Go có thể import nhiều module ngoài. Tương tự mỗi module ngoài có thể chứa 1-nhiều package

## 2. Hàm trả về error
Do Go không có try/catch nên các bạn luôn phải kiếm tra lỗi trả về từ mỗi func.

Quy ước tham số trả về cuối cùng luôn là error

Golang có cú pháp đặc biệt gọi là if assignment gồm 2 vế:
1. Assignment
2. Condition

```
> Luôn bắt lỗi, kiểm tra và xử lý.
```

## 3. Khai báo map
Syntax: 
```go
// dictionary := make(map[string]string)
dictionary := map[string]string{}      --- recommended
```

## 4. Golang cung cấp 2 kỹ thuật Test
1. Unit Test kiểm thử logic chạy có theo ý đồ lập trình viên hay không
2. Benchmark kiểm thử tốc độ thực thi

Chúng ta có thể viết file *test.go nằm trong cùng package hoặc package riêng.

File unit test hay benchmark luôn phải kết thúc bằng `test.go`

Hàm benchmark luôn phải bắt đầu bằng `func Benchmark`

## 5. Khai báo mảng sử dụng anonymous struct
Cách này vừa khai báo mảng chứa các struct. Không cần ghi rõ tên struct mà chỉ cần các thuộc tính bên trong struct. Sau khi khai báo xong, thì khởi tạo dữ liệu luôn