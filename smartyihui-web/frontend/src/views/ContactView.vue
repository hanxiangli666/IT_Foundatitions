<template>
  <div class="page-top-pad">
    <section class="contact-hero">
      <div class="ch-inner">
        <div class="section-tag">联系我们</div>
        <h1 class="ch-title">有项目想法？<br/><span class="gold">告诉我们</span></h1>
        <p class="ch-sub">我们通常在24小时内回复，期待与您的合作</p>
      </div>
    </section>

    <section class="contact-main">
      <div class="cm-inner">
        <!-- 左侧信息 -->
        <div class="contact-info">
          <h2>联系方式</h2>
          <div class="info-list">
            <div class="info-item" v-for="item in infos" :key="item.id">
              <div class="info-icon">{{ item.icon }}</div>
              <div>
                <div class="info-label">{{ item.label }}</div>
                <div class="info-value">{{ item.value }}</div>
              </div>
            </div>
          </div>

          <div class="map-placeholder">
            <div class="map-pin">📍</div>
            <p>广州市天河区<br/>中山大道建中路5号</p>
          </div>
        </div>

        <!-- 右侧表单 -->
        <div class="contact-form-wrap">
          <h2>发送需求</h2>
          <form class="contact-form" @submit.prevent="handleSubmit">
            <div class="form-row">
              <div class="field">
                <label>您的姓名 *</label>
                <input v-model="form.name" type="text" placeholder="请输入您的姓名" required />
              </div>
              <div class="field">
                <label>联系电话 *</label>
                <input v-model="form.phone" type="tel" placeholder="请输入您的手机号" required />
              </div>
            </div>
            <div class="field">
              <label>项目类型</label>
              <select v-model="form.type">
                <option value="">请选择项目类型</option>
                <option v-for="t in projectTypes" :key="t" :value="t">{{ t }}</option>
              </select>
            </div>
            <div class="field">
              <label>需求描述 *</label>
              <textarea v-model="form.desc" rows="5" placeholder="请简单描述您的项目需求、功能点、预算等..." required></textarea>
            </div>
            <button type="submit" class="btn-submit-form" :disabled="submitting">
              {{ submitting ? '发送中...' : '发送需求 →' }}
            </button>
            <p v-if="submitMsg" class="submit-msg" :class="{ success: submitOk }">{{ submitMsg }}</p>
          </form>
        </div>
      </div>
    </section>

    <section class="faq-section">
      <div class="section-inner">
        <div class="section-header">
          <div class="section-tag">常见问题</div>
          <h2 class="section-title">您可能想问的问题</h2>
        </div>
        <div class="faq-list">
          <div class="faq-item" v-for="faq in faqs" :key="faq.id"
            @click="faq.open = !faq.open" :class="{ open: faq.open }">
            <div class="faq-q">
              <span>{{ faq.q }}</span>
              <span class="faq-arrow">{{ faq.open ? '▲' : '▼' }}</span>
            </div>
            <div class="faq-a" v-if="faq.open">{{ faq.a }}</div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

const form = reactive({ name: '', phone: '', type: '', desc: '' })
const submitting = ref(false)
const submitMsg = ref('')
const submitOk = ref(false)

const projectTypes = [
  '企业官网', '电商平台', '管理后台系统', '微信小程序', '微信小游戏', 'H5页面', 'App开发', '其他'
]

const infos = [
  { id:1, icon:'📍', label:'公司地址', value:'广州市天河区中山大道建中路5号' },
  { id:2, icon:'🏢', label:'公司全称', value:'广州熠辉智能科技有限公司' },
  { id:3, icon:'🌐', label:'官方网站', value:'www.smartyihui.com' },
]

const faqs = ref([
  { id:1, q:'开发一个网站大概需要多少钱？', open: false,
    a:'价格根据功能复杂度而定，简单的企业官网通常在几千元，带后台管理系统的网站在1万元左右。欢迎联系我们获取具体报价。' },
  { id:2, q:'开发周期一般是多久？', open: false,
    a:'简单官网1-2周，小程序2-4周，复杂系统视功能量而定。我们会在签合同前给出明确的排期计划。' },
  { id:3, q:'你们是否可以提供合同和发票？', open: false,
    a:'是的，我们是正规注册的有限责任公司，可以签署正式合同并开具增值税发票，满足企业报账需求。' },
  { id:4, q:'代码和版权归谁所有？', open: false,
    a:'项目验收交付后，代码版权完全归属客户，我们会交付完整源代码，不附加任何隐形费用或版权限制。' },
  { id:5, q:'上线后还需要付费维护吗？', open: false,
    a:'合同范围内的Bug修复是免费的。后续如有新功能迭代，我们会提供优惠的维护服务价格。' },
])

async function handleSubmit() {
  submitting.value = true
  submitMsg.value = ''
  // 模拟发送（实际可接入邮件或企业微信机器人）
  await new Promise(r => setTimeout(r, 1000))
  submitOk.value = true
  submitMsg.value = '✅ 需求已发送！我们将在24小时内联系您'
  submitting.value = false
  form.name = ''; form.phone = ''; form.type = ''; form.desc = ''
}
</script>

<style scoped>
.page-top-pad { padding-top: 80px; }

.contact-hero {
  background: linear-gradient(135deg, var(--ink) 0%, #1a2d4a 100%);
  padding: 80px 32px;
}
.ch-inner { max-width: 700px; margin: 0 auto; text-align: center; }
.ch-inner .section-tag { background: rgba(200,151,58,0.2); }
.ch-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(32px, 4vw, 52px);
  font-weight: 700; color: #fff;
  margin: 16px 0 16px; line-height: 1.3;
}
.ch-sub { color: rgba(255,255,255,0.6); font-size: 16px; }

.contact-main { background: #f8fafc; }
.cm-inner {
  max-width: 1100px; margin: 0 auto;
  padding: 80px 32px;
  display: grid;
  grid-template-columns: 1fr 1.5fr;
  gap: 60px;
  align-items: start;
}

.contact-info h2, .contact-form-wrap h2 {
  font-family: 'Noto Serif SC', serif;
  font-size: 22px; font-weight: 700; margin-bottom: 28px;
}
.info-list { display: flex; flex-direction: column; gap: 20px; margin-bottom: 32px; }
.info-item {
  display: flex; align-items: flex-start; gap: 14px;
  padding: 16px;
  background: #fff; border-radius: 12px;
  border: 1px solid var(--border);
}
.info-icon { font-size: 22px; flex-shrink: 0; }
.info-label { font-size: 12px; color: var(--ink-mute); margin-bottom: 3px; font-weight: 600; }
.info-value { font-size: 14px; color: var(--ink); font-weight: 500; }

.map-placeholder {
  background: var(--ink);
  border-radius: 16px;
  padding: 32px;
  text-align: center;
  color: rgba(255,255,255,0.7);
}
.map-pin { font-size: 36px; margin-bottom: 12px; }
.map-placeholder p { font-size: 14px; line-height: 1.8; color: rgba(255,255,255,0.6); }

.contact-form-wrap {
  background: #fff;
  border-radius: 20px;
  padding: 36px;
  border: 1px solid var(--border);
  box-shadow: 0 8px 24px rgba(0,0,0,0.05);
}
.contact-form { display: flex; flex-direction: column; gap: 20px; }
.form-row { display: grid; grid-template-columns: 1fr 1fr; gap: 16px; }
.field { display: flex; flex-direction: column; gap: 6px; }
.field label { font-size: 13px; font-weight: 700; color: var(--ink-soft); }
.field input, .field select, .field textarea {
  padding: 11px 14px; border-radius: 10px;
  border: 1.5px solid var(--border);
  font-size: 14px; font-family: inherit;
  transition: border-color 0.2s; outline: none;
  background: #fff;
}
.field textarea { resize: vertical; }
.field input:focus, .field select:focus, .field textarea:focus { border-color: var(--gold); }

.btn-submit-form {
  background: var(--gold); color: #fff;
  border: none; cursor: pointer;
  padding: 14px; border-radius: 10px;
  font-size: 16px; font-weight: 700; font-family: inherit;
  transition: all 0.2s;
}
.btn-submit-form:hover:not(:disabled) { background: var(--gold-lt); transform: translateY(-1px); }
.btn-submit-form:disabled { opacity: 0.6; cursor: not-allowed; }
.submit-msg { text-align: center; font-size: 14px; color: var(--ink-mute); }
.submit-msg.success { color: #38a169; }

/* FAQ */
.faq-section { background: #fff; }
.faq-list { max-width: 760px; margin: 0 auto; display: flex; flex-direction: column; gap: 12px; }
.faq-item {
  border: 1.5px solid var(--border); border-radius: 14px;
  overflow: hidden; cursor: pointer;
  transition: border-color 0.2s;
}
.faq-item.open { border-color: var(--gold); }
.faq-q {
  display: flex; justify-content: space-between; align-items: center;
  padding: 18px 20px;
  font-size: 15px; font-weight: 600;
  user-select: none;
}
.faq-arrow { color: var(--gold); font-size: 12px; }
.faq-a {
  padding: 0 20px 18px;
  font-size: 14px; color: var(--ink-soft); line-height: 1.8;
}

@media (max-width: 768px) {
  .cm-inner { grid-template-columns: 1fr; gap: 32px; padding: 40px 20px; }
  .form-row { grid-template-columns: 1fr; }
  .section-inner { padding: 60px 20px; }
}
</style>
