<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const username = ref('')
const password = ref('')
const message = ref('')
const auth = useAuthStore()
const router = useRouter()

async function submit() {
  message.value = ''
  try {
    await auth.login(username.value, password.value)
    router.push('/')
  } catch (error) {
    message.value = error.message
  }
}
</script>

<template>
  <section class="login-page">
    <div class="login-card surface">
      <div>
        <RouterLink to="/" class="brand">Curly</RouterLink>
        <span class="eyebrow">Account</span>
        <h1>欢迎回来。</h1>
        <p>登录用于未来的个人中心、收藏和后台入口。权限仍必须由后端验证。</p>
      </div>

      <form @submit.prevent="submit">
        <label>
          <span>用户名</span>
          <input v-model.trim="username" autocomplete="username" required />
        </label>

        <label>
          <span>密码</span>
          <input
            v-model="password"
            type="password"
            autocomplete="current-password"
            required
          />
        </label>

        <button class="primary-button pressable" type="submit" :disabled="auth.loading">
          {{ auth.loading ? '正在登录…' : '登录' }}
        </button>

        <p v-if="message" class="error" role="alert">{{ message }}</p>
      </form>
    </div>
  </section>
</template>

<style scoped>
.login-page {
  min-height: calc(100svh - var(--nav-height));
  padding: 70px 20px;
  display: grid;
  place-items: center;
}

.login-card {
  width: min(100%, 880px);
  padding: clamp(28px, 6vw, 64px);
  border-radius: var(--radius-lg);
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 58px;
}

.brand {
  display: block;
  margin-bottom: 64px;
  font-weight: 800;
  text-decoration: none;
}

h1 {
  margin: 13px 0;
  font-size: clamp(2.5rem, 6vw, 4.6rem);
  line-height: .98;
  letter-spacing: -0.055em;
}

p {
  color: var(--text-2);
  line-height: 1.7;
}

form {
  display: grid;
  gap: 15px;
  align-content: center;
}

label {
  display: grid;
  gap: 8px;
  color: var(--text-2);
  font-size: .9rem;
}

input {
  width: 100%;
  min-height: 50px;
  border: 1px solid var(--border);
  border-radius: 14px;
  background: var(--surface-soft);
  color: var(--text-1);
  padding: 0 14px;
}

button:disabled {
  opacity: .58;
  cursor: wait;
}

.error {
  margin: 0;
  color: #d94141;
  font-size: .9rem;
}

@media (max-width: 760px) {
  .login-card {
    grid-template-columns: 1fr;
  }

  .brand {
    margin-bottom: 36px;
  }
}
</style>
