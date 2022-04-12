# Lesson 03

## 1. Struct
Golang không có class chỉ có Struct
- Struct không có tính kế thừa
- Go tổ hợp các struct (composition) thay vì kế thừa (inheritance)
- Struct Golang không có method nhưng có khái niệm receiver. Có 2 loại receiver là Pointer Receiver và Value Receiver. Receiver bản chất là `func`, được viết bên ngoài struct.

## 2. Pointer
Syntax:
`test : = new(int) // <=> var test *int`
