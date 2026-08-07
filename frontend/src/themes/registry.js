const themes = new Map([
  ['default', {
    id: 'default',
    name: 'Curly Default',
    author: 'Curly',
    version: '0.5.0',
    description: '清晰、克制、偏阅读。',
    capabilities: ['light', 'dark', 'glass']
  }],
  ['midnight', {
    id: 'midnight',
    name: 'Midnight',
    author: 'Curly',
    version: '0.5.0',
    description: '蓝黑色调的主题示例。',
    capabilities: ['dark', 'glass']
  }],
])

export function registerTheme(theme) {
  if (!theme?.id || !theme?.name) throw new Error('Theme requires id and name')
  themes.set(theme.id, theme)
}
export const getThemes = () => [...themes.values()]
export const hasTheme = (id) => themes.has(id)
