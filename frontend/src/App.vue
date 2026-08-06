<script setup>

import { ref } from "vue"
import axios from "axios"


const API = "http://192.168.8.48:8080/api/v1"


const username = ref("admin")
const password = ref("123456")
const email = ref("admin@test.com")


const token = ref(
    localStorage.getItem("token") || ""
)


const result = ref("")


// 注册

async function register(){

    try{

        const res = await axios.post(
            `${API}/users/register`,
            {
                username:username.value,
                password:password.value,
                email:email.value
            }
        )


        result.value =
            JSON.stringify(
                res.data,
                null,
                2
            )

    }catch(err){

        result.value =
            err.response?.data ||
            err.message

    }

}



// 登录

async function login(){

    try{

        const res = await axios.post(
            `${API}/users/login`,
            {
                username:username.value,
                password:password.value
            }
        )


        token.value =
            res.data.data.token


        localStorage.setItem(
            "token",
            token.value
        )


        result.value =
            "登录成功\nToken已保存"


    }catch(err){

        result.value =
            err.response?.data ||
            err.message

    }

}



// 获取用户

async function getMe(){

    try{

        const res = await axios.get(
            `${API}/users/me`,
            {
                headers:{
                    Authorization:
                    `Bearer ${token.value}`
                }
            }
        )


        result.value =
            JSON.stringify(
                res.data,
                null,
                2
            )


    }catch(err){

        result.value =
            err.response?.data ||
            err.message

    }

}



// 创建文章

async function createPost(){

    try{

        const res = await axios.post(
            `${API}/posts`,
            {
                title:"Vue测试文章",
                content:"来自前端测试页面"
            },
            {
                headers:{
                    Authorization:
                    `Bearer ${token.value}`
                }
            }
        )


        result.value =
            JSON.stringify(
                res.data,
                null,
                2
            )


    }catch(err){

        result.value =
            err.response?.data ||
            err.message

    }

}


</script>



<template>

<div class="container">

<h1>
Curly API Test
</h1>


<div>

<label>
用户名
</label>

<input v-model="username">


<label>
密码
</label>

<input
type="password"
v-model="password"
>


<label>
邮箱
</label>

<input v-model="email">


</div>



<button @click="register">
注册
</button>


<button @click="login">
登录
</button>


<button @click="getMe">
获取用户
</button>


<button @click="createPost">
创建文章
</button>



<h3>
返回结果
</h3>


<pre>
{{ result }}
</pre>


</div>


</template>



<style>

.container{

    width:600px;

    margin:50px auto;

    font-family:sans-serif;

}


input{

    display:block;

    width:100%;

    margin:8px 0;

    padding:8px;

}


button{

    margin:8px;

    padding:10px 20px;

}


pre{

    background:#eee;

    padding:15px;

    white-space:pre-wrap;

}


</style>