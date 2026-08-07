<script setup>
import { computed, onMounted, ref } from 'vue'
import { getPosts } from '../api/content'
import { demoPosts, normalizePosts } from '../utils/content'

const posts = ref([])

onMounted(async () => {
  try {
    posts.value = normalizePosts(await getPosts())
    if (!posts.value.length) posts.value = demoPosts
  } catch {
    posts.value = demoPosts
  }
})

const tags = computed(() => {
  const map = new Map()
  for (const post of posts.value) {
    for (const tag of post.tags || []) {
      map.set(tag, (map.get(tag) || 0) + 1)
    }
  }
  return [...map.entries()].map(([name, count]) => ({ name, count }))
})
</script>

<template>
  <section class="section">
    <div class="page-shell">
      <div class="page-heading">
        <span class="eyebrow">Tags</span>
        <h1 class="page-title">标签</h1>
        <p class="lead">比分类更轻、更自由的内容连接方式。</p>
      </div>

      <div class="tag-cloud surface">
        <RouterLink
          v-for="item in tags"
          :key="item.name"
          class="tag pressable"
          :to="`/tags/${encodeURIComponent(item.name)}`"
        >
          <span>#{{ item.name }}</span>
          <small>{{ item.count }}</small>
        </RouterLink>
      </div>
    </div>
  </section>
</template>

<style scoped>
.page-heading {
  padding: 40px 0 34px;
}

.tag-cloud {
  padding: 24px;
  border-radius: var(--radius-md);
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.tag {
  display: inline-flex;
  align-items: center;
  gap: 9px;
  min-height: 46px;
  padding: 0 16px;
  border-radius: 999px;
  background: var(--surface-soft);
  border: 1px solid var(--border);
  text-decoration: none;
}

.tag small {
  color: var(--text-3);
}
</style>
