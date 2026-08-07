import { shallowReactive } from 'vue'

const plugins = []
const slots = shallowReactive(new Map())

export function registerPlugin(plugin) {
  if (!plugin?.id) throw new Error('Plugin must provide id')
  plugins.push(plugin)

  for (const entry of plugin.slots || []) {
    if (!slots.has(entry.name)) slots.set(entry.name, [])
    slots.get(entry.name).push({
      pluginId: plugin.id,
      component: entry.component,
      props: entry.props || {},
    })
  }
}

export function getPlugins() {
  return [...plugins]
}

export function getSlotEntries(name) {
  return slots.get(name) || []
}

export function installRegisteredPlugins(app, context) {
  for (const plugin of plugins) {
    plugin.install?.(app, context)
  }
}

/*
第三方插件示例：

registerPlugin({
  id: 'comments',
  name: 'Comments',
  version: '1.0.0',
  slots: [
    {
      name: 'post.after',
      component: CommentsWidget,
    }
  ],
  routes: [],
  settings: [],
  adminPages: [],
})
*/
