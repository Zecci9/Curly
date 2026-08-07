<script setup>
defineProps({
  title: { type: String, default: '这里还没有内容' },
  description: { type: String, default: '等内容准备好后，它会出现在这里。' },
})
</script>

<template>
  <div class="empty-state surface">
    <div class="dot">·</div>
    <h3>{{ title }}</h3>
    <p>{{ description }}</p>
    <slot />
  </div>
</template>

<style scoped>
.dot {
  width: 44px;
  height: 44px;
  margin: 0 auto 15px;
  display: grid;
  place-items: center;
  border-radius: 14px;
  background: var(--accent);
  color: var(--accent-text);
  font-size: 2rem;
  line-height: 1;
}

h3 {
  margin: 0 0 8px;
}

p {
  margin: 0;
  color: var(--text-2);
  line-height: 1.65;
}
</style>
