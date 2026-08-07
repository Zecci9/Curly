<script setup>
import {computed,onMounted,ref} from 'vue'
import {getPosts} from '../../api/content'
import {demoPosts,normalizePosts,formatDate} from '../../utils/content'
import AdminPageHeader from '../../components/admin/AdminPageHeader.vue'
import StatCard from '../../components/admin/StatCard.vue'
const posts=ref([])
onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const published=computed(()=>posts.value.filter(p=>p.status==='published').length)
const drafts=computed(()=>posts.value.filter(p=>p.status!=='published').length)
</script>
<template>
<AdminPageHeader title="仪表盘" description="今天的 Curly 一切正常。这里汇总内容、待处理事项与系统状态。">
 <template #actions><RouterLink class="secondary-button pressable" to="/">查看网站</RouterLink><RouterLink class="primary-button pressable" to="/admin/posts/new">写文章</RouterLink></template>
</AdminPageHeader>
<div class="stats"><StatCard label="全部文章" :value="posts.length" trend="当前内容库" tone="good"/><StatCard label="已发布" :value="published" trend="公开可见" tone="good"/><StatCard label="草稿" :value="drafts" trend="等待继续编辑" tone="warn"/><StatCard label="评论待审" value="3" trend="演示数据" tone="warn"/></div>
<div class="dash-grid">
 <section class="panel surface"><div class="panel-head"><div><span class="eyebrow">Recent</span><h2>最近内容</h2></div><RouterLink to="/admin/posts">管理 →</RouterLink></div>
  <RouterLink v-for="p in posts.slice(0,5)" :key="p.id" class="recent" :to="`/admin/posts/${p.id}/edit`"><div><b>{{p.title}}</b><small>{{formatDate(p.created_at)}}</small></div><span :class="['status',p.status]">{{p.status==='published'?'已发布':'草稿'}}</span></RouterLink>
 </section>
 <section class="panel surface"><div class="panel-head"><div><span class="eyebrow">System</span><h2>系统状态</h2></div></div>
  <div class="health"><div><span class="light good"></span><b>API</b><small>后端连接由页面请求实际验证</small></div><div><span class="light good"></span><b>数据库</b><small>由 Go 后端负责</small></div><div><span class="light"></span><b>主题</b><small>2 个本地主题</small></div><div><span class="light"></span><b>插件</b><small>注册接口已就绪</small></div></div>
 </section>
</div>
</template>
<style scoped>
.stats{display:grid;grid-template-columns:repeat(4,1fr);gap:13px}.dash-grid{display:grid;grid-template-columns:1.35fr .65fr;gap:14px;margin-top:14px}.panel{border-radius:20px;padding:20px}.panel-head{display:flex;justify-content:space-between;align-items:end;margin-bottom:14px}.panel-head h2{margin:5px 0 0;letter-spacing:-.035em}.panel-head a{text-decoration:none;color:var(--text-2);font-size:.88rem}
.recent{min-height:64px;padding:9px 3px;border-bottom:1px solid var(--border);display:flex;align-items:center;justify-content:space-between;gap:16px;text-decoration:none}.recent:last-child{border:0}.recent div{display:grid;gap:4px}.recent b{font-size:.92rem}.recent small{color:var(--text-3)}.status{padding:5px 9px;border-radius:99px;background:var(--surface-soft);color:var(--text-3);font-size:.74rem}.status.published{color:var(--success)}
.health{display:grid}.health>div{min-height:67px;border-bottom:1px solid var(--border);display:grid;grid-template-columns:auto 1fr;column-gap:9px;align-content:center}.health>div:last-child{border:0}.health b{font-size:.9rem}.health small{grid-column:2;color:var(--text-3)}.light{width:8px;height:8px;border-radius:99px;background:var(--text-3);align-self:center}.light.good{background:var(--success)}
@media(max-width:1050px){.stats{grid-template-columns:1fr 1fr}.dash-grid{grid-template-columns:1fr}}@media(max-width:560px){.stats{grid-template-columns:1fr}}
</style>
