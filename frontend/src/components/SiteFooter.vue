<script setup>import ThemeSwitcher from './ThemeSwitcher.vue'</script>
<template>
<footer class="footer">
 <div class="page-shell grid">
  <div><h3>Curly</h3><p>一个开放、可主题化、可插件化的内容系统。</p></div>
  <div class="links"><RouterLink to="/posts">文章</RouterLink><RouterLink to="/categories">分类</RouterLink><RouterLink to="/tags">标签</RouterLink><RouterLink to="/about">关于</RouterLink></div>
  <ThemeSwitcher/>
 </div>
 <div class="page-shell copy"><span>© 2026 Curly</span></div>
</footer>
</template>
<style scoped>
.footer{padding:62px 0 28px;border-top:1px solid var(--border);background:linear-gradient(to bottom,transparent,var(--surface-soft))}
.grid{display:grid;grid-template-columns:1.3fr .7fr 1fr;gap:30px}.footer h3{font-size:1.4rem;margin:0}.footer p{color:var(--text-2);line-height:1.7;max-width:420px}.links{display:grid;gap:10px}.links a,.copy a{text-decoration:none;color:var(--text-2)}
.copy{margin-top:34px;padding-top:18px;border-top:1px solid var(--border);display:flex;justify-content:space-between;color:var(--text-3);font-size:.84rem}
@media(max-width:780px){.grid{grid-template-columns:1fr}}
</style>
