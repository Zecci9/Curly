<script setup>
defineProps({label:String,value:[String,Number],trend:String,tone:{type:String,default:'normal'}})
</script>
<template><article class="stat surface"><div class="top"><span>{{label}}</span><span class="dot" :class="tone"></span></div><strong>{{value}}</strong><small>{{trend}}</small></article></template>
<style scoped>.stat{padding:20px;border-radius:18px;min-height:148px;display:flex;flex-direction:column}.top{display:flex;justify-content:space-between;color:var(--text-2);font-size:.86rem}.dot{width:8px;height:8px;border-radius:99px;background:var(--text-3)}.dot.good{background:var(--success)}.dot.warn{background:var(--warning)}strong{font-size:2.4rem;letter-spacing:-.05em;margin:auto 0 4px}small{color:var(--text-3)}</style>
