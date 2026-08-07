<script setup>
import {onMounted,ref} from 'vue'
import {useRoute,useRouter} from 'vue-router'
import {getPost} from '../../api/content'
import {createPost} from '../../api/admin'
import {demoPosts,normalizePost} from '../../utils/content'
import AdminPageHeader from '../../components/admin/AdminPageHeader.vue'
const route=useRoute(),router=useRouter(),saving=ref(false),message=ref('')
const form=ref({title:'',content:'',status:'draft',category:'未分类',tags:''})
const editing=route.params.id!==undefined
onMounted(async()=>{if(!editing)return;try{const p=normalizePost(await getPost(route.params.id));if(!p)throw 0;form.value={...form.value,...p,tags:(p.tags||[]).join(', ')}}catch{const p=demoPosts.find(x=>String(x.id)===String(route.params.id));if(p)form.value={...form.value,...p,tags:(p.tags||[]).join(', ')}}})
async function save(){
 saving.value=true;message.value=''
 try{
  if(!editing){await createPost({title:form.value.title,content:form.value.content,status:form.value.status});message.value='已创建文章';setTimeout(()=>router.push('/admin/posts'),500)}
  else{message.value='编辑 UI 已完成；你的后端 PUT /posts/:id 完成后即可真正保存。'}
 }catch(e){message.value=e.message}finally{saving.value=false}
}
</script>
<template>
<AdminPageHeader :title="editing?'编辑文章':'新建文章'" :description="editing?'修改内容并准备发布。':'先写清楚，再决定什么时候发布。'">
 <template #actions><button class="secondary-button pressable" @click="$router.back()">返回</button><button class="primary-button pressable" @click="save" :disabled="saving">{{saving?'保存中…':'保存'}}</button></template>
</AdminPageHeader>
<div class="editor-grid">
 <section class="editor surface"><input class="title" v-model="form.title" placeholder="文章标题"><textarea v-model="form.content" placeholder="开始写作…"></textarea></section>
 <aside class="settings surface"><h3>发布</h3><label>状态<select v-model="form.status"><option value="draft">草稿</option><option value="published">发布</option></select></label><label>分类<input v-model="form.category"></label><label>标签<input v-model="form.tags" placeholder="Go, Vue, CMS"></label><div class="hint">分类/标签后端接口尚未实现，这里先保留编辑能力。</div><p v-if="message" class="message">{{message}}</p></aside>
</div>
</template>
<style scoped>.editor-grid{display:grid;grid-template-columns:minmax(0,1fr) 300px;gap:14px}.editor,.settings{border-radius:20px}.editor{padding:18px}.title{width:100%;border:0;background:transparent;color:var(--text-1);font-size:clamp(1.8rem,4vw,3.2rem);font-weight:760;letter-spacing:-.045em;outline:0;padding:12px}.editor textarea{width:100%;min-height:58vh;resize:vertical;border:0;background:transparent;color:var(--text-1);outline:0;padding:20px 12px;font-size:1.04rem;line-height:1.85}.settings{padding:18px;height:max-content;position:sticky;top:82px}.settings h3{margin:0 0 18px}.settings label{display:grid;gap:7px;margin-bottom:14px;color:var(--text-2);font-size:.86rem}.settings input,.settings select{min-height:42px;border:1px solid var(--border);border-radius:12px;background:var(--surface-soft);color:var(--text-1);padding:0 11px}.hint{padding:12px;border-radius:12px;background:var(--accent-soft);color:var(--text-2);font-size:.8rem;line-height:1.6}.message{color:var(--success);font-size:.85rem}@media(max-width:850px){.editor-grid{grid-template-columns:1fr}.settings{position:static}}</style>
