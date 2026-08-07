<script setup>
import { computed, onMounted, ref } from 'vue'
import { getPosts } from '../api/content'
import { demoPosts, normalizePosts } from '../utils/content'
import PostCard from '../components/PostCard.vue'

const query = ref('')
const posts = ref([])

onMounted(async () => {
  try {
    posts.value = normalizePosts(await getPosts())
    if (!posts.value.length) posts.value = demoPosts
  } catch {
    posts.value = demoPosts
  }
})

const results = computed(() => {
  const key = query.value.trim().toLowerCase()
  if (!key) return []
  return posts.value.filter((post) =>
    `${post.title} ${post.content} ${post.excerpt} ${post.category} ${(post.tags || []).join(' ')}`
      .toLowerCase()
      .includes(key)
  )
})
</script>

<template>
  <section class="search-page">
    <div class="page-shell search-shell">
      <span class="eyebrow">Search</span>
      <h1 class="page-title">找到你要的内容。</h1>

      <label class="big-search surface">
        <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
          <circle cx="11" cy="11" r="6.5" />
          <path d="m16 16 4.2 4.2" />
        </svg>
        <input v-model="query" autofocus type="search" placeholder="输入标题、正文、分类或标签" />
      </label>

      <div v-if="query" class="result-count">
        {{ results.length }} 个结果
      </div>

      <div class="results">
        <PostCard v-for="post in results" :key="post.id" :post="post" />
      </div>
    </div>
  </section>
</template>

<style scoped>
.search-page {
  min-height: 72vh;
  padding: 88px 0;
}

.search-shell {
  max-width: 920px;
}

.big-search {
  margin-top: 34px;
  min-height: 76px;
  padding: 0 22px;
  border-radius: 22px;
  display: flex;
  align-items: center;
  gap: 15px;
}

.big-search svg {
  width: 24px;
  height: 24px;
  color: var(--text-3);
}

.big-search input {
  width: 100%;
  border: 0;
  outline: 0;
  background: transparent;
  color: var(--text-1);
  font-size: clamp(1rem, 2vw, 1.25rem);
}

.result-count {
  margin: 24px 0 12px;
  color: var(--text-3);
}

.results {
  display: grid;
  gap: 16px;
}
</style>
