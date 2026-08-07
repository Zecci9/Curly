<script setup>
import { ref } from 'vue'
import { useThemeStore } from '../stores/theme'
const open=ref(false)
const theme=useThemeStore()
const nav=[['/','首页'],['/posts','文章'],['/categories','分类'],['/tags','标签'],['/about','关于']]
</script>
<template>
<header class="header">
 <div class="page-shell">
  <div class="nav glass">
   <RouterLink class="brand pressable" to="/" @click="open=false"><b>C</b><span>Curly</span></RouterLink>
   <nav>
    <RouterLink v-for="item in nav" :key="item[0]" :to="item[0]">{{item[1]}}</RouterLink>
   </nav>
   <div class="actions">
    <RouterLink class="round pressable" to="/search" aria-label="搜索">⌕</RouterLink>
    <button class="round pressable" @click="theme.cycleMode()" aria-label="切换明暗">◐</button>
    <button class="round pressable menu" @click="open=!open" aria-label="菜单">☰</button>
   </div>
  </div>
  <Transition name="drop"><div v-if="open" class="mobile glass">
    <RouterLink v-for="item in nav" :key="item[0]" :to="item[0]" @click="open=false">{{item[1]}}</RouterLink>
    <RouterLink to="/login" @click="open=false">登录</RouterLink>
  </div></Transition>
 </div>
</header>
</template>
<style scoped>
.header{height:var(--nav-height);position:sticky;top:0;z-index:60;padding-top:10px;pointer-events:none}.nav{min-height:52px;padding:6px 8px 6px 10px;border-radius:18px;display:grid;grid-template-columns:auto 1fr auto;align-items:center;pointer-events:auto;box-shadow:var(--shadow-sm)}
.brand{display:flex;align-items:center;gap:9px;text-decoration:none;font-weight:780}.brand b{width:31px;height:31px;border-radius:10px;display:grid;place-items:center;background:var(--accent);color:var(--accent-text)}
nav{justify-self:center;display:flex;gap:4px}nav a{padding:9px 12px;border-radius:11px;text-decoration:none;color:var(--text-2);font-size:.91rem}nav a.router-link-active{color:var(--text-1);background:var(--surface-soft)}
.actions{display:flex;gap:6px}.round{width:40px;height:40px;border-radius:999px;border:1px solid var(--border);background:var(--surface-soft);display:grid;place-items:center;text-decoration:none;cursor:pointer}.menu{display:none}
.mobile{pointer-events:auto;margin-top:8px;border-radius:18px;padding:9px;display:grid}.mobile a{padding:13px;border-radius:12px;text-decoration:none;color:var(--text-2)}
.drop-enter-active,.drop-leave-active{transition:opacity .18s ease,transform .25s var(--ease-fluid);transform-origin:top right}.drop-enter-from,.drop-leave-to{opacity:0;transform:translateY(-7px) scale(.98)}
@media(min-width:761px){.mobile{display:none}}@media(max-width:760px){nav{display:none}.menu{display:grid}.nav{grid-template-columns:1fr auto}}
</style>
