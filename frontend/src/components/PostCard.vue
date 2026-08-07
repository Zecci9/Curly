<script setup>
import { formatDate } from '../utils/content'
defineProps({post:{type:Object,required:true}})
</script>
<template>
<article class="post-card surface pressable">
  <div class="meta"><span>{{post.category||'未分类'}}</span><span>{{formatDate(post.created_at)}}</span></div>
  <div>
    <h2><RouterLink :to="`/posts/${post.id}`">{{post.title}}</RouterLink></h2>
    <p>{{post.excerpt||'暂无摘要'}}</p>
  </div>
  <div class="bottom">
    <div><RouterLink v-for="tag in post.tags||[]" :key="tag" :to="`/tags/${encodeURIComponent(tag)}`">#{{tag}}</RouterLink></div>
    <RouterLink :to="`/posts/${post.id}`">阅读 ↗</RouterLink>
  </div>
</article>
</template>
<style scoped>
.post-card{min-height:260px;padding:24px;border-radius:var(--radius-md);display:flex;flex-direction:column;justify-content:space-between}
.meta,.bottom{display:flex;justify-content:space-between;gap:14px;color:var(--text-3);font-size:.84rem}
h2{font-size:clamp(1.45rem,3vw,2.25rem);line-height:1.08;letter-spacing:-.04em;margin:30px 0 12px}h2 a,.bottom a{text-decoration:none}
p{color:var(--text-2);line-height:1.7}.bottom>div{display:flex;gap:8px;flex-wrap:wrap}
</style>
