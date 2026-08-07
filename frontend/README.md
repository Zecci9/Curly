# Curly Frontend v0.5.1

这一版包含 **完整前台 + 管理后台**，并重新处理了本地 HTTPS 后端连接。

## 为什么这版不会再把手机带到 `localhost:8080`

浏览器只请求：

```text
/api/v1/...
```

开发时由 Vite 代理：

```text
手机/电脑浏览器
  ↓ HTTP http://电脑IP:5173/api/v1/...
Vite :5173
  ↓ HTTPS（secure:false，仅开发环境接受自签名证书）
Go :8080
```

因此：

- 手机不需要知道 Go 后端 IP。
- 浏览器不直接连接自签名 HTTPS 后端。
- 开发环境避免 CORS。
- 你的 Go 后端仍然可以保持 `PORT=8080` + `RunTLS(":" + cfg.Port, ...)`。

## 启动

1. 后端：

```bash
go run ./cmd/server
```

后端应监听：

```text
https://127.0.0.1:8080
```

2. 前端：

复制：

```text
.env.example -> .env
```

默认配置已经是：

```env
VITE_API_BASE_URL=/api/v1
VITE_BACKEND_TARGET=https://127.0.0.1:8080
```

然后：

```bash
npm install
npm run dev
```

电脑访问：

```text
http://localhost:5173
```

手机与电脑同一局域网时访问 Vite 输出的 Network 地址，例如：

```text
http://192.168.8.100:5173
```

API 仍然由电脑上的 Vite 转发到 Go。

## 页面

### 前台
- `/`
- `/posts`
- `/posts/:id`
- `/categories`
- `/categories/:name`
- `/tags`
- `/tags/:name`
- `/search`
- `/about`
- `/login`

### 后台
- `/admin` 仪表盘
- `/admin/posts` 文章
- `/admin/posts/new` 新建文章
- `/admin/posts/:id/edit` 编辑文章
- `/admin/pages` 页面
- `/admin/categories` 分类
- `/admin/tags` 标签
- `/admin/comments` 评论
- `/admin/media` 媒体
- `/admin/users` 用户
- `/admin/themes` 主题
- `/admin/plugins` 插件
- `/admin/settings` 设置
- `/admin/profile` 个人资料

## 当前后端能力说明

已经可以接：登录、当前用户、文章列表、文章详情、创建文章。
页面、分类、标签、评论、媒体、用户管理、主题、插件和设置等后台 API 还需要 Go 后端继续实现，因此对应页面暂时使用演示数据或接口占位。

前端的 `/admin` 路由守卫只是体验层。真正后台权限必须由 Go 的认证/RBAC Middleware 判断。
