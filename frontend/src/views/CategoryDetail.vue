<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { getPosts } from '../api/content'
import { demoPosts, normalizePosts } from '../utils/content'
import PostCard from '../components/PostCard.vue'
import EmptyState from '../components/EmptyState.vue'

const route = useRoute()
const posts = ref([])

onMounted(async () => {
  try {
    posts.value = normalizePosts(await getPosts())
    if (!posts.value.length) posts.value = demoPosts
  } catch {
    posts.value = demoPosts
  }
})

const name = computed(() => decodeURIComponent(String(route.params.name)))
const filtered = computed(() => posts.value.filter((post) => (post.category || '未分类') === name.value))
</script>

<template>
  <section class="section">
    <div class="page-shell">
      <div class="page-heading">
        <span class="eyebrow">Category</span>
        <h1 class="page-title">{{ name }}</h1>
        <p class="lead">{{ filtered.length }} 篇内容</p>
      </div>

      <div v-if="filtered.length" class="grid">
        <PostCard v-for="post in filtered" :key="post.id" :post="post" />
      </div>

      <EmptyState v-else title="这个分类还没有文章" />
    </div>
  </section>
</template>

<style scoped>
.page-heading {
  padding: 40px 0 34px;
}
.grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 18px;
}
@media (max-width: 720px) {
  .grid { grid-template-columns: 1fr; }
}
</style>
