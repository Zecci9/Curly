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

const categories = computed(() => {
  const map = new Map()
  for (const post of posts.value) {
    const name = post.category || '未分类'
    map.set(name, (map.get(name) || 0) + 1)
  }
  return [...map.entries()].map(([name, count]) => ({ name, count }))
})
</script>

<template>
  <section class="section">
    <div class="page-shell">
      <div class="page-heading">
        <span class="eyebrow">Categories</span>
        <h1 class="page-title">分类</h1>
        <p class="lead">用更稳定的主题脉络组织内容。</p>
      </div>

      <div class="category-grid">
        <RouterLink
          v-for="item in categories"
          :key="item.name"
          class="category-card surface pressable"
          :to="`/categories/${encodeURIComponent(item.name)}`"
        >
          <div>
            <span class="count">{{ item.count }} 篇</span>
            <h2>{{ item.name }}</h2>
          </div>
          <span class="arrow">↗</span>
        </RouterLink>
      </div>
    </div>
  </section>
</template>

<style scoped>
.page-heading {
  padding: 40px 0 34px;
}

.category-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.category-card {
  min-height: 200px;
  padding: 24px;
  border-radius: var(--radius-md);
  display: flex;
  align-items: end;
  justify-content: space-between;
  text-decoration: none;
}

.count {
  color: var(--text-3);
  font-size: .86rem;
}

h2 {
  margin: 12px 0 0;
  font-size: clamp(1.5rem, 3vw, 2.3rem);
  letter-spacing: -0.04em;
}

.arrow {
  font-size: 1.5rem;
}

@media (max-width: 760px) {
  .category-grid {
    grid-template-columns: 1fr;
  }
}
</style>
