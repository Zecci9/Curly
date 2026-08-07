<template>

<div class="layout">

<header>

<div class="logo">
Curly
</div>


<nav>

<router-link to="/">
首页
</router-link>


<router-link to="/posts">
文章
</router-link>


<router-link to="/login">
登录
</router-link>

</nav>

</header>


<main>

<router-view />

</main>


<footer>

Curly CMS © 2026

</footer>


</div>

</template>


<style scoped>

.layout{

min-height:100vh;

background:#f7f7f7;

}


header{

height:64px;

display:flex;

align-items:center;

justify-content:space-between;

padding:0 40px;

background:white;

border-bottom:1px solid #eee;

}


.logo{

font-size:26px;

font-weight:bold;

}


nav a{

margin-left:25px;

text-decoration:none;

color:#333;

}


main{

padding:40px;

}


footer{

text-align:center;

padding:30px;

color:#888;

}

</style>