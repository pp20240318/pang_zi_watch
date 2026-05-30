const API_BASE = import.meta.env.VITE_API_BASE || ''

async function request(path, options = {}) {
  const res = await fetch(`${API_BASE}${path}`, {
    headers: { 'Content-Type': 'application/json', ...options.headers },
    ...options,
  })
  if (!res.ok) {
    const err = await res.json().catch(() => ({}))
    throw new Error(err.error || res.statusText)
  }
  return res.json()
}

export function fetchBanners() {
  return request('/api/public/banners').then((r) => r.data)
}

export function fetchBrands() {
  return request('/api/public/brands').then((r) => r.data)
}

export function fetchProducts(brand) {
  const q = brand && brand !== 'all' ? `?brand=${encodeURIComponent(brand)}` : ''
  return request(`/api/public/products${q}`).then((r) => r.data)
}

export function fetchProduct(id) {
  return request(`/api/public/products/${id}`).then((r) => r.data)
}

export function createOrder(payload) {
  return request('/api/public/orders', {
    method: 'POST',
    body: JSON.stringify(payload),
  }).then((r) => r.data)
}

export function payOrder(orderNo, paymentMethod) {
  return request(`/api/public/orders/${orderNo}/pay`, {
    method: 'POST',
    body: JSON.stringify({ paymentMethod }),
  }).then((r) => r.data)
}

export function fetchPages() {
  return request('/api/public/pages').then((r) => r.data)
}

export function fetchPage(slug) {
  return request(`/api/public/pages/${slug}`).then((r) => r.data)
}
