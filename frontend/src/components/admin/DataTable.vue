<script setup>
defineProps({columns:{type:Array,required:true},rows:{type:Array,required:true},empty:{type:String,default:'暂无数据'}})
</script>
<template>
<div class="table surface">
 <div class="scroll">
  <table><thead><tr><th v-for="c in columns" :key="c.key">{{c.label}}</th><th v-if="$slots.actions"></th></tr></thead>
  <tbody><tr v-for="row in rows" :key="row.id??JSON.stringify(row)"><td v-for="c in columns" :key="c.key"><slot :name="`cell-${c.key}`" :row="row">{{row[c.key]}}</slot></td><td v-if="$slots.actions" class="actions"><slot name="actions" :row="row"/></td></tr></tbody></table>
 </div>
 <div v-if="!rows.length" class="empty">{{empty}}</div>
</div>
</template>
<style scoped>.table{border-radius:18px;overflow:hidden}.scroll{overflow:auto}table{width:100%;border-collapse:collapse;min-width:680px}th,td{padding:14px 16px;text-align:left;border-bottom:1px solid var(--border);font-size:.9rem}th{color:var(--text-3);font-size:.78rem;font-weight:700;background:var(--surface-soft)}td{color:var(--text-2)}tbody tr:last-child td{border-bottom:0}.actions{text-align:right}.empty{padding:40px;text-align:center;color:var(--text-3)}</style>
