<script setup>
defineProps({name:{type:String,required:true}})
</script>
<template>
<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
 <template v-if="name==='dashboard'"><rect x="4" y="4" width="6" height="6" rx="1"/><rect x="14" y="4" width="6" height="6" rx="1"/><rect x="4" y="14" width="6" height="6" rx="1"/><rect x="14" y="14" width="6" height="6" rx="1"/></template>
 <template v-else-if="name==='post'"><path d="M6 4h9l3 3v13H6z"/><path d="M9 11h6M9 15h6"/></template>
 <template v-else-if="name==='page'"><path d="M7 3h8l3 3v15H7z"/><path d="M15 3v4h4"/></template>
 <template v-else-if="name==='folder'"><path d="M3 7h7l2 2h9v10H3z"/></template>
 <template v-else-if="name==='tag'"><path d="M4 12V5h7l9 9-7 7z"/><circle cx="8" cy="9" r="1"/></template>
 <template v-else-if="name==='comment'"><path d="M4 5h16v11H9l-5 4z"/></template>
 <template v-else-if="name==='media'"><rect x="3" y="4" width="18" height="16" rx="2"/><circle cx="9" cy="9" r="2"/><path d="m4 17 5-5 4 4 2-2 5 5"/></template>
 <template v-else-if="name==='user'"><circle cx="12" cy="8" r="4"/><path d="M4 21c1-5 4-7 8-7s7 2 8 7"/></template>
 <template v-else-if="name==='theme'"><circle cx="12" cy="12" r="8"/><path d="M12 4a8 8 0 0 0 0 16z"/></template>
 <template v-else-if="name==='plugin'"><path d="M8 3v5H3v8h5v5h8v-5h5V8h-5V3z"/></template>
 <template v-else-if="name==='settings'"><circle cx="12" cy="12" r="3"/><path d="M12 3v3M12 18v3M3 12h3M18 12h3M5.6 5.6l2.1 2.1M16.3 16.3l2.1 2.1M18.4 5.6l-2.1 2.1M7.7 16.3l-2.1 2.1"/></template>
 <template v-else><circle cx="12" cy="12" r="8"/></template>
</svg>
</template>
