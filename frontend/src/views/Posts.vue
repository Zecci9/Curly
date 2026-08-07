<script setup>
import {computed,onMounted,ref} from 'vue'
import {getPosts} from '../api/content'
import {demoPosts,normalizePosts,uniqueCategories} from '../utils/content'
import PostCard from '../components/PostCard.vue'
const posts=ref([]),q=ref(''),category=ref('全部')
onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const categories=computed(()=>['全部',...uniqueCategories(posts.value)])
const filtered=computed(()=>posts.value.filter(p=>(category.value==='全部'||p.category===category.value)&&(!q.value||`${p.title} ${p.excerpt} ${(p.tags||[]).join(' ')}`.toLowerCase().includes(q.value.toLowerCase()))))
</script>
<template><section class="section"><div class="page-shell"><div class="intro"><span class="eyebrow">Archive</span><h1 class="page-title">文章</h1><p class="lead">全部公开内容。</p></div>
<div class="tools glass"><input v-model="q" placeholder="搜索标题、摘要或标签"><div><button v-for="c in categories" :key="c" class="chip pressable" :class="{active:category===c}" @click="category=c">{{c}}</button></div></div>
<div class="grid"><PostCard v-for="p in filtered" :key="p.id" :post="p"/></div></div></section></template>
<style scoped>.intro{padding:38px 0 30px}.tools{position:sticky;top:calc(var(--nav-height) + 12px);z-index:20;padding:10px;border-radius:18px;margin-bottom:20px}.tools input{width:100%;min-height:48px;border:0;border-radius:13px;background:var(--surface-soft);color:var(--text-1);padding:0 14px;outline:0}.tools div{display:flex;gap:7px;flex-wrap:wrap;margin-top:9px}.tools button{cursor:pointer}.active{background:var(--accent);color:var(--accent-text)}.grid{display:grid;grid-template-columns:1fr 1fr;gap:17px}@media(max-width:720px){.grid{grid-template-columns:1fr}}</style>
