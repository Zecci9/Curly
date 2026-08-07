import { createRouter, createWebHistory } from 'vue-router'
import FrontLayout from '../layouts/FrontLayout.vue'

const routes = [
  {
    path: '/',
    component: FrontLayout,
    children: [
      { path: '', name: 'home', component: () => import('../views/Home.vue') },
      { path: 'posts', name: 'posts', component: () => import('../views/Posts.vue') },
      { path: 'posts/:id', name: 'post-detail', component: () => import('../views/PostDetail.vue') },
      { path: 'categories', name: 'categories', component: () => import('../views/Categories.vue') },
      { path: 'categories/:name', name: 'category-detail', component: () => import('../views/CategoryDetail.vue') },
      { path: 'tags', name: 'tags', component: () => import('../views/Tags.vue') },
      { path: 'tags/:name', name: 'tag-detail', component: () => import('../views/TagDetail.vue') },
      { path: 'search', name: 'search', component: () => import('../views/Search.vue') },
      { path: 'about', name: 'about', component: () => import('../views/About.vue') },
      { path: 'login', name: 'login', component: () => import('../views/Login.vue') },
    ],
  },
  { path: '/:pathMatch(.*)*', component: () => import('../views/NotFound.vue') },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition
    return { top: 0 }
  },
})

export default router
