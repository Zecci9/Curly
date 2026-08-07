<script setup>
import { computed, onMounted, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { getPost } from '../api/content'
import { demoPosts, normalizePost, formatDate } from '../utils/content'
import PluginSlot from '../components/PluginSlot.vue'

const route = useRoute()
const post = ref(null)
const loading = ref(true)
const error = ref('')

async function load() {
  loading.value = true
  error.value = ''

  try {
    post.value = normalizePost(await getPost(route.params.id))
    if (!post.value) throw new Error('文章不存在')
  } catch (err) {
    post.value = demoPosts.find((item) => String(item.id) === String(route.params.id)) || null
    if (!post.value) error.value = err.message || '文章不存在'
  } finally {
    loading.value = false
  }
}

onMounted(load)
watch(() => route.params.id, load)

const paragraphs = computed(() =>
  String(post.value?.content || '')
    .split(/\n{2,}/)
    .map((item) => item.trim())
    .filter(Boolean)
)
</script>

<template>
  <section class="article-page">
    <div v-if="loading" class="page-shell loading-state">正在打开文章…</div>

    <div v-else-if="error" class="page-shell error-state surface">
      {{ error }}
    </div>

    <template v-else-if="post">
      <PluginSlot name="post.before" :context="{ post }" />

      <header class="article-hero">
        <div class="page-shell article-head-inner">
          <RouterLink class="back-link" to="/posts">← 返回文章</RouterLink>

          <div class="article-meta">
            <RouterLink :to="`/categories/${encodeURIComponent(post.category || '未分类')}`">
              {{ post.category || '未分类' }}
            </RouterLink>
            <span>{{ formatDate(post.created_at) }}</span>
            <span>{{ post.author || 'Curly' }}</span>
          </div>

          <h1>{{ post.title }}</h1>
          <p v-if="post.excerpt" class="article-deck">{{ post.excerpt }}</p>

          <div class="tags">
            <RouterLink
              v-for="tag in post.tags || []"
              :key="tag"
              class="chip pressable"
              :to="`/tags/${encodeURIComponent(tag)}`"
            >
              #{{ tag }}
            </RouterLink>
          </div>
        </div>
      </header>

      <article class="article-body">
        <div class="reading-column">
          <p v-for="(paragraph, index) in paragraphs" :key="index">
            {{ paragraph }}
          </p>

          <div class="endmark">C</div>
        </div>
      </article>

      <PluginSlot name="post.after" :context="{ post }" />

      <section class="article-next">
        <div class="page-shell">
          <RouterLink class="next-card surface pressable" to="/posts">
            <span class="eyebrow">Continue</span>
            <strong>继续浏览全部文章</strong>
            <span>→</span>
          </RouterLink>
        </div>
      </section>
    </template>
  </section>
</template>

<style scoped>
.article-page {
  min-height: 75vh;
}

.article-hero {
  padding: 62px 0 54px;
  border-bottom: 1px solid var(--border);
}

.article-head-inner {
  max-width: 980px;
}

.back-link {
  display: inline-block;
  margin-bottom: 46px;
  text-decoration: none;
  color: var(--text-2);
}

.article-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 12px 18px;
  color: var(--text-3);
  font-size: .88rem;
}

.article-meta a {
  color: var(--text-1);
  text-decoration: none;
  font-weight: 680;
}

h1 {
  max-width: 920px;
  margin: 20px 0;
  font-size: clamp(2.7rem, 8vw, 6rem);
  line-height: .98;
  letter-spacing: -0.06em;
}

.article-deck {
  max-width: 760px;
  color: var(--text-2);
  font-size: clamp(1.05rem, 2.4vw, 1.35rem);
  line-height: 1.75;
}

.tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-top: 26px;
}

.article-body {
  padding: 70px 20px;
}

.reading-column {
  max-width: var(--reading-width);
  margin: 0 auto;
  font-family: ui-serif, "Songti SC", "Noto Serif CJK SC", serif;
}

.reading-column p {
  margin: 0 0 1.55em;
  font-size: clamp(1.05rem, 2vw, 1.18rem);
  line-height: 1.95;
  letter-spacing: .01em;
}

.endmark {
  width: 42px;
  height: 42px;
  margin-top: 50px;
  display: grid;
  place-items: center;
  border-radius: 13px;
  background: var(--accent);
  color: var(--accent-text);
  font-family: system-ui, sans-serif;
  font-weight: 800;
}

.article-next {
  padding: 20px 0 80px;
}

.next-card {
  min-height: 135px;
  padding: 28px;
  border-radius: var(--radius-md);
  display: grid;
  grid-template-columns: 1fr auto;
  gap: 10px;
  text-decoration: none;
}

.next-card .eyebrow {
  grid-column: 1 / -1;
}

.next-card strong {
  font-size: clamp(1.3rem, 3vw, 2rem);
  letter-spacing: -0.03em;
}
</style>
