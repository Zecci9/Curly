<script setup>
import {computed,onMounted,ref} from 'vue';import {useRoute} from 'vue-router';import {getPosts} from '../api/content';import {demoPosts,normalizePosts} from '../utils/content';import PostCard from '../components/PostCard.vue'
const route=useRoute(),posts=ref([]);onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const name=computed(()=>decodeURIComponent(String(route.params.name))),filtered=computed(()=>posts.value.filter(p=>(p.category||'未分类')===name.value))
</script><template><section class="section"><div class="page-shell"><span class="eyebrow">Category</span><h1 class="page-title">{{name}}</h1><div class="grid"><PostCard v-for="p in filtered" :key="p.id" :post="p"/></div></div></section></template>
<style scoped>.grid{display:grid;grid-template-columns:1fr 1fr;gap:17px;margin-top:32px}@media(max-width:720px){.grid{grid-template-columns:1fr}}</style>
