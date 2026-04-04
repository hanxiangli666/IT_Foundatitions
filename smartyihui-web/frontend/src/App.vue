<template>
  <div class="site-root">
    <!-- 顶部导航 -->
    <header class="navbar" :class="{ scrolled: isScrolled, 'menu-open': menuOpen }">
      <div class="navbar-inner">
        <router-link to="/" class="brand">
          <span class="brand-icon">熠</span>
          <span class="brand-text">广州熠辉智能科技</span>
        </router-link>

        <nav class="nav-links">
          <router-link to="/">首页</router-link>
          <router-link to="/about">关于我们</router-link>
          <router-link to="/services">服务项目</router-link>
          <router-link to="/contact">联系方式</router-link>
        </nav>

        <div class="nav-actions">
          <button class="btn-login" @click="showLogin = true">登录 / 注册</button>
          <button
            class="hamburger"
            @click="menuOpen = !menuOpen"
            :aria-expanded="menuOpen"
            aria-controls="mobile-menu"
            aria-label="菜单"
          >
            <span></span><span></span><span></span>
          </button>
        </div>
      </div>

      <!-- 移动端菜单 -->
      <div id="mobile-menu" class="mobile-menu" v-if="menuOpen">
        <router-link to="/" @click="menuOpen = false">首页</router-link>
        <router-link to="/about" @click="menuOpen = false">关于我们</router-link>
        <router-link to="/services" @click="menuOpen = false">服务项目</router-link>
        <router-link to="/contact" @click="menuOpen = false">联系方式</router-link>
        <button class="btn-login-mobile" @click="showLogin = true; menuOpen = false">登录 / 注册</button>
      </div>
    </header>

    <main>
      <router-view v-slot="{ Component }">
        <transition name="page" mode="out-in">
          <component :is="Component" :key="$route.path" />
        </transition>
      </router-view>
    </main>

    <footer class="site-footer">
      <div class="footer-inner">
        <div class="footer-brand">
          <span class="brand-icon sm">熠</span>
          <span>广州熠辉智能科技有限公司</span>
        </div>
        <p class="footer-addr">广州市天河区中山大道建中路5号</p>
        <p class="footer-copy">© 2026 广州熠辉智能科技有限公司. All rights reserved.</p>
      </div>
    </footer>

    <!-- 登录/注册弹窗 -->
    <div
      class="modal-overlay"
      v-if="showLogin"
      @click.self="showLogin = false"
      role="dialog"
      aria-modal="true"
      aria-labelledby="modal-title"
    >
      <div class="modal-box">
        <button class="modal-close" @click="showLogin = false" aria-label="关闭">
          <SvgIcon name="x" :size="16" />
        </button>
        <div class="modal-tabs" id="modal-title">
          <button :class="{ active: authTab === 'login' }" @click="authTab = 'login'">登录</button>
          <button :class="{ active: authTab === 'register' }" @click="authTab = 'register'">注册</button>
        </div>

        <!-- 登录 -->
        <form v-if="authTab === 'login'" class="auth-form" @submit.prevent="handleLogin">
          <div class="field">
            <label>用户名</label>
            <input v-model="loginForm.username" type="text" placeholder="请输入用户名" required />
          </div>
          <div class="field">
            <label>密码</label>
            <input v-model="loginForm.password" type="password" placeholder="请输入密码" required />
          </div>
          <button type="submit" class="btn-submit" :disabled="authLoading">
            {{ authLoading ? '登录中...' : '登 录' }}
          </button>
          <p v-if="authError" class="auth-error">{{ authError }}</p>
        </form>

        <!-- 注册 -->
        <form v-else class="auth-form" @submit.prevent="handleRegister">
          <div class="field">
            <label>用户名</label>
            <input v-model="registerForm.username" type="text" placeholder="请输入用户名" required />
          </div>
          <div class="field">
            <label>密码</label>
            <input v-model="registerForm.password" type="password" placeholder="至少6位密码" required />
          </div>
          <div class="field">
            <label>手机号</label>
            <div class="phone-row">
              <input v-model="registerForm.phone" type="tel" placeholder="请输入手机号" required />
              <button type="button" class="btn-sms" @click="sendSms" :disabled="smsCooldown > 0">
                {{ smsCooldown > 0 ? `${smsCooldown}s` : '获取验证码' }}
              </button>
            </div>
          </div>
          <div class="field">
            <label>验证码</label>
            <input v-model="registerForm.verifyCode" type="text" placeholder="请输入短信验证码" required />
          </div>
          <button type="submit" class="btn-submit" :disabled="authLoading">
            {{ authLoading ? '注册中...' : '注 册' }}
          </button>
          <p v-if="authError" class="auth-error">{{ authError }}</p>
          <p v-if="authSuccess" class="auth-success">{{ authSuccess }}</p>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import axios from 'axios'
import SvgIcon from '@/components/SvgIcon.vue'

const isScrolled = ref(false)
const menuOpen = ref(false)
const showLogin = ref(false)
const authTab = ref('login')
const authLoading = ref(false)
const authError = ref('')
const authSuccess = ref('')
const smsCooldown = ref(0)

const loginForm = ref({ username: '', password: '' })
const registerForm = ref({ username: '', password: '', phone: '', verifyCode: '' })

const SSO_BASE = 'https://www.smartyihui.com/api/sso'

function handleScroll() {
  isScrolled.value = window.scrollY > 40
}

onMounted(() => window.addEventListener('scroll', handleScroll))
onUnmounted(() => window.removeEventListener('scroll', handleScroll))

async function handleLogin() {
  authError.value = ''
  authLoading.value = true
  try {
    await axios.post(`${SSO_BASE}/auth/login`, loginForm.value)
    authSuccess.value = '登录成功！'
    setTimeout(() => { showLogin.value = false; authSuccess.value = '' }, 1200)
  } catch (e) {
    authError.value = e.response?.data?.message || '登录失败，请检查账号密码'
  } finally {
    authLoading.value = false
  }
}

async function handleRegister() {
  authError.value = ''
  authLoading.value = true
  try {
    await axios.post(`${SSO_BASE}/auth/register`, registerForm.value)
    authSuccess.value = '注册成功！请登录'
    authTab.value = 'login'
  } catch (e) {
    authError.value = e.response?.data?.message || '注册失败，请稍后重试'
  } finally {
    authLoading.value = false
  }
}

async function sendSms() {
  if (!registerForm.value.phone) {
    authError.value = '请先填写手机号'
    return
  }
  try {
    await axios.post(`${SSO_BASE}/auth/send-code`, { phone: registerForm.value.phone })
    smsCooldown.value = 60
    const timer = setInterval(() => {
      smsCooldown.value--
      if (smsCooldown.value <= 0) clearInterval(timer)
    }, 1000)
  } catch (e) {
    authError.value = e.response?.data?.message || '发送失败'
  }
}
</script>

<style>
@import url('https://fonts.googleapis.com/css2?family=Noto+Serif+SC:wght@400;600;700&family=DM+Sans:wght@300;400;500;600&display=swap');

*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

:root {
  --ink: #0d1117;
  --ink-soft: #3d4450;
  --ink-mute: #7a8394;
  --gold: #c8973a;
  --gold-lt: #e8b84b;
  --gold-bg: #fdf6e8;
  --surface: #ffffff;
  --border: #e8eaed;
  --radius: 12px;
}

html { scroll-behavior: smooth; }
body {
  font-family: 'DM Sans', 'PingFang SC', 'Microsoft YaHei', sans-serif;
  color: var(--ink);
  background: var(--surface);
  line-height: 1.6;
}

.site-root { min-height: 100vh; display: flex; flex-direction: column; }
main { flex: 1; }

/* ── NAVBAR ── */
.navbar {
  position: fixed; top: 0; left: 0; right: 0; z-index: 100;
  transition: background 0.3s, box-shadow 0.3s;
  background: transparent;
}
.navbar.scrolled {
  background: rgba(255,255,255,0.96);
  box-shadow: 0 1px 0 var(--border);
  backdrop-filter: blur(12px);
}
.navbar-inner {
  max-width: 1200px; margin: 0 auto;
  padding: 20px 32px;
  display: flex; align-items: center; gap: 40px;
}
.brand {
  display: flex; align-items: center; gap: 10px;
  text-decoration: none; color: var(--ink);
  font-family: 'Noto Serif SC', serif;
  font-weight: 700; font-size: 17px;
  white-space: nowrap;
}
.brand-icon {
  width: 36px; height: 36px;
  background: var(--gold); color: #fff;
  border-radius: 8px;
  display: flex; align-items: center; justify-content: center;
  font-size: 18px; font-weight: 700;
  font-family: 'Noto Serif SC', serif;
  flex-shrink: 0;
}
.brand-icon.sm { width: 28px; height: 28px; font-size: 14px; border-radius: 6px; }
.nav-links {
  display: flex; gap: 8px; align-items: center; flex: 1;
}
.nav-links a {
  text-decoration: none; color: var(--ink-soft);
  padding: 8px 14px; border-radius: 8px;
  font-size: 15px; font-weight: 500;
  transition: all 0.2s;
}
.nav-links a:hover, .nav-links a.router-link-active {
  color: var(--gold); background: var(--gold-bg);
}
.nav-actions { margin-left: auto; display: flex; align-items: center; gap: 12px; }
.btn-login {
  background: var(--gold); color: #fff;
  border: none; cursor: pointer;
  padding: 9px 22px; border-radius: 8px;
  font-size: 14px; font-weight: 600;
  font-family: inherit;
  transition: background 0.2s, transform 0.15s;
}
.btn-login:hover { background: var(--gold-lt); transform: translateY(-1px); }
.hamburger {
  display: none; flex-direction: column; gap: 5px;
  background: none; border: none; cursor: pointer; padding: 4px;
}
.hamburger span {
  display: block; width: 22px; height: 2px;
  background: var(--ink); border-radius: 2px; transition: 0.3s;
}
.mobile-menu {
  display: flex; flex-direction: column; gap: 4px;
  padding: 12px 24px 20px;
  background: rgba(255,255,255,0.98);
  border-top: 1px solid var(--border);
}
.mobile-menu a {
  text-decoration: none; color: var(--ink-soft);
  padding: 12px 16px; border-radius: 8px;
  font-size: 16px; font-weight: 500;
}
.mobile-menu a:hover { background: var(--gold-bg); color: var(--gold); }
.btn-login-mobile {
  margin-top: 8px; background: var(--gold); color: #fff;
  border: none; cursor: pointer;
  padding: 12px; border-radius: 8px;
  font-size: 15px; font-weight: 600; font-family: inherit;
}

/* ── FOOTER ── */
.site-footer {
  background: var(--ink); color: rgba(255,255,255,0.6);
  padding: 40px 32px;
}
.footer-inner {
  max-width: 1200px; margin: 0 auto;
  display: flex; flex-direction: column; gap: 8px; align-items: center;
  text-align: center;
}
.footer-brand {
  display: flex; align-items: center; gap: 10px;
  color: rgba(255,255,255,0.9); font-size: 15px; font-weight: 600;
  margin-bottom: 4px;
}
.footer-addr { font-size: 14px; }
.footer-copy { font-size: 13px; opacity: 0.7; }

/* ── MODAL ── */
.modal-overlay {
  position: fixed; inset: 0; z-index: 200;
  background: rgba(13,17,23,0.6);
  display: flex; align-items: center; justify-content: center;
  padding: 20px;
  backdrop-filter: blur(4px);
  animation: fadeIn 0.2s ease;
}
@keyframes fadeIn { from { opacity: 0; } to { opacity: 1; } }

.modal-box {
  background: #fff; border-radius: 20px;
  padding: 36px; width: 100%; max-width: 420px;
  position: relative;
  box-shadow: 0 24px 60px rgba(0,0,0,0.2);
  animation: slideUp 0.25s ease;
}
@keyframes slideUp { from { transform: translateY(20px); opacity: 0; } to { transform: translateY(0); opacity: 1; } }

.modal-close {
  position: absolute; top: 16px; right: 16px;
  background: none; border: none; cursor: pointer;
  font-size: 18px; color: var(--ink-mute);
  width: 32px; height: 32px; border-radius: 8px;
  display: flex; align-items: center; justify-content: center;
  transition: background 0.2s;
}
.modal-close:hover { background: #f1f3f5; }

.modal-tabs {
  display: flex; gap: 4px; margin-bottom: 28px;
  background: #f1f3f5; border-radius: 10px; padding: 4px;
}
.modal-tabs button {
  flex: 1; padding: 9px; border-radius: 8px;
  border: none; cursor: pointer; font-size: 15px;
  font-weight: 600; font-family: inherit;
  color: var(--ink-mute); background: transparent;
  transition: all 0.2s;
}
.modal-tabs button.active {
  background: #fff; color: var(--gold);
  box-shadow: 0 1px 4px rgba(0,0,0,0.1);
}

.auth-form { display: flex; flex-direction: column; gap: 18px; }
.field { display: flex; flex-direction: column; gap: 6px; }
.field label { font-size: 14px; font-weight: 600; color: var(--ink-soft); }
.field input {
  padding: 11px 14px; border-radius: 10px;
  border: 1.5px solid var(--border);
  font-size: 15px; font-family: inherit;
  transition: border-color 0.2s;
  outline: none;
}
.field input:focus { border-color: var(--gold); }
.phone-row { display: flex; gap: 8px; }
.phone-row input { flex: 1; }
.btn-sms {
  padding: 0 14px; border-radius: 10px;
  border: 1.5px solid var(--gold); color: var(--gold);
  background: transparent; cursor: pointer;
  font-size: 13px; font-weight: 600; font-family: inherit;
  white-space: nowrap; transition: all 0.2s;
}
.btn-sms:hover:not(:disabled) { background: var(--gold-bg); }
.btn-sms:disabled { opacity: 0.5; cursor: not-allowed; }

.btn-submit {
  background: var(--gold); color: #fff;
  border: none; cursor: pointer;
  padding: 13px; border-radius: 10px;
  font-size: 16px; font-weight: 700; font-family: inherit;
  transition: all 0.2s; margin-top: 4px;
}
.btn-submit:hover:not(:disabled) { background: var(--gold-lt); transform: translateY(-1px); }
.btn-submit:disabled { opacity: 0.6; cursor: not-allowed; }
.auth-error { color: #e53e3e; font-size: 14px; text-align: center; }
.auth-success { color: #38a169; font-size: 14px; text-align: center; }

@media (max-width: 768px) {
  .nav-links { display: none; }
  .hamburger { display: flex; }
  .navbar-inner { padding: 16px 20px; }
  .btn-login { display: none; }
}

/* ── PAGE TRANSITION ── */
.page-enter-active,
.page-leave-active {
  transition: opacity 0.18s ease, transform 0.18s ease;
}
.page-enter-from {
  opacity: 0;
  transform: translateY(8px);
}
.page-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
</style>
