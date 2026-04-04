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
              <div class="info-icon">
                <SvgIcon :name="item.icon" :size="22" />
              </div>
              <div>
                <div class="info-label">{{ item.label }}</div>
                <div class="info-value">{{ item.value }}</div>
              </div>
            </div>
          </div>

          <div class="map-placeholder">
            <div class="map-pin-icon">
              <SvgIcon name="map-pin" :size="36" />
            </div>
            <p>广州市天河区<br/>中山大道建中路5号</p>
          </div>
        </div>

        <!-- 右侧表单 -->
        <div class="contact-form-wrap">
          <h2>发送需求</h2>
          <form class="contact-form" @submit.prevent="handleSubmit">
            <div class="form-row">
              <div class="field">
                <label for="contact-name">您的姓名 *</label>
                <input id="contact-name" v-model="form.name" type="text" placeholder="请输入您的姓名" required />
              </div>
              <div class="field">
                <label for="contact-phone">联系电话 *</label>
                <input
                  id="contact-phone"
                  v-model="form.phone"
                  type="tel"
                  placeholder="请输入您的手机号"
                  :class="{ 'input-error': phoneError }"
                  required
                />
                <span v-if="phoneError" class="field-error" role="alert">{{ phoneError }}</span>
              </div>
            </div>
            <div class="field">
              <label for="contact-type">项目类型</label>
              <select id="contact-type" v-model="form.type">
                <option value="">请选择项目类型</option>
                <option v-for="t in projectTypes" :key="t" :value="t">{{ t }}</option>
              </select>
            </div>
            <div class="field">
              <label for="contact-desc">需求描述 *</label>
              <textarea id="contact-desc" v-model="form.desc" rows="5" placeholder="请简单描述您的项目需求、功能点、预算等..." required></textarea>
            </div>
            <button type="submit" class="btn-submit-form" :disabled="submitting || submitOk">
              {{ submitting ? '发送中...' : submitOk ? '已发送 ✓' : '发送需求 →' }}
            </button>
            <p v-if="submitMsg" class="submit-msg" :class="{ success: submitOk }" role="alert">{{ submitMsg }}</p>
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
          <div
            class="faq-item"
            v-for="faq in faqs"
            :key="faq.id"
            role="button"
            tabindex="0"
            :aria-expanded="faq.open"
            @click="faq.open = !faq.open"
            @keydown.enter.prevent="faq.open = !faq.open"
            @keydown.space.prevent="faq.open = !faq.open"
            :class="{ open: faq.open }"
          >
            <div class="faq-q">
              <span>{{ faq.q }}</span>
              <SvgIcon name="chevron-down" :size="16" class="faq-arrow" :class="{ 'is-open': faq.open }" />
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
import axios from 'axios'
import SvgIcon from '@/components/SvgIcon.vue'

const form = reactive({ name: '', phone: '', type: '', desc: '' })
const submitting = ref(false)
const submitMsg = ref('')
const submitOk = ref(false)
const phoneError = ref('')

// 中国大陆手机号格式：1 开头，第二位 3-9，共 11 位
const PHONE_RE = /^1[3-9]\d{9}$/

function validatePhone() {
  if (!form.phone) {
    phoneError.value = '请输入手机号'
    return false
  }
  if (!PHONE_RE.test(form.phone)) {
    phoneError.value = '请输入有效的 11 位手机号'
    return false
  }
  phoneError.value = ''
  return true
}

const projectTypes = [
  '企业官网', '电商平台', '管理后台系统', '微信小程序', '微信小游戏', 'H5页面', 'App开发', '其他'
]

const infos = [
  { id:1, icon:'map-pin', label:'公司地址', value:'广州市天河区中山大道建中路5号' },
  { id:2, icon:'building', label:'公司全称', value:'广州熠辉智能科技有限公司' },
  { id:3, icon:'globe', label:'官方网站', value:'www.smartyihui.com' },
]

const faqs = ref([
  { id:1, q:'开发一个网站大概需要多少钱？', open: false,
    a:'价格根据功能复杂度而定，简单的企业官网通常在几千元，带后台管理系统的网站在1万元左右。欢迎联系我们获取具体报价。' },
  { id:2, q:'开发周期一般是多久？', open: false,
    a:'简单官网1-2周，小程序2-5周，复杂系统视功能量而定。我们会在签合同前给出明确的排期计划。' },
  { id:3, q:'你们是否可以提供合同和发票？', open: false,
    a:'是的，我们是正规注册的有限责任公司，可以签署正式合同并开具增值税发票，满足企业报账需求。' },
  { id:4, q:'代码和版权归谁所有？', open: false,
    a:'项目验收交付后，代码版权完全归属客户，我们会交付完整源代码，不附加任何隐形费用或版权限制。' },
  { id:5, q:'上线后还需要付费维护吗？', open: false,
    a:'合同范围内的Bug修复是免费的。后续如有新功能迭代，我们会提供优惠的维护服务价格。' },
])

async function handleSubmit() {
  // 手机号格式校验（在浏览器 required 之后额外校验）
  if (!validatePhone()) return

  // 防重复提交：提交中或已成功时直接返回
  if (submitting.value || submitOk.value) return

  submitting.value = true
  submitMsg.value = ''

  try {
    // ── 对接后端接口(这里需要把接口按照格式填写好,我再接上去) ────────────────────────────────────────────
    // 将下方 URL 替换为实际接口地址，例如：
    //   POST https://www.smartyihui.com/api/contact
    // 请求体字段：{ name, phone, type, desc }
    // ────────────────────────────────────────────────────────────
    await axios.post('/* TODO: 填入后端接口地址 */', {
      name: form.name,
      phone: form.phone,
      type: form.type,
      desc: form.desc,
    })

    submitOk.value = true
    submitMsg.value = '需求已发送！我们将在24小时内联系您'
    form.name = ''; form.phone = ''; form.type = ''; form.desc = ''
  } catch (e) {
    submitOk.value = false
    submitMsg.value = e.response?.data?.message || '发送失败，请稍后重试或直接联系我们'
  } finally {
    submitting.value = false
  }
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
.gold { color: var(--gold); }
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
.info-icon { color: var(--gold); flex-shrink: 0; display: flex; }
.info-label { font-size: 12px; color: var(--ink-mute); margin-bottom: 3px; font-weight: 600; }
.info-value { font-size: 14px; color: var(--ink); font-weight: 500; }

.map-placeholder {
  background: var(--ink);
  border-radius: 16px;
  padding: 32px;
  text-align: center;
  color: rgba(255,255,255,0.7);
}
.map-pin-icon {
  color: var(--gold-lt);
  margin-bottom: 12px;
  display: flex;
  justify-content: center;
}
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
.field .input-error { border-color: #e53e3e; }
.field .input-error:focus { border-color: #e53e3e; }
.field-error { font-size: 12px; color: #e53e3e; margin-top: 2px; }

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
.section-inner {
  max-width: 1200px; margin: 0 auto;
  padding: 100px 32px;
}
.section-header { text-align: center; margin-bottom: 60px; }
.section-tag {
  display: inline-block;
  padding: 5px 14px; border-radius: 999px;
  background: var(--gold-bg); color: var(--gold);
  font-size: 13px; font-weight: 700;
  margin-bottom: 16px;
}
.section-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(28px, 3vw, 40px);
  font-weight: 700; margin-bottom: 16px;
}
.faq-section { background: #fff; }
.faq-list { max-width: 760px; margin: 0 auto; display: flex; flex-direction: column; gap: 12px; }
.faq-item {
  border: 1.5px solid var(--border); border-radius: 14px;
  overflow: hidden; cursor: pointer;
  transition: border-color 0.2s;
  outline: none;
}
.faq-item:focus-visible { box-shadow: 0 0 0 3px rgba(200,151,58,0.35); }
.faq-item.open { border-color: var(--gold); }
.faq-q {
  display: flex; justify-content: space-between; align-items: center;
  padding: 18px 20px;
  font-size: 15px; font-weight: 600;
  user-select: none;
}
.faq-arrow {
  color: var(--gold);
  flex-shrink: 0;
  transition: transform 0.2s ease;
}
.faq-arrow.is-open { transform: rotate(180deg); }
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
