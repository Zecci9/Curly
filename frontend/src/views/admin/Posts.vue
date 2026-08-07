<script setup>
import {computed,onMounted,ref} from 'vue'
import {getPosts} from '../../api/content'
import {demoPosts,normalizePosts,formatDate} from '../../utils/content'
import AdminPageHeader from '../../components/admin/AdminPageHeader.vue'
import DataTable from '../../components/admin/DataTable.vue'
const posts=ref([]),q=ref(''),status=ref('all')
onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}})
const rows=computed(()=>posts.value.filter(p=>(status.value==='all'||p.status===status.value)&&(!q.value||p.title.toLowerCase().includes(q.value.toLowerCase()))))
const columns=[{key:'title',label:'标题'},{key:'status',label:'状态'},{key:'category',label:'分类'},{key:'created_at',label:'日期'}]
</script>
<template>
<AdminPageHeader title="文章" description="创建、编辑、发布与整理文章。">
 <template #actions><RouterLink class="primary-button pressable" to="/admin/posts/new">＋ 新建文章</RouterLink></template>
</AdminPageHeader>
<div class="toolbar glass"><input v-model="q" placeholder="搜索文章"><select v-model="status"><option value="all">全部状态</option><option value="published">已发布</option><option value="draft">草稿</option></select></div>
<DataTable :columns="columns" :rows="rows">
 <template #cell-title="{row}"><div class="title"><b>{{row.title}}</b><small>ID {{row.id}}</small></div></template>
 <template #cell-status="{row}"><span :class="['status',row.status]">{{row.status==='published'?'已发布':'草稿'}}</span></template>
 <template #cell-created_at="{row}">{{formatDate(row.created_at)}}</template>
 <template #actions="{row}"><RouterLink class="edit" :to="`/admin/posts/${row.id}/edit`">编辑</RouterLink></template>
</DataTable>
</template>
<style scoped>.toolbar{padding:9px;border-radius:17px;display:flex;gap:8px;margin-bottom:14px}.toolbar input,.toolbar select{min-height:43px;border:1px solid var(--border);background:var(--surface-soft);color:var(--text-1);border-radius:12px;padding:0 12px}.toolbar input{flex:1}.title{display:grid;gap:3px}.title b{color:var(--text-1)}.title small{color:var(--text-3)}.status{font-size:.76rem;padding:5px 9px;border-radius:99px;background:var(--surface-soft)}.status.published{color:var(--success)}.edit{text-decoration:none;font-weight:700;color:var(--text-1)}</style>
