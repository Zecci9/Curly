import request from './request'

export function login(payload) {
  return request.post('/users/login', payload)
}

export function getMe() {
  return request.get('/users/me')
}
