<script setup>
import AdminIcon from './AdminIcon.vue'
import { getPluginAdminItems } from '../../plugins/registry'
const props=defineProps({open:Boolean})
const emit=defineEmits(['close'])
const groups=[
 {label:'概览',items:[['/admin','dashboard','仪表盘']]},
 {label:'内容',items:[['/admin/posts','post','文章'],['/admin/pages','page','页面'],['/admin/categories','folder','分类'],['/admin/tags','tag','标签'],['/admin/comments','comment','评论'],['/admin/media','media','媒体']]},
 {label:'外观',items:[['/admin/themes','theme','主题'],['/admin/plugins','plugin','插件']]},
 {label:'系统',items:[['/admin/users','user','用户'],['/admin/settings','settings','设置']]},
]
const pluginItems=getPluginAdminItems()
</script>
<template>
<aside class="sidebar glass" :class="{open}">
 <div class="brand-row"><RouterLink to="/" class="brand"><b>C</b><span>Curly</span></RouterLink><span class="badge">Admin</span></div>
 <nav>
  <section v-for="group in groups" :key="group.label"><div class="label">{{group.label}}</div>
   <RouterLink v-for="item in group.items" :key="item[0]" :to="item[0]" @click="emit('close')"><AdminIcon :name="item[1]"/><span>{{item[2]}}</span></RouterLink>
  </section>
  <section v-if="pluginItems.length"><div class="label">插件</div><RouterLink v-for="item in pluginItems" :key="item.path" :to="item.path" @click="emit('close')"><AdminIcon name="plugin"/><span>{{item.label}}</span></RouterLink></section>
 </nav>
 <div class="bottom"><RouterLink to="/">← 查看网站</RouterLink></div>
</aside>
<div v-if="open" class="scrim" @click="emit('close')"></div>
</template>
<style scoped>
.sidebar{position:fixed;left:12px;top:12px;bottom:12px;width:244px;border-radius:22px;padding:14px;z-index:80;display:flex;flex-direction:column;box-shadow:var(--shadow-sm)}
.brand-row{display:flex;align-items:center;justify-content:space-between;padding:4px 4px 18px}.brand{display:flex;align-items:center;gap:9px;text-decoration:none;font-weight:790}.brand b{width:31px;height:31px;border-radius:10px;background:var(--accent);color:var(--accent-text);display:grid;place-items:center}.badge{font-size:.72rem;color:var(--text-3)}
nav{overflow:auto;display:grid;gap:16px}section{display:grid;gap:3px}.label{padding:0 10px 6px;color:var(--text-3);font-size:.72rem;font-weight:700;letter-spacing:.07em;text-transform:uppercase}
nav a{min-height:42px;border-radius:12px;padding:0 10px;display:flex;align-items:center;gap:11px;text-decoration:none;color:var(--text-2);font-size:.92rem;font-weight:600;transition:background .15s ease,color .15s ease,transform .15s ease}
nav a:hover{background:var(--surface-soft);color:var(--text-1)}nav a:active{transform:scale(.98)}nav a.router-link-exact-active,nav a.router-link-active{background:var(--surface-soft);color:var(--text-1)}
nav svg{width:18px;height:18px}.bottom{margin-top:auto;padding-top:14px;border-top:1px solid var(--border)}.bottom a{display:block;padding:9px;color:var(--text-2);text-decoration:none;font-size:.88rem}
.scrim{display:none}
@media(max-width:900px){.sidebar{transform:translateX(calc(-100% - 24px));transition:transform .3s var(--ease-fluid)}.sidebar.open{transform:translateX(0)}.scrim{display:block;position:fixed;inset:0;background:rgba(0,0,0,.28);z-index:70;backdrop-filter:blur(3px)}}
</style>
