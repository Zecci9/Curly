<script setup>
import { onMounted, ref, computed } from 'vue'
import { getPosts } from '../api/content'
import { demoPosts, normalizePosts, uniqueCategories } from '../utils/content'
import PostCard from '../components/PostCard.vue'
import PluginSlot from '../components/PluginSlot.vue'

const posts = ref([])
const source = ref('api')

onMounted(async () => {
  try {
    posts.value = normalizePosts(await getPosts())
    if (!posts.value.length) throw new Error('empty')
  } catch {
    posts.value = demoPosts
    source.value = 'demo'
  }
})

const featured = computed(() => posts.value[0])
const rest = computed(() => posts.value.slice(1, 5))
const categories = computed(() => uniqueCategories(posts.value).slice(0, 6))
</script>

<template>
  <div>
    <section class="hero">
      <div class="page-shell hero-inner">
        <div class="hero-copy">
          <div class="eyebrow">Curly CMS · Open by design</div>
          <h1 class="display">内容应该自由地<br />生长。</h1>
          <p class="lead">
            一个为文章、主题与插件留出空间的内容系统。
            默认足够简洁，也允许开发者把它变成完全不同的样子。
          </p>

          <div class="hero-actions">
            <RouterLink class="primary-button pressable" to="/posts">开始阅读</RouterLink>
            <RouterLink class="secondary-button pressable" to="/about">了解 Curly</RouterLink>
          </div>
        </div>

        <div class="hero-object surface" aria-hidden="true">
          <div class="orbit orbit-a"></div>
          <div class="orbit orbit-b"></div>
          <div class="curly-c">C</div>
          <div class="float-card card-a">Theme</div>
          <div class="float-card card-b">Plugin</div>
          <div class="float-card card-c">Content</div>
        </div>
      </div>

      <PluginSlot name="home.hero.after" />
    </section>

    <PluginSlot name="home.feed.before" />

    <section class="section">
      <div class="page-shell">
        <div class="section-head">
          <div>
            <div class="eyebrow">Latest</div>
            <h2 class="section-title">最近发布</h2>
          </div>
          <RouterLink class="chip pressable" to="/posts">查看全部文章 →</RouterLink>
        </div>

        <div v-if="featured" class="feed-grid">
          <PostCard :post="featured" featured />
          <div class="side-feed">
            <PostCard v-for="post in rest" :key="post.id" :post="post" />
          </div>
        </div>

        <p v-if="source === 'demo'" class="demo-note">
          当前后端未返回文章，正在显示前台演示内容。
        </p>
      </div>
    </section>

    <section class="section categories-section">
      <div class="page-shell">
        <div class="section-head">
          <div>
            <div class="eyebrow">Explore</div>
            <h2 class="section-title">按主题发现内容</h2>
          </div>
        </div>

        <div class="category-row">
          <RouterLink
            v-for="category in categories"
            :key="category"
            class="category-card surface pressable"
            :to="`/categories/${encodeURIComponent(category)}`"
          >
            <span>{{ category }}</span>
            <span>↗</span>
          </RouterLink>
        </div>
      </div>
    </section>

    <section class="section manifesto">
      <div class="page-shell manifesto-card surface">
        <span class="eyebrow">Built to change</span>
        <h2>默认主题只是开始，不是边界。</h2>
        <p>
          页面只依赖稳定的数据与组件接口。主题可以重新定义视觉，插件可以插入新能力，
          而核心内容仍然保持清晰。
        </p>
        <div class="manifesto-links">
          <RouterLink to="/about">查看设计理念</RouterLink>
          <RouterLink to="/tags">浏览标签</RouterLink>
        </div>
      </div>
    </section>

    <PluginSlot name="home.feed.after" />
  </div>
</template>

<style scoped>
.hero {
  padding: 72px 0 48px;
  min-height: calc(100svh - var(--nav-height));
  display: flex;
  align-items: center;
}

.hero-inner {
  display: grid;
  grid-template-columns: minmax(0, 1.25fr) minmax(340px, .75fr);
  gap: 70px;
  align-items: center;
}

.hero-copy {
  max-width: 780px;
}

.hero-copy .lead {
  max-width: 660px;
  margin: 28px 0 0;
}

.hero-actions {
  margin-top: 30px;
  display: flex;
  gap: 10px;
  flex-wrap: wrap;
}

.hero-object {
  min-height: 480px;
  border-radius: 40px;
  position: relative;
  overflow: hidden;
  display: grid;
  place-items: center;
  background:
    radial-gradient(circle at 50% 35%, var(--accent-soft), transparent 42%),
    var(--surface);
}

.curly-c {
  width: 145px;
  height: 145px;
  display: grid;
  place-items: center;
  border-radius: 42px;
  color: var(--accent-text);
  background: var(--accent);
  font-size: 5rem;
  font-weight: 800;
  letter-spacing: -0.08em;
  box-shadow: 0 30px 90px var(--accent-shadow);
  z-index: 3;
}

.orbit {
  position: absolute;
  border: 1px solid var(--border);
  border-radius: 50%;
}

.orbit-a {
  width: 320px;
  height: 320px;
}

.orbit-b {
  width: 420px;
  height: 420px;
}

.float-card {
  position: absolute;
  padding: 12px 15px;
  border-radius: 15px;
  background: var(--glass);
  border: 1px solid var(--glass-border);
  backdrop-filter: blur(16px);
  box-shadow: var(--shadow-sm);
  color: var(--text-2);
  font-size: .86rem;
  font-weight: 700;
}

.card-a { transform: translate(118px, -116px) rotate(5deg); }
.card-b { transform: translate(-132px, 108px) rotate(-6deg); }
.card-c { transform: translate(130px, 126px) rotate(3deg); }

.feed-grid {
  display: grid;
  grid-template-columns: minmax(0, 1.1fr) minmax(0, .9fr);
  gap: 18px;
}

.side-feed {
  display: grid;
  grid-template-columns: 1fr;
  gap: 18px;
}

.demo-note {
  margin: 14px 0 0;
  color: var(--text-3);
  font-size: .82rem;
}

.category-row {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 14px;
}

.category-card {
  min-height: 110px;
  padding: 22px;
  border-radius: var(--radius-md);
  display: flex;
  justify-content: space-between;
  align-items: end;
  text-decoration: none;
  font-size: 1.12rem;
  font-weight: 700;
}

.manifesto-card {
  padding: clamp(28px, 6vw, 70px);
  border-radius: var(--radius-lg);
}

.manifesto h2 {
  max-width: 880px;
  margin: 18px 0;
  font-size: clamp(2.2rem, 6vw, 5rem);
  line-height: .98;
  letter-spacing: -0.055em;
}

.manifesto p {
  max-width: 690px;
  color: var(--text-2);
  line-height: 1.8;
}

.manifesto-links {
  display: flex;
  gap: 20px;
  margin-top: 28px;
}

.manifesto-links a {
  font-weight: 700;
}

@media (max-width: 900px) {
  .hero-inner,
  .feed-grid {
    grid-template-columns: 1fr;
  }

  .hero-object {
    min-height: 360px;
  }

  .category-row {
    grid-template-columns: 1fr 1fr;
  }
}

@media (max-width: 560px) {
  .hero {
    padding-top: 44px;
  }

  .hero-object {
    min-height: 300px;
    border-radius: 28px;
  }

  .curly-c {
    width: 110px;
    height: 110px;
    border-radius: 34px;
    font-size: 4rem;
  }

  .orbit-a {
    width: 230px;
    height: 230px;
  }

  .orbit-b {
    width: 290px;
    height: 290px;
  }

  .category-row {
    grid-template-columns: 1fr;
  }
}
</style>
