import axios from 'axios'

const baseURL = import.meta.env.VITE_API_BASE_URL || 'http://127.0.0.1:31098'

export const apiClient = axios.create({
  baseURL,
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json',
  },
})

export async function fetchMessages() {
  const { data } = await apiClient.get('/api/messages')
  return data
}

export async function createMessage(content) {
  const { data } = await apiClient.post('/api/messages', { content })
  return data
}
