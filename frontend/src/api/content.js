import request from './request'
export const getPosts=()=>request.get('/posts')
export const getPost=(id)=>request.get(`/posts/${id}`)
