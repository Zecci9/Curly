<script setup>
import {computed,onMounted,ref} from 'vue';import {getPosts} from '../api/content';import {demoPosts,normalizePosts} from '../utils/content';import PostCard from '../components/PostCard.vue'
const q=ref(''),posts=ref([]);onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const results=computed(()=>{const k=q.value.trim().toLowerCase();return k?posts.value.filter(p=>`${p.title} ${p.content} ${p.category} ${(p.tags||[]).join(' ')}`.toLowerCase().includes(k)):[]})
</script><template><section class="section"><div class="page-shell search"><span class="eyebrow">Search</span><h1 class="page-title">找到你要的内容。</h1><input class="surface" v-model="q" autofocus placeholder="搜索标题、正文、分类或标签"><p class="muted" v-if="q">{{results.length}} 个结果</p><div class="results"><PostCard v-for="p in results" :key="p.id" :post="p"/></div></div></section></template>
<style scoped>.search{max-width:900px}.search>input{width:100%;min-height:72px;border-radius:22px;padding:0 20px;color:var(--text-1);font-size:1.15rem;margin:32px 0 10px;outline:0}.results{display:grid;gap:16px;margin-top:20px}</style>
