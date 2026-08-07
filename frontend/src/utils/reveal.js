export function observeReveal(root = document) {
  const nodes = [...root.querySelectorAll('[data-reveal]')]
  if (!nodes.length) return null

  if (matchMedia('(prefers-reduced-motion: reduce)').matches) {
    nodes.forEach(el => el.classList.add('is-visible'))
    return null
  }

  const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      if (!entry.isIntersecting) return
      entry.target.classList.add('is-visible')
      observer.unobserve(entry.target)
    })
  }, { threshold: .12, rootMargin: '0px 0px -7% 0px' })

  nodes.forEach(el => observer.observe(el))
  return observer
}
