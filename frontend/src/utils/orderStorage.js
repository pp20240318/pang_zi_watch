const ORDER_KEY = 'pz_pending_order'

export function savePendingOrder(order) {
  sessionStorage.setItem(ORDER_KEY, JSON.stringify(order))
}

export function loadPendingOrder() {
  const raw = sessionStorage.getItem(ORDER_KEY)
  if (!raw) return null
  try {
    return JSON.parse(raw)
  } catch {
    return null
  }
}

export function clearPendingOrder() {
  sessionStorage.removeItem(ORDER_KEY)
}
