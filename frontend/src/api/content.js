import request from './request'

export function getPosts() {
  return request.get('/posts')
}

export function getPost(id) {
  return request.get(`/posts/${id}`)
}

/*
 * 后端未来建议补充：
 *
 * GET /categories
 * GET /categories/:slug/posts
 * GET /tags
 * GET /tags/:slug/posts
 * GET /search?q=...
 *
 * 页面层不直接写死请求地址，后续只改这里即可。
 */
