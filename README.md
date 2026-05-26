# WinDivert Go Bindings

基于 [WinDivert](https://reqrypt.org/windivert.html) 的 Go 语言绑定，**纯 Go 实现，无 CGO 依赖**。

## 特性

- **🚀 无 CGO**: 纯 Go 实现
- **DLL 嵌入**: `WinDivert64.dll` 嵌入到 Go 二进制
- **网络包拦截**: Windows 用户态网络包拦截

## 使用方法

```go
package main

import "github.com/ddkwork/WinDivert"

func main() {
    w := &windivert.WinDivert{}
    handle := w.Open("tcp.DstPort == 80", 0, 0)
    defer w.Close(handle)
    // 拦截和修改网络包...
}
```

## 测试

```bash
go test -v
```

## 许可证

MIT License
