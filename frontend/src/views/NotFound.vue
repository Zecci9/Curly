<template>
  <section class="not-found">
    <div>
      <span class="eyebrow">404</span>
      <h1>这个页面不在这里。</h1>
      <p>它可能被移动、删除，或者还没有被创建。</p>
      <RouterLink class="primary-button pressable" to="/">回到首页</RouterLink>
    </div>
  </section>
</template>

<style scoped>
.not-found {
  min-height: 100vh;
  padding: 30px;
  display: grid;
  place-items: center;
  text-align: center;
}

h1 {
  margin: 12px auto;
  max-width: 800px;
  font-size: clamp(3rem, 9vw, 7rem);
  line-height: .94;
  letter-spacing: -0.06em;
}

p {
  color: var(--text-2);
  margin-bottom: 28px;
}
</style>
