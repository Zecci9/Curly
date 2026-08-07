<script setup>
import {computed,onMounted,ref} from 'vue'
import {useRoute} from 'vue-router'
import {getPost} from '../api/content'
import {demoPosts,normalizePost,formatDate} from '../utils/content'
const route=useRoute(),post=ref(null)
onMounted(async()=>{try{post.value=normalizePost(await getPost(route.params.id));if(!post.value)throw 0}catch{post.value=demoPosts.find(p=>String(p.id)===String(route.params.id))}})
const paragraphs=computed(()=>String(post.value?.content||'').split(/\n{2,}/).filter(Boolean))
</script>
<template><section v-if="post" class="article"><header><div class="page-shell h"><RouterLink to="/posts">← 返回文章</RouterLink><div class="meta"><span>{{post.category}}</span><span>{{formatDate(post.created_at)}}</span><span>{{post.author}}</span></div><h1>{{post.title}}</h1><p>{{post.excerpt}}</p></div></header><article><div class="reading"><p v-for="(p,i) in paragraphs" :key="i">{{p}}</p><div class="end">C</div></div></article></section><section v-else class="section"><div class="page-shell surface empty">文章不存在。</div></section></template>
<style scoped>header{padding:56px 0;border-bottom:1px solid var(--border)}.h{max-width:980px}.h>a{text-decoration:none;color:var(--text-2)}.meta{display:flex;gap:16px;flex-wrap:wrap;color:var(--text-3);font-size:.86rem;margin-top:42px}.h h1{font-size:clamp(2.7rem,8vw,5.8rem);line-height:.98;letter-spacing:-.06em;margin:18px 0}.h p{max-width:760px;color:var(--text-2);font-size:1.15rem;line-height:1.75}article{padding:68px 20px}.reading{max-width:var(--reading-width);margin:auto;font-family:ui-serif,"Songti SC",serif}.reading p{font-size:1.1rem;line-height:1.95;margin:0 0 1.5em}.end{width:42px;height:42px;border-radius:13px;display:grid;place-items:center;background:var(--accent);color:var(--accent-text);font-family:system-ui;font-weight:800}.empty{padding:40px;border-radius:20px}</style>
