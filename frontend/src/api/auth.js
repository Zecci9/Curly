import request from './request'
export const login=(payload)=>request.post('/users/login',payload)
export const getMe=()=>request.get('/users/me')
