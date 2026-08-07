<script setup>
import { ref } from 'vue'
import { RouterLink } from 'vue-router'
import { useThemeStore } from '../stores/theme'
import IconButton from './IconButton.vue'

const open = ref(false)
const theme = useThemeStore()

const nav = [
  { to: '/', label: '首页' },
  { to: '/posts', label: '文章' },
  { to: '/categories', label: '分类' },
  { to: '/tags', label: '标签' },
  { to: '/about', label: '关于' },
]
</script>

<template>
  <header class="site-header">
    <div class="page-shell">
      <div class="nav glass">
        <RouterLink class="brand pressable" to="/" @click="open = false">
          <span class="brand-mark">C</span>
          <span>Curly</span>
        </RouterLink>

        <nav class="desktop-nav" aria-label="主导航">
          <RouterLink v-for="item in nav" :key="item.to" :to="item.to">
            {{ item.label }}
          </RouterLink>
        </nav>

        <div class="actions">
          <RouterLink class="search-link pressable" to="/search" aria-label="搜索">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <circle cx="11" cy="11" r="6.5" />
              <path d="m16 16 4.2 4.2" />
            </svg>
            <span>搜索</span>
          </RouterLink>

          <IconButton label="切换明暗模式" :title="`当前：${theme.mode}`" @click="theme.cycleMode()">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <path d="M20.4 14.7A8.3 8.3 0 0 1 9.3 3.6 8.5 8.5 0 1 0 20.4 14.7Z" />
            </svg>
          </IconButton>

          <IconButton class="menu-button" label="打开菜单" @click="open = !open">
            <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
              <path d="M5 8h14M5 16h14" />
            </svg>
          </IconButton>
        </div>
      </div>

      <Transition name="menu">
        <div v-if="open" class="mobile-menu glass">
          <RouterLink v-for="item in nav" :key="item.to" :to="item.to" @click="open = false">
            {{ item.label }}
          </RouterLink>
          <RouterLink to="/search" @click="open = false">搜索</RouterLink>
          <RouterLink to="/login" @click="open = false">登录</RouterLink>
        </div>
      </Transition>
    </div>
  </header>
</template>

<style scoped>
.site-header {
  position: sticky;
  top: 0;
  z-index: 50;
  height: var(--nav-height);
  padding-top: 10px;
  pointer-events: none;
}

.nav {
  min-height: 52px;
  padding: 6px 8px 6px 10px;
  display: grid;
  grid-template-columns: auto 1fr auto;
  align-items: center;
  border-radius: 18px;
  pointer-events: auto;
  box-shadow: 0 12px 44px rgba(18, 20, 23, 0.08);
}

.brand {
  display: inline-flex;
  align-items: center;
  gap: 9px;
  text-decoration: none;
  font-weight: 760;
  letter-spacing: -0.02em;
}

.brand-mark {
  width: 31px;
  height: 31px;
  display: grid;
  place-items: center;
  border-radius: 10px;
  color: var(--accent-text);
  background: var(--accent);
  font-size: 0.88rem;
}

.desktop-nav {
  justify-self: center;
  display: flex;
  align-items: center;
  gap: 5px;
}

.desktop-nav a {
  padding: 9px 12px;
  border-radius: 11px;
  text-decoration: none;
  color: var(--text-2);
  font-size: 0.92rem;
  transition: background-color 150ms ease, color 150ms ease, transform 150ms ease;
}

.desktop-nav a:hover,
.desktop-nav a.router-link-active {
  background: var(--surface-soft);
  color: var(--text-1);
}

.desktop-nav a:active {
  transform: scale(0.97);
}

.actions {
  display: flex;
  gap: 6px;
  align-items: center;
}

.search-link {
  min-height: 42px;
  padding: 0 14px;
  border: 1px solid var(--border);
  border-radius: 999px;
  display: inline-flex;
  align-items: center;
  gap: 7px;
  text-decoration: none;
  color: var(--text-2);
  background: var(--surface-soft);
  font-size: 0.9rem;
}

.search-link svg {
  width: 17px;
  height: 17px;
}

.menu-button {
  display: none;
}

.mobile-menu {
  pointer-events: auto;
  margin-top: 8px;
  border-radius: 18px;
  padding: 10px;
  display: grid;
  gap: 4px;
}

.mobile-menu a {
  padding: 13px 14px;
  border-radius: 12px;
  text-decoration: none;
  color: var(--text-2);
}

.mobile-menu a.router-link-active {
  color: var(--text-1);
  background: var(--surface-soft);
}

.menu-enter-active,
.menu-leave-active {
  transition: opacity 180ms ease, transform 240ms var(--spring-standard);
  transform-origin: top right;
}

.menu-enter-from,
.menu-leave-to {
  opacity: 0;
  transform: translateY(-8px) scale(0.98);
}

@media (min-width: 761px) {
  .mobile-menu {
    display: none;
  }
}

@media (max-width: 760px) {
  .desktop-nav,
  .search-link {
    display: none;
  }

  .menu-button {
    display: inline-grid;
  }

  .nav {
    grid-template-columns: 1fr auto;
  }
}
</style>
