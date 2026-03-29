# bbgwire

一个简单的 Go 程序，用于从网站抓取文章数据并输出为 JSON 格式。

## 功能

- 抓取网站的前 50 篇文章
- 解析文章内容、链接和发布时间
- 输出 JSON 格式的结构化数据

## 安装

```bash
go install github.com/8treenet/bbgwire@latest
```

## 使用

```bash
bbgwire
```

## 输出示例

程序会输出 JSON 格式的文章列表：

```json
[
  {
    "Content": "文章标题内容",
    "URL": "https://example.com/article",
    "Published": "2026-03-29T00:00:00Z"
  }
]
```

## 数据结构

```go
type Article struct {
    Content   string // 文章内容
    URL       string // 文章链接
    Published string // 发布时间
}
```

## 依赖

- [github.com/8treenet/freedom/infra/requests](https://github.com/8treenet/freedom) - HTTP 请求库
- [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) - HTML 解析库

## 许可证

MIT License
