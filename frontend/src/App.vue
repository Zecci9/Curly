<script setup>

import { ref, onMounted } from "vue"
import axios from "axios"


const backend = ref(null)


onMounted(async()=>{

    const res = await axios.get(
        `${import.meta.env.VITE_API_URL}/health`
    )

    backend.value = res.data.data

})


</script>


<template>

<main>

<h1>Curly</h1>

<p>
一个使用 Go + Vue 开发的 CMS 系统
</p>


<div v-if="backend">

<p>
后端状态：
{{ backend.status }}
</p>

<p>
版本：
{{ backend.version }}
</p>

</div>


</main>

</template>