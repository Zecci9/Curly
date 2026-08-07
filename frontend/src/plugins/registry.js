import { shallowReactive } from 'vue'

const plugins=[]
const slots=shallowReactive(new Map())
const adminItems=shallowReactive([])

export function registerPlugin(plugin){
  if(!plugin?.id) throw new Error('Plugin requires id')
  plugins.push(plugin)

  for(const entry of plugin.slots||[]){
    if(!slots.has(entry.name)) slots.set(entry.name,[])
    slots.get(entry.name).push({pluginId:plugin.id,...entry})
  }

  for(const item of plugin.adminMenu||[]){
    adminItems.push({...item,pluginId:plugin.id})
  }
}
export const getPlugins=()=>[...plugins]
export const getSlotEntries=(name)=>slots.get(name)||[]
export const getPluginAdminItems=()=>[...adminItems]
export function installRegisteredPlugins(app,ctx){ plugins.forEach(p=>p.install?.(app,ctx)) }
