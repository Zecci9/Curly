<script setup>
import {ref} from 'vue';import AdminPageHeader from '../../components/admin/AdminPageHeader.vue';import DataTable from '../../components/admin/DataTable.vue'
const rows=ref([{id:1,title:'关于',slug:'about',status:'published'},{id:2,title:'友情链接',slug:'links',status:'draft'}]);const columns=[{key:'title',label:'页面'},{key:'slug',label:'Slug'},{key:'status',label:'状态'}]
</script><template><AdminPageHeader title="页面" description="用于关于、友链、隐私政策等独立内容。"><template #actions><button class="primary-button pressable">＋ 新建页面</button></template></AdminPageHeader><DataTable :columns="columns" :rows="rows"><template #cell-status="{row}"><span>{{row.status==='published'?'已发布':'草稿'}}</span></template><template #actions>编辑</template></DataTable><p class="note">后端 Page 模型/API 尚未实现，当前是前端设计预览。</p></template>
<style scoped>.note{color:var(--text-3);font-size:.82rem}</style>
