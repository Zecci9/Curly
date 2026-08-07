<script setup>
defineProps({
  label: { type: String, required: true },
  title: { type: String, default: '' },
})
</script>

<template>
  <button class="icon-button pressable" type="button" :aria-label="label" :title="title || label">
    <slot />
  </button>
</template>

<style scoped>
.icon-button {
  width: 42px;
  height: 42px;
  padding: 0;
  display: inline-grid;
  place-items: center;
  border-radius: 999px;
  border: 1px solid var(--border);
  background: var(--surface-soft);
  cursor: pointer;
}

.icon-button :deep(svg) {
  width: 19px;
  height: 19px;
}
</style>
