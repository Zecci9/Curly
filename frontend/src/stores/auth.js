import { defineStore } from 'pinia'
import { login, getMe } from '../api/auth'

const KEY='curly.token'

export const useAuthStore=defineStore('auth',{
  state:()=>({ token:'', user:null, loading:false }),
  getters:{
    isLoggedIn:s=>Boolean(s.token),
    displayName:s=>s.user?.username||'管理员'
  },
  actions:{
    restore(){ this.token=localStorage.getItem(KEY)||'' },
    async signIn(username,password){
      this.loading=true
      try{
        const res=await login({username,password})
        const token=res?.data?.token??res?.token
        if(!token) throw new Error('后端没有返回 token')
        this.token=token
        localStorage.setItem(KEY,token)
        try{
          const me=await getMe()
          this.user=me?.data??me
        }catch{
          this.user={username}
        }
      }finally{ this.loading=false }
    },
    logout(){ this.token=''; this.user=null; localStorage.removeItem(KEY) }
  }
})
