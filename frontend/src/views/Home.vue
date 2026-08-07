<script setup>
import {onMounted,onBeforeUnmount,ref,computed,nextTick} from 'vue'
import {getPosts} from '../api/content'
import {demoPosts,normalizePosts,uniqueCategories} from '../utils/content'
import {observeReveal} from '../utils/reveal'
import PostCard from '../components/PostCard.vue'
const posts=ref([]);let observer
onMounted(async()=>{try{posts.value=normalizePosts(await getPosts());if(!posts.value.length)throw 0}catch{posts.value=demoPosts}await nextTick();observer=observeReveal()})
onBeforeUnmount(()=>observer?.disconnect())
const categories=computed(()=>uniqueCategories(posts.value))
</script>
<template>
<section class="hero">
 <div class="page-shell hero-grid">
  <div>
   <span class="eyebrow" data-reveal>Curly CMS · Open by design</span>
   <h1 class="display" data-reveal data-delay="1">内容应该自由地<br>生长。</h1>
   <p class="lead" data-reveal data-delay="2">一个为文章、主题与插件留出空间的内容系统。默认足够清晰，也允许开发者让它完全变样。</p>
   <div class="buttons" data-reveal data-delay="3"><RouterLink class="primary-button pressable" to="/posts">开始阅读</RouterLink><RouterLink class="secondary-button pressable" to="/about">了解 Curly</RouterLink></div>
  </div>
  <div class="object surface" data-reveal data-delay="2"><div class="orb"></div><div class="logo">C</div><span class="pill p1">Theme</span><span class="pill p2">Plugin</span><span class="pill p3">Content</span></div>
 </div>
</section>
<section class="section"><div class="page-shell">
 <div class="head" data-reveal><div><span class="eyebrow">Latest</span><h2>最近发布</h2></div><RouterLink class="chip pressable" to="/posts">全部文章 →</RouterLink></div>
 <div class="posts"><div v-for="(post,i) in posts" :key="post.id" data-reveal :data-delay="Math.min(i,3)"><PostCard :post="post"/></div></div>
</div></section>
<section class="section"><div class="page-shell">
 <div class="head" data-reveal><div><span class="eyebrow">Explore</span><h2>按主题发现</h2></div></div>
 <div class="categories"><RouterLink v-for="(c,i) in categories" :key="c" class="category surface pressable" :to="`/categories/${encodeURIComponent(c)}`" data-reveal :data-delay="i%3"><span>{{c}}</span><span>↗</span></RouterLink></div>
</div></section>
<section class="section"><div class="page-shell manifesto surface" data-reveal><span class="eyebrow">Built to change</span><h2>默认主题只是开始，<br>不是边界。</h2><p>页面依赖稳定的数据与组件接口。主题重新定义视觉，插件加入能力，而内容本身保持清晰。</p></div></section>
</template>
<style scoped>
.hero{min-height:calc(100svh - var(--nav-height));padding:70px 0;display:grid;align-items:center}.hero-grid{display:grid;grid-template-columns:1.2fr .8fr;gap:60px;align-items:center}.lead{max-width:650px;margin-top:25px}.buttons{display:flex;gap:10px;flex-wrap:wrap;margin-top:28px}
.object{min-height:430px;border-radius:38px;display:grid;place-items:center;position:relative;overflow:hidden;background:radial-gradient(circle at 50% 40%,var(--accent-soft),transparent 42%),var(--surface)}
.orb{position:absolute;width:330px;height:330px;border:1px solid var(--border);border-radius:50%;box-shadow:0 0 0 55px transparent,0 0 0 56px var(--border)}.logo{width:130px;height:130px;border-radius:38px;display:grid;place-items:center;background:var(--accent);color:var(--accent-text);font-size:4.8rem;font-weight:800;z-index:2}
.pill{position:absolute;padding:10px 13px;border-radius:14px;background:var(--glass);border:1px solid var(--glass-border);backdrop-filter:blur(16px);color:var(--text-2);font-size:.84rem;font-weight:700}.p1{transform:translate(120px,-120px) rotate(5deg)}.p2{transform:translate(-130px,100px) rotate(-6deg)}.p3{transform:translate(115px,120px) rotate(3deg)}
.head{display:flex;justify-content:space-between;align-items:end;gap:20px;margin-bottom:24px}.head h2{font-size:clamp(1.9rem,4vw,3rem);letter-spacing:-.04em;margin:6px 0 0}.posts{display:grid;grid-template-columns:repeat(2,1fr);gap:17px}.categories{display:grid;grid-template-columns:repeat(3,1fr);gap:14px}.category{min-height:110px;border-radius:var(--radius-md);padding:22px;display:flex;justify-content:space-between;align-items:end;text-decoration:none;font-weight:700}
.manifesto{padding:clamp(30px,6vw,70px);border-radius:var(--radius-lg)}.manifesto h2{font-size:clamp(2.2rem,6vw,4.8rem);line-height:.98;letter-spacing:-.055em;margin:16px 0}.manifesto p{max-width:700px;color:var(--text-2);line-height:1.8}
@media(max-width:880px){.hero-grid,.posts{grid-template-columns:1fr}.object{min-height:330px}.categories{grid-template-columns:1fr 1fr}}@media(max-width:560px){.categories{grid-template-columns:1fr}}
</style>
