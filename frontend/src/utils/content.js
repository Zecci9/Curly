export const demoPosts = [
  {
    id: 1,
    title: 'Curly：从一个接口开始',
    excerpt: '把内容、主题与插件拆开，让 CMS 从第一天就为变化留出空间。',
    content: `Curly 的目标不是把所有能力一次性写死，而是建立一个清晰的内容核心。

页面负责展示，主题负责视觉，插件负责扩展，API 负责数据。这样，当项目继续长大时，我们仍然知道每一部分该放在哪里。`,
    author: 'Curly',
    category: '开发',
    tags: ['Go', 'Vue', 'CMS'],
    status: 'published',
    created_at: '2026-08-06T10:00:00+08:00',
  },
  {
    id: 2,
    title: '设计不是把东西变漂亮',
    excerpt: '一个好的界面首先应该让人知道自己在哪、能做什么，以及下一步会发生什么。',
    content: `设计的价值来自秩序、反馈与可预测性。

漂亮的背景和动画可以加分，但它们应该服务内容，而不是抢走内容。Curly 的默认主题会保持克制，同时把主题接口开放给其他开发者。`,
    author: 'Curly',
    category: '设计',
    tags: ['Design', 'Theme'],
    status: 'published',
    created_at: '2026-08-05T20:10:00+08:00',
  },
  {
    id: 3,
    title: '为什么要给插件留位置',
    excerpt: '评论、SEO、统计甚至特殊页面，都不应该强迫核心代码不断膨胀。',
    content: `插件系统的第一步不是做插件市场，而是先定义边界。

Curly 前台先预留稳定的插槽和注册入口。未来评论、SEO、统计以及主题增强，都可以从这些位置进入。`,
    author: 'Curly',
    category: '架构',
    tags: ['Plugin', 'Architecture'],
    status: 'published',
    created_at: '2026-08-04T18:30:00+08:00',
  },
]

export function normalizePosts(payload) {
  const value = payload?.data ?? payload
  if (!Array.isArray(value)) return []
  return value.map((post) => ({
    category: '未分类',
    tags: [],
    author: 'Curly',
    excerpt: makeExcerpt(post.content),
    ...post,
  }))
}

export function normalizePost(payload) {
  const post = payload?.data ?? payload
  if (!post || typeof post !== 'object') return null
  return {
    category: '未分类',
    tags: [],
    author: 'Curly',
    excerpt: makeExcerpt(post.content),
    ...post,
  }
}

export function makeExcerpt(content = '', max = 100) {
  const plain = String(content).replace(/\s+/g, ' ').trim()
  return plain.length > max ? `${plain.slice(0, max)}…` : plain
}

export function formatDate(value) {
  if (!value) return '暂无日期'
  const date = new Date(value)
  if (Number.isNaN(date.getTime())) return value
  return new Intl.DateTimeFormat('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  }).format(date)
}

export function uniqueCategories(posts) {
  return [...new Set(posts.map((p) => p.category || '未分类'))]
}

export function uniqueTags(posts) {
  return [...new Set(posts.flatMap((p) => p.tags || []))]
}
