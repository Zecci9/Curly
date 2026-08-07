import defaultTheme from './default'
import midnightTheme from './midnight'

const themes = new Map([
  [defaultTheme.id, defaultTheme],
  [midnightTheme.id, midnightTheme],
])

export function registerTheme(theme) {
  if (!theme?.id || !theme?.name) {
    throw new Error('Theme must provide id and name')
  }
  themes.set(theme.id, theme)
}

export function getTheme(id) {
  return themes.get(id)
}

export function getThemes() {
  return [...themes.values()]
}

export function hasTheme(id) {
  return themes.has(id)
}
