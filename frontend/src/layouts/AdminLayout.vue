<script setup>
import {ref} from 'vue'
import AdminSidebar from '../components/admin/AdminSidebar.vue'
import AdminTopbar from '../components/admin/AdminTopbar.vue'
const sidebar=ref(false)
</script>
<template>
<div class="admin-layout">
 <AdminSidebar :open="sidebar" @close="sidebar=false"/>
 <div class="main">
  <AdminTopbar @menu="sidebar=true"/>
  <main><RouterView/></main>
 </div>
</div>
</template>
<style scoped>
.admin-layout{min-height:100vh}.main{margin-left:268px;padding:12px 16px 32px 0}.main>main{max-width:1500px;margin:0 auto}
@media(max-width:900px){.main{margin-left:0;padding:12px}}
</style>
