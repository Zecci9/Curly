<script setup>
import SiteHeader from '../components/SiteHeader.vue'
import SiteFooter from '../components/SiteFooter.vue'
</script>

<template>
  <div class="front-layout">
    <SiteHeader />

    <main>
      <RouterView />
    </main>

    <SiteFooter />
  </div>
</template>

<style scoped>
.front-layout {
  min-height: 100vh;
}
</style>
