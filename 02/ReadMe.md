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