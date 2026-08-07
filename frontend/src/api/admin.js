import request from './request'

// 已有/即将对接的文章管理接口
export const createPost=(payload)=>request.post('/posts',payload)

// 后端实现后直接启用：
// export const updatePost=(id,payload)=>request.put(`/posts/${id}`,payload)
// export const deletePost=(id)=>request.delete(`/posts/${id}`)
// export const getAdminUsers=()=>request.get('/admin/users')
// export const getAdminComments=()=>request.get('/admin/comments')
// export const getAdminThemes=()=>request.get('/admin/themes')
// export const getAdminPlugins=()=>request.get('/admin/plugins')
// export const getSettings=()=>request.get('/admin/settings')
// export const updateSettings=(payload)=>request.put('/admin/settings',payload)
