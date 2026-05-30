export interface JWTPayload {
  admin_id: number
  username: string
  role: string
  permissions: string
  exp?: number
}

export function parseJWT(token: string): JWTPayload | null {
  try {
    const payload = token.split('.')[1]
    const json = atob(payload.replace(/-/g, '+').replace(/_/g, '/'))
    return JSON.parse(json) as JWTPayload
  } catch {
    return null
  }
}
