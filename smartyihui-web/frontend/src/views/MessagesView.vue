<script setup>
import { onMounted, ref } from 'vue'
import { useMessage } from 'naive-ui'
import { createMessage, fetchMessages } from '@/lib/api'

const message = useMessage()
const loading = ref(false)
const saving = ref(false)
const content = ref('')
const items = ref([])

async function loadMessages() {
  loading.value = true
  try {
    const data = await fetchMessages()
    items.value = data.items || []
  } catch (error) {
    message.error(error.response?.data?.error || error.message || '加载失败')
  } finally {
    loading.value = false
  }
}

async function handleSubmit() {
  if (!content.value.trim()) {
    message.warning('先输入一点内容')
    return
  }

  saving.value = true
  try {
    const created = await createMessage(content.value.trim())
    items.value = [created, ...items.value]
    content.value = ''
    message.success('已写入 MySQL')
  } catch (error) {
    message.error(error.response?.data?.error || error.message || '保存失败')
  } finally {
    saving.value = false
  }
}

onMounted(loadMessages)
</script>

<template>
  <section class="message-page">
    <div class="editor-card">
      <h2>内容管理示例</h2>
      <p>通过表单提交一条内容到 Go Fiber API，后端写入 MySQL 后再回显到列表，可作为企业门户内容管理的基础能力示例。</p>
      <n-input
        v-model:value="content"
        type="textarea"
        :autosize="{ minRows: 4, maxRows: 8 }"
        placeholder="比如：下周三下午三点开项目例会"
      />
      <div class="toolbar">
        <n-button type="primary" :loading="saving" @click="handleSubmit">
          提交消息
        </n-button>
        <n-button secondary :loading="loading" @click="loadMessages">
          刷新列表
        </n-button>
      </div>
    </div>

    <div class="list-card">
      <div class="list-header">
        <h3>已保存消息</h3>
        <span>{{ items.length }} 条</span>
      </div>

      <n-spin :show="loading">
        <n-empty v-if="!items.length" description="还没有数据，先新增一条" />
        <div v-else class="message-list">
          <article v-for="item in items" :key="item.id" class="message-item">
            <div class="message-meta">
              <strong>#{{ item.id }}</strong>
              <span>{{ item.createdAt }}</span>
            </div>
            <p>{{ item.content }}</p>
          </article>
        </div>
      </n-spin>
    </div>
  </section>
</template>

<style scoped>
.message-page {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
}

.editor-card,
.list-card {
  padding: 24px;
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.96);
  box-shadow: 0 18px 50px rgba(15, 23, 42, 0.08);
}

.editor-card h2,
.list-card h3 {
  margin-top: 0;
}

.editor-card p {
  color: #64748b;
}

.toolbar {
  display: flex;
  gap: 12px;
  margin-top: 16px;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.message-list {
  display: grid;
  gap: 12px;
}

.message-item {
  padding: 16px;
  border-radius: 16px;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
}

.message-meta {
  display: flex;
  justify-content: space-between;
  gap: 16px;
  margin-bottom: 8px;
  color: #64748b;
  font-size: 14px;
}

.message-item p {
  margin: 0;
  white-space: pre-wrap;
}

@media (max-width: 900px) {
  .message-page {
    grid-template-columns: 1fr;
  }
}
</style>
