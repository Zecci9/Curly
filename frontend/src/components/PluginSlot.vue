<script setup>
import { computed } from 'vue'
import { getSlotEntries } from '../plugins/registry'
const props=defineProps({name:{type:String,required:true},context:{type:Object,default:()=>({})}})
const entries=computed(()=>getSlotEntries(props.name))
</script>
<template>
  <component v-for="entry in entries" :key="`${name}:${entry.pluginId}`" :is="entry.component" v-bind="{...(entry.props||{}),...context}" />
</template>
