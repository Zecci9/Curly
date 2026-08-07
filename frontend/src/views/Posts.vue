<script setup>
import { computed, onMounted, ref } from 'vue'
import { getPosts } from '../api/content'
import { demoPosts, normalizePosts, uniqueCategories } from '../utils/content'
import PostCard from '../components/PostCard.vue'

const posts = ref([])
const query = ref('')
const category = ref('全部')
const loading = ref(true)

onMounted(async () => {
  try {
    posts.value = normalizePosts(await getPosts())
    if (!posts.value.length) posts.value = demoPosts
  } catch {
    posts.value = demoPosts
  } finally {
    loading.value = false
  }
})

const categories = computed(() => ['全部', ...uniqueCategories(posts.value)])

const filtered = computed(() => {
  const keyword = query.value.trim().toLowerCase()
  return posts.value.filter((post) => {
    const matchCategory = category.value === '全部' || post.category === category.value
    const matchQuery =
      !keyword ||
      `${post.title} ${post.excerpt} ${(post.tags || []).join(' ')}`
        .toLowerCase()
        .includes(keyword)
    return matchCategory && matchQuery
  })
})
</script>

<template>
  <section class="section posts-page">
    <div class="page-shell">
      <div class="page-heading">
        <span class="eyebrow">Archive</span>
        <h1 class="page-title">文章</h1>
        <p class="lead">所有公开内容都应该容易被找到，而不是被埋在导航里。</p>
      </div>

      <div class="tools glass">
        <label class="search-box">
          <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8">
            <circle cx="11" cy="11" r="6.5" />
            <path d="m16 16 4.2 4.2" />
          </svg>
          <input v-model="query" type="search" placeholder="搜索标题、摘要或标签" />
        </label>

        <div class="filters">
          <button
            v-for="item in categories"
            :key="item"
            type="button"
            class="chip pressable"
            :class="{ active: category === item }"
            @click="category = item"
          >
            {{ item }}
          </button>
        </div>
      </div>

      <div v-if="loading" class="loading-state">正在加载文章…</div>

      <div v-else class="posts-grid">
        <PostCard v-for="post in filtered" :key="post.id" :post="post" />
      </div>

      <div v-if="!loading && !filtered.length" class="empty-state">
        没找到符合条件的文章。
      </div>
    </div>
  </section>
</template>

<style scoped>
.posts-page {
  min-height: 75vh;
}

.page-heading {
  max-width: 760px;
  padding: 40px 0 34px;
}

.page-heading .lead {
  margin-bottom: 0;
}

.tools {
  position: sticky;
  top: calc(var(--nav-height) + 12px);
  z-index: 20;
  padding: 12px;
  border-radius: 18px;
  margin-bottom: 24px;
}

.search-box {
  min-height: 48px;
  border-radius: 13px;
  background: var(--surface-soft);
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 0 14px;
}

.search-box svg {
  width: 18px;
  height: 18px;
  color: var(--text-3);
}

.search-box input {
  width: 100%;
  border: 0;
  outline: 0;
  color: var(--text-1);
  background: transparent;
}

.filters {
  margin-top: 10px;
  display: flex;
  flex-wrap: wrap;
  gap: 7px;
}

.filters button {
  cursor: pointer;
}

.filters .active {
  background: var(--accent);
  color: var(--accent-text);
  border-color: transparent;
}

.posts-grid {
  display: grid;
  grid-template-columns: repeat(2, minmax(0, 1fr));
  gap: 18px;
}

@media (max-width: 720px) {
  .posts-grid {
    grid-template-columns: 1fr;
  }
}
</style>
