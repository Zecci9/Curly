<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../../stores/auth'
import { useThemeStore } from '../../stores/theme'

defineEmits(['menu'])

const auth = useAuthStore()
const theme = useThemeStore()
const router = useRouter()

function logout() {
  auth.logout()

  router.replace('/login')
}
</script>

<template>
  <header class="topbar glass">

    <button
      class="menu pressable"
      @click="$emit('menu')"
      aria-label="打开侧栏"
    >
      ☰
    </button>


    <div class="search">
      ⌕

      <input
        placeholder="搜索后台（界面预留）"
      />
    </div>


    <div class="right">

      <button
        class="round pressable"
        @click="theme.cycleMode()"
        title="切换明暗"
      >
        ◐
      </button>


      <RouterLink
        class="profile pressable"
        to="/admin/profile"
      >
        <span class="avatar">
          {{ auth.displayName.slice(0, 1).toUpperCase() }}
        </span>

        <span>
          {{ auth.displayName }}
        </span>
      </RouterLink>


      <button
        class="logout pressable"
        type="button"
        @click="logout"
      >
        退出
      </button>

    </div>

  </header>
</template>


<style scoped>
.topbar {
  position: sticky;
  top: 12px;
  z-index: 50;

  min-height: 56px;

  border-radius: 18px;

  padding: 7px 9px;

  display: grid;
  grid-template-columns: 1fr auto;

  align-items: center;

  box-shadow: var(--shadow-sm);
}


.search {
  display: flex;
  align-items: center;

  gap: 8px;

  color: var(--text-3);

  padding-left: 8px;
}


.search input {
  width: min(36vw, 420px);

  border: 0;
  outline: 0;

  background: transparent;

  color: var(--text-1);
}


.right {
  display: flex;
  align-items: center;

  gap: 7px;
}


.round {
  width: 40px;
  height: 40px;

  border-radius: 999px;

  border: 1px solid var(--border);

  background: var(--surface-soft);

  cursor: pointer;
}


.profile {
  display: flex;
  align-items: center;

  gap: 8px;

  min-height: 42px;

  padding: 0 10px;

  border-radius: 999px;

  text-decoration: none;

  font-size: 0.88rem;
}


.avatar {
  width: 30px;
  height: 30px;

  border-radius: 999px;

  background: var(--accent);

  color: var(--accent-text);

  display: grid;
  place-items: center;

  font-weight: 800;
}


.logout {
  min-height: 40px;

  padding: 0 14px;

  border-radius: 999px;

  border: 1px solid var(--border);

  background: var(--surface-soft);

  color: var(--text-2);

  cursor: pointer;
}


.logout:hover {
  color: var(--danger);
}


.menu {
  display: none;

  width: 40px;
  height: 40px;

  border-radius: 12px;

  border: 1px solid var(--border);

  background: var(--surface-soft);
}


@media (max-width: 900px) {

  .topbar {
    grid-template-columns: auto 1fr auto;
  }


  .menu {
    display: block;
  }


  .search input {
    width: 100%;
  }

}


@media (max-width: 560px) {

  .search {
    display: none;
  }


  .topbar {
    grid-template-columns: auto 1fr;
  }


  .right {
    justify-self: end;
  }


  .profile > span:last-child {
    display: none;
  }


  .logout {
    padding: 0 11px;
  }

}
</style>