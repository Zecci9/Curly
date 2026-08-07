import { createRouter, createWebHistory } from 'vue-router'
import FrontLayout from '../layouts/FrontLayout.vue'
import AdminLayout from '../layouts/AdminLayout.vue'

const routes=[
 {
  path:'/',component:FrontLayout,children:[
   {path:'',name:'home',component:()=>import('../views/Home.vue')},
   {path:'posts',name:'posts',component:()=>import('../views/Posts.vue')},
   {path:'posts/:id',name:'post-detail',component:()=>import('../views/PostDetail.vue')},
   {path:'categories',component:()=>import('../views/Categories.vue')},
   {path:'categories/:name',component:()=>import('../views/CategoryDetail.vue')},
   {path:'tags',component:()=>import('../views/Tags.vue')},
   {path:'tags/:name',component:()=>import('../views/TagDetail.vue')},
   {path:'search',component:()=>import('../views/Search.vue')},
   {path:'about',component:()=>import('../views/About.vue')},
   {path:'login',name:'login',component:()=>import('../views/Login.vue')},
  ]
 },
 {
  path:'/admin',component:AdminLayout,meta:{requiresAuth:true},children:[
   {path:'',name:'admin-dashboard',component:()=>import('../views/admin/Dashboard.vue')},
   {path:'posts',component:()=>import('../views/admin/Posts.vue')},
   {path:'posts/new',component:()=>import('../views/admin/PostEditor.vue')},
   {path:'posts/:id/edit',component:()=>import('../views/admin/PostEditor.vue')},
   {path:'pages',component:()=>import('../views/admin/Pages.vue')},
   {path:'categories',component:()=>import('../views/admin/Categories.vue')},
   {path:'tags',component:()=>import('../views/admin/Tags.vue')},
   {path:'comments',component:()=>import('../views/admin/Comments.vue')},
   {path:'media',component:()=>import('../views/admin/Media.vue')},
   {path:'users',component:()=>import('../views/admin/Users.vue')},
   {path:'themes',component:()=>import('../views/admin/Themes.vue')},
   {path:'plugins',component:()=>import('../views/admin/Plugins.vue')},
   {path:'settings',component:()=>import('../views/admin/Settings.vue')},
   {path:'profile',component:()=>import('../views/admin/Profile.vue')},
  ]
 },
 {path:'/:pathMatch(.*)*',component:()=>import('../views/NotFound.vue')}
]

const router=createRouter({
 history:createWebHistory(),
 routes,
 scrollBehavior(to,from,saved){return saved||{top:0}}
})

router.beforeEach((to)=>{
 if(!to.matched.some(r=>r.meta.requiresAuth)) return true
 const token=localStorage.getItem('curly.token')
 if(token) return true
 return {name:'login',query:{redirect:to.fullPath}}
})

export default router
