<template>
  <section class="section">
    <div class="page-shell about">
      <header>
        <span class="eyebrow">About Curly</span>
        <h1 class="page-title">一个愿意给未来留位置的 CMS。</h1>
        <p class="lead">
          Curly 把内容作为核心，把主题与插件作为扩展层。
          默认体验应该清晰、克制、可靠，但不会阻止开发者创造完全不同的视觉与交互。
        </p>
      </header>

      <div class="principles">
        <article class="surface">
          <span>01</span>
          <h2>内容优先</h2>
          <p>页面结构围绕阅读、发现与导航组织，而不是围绕组件库组织。</p>
        </article>
        <article class="surface">
          <span>02</span>
          <h2>用户掌控</h2>
          <p>主题、明暗、动效偏好与未来布局选项应该允许用户自己决定。</p>
        </article>
        <article class="surface">
          <span>03</span>
          <h2>开放扩展</h2>
          <p>评论、SEO、统计、特殊页面与后台工具通过插件接口加入，而不是不断侵入核心。</p>
        </article>
      </div>

      <div class="extension surface">
        <div>
          <span class="eyebrow">Extension points</span>
          <h2>主题、插件、页面、设置、后台管理。</h2>
        </div>
        <p>
          当前前台已经提供主题注册表、插件注册表和页面路由边界。
          后续后端实现插件/主题元数据 API 后，可以直接把这些本地注册能力升级为动态安装能力。
        </p>
      </div>
    </div>
  </section>
</template>

<style scoped>
.about header {
  max-width: 850px;
  padding: 40px 0 48px;
}

.principles {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.principles article {
  padding: 26px;
  border-radius: var(--radius-md);
}

.principles article > span {
  color: var(--text-3);
  font-size: .84rem;
}

h2 {
  letter-spacing: -0.035em;
}

.principles p,
.extension p {
  color: var(--text-2);
  line-height: 1.75;
}

.extension {
  margin-top: 18px;
  padding: clamp(28px, 5vw, 56px);
  border-radius: var(--radius-lg);
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 44px;
}

.extension h2 {
  margin: 14px 0 0;
  font-size: clamp(2rem, 4vw, 3.6rem);
  line-height: 1;
}

@media (max-width: 800px) {
  .principles,
  .extension {
    grid-template-columns: 1fr;
  }
}
</style>
