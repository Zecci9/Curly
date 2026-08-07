<script setup>
import {computed,onMounted,ref} from 'vue';import {getPosts} from '../api/content';import {demoPosts,normalizePosts} from '../utils/content'
const posts=ref([]);onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const tags=computed(()=>{const m=new Map();posts.value.flatMap(p=>p.tags||[]).forEach(t=>m.set(t,(m.get(t)||0)+1));return [...m].map(([name,count])=>({name,count}))})
</script><template><section class="section"><div class="page-shell"><span class="eyebrow">Tags</span><h1 class="page-title">标签</h1><div class="cloud surface"><RouterLink v-for="t in tags" :key="t.name" class="chip pressable" :to="`/tags/${encodeURIComponent(t.name)}`">#{{t.name}} <small>{{t.count}}</small></RouterLink></div></div></section></template>
<style scoped>.cloud{margin-top:32px;padding:24px;border-radius:20px;display:flex;gap:10px;flex-wrap:wrap}.cloud small{margin-left:6px;color:var(--text-3)}</style>
