<script setup>
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'
import { useThemeStore } from './stores/theme'
import { useAuthStore } from './stores/auth'

const theme = useThemeStore()
const auth = useAuthStore()

onMounted(() => {
  theme.init()
  auth.restore()
})
</script>

<template>
  <RouterView />
</template>
