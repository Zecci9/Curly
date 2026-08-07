<script setup>
defineProps({eyebrow:{type:String,default:'Curly Admin'},title:{type:String,required:true},description:{type:String,default:''}})
</script>
<template>
<div class="head">
 <div><span class="eyebrow">{{eyebrow}}</span><h1>{{title}}</h1><p v-if="description">{{description}}</p></div>
 <div class="actions"><slot name="actions"/></div>
</div>
</template>
<style scoped>
.head{display:flex;justify-content:space-between;align-items:end;gap:24px;margin:28px 2px 26px}.head h1{margin:7px 0 0;font-size:clamp(2rem,4vw,3.2rem);letter-spacing:-.045em;line-height:1}.head p{margin:10px 0 0;color:var(--text-2);line-height:1.65;max-width:700px}.actions{display:flex;gap:8px;flex-wrap:wrap}
@media(max-width:650px){.head{align-items:flex-start;flex-direction:column}.actions{width:100%}}
</style>
