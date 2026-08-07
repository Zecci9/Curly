<script setup>
import { formatDate } from '../utils/content'

defineProps({
  post: { type: Object, required: true },
  featured: { type: Boolean, default: false },
})
</script>

<template>
  <article class="post-card surface pressable" :class="{ featured }">
    <div class="meta">
      <RouterLink class="category" :to="`/categories/${encodeURIComponent(post.category || '未分类')}`">
        {{ post.category || '未分类' }}
      </RouterLink>
      <span>{{ formatDate(post.created_at) }}</span>
    </div>

    <div class="copy">
      <h2>
        <RouterLink :to="`/posts/${post.id}`">{{ post.title }}</RouterLink>
      </h2>
      <p>{{ post.excerpt || '这篇文章暂时没有摘要。' }}</p>
    </div>

    <div class="bottom">
      <div class="tags">
        <RouterLink
          v-for="tag in (post.tags || []).slice(0, 3)"
          :key="tag"
          class="tag"
          :to="`/tags/${encodeURIComponent(tag)}`"
        >
          #{{ tag }}
        </RouterLink>
      </div>

      <RouterLink class="read" :to="`/posts/${post.id}`" :aria-label="`阅读 ${post.title}`">
        阅读
        <span aria-hidden="true">↗</span>
      </RouterLink>
    </div>
  </article>
</template>

<style scoped>
.post-card {
  min-height: 260px;
  padding: 25px;
  border-radius: var(--radius-md);
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.post-card.featured {
  min-height: 360px;
}

.meta,
.bottom {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  align-items: center;
  color: var(--text-3);
  font-size: 0.84rem;
}

.category,
.tag,
.read,
h2 a {
  text-decoration: none;
}

.category {
  color: var(--text-1);
  font-weight: 660;
}

.copy {
  margin: 35px 0;
}

h2 {
  margin: 0 0 12px;
  font-size: clamp(1.45rem, 3vw, 2.35rem);
  line-height: 1.08;
  letter-spacing: -0.04em;
}

p {
  margin: 0;
  color: var(--text-2);
  line-height: 1.75;
}

.tags {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.tag {
  color: var(--text-3);
}

.read {
  color: var(--text-1);
  font-weight: 690;
  white-space: nowrap;
}
</style>
