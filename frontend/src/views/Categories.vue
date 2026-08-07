<script setup>
import {computed,onMounted,ref} from 'vue';import {getPosts} from '../api/content';import {demoPosts,normalizePosts} from '../utils/content'
const posts=ref([]);onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const items=computed(()=>{const m=new Map();posts.value.forEach(p=>m.set(p.category||'未分类',(m.get(p.category||'未分类')||0)+1));return [...m].map(([name,count])=>({name,count}))})
</script>
<template><section class="section"><div class="page-shell"><span class="eyebrow">Categories</span><h1 class="page-title">分类</h1><div class="grid"><RouterLink v-for="i in items" :key="i.name" class="surface pressable" :to="`/categories/${encodeURIComponent(i.name)}`"><small>{{i.count}} 篇</small><h2>{{i.name}}</h2><b>↗</b></RouterLink></div></div></section></template>
<style scoped>.grid{display:grid;grid-template-columns:repeat(3,1fr);gap:15px;margin-top:32px}.grid a{min-height:190px;padding:24px;border-radius:20px;text-decoration:none;display:flex;flex-direction:column;justify-content:end;position:relative}.grid small{color:var(--text-3)}h2{font-size:2rem;margin:8px 0 0;letter-spacing:-.04em}.grid b{position:absolute;right:24px;top:24px}@media(max-width:760px){.grid{grid-template-columns:1fr}}</style>
