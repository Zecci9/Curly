export const demoPosts=[
 {id:1,title:'Curly：从一个接口开始',excerpt:'把内容、主题与插件拆开，让 CMS 从第一天就为变化留出空间。',content:'Curly 的目标不是一次写死所有功能，而是建立清晰边界。\n\n页面负责展示，主题负责视觉，插件负责扩展，API 负责数据。',author:'Curly',category:'开发',tags:['Go','Vue','CMS'],status:'published',created_at:'2026-08-06T10:00:00+08:00'},
 {id:2,title:'设计不是把东西变漂亮',excerpt:'一个好的界面首先应该让人知道自己在哪、能做什么。',content:'设计的价值来自秩序、反馈与可预测性。',author:'Curly',category:'设计',tags:['Design','Theme'],status:'published',created_at:'2026-08-05T20:10:00+08:00'},
 {id:3,title:'为什么要给插件留位置',excerpt:'评论、SEO、统计和特殊页面不应该让核心无限膨胀。',content:'插件系统的第一步不是做市场，而是定义边界。',author:'Curly',category:'架构',tags:['Plugin','Architecture'],status:'draft',created_at:'2026-08-04T18:30:00+08:00'},
]
export const normalizePosts=(payload)=>{
  const v=payload?.data??payload
  if(!Array.isArray(v)) return []
  return v.map(p=>({category:'未分类',tags:[],author:'Curly',excerpt:excerpt(p.content),...p}))
}
export const normalizePost=(payload)=>{
  const p=payload?.data??payload
  if(!p||typeof p!=='object') return null
  return {category:'未分类',tags:[],author:'Curly',excerpt:excerpt(p.content),...p}
}
export function excerpt(v='',max=100){ const s=String(v).replace(/\s+/g,' ').trim(); return s.length>max?s.slice(0,max)+'…':s }
export function formatDate(v){
  if(!v)return '暂无日期'
  const d=new Date(v)
  return Number.isNaN(d.getTime())?v:new Intl.DateTimeFormat('zh-CN',{year:'numeric',month:'long',day:'numeric'}).format(d)
}
export const uniqueCategories=(posts)=>[...new Set(posts.map(p=>p.category||'未分类'))]
export const uniqueTags=(posts)=>[...new Set(posts.flatMap(p=>p.tags||[]))]
