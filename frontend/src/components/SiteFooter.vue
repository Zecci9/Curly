<script setup>
import ThemeSwitcher from './ThemeSwitcher.vue'
import PluginSlot from './PluginSlot.vue'
</script>

<template>
  <footer class="footer">
    <PluginSlot name="footer.before" />

    <div class="page-shell footer-grid">
      <div>
        <div class="footer-brand">Curly</div>
        <p>一个开放、可主题化、可插件化的内容系统。</p>
      </div>

      <div class="footer-links">
        <RouterLink to="/posts">文章</RouterLink>
        <RouterLink to="/categories">分类</RouterLink>
        <RouterLink to="/tags">标签</RouterLink>
        <RouterLink to="/about">关于</RouterLink>
      </div>

      <ThemeSwitcher />
    </div>

    <div class="page-shell copyright">
      <span>© 2026 Curly</span>
      <span>Powered by Go + Vue</span>
    </div>
  </footer>
</template>

<style scoped>
.footer {
  padding: 64px 0 28px;
  border-top: 1px solid var(--border);
  background: linear-gradient(to bottom, transparent, var(--surface-soft));
}

.footer-grid {
  display: grid;
  grid-template-columns: 1.3fr .7fr 1fr;
  gap: 32px;
  align-items: start;
}

.footer-brand {
  font-weight: 780;
  font-size: 1.4rem;
  letter-spacing: -0.03em;
}

.footer p {
  max-width: 400px;
  line-height: 1.7;
  color: var(--text-2);
}

.footer-links {
  display: grid;
  gap: 10px;
}

.footer-links a {
  color: var(--text-2);
  text-decoration: none;
}

.copyright {
  margin-top: 36px;
  padding-top: 20px;
  border-top: 1px solid var(--border);
  display: flex;
  justify-content: space-between;
  color: var(--text-3);
  font-size: 0.85rem;
}

@media (max-width: 780px) {
  .footer-grid {
    grid-template-columns: 1fr;
  }
}
</style>
