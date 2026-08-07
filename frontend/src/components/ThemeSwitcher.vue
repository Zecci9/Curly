<script setup>
import { useThemeStore } from '../stores/theme'

const theme = useThemeStore()
</script>

<template>
  <div class="theme-switcher">
    <div class="label">外观</div>

    <label>
      <span>主题</span>
      <select :value="theme.themeId" @change="theme.setTheme($event.target.value)">
        <option v-for="item in theme.themes" :key="item.id" :value="item.id">
          {{ item.name }}
        </option>
      </select>
    </label>

    <label>
      <span>明暗</span>
      <select :value="theme.mode" @change="theme.setMode($event.target.value)">
        <option value="system">跟随系统</option>
        <option value="light">浅色</option>
        <option value="dark">深色</option>
      </select>
    </label>
  </div>
</template>

<style scoped>
.theme-switcher {
  display: grid;
  gap: 10px;
  min-width: 220px;
}

.label {
  font-weight: 700;
}

label {
  display: grid;
  grid-template-columns: 56px 1fr;
  gap: 10px;
  align-items: center;
  color: var(--text-2);
  font-size: 0.9rem;
}

select {
  width: 100%;
  border-radius: 12px;
  border: 1px solid var(--border);
  background: var(--surface-soft);
  color: var(--text-1);
  padding: 10px 12px;
}
</style>
