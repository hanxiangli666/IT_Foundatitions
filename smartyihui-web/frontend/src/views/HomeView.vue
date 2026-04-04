<template>
  <div class="home-root">
  <!-- Hero 封面 -->
  <section class="hero">
    <div class="hero-bg">
      <div class="hero-blob b1"></div>
      <div class="hero-blob b2"></div>
      <div class="hero-grid"></div>
    </div>
    <div class="hero-inner">
      <div class="hero-text">
        <div class="hero-badge">专业软件开发服务商</div>
        <h1 class="hero-title">
          广州熠辉<br/>
          <span class="gold">智能科技</span>
        </h1>
        <p class="hero-sub">
          专注网站开发、软件定制、微信小程序与小游戏开发<br/>
          用技术赋能您的业务增长
        </p>
        <div class="hero-actions">
          <router-link to="/services" class="btn-primary">了解服务</router-link>
          <router-link to="/contact" class="btn-outline">立即咨询</router-link>
        </div>
      </div>
      <div class="hero-visual">
        <div class="handshake-card">
          <div class="hc-img">
            <!-- 西装男握手 SVG 插图 -->
            <svg viewBox="0 0 320 280" xmlns="http://www.w3.org/2000/svg" class="biz-svg">
              <!-- 背景圆 -->
              <circle cx="160" cy="140" r="120" fill="#fdf6e8" opacity="0.8"/>
              <!-- 左人 -->
              <rect x="40" y="100" width="60" height="90" rx="8" fill="#1a2332"/>
              <rect x="48" y="108" width="44" height="12" rx="3" fill="#c8973a" opacity="0.6"/>
              <circle cx="70" cy="88" r="22" fill="#f5d5b8"/>
              <rect x="55" y="82" width="30" height="10" rx="5" fill="#1a2332"/>
              <!-- 右人 -->
              <rect x="220" y="100" width="60" height="90" rx="8" fill="#2d4a6e"/>
              <rect x="228" y="108" width="44" height="12" rx="3" fill="#c8973a" opacity="0.6"/>
              <circle cx="250" cy="88" r="22" fill="#f5d5b8"/>
              <rect x="235" y="82" width="30" height="10" rx="5" fill="#2d4a6e"/>
              <!-- 握手臂 -->
              <path d="M100 155 Q130 145 160 150 Q190 155 220 145" stroke="#1a2332" stroke-width="18" fill="none" stroke-linecap="round"/>
              <path d="M100 155 Q130 145 160 150 Q190 155 220 145" stroke="#f5d5b8" stroke-width="10" fill="none" stroke-linecap="round"/>
              <!-- 握手拳 -->
              <ellipse cx="160" cy="150" rx="24" ry="16" fill="#f5d5b8"/>
              <ellipse cx="160" cy="150" rx="24" ry="16" fill="none" stroke="#e8c49a" stroke-width="2"/>
              <!-- 金色光芒 -->
              <circle cx="160" cy="150" r="6" fill="#c8973a" opacity="0.9"/>
              <line x1="160" y1="126" x2="160" y2="118" stroke="#c8973a" stroke-width="2" opacity="0.6"/>
              <line x1="160" y1="174" x2="160" y2="182" stroke="#c8973a" stroke-width="2" opacity="0.6"/>
              <line x1="136" y1="150" x2="128" y2="150" stroke="#c8973a" stroke-width="2" opacity="0.6"/>
              <line x1="184" y1="150" x2="192" y2="150" stroke="#c8973a" stroke-width="2" opacity="0.6"/>
              <!-- 装饰文字 -->
              <text x="160" y="220" text-anchor="middle" fill="#c8973a" font-size="13" font-family="serif" font-weight="600">携手共创 · 共赢未来</text>
            </svg>
          </div>
          <div class="hc-footer">
            <div class="hc-stat"><span class="stat-num">50+</span><span class="stat-lbl">成功案例</span></div>
            <div class="hc-divider"></div>
            <div class="hc-stat"><span class="stat-num">100%</span><span class="stat-lbl">自主研发</span></div>
            <div class="hc-divider"></div>
            <div class="hc-stat"><span class="stat-num">7×24</span><span class="stat-lbl">技术支持</span></div>
          </div>
        </div>
      </div>
    </div>
    <!-- 滚动提示 -->
    <div class="scroll-hint" aria-hidden="true">
      <span>向下滚动</span>
      <div class="scroll-arrow"></div>
    </div>
  </section>

  <!-- 服务概览 -->
  <section class="services-preview">
    <div class="section-inner">
      <div class="section-header">
        <div class="section-tag">我们的服务</div>
        <h2 class="section-title">一站式软件开发解决方案</h2>
        <p class="section-desc">从创意到落地，我们提供全流程技术服务，助力您的产品快速触达用户</p>
      </div>
      <div class="services-grid">
        <div class="srv-card" v-for="svc in services" :key="svc.id" :style="`--acc: ${svc.color}`">
          <div class="srv-icon">
            <SvgIcon :name="svc.icon" :size="36" />
          </div>
          <h3>{{ svc.title }}</h3>
          <p>{{ svc.desc }}</p>
          <ul class="srv-tags">
            <li v-for="t in svc.tags" :key="t">{{ t }}</li>
          </ul>
        </div>
      </div>
      <div class="srv-cta">
        <router-link to="/services" class="btn-primary">查看全部服务</router-link>
      </div>
    </div>
  </section>

  <!-- 多端产品展示 -->
  <section class="multiplatform">
    <div class="section-inner">
      <div class="section-header">
        <div class="section-tag">产品体验</div>
        <h2 class="section-title">一套产品，多端覆盖</h2>
        <p class="section-desc">我们的产品支持全平台部署，让您的用户随时随地享受一致体验</p>
      </div>
      <div class="platform-list">
        <div class="platform-item" v-for="p in platforms" :key="p.id">
          <div class="platform-icon">
            <SvgIcon :name="p.icon" :size="32" />
          </div>
          <div class="platform-info">
            <h4>{{ p.name }}</h4>
            <p>{{ p.desc }}</p>
          </div>
          <a v-if="p.link" :href="p.link" target="_blank" rel="noopener noreferrer" class="platform-link">体验</a>
        </div>
      </div>
    </div>
  </section>

  <!-- 为什么选择我们 -->
  <section class="why-us">
    <div class="section-inner">
      <div class="section-header light">
        <div class="section-tag">我们的优势</div>
        <h2 class="section-title">为什么选择熠辉</h2>
      </div>
      <div class="why-grid">
        <div class="why-card" v-for="w in whys" :key="w.id">
          <div class="why-num">{{ w.num }}</div>
          <h4>{{ w.title }}</h4>
          <p>{{ w.desc }}</p>
        </div>
      </div>
    </div>
  </section>

  <!-- CTA -->
  <section class="cta-band">
    <div class="cta-inner">
      <h2>有项目想法？<span class="gold">让我们来实现它</span></h2>
      <p>无论是企业官网、管理系统还是微信小程序，我们都能快速响应</p>
      <router-link to="/contact" class="btn-primary lg">立即联系我们</router-link>
    </div>
  </section>
  </div>
</template>

<script setup>
import SvgIcon from '@/components/SvgIcon.vue'

const services = [
  {
    id: 1, icon: 'globe', color: '#3b82f6',
    title: '网站开发',
    desc: '企业官网、商城、管理后台，响应式设计适配所有终端',
    tags: ['Vue / React', '响应式', 'SEO优化', '后台管理']
  },
  {
    id: 2, icon: 'settings', color: '#c8973a',
    title: '软件定制开发',
    desc: '根据业务需求量身定制，提供完整的前后端一体化解决方案',
    tags: ['需求分析', '架构设计', 'API开发', '持续维护']
  },
  {
    id: 3, icon: 'smartphone', color: '#10b981',
    title: '微信小程序',
    desc: '覆盖电商、预约、社区等场景，快速触达微信生态10亿用户',
    tags: ['原生开发', 'Taro跨端', '云开发', '支付接入']
  },
  {
    id: 4, icon: 'gamepad', color: '#8b5cf6',
    title: '小游戏开发',
    desc: '休闲益智、H5小游戏，支持微信小游戏平台上线与变现',
    tags: ['Cocos', 'Phaser', '微信平台', '广告变现']
  },
]

const platforms = [
  { id: 1, icon: 'monitor', name: '网页版', desc: '无需下载，打开浏览器即可使用，支持PC与移动端', link: 'https://www.smartyihui.com' },
  { id: 2, icon: 'message-circle', name: '微信小游戏', desc: '在微信内直接体验，分享传播方便', link: null },
  { id: 3, icon: 'smartphone', name: '手机App', desc: '打包为Android / iOS 应用，原生性能体验', link: null },
  { id: 4, icon: 'mouse', name: '桌面端', desc: '打包为Windows / macOS 桌面应用', link: null },
]

const whys = [
  { id: 1, num: '01', title: '完全自主研发', desc: '所有代码由团队独立开发，无外包，质量可控、代码可交付' },
  { id: 2, num: '02', title: '快速交付', desc: '敏捷开发流程，MVP版本最快2周上线，快速验证您的商业想法' },
  { id: 3, num: '03', title: '合法合规', desc: '正规注册企业，可签合同、开发票，提供实习/工作证明' },
  { id: 4, num: '04', title: '长期维护', desc: '上线不是终点，持续提供技术支持、Bug修复和版本迭代服务' },
]
</script>

<style scoped>
.home-root { display: block; }

/* ── HERO ── */
.hero {
  min-height: 100vh;
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: center;
  overflow: hidden;
  padding-top: 80px;
}
.hero-bg {
  position: absolute; inset: 0; z-index: 0; pointer-events: none;
}
.hero-blob {
  position: absolute; border-radius: 50%;
  filter: blur(80px); opacity: 0.18;
}
.b1 {
  width: 600px; height: 600px;
  background: radial-gradient(circle, #c8973a, transparent);
  top: -100px; right: -100px;
}
.b2 {
  width: 400px; height: 400px;
  background: radial-gradient(circle, #3b82f6, transparent);
  bottom: 0; left: -80px;
}
.hero-grid {
  position: absolute; inset: 0;
  background-image:
    linear-gradient(rgba(200,151,58,0.05) 1px, transparent 1px),
    linear-gradient(90deg, rgba(200,151,58,0.05) 1px, transparent 1px);
  background-size: 60px 60px;
}

.hero-inner {
  position: relative; z-index: 1;
  max-width: 1200px; margin: 0 auto;
  padding: 80px 32px;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
}
.hero-badge {
  display: inline-block;
  padding: 6px 14px;
  border-radius: 999px;
  background: var(--gold-bg);
  border: 1px solid rgba(200,151,58,0.3);
  color: var(--gold);
  font-size: 13px; font-weight: 600;
  margin-bottom: 20px;
}
.hero-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(40px, 5vw, 64px);
  font-weight: 700;
  line-height: 1.2;
  margin-bottom: 20px;
  color: var(--ink);
}
.gold { color: var(--gold); }
.hero-sub {
  font-size: 17px; color: var(--ink-soft);
  line-height: 1.8; margin-bottom: 36px;
}
.hero-actions { display: flex; gap: 14px; flex-wrap: wrap; }
.btn-primary {
  display: inline-flex; align-items: center;
  background: var(--gold); color: #fff;
  padding: 12px 28px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 15px;
  transition: all 0.2s;
  border: 2px solid var(--gold);
}
.btn-primary:hover { background: var(--gold-lt); border-color: var(--gold-lt); transform: translateY(-2px); }
.btn-primary.lg { padding: 15px 36px; font-size: 16px; }
.btn-outline {
  display: inline-flex; align-items: center;
  padding: 12px 28px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 15px;
  border: 2px solid var(--border);
  color: var(--ink-soft);
  transition: all 0.2s;
}
.btn-outline:hover { border-color: var(--gold); color: var(--gold); transform: translateY(-2px); }

/* Handshake card */
.handshake-card {
  background: #fff;
  border-radius: 24px;
  box-shadow: 0 24px 60px rgba(0,0,0,0.1), 0 0 0 1px var(--border);
  overflow: hidden;
}
.hc-img {
  padding: 32px;
  background: linear-gradient(135deg, #fdf6e8 0%, #fff8f0 100%);
  display: flex; align-items: center; justify-content: center;
}
.biz-svg { width: 100%; max-width: 300px; }
.hc-footer {
  display: flex; align-items: center; justify-content: space-around;
  padding: 20px 24px;
  border-top: 1px solid var(--border);
}
.hc-stat { display: flex; flex-direction: column; align-items: center; gap: 2px; }
.stat-num { font-size: 22px; font-weight: 700; color: var(--gold); font-family: 'Noto Serif SC', serif; }
.stat-lbl { font-size: 12px; color: var(--ink-mute); }
.hc-divider { width: 1px; height: 32px; background: var(--border); }

/* Scroll hint */
.scroll-hint {
  position: absolute; bottom: 32px; left: 50%; transform: translateX(-50%);
  display: flex; flex-direction: column; align-items: center; gap: 8px;
  color: var(--ink-mute); font-size: 12px;
  animation: bounce 2s infinite;
}
.scroll-arrow {
  width: 20px; height: 20px;
  border-right: 2px solid var(--ink-mute);
  border-bottom: 2px solid var(--ink-mute);
  transform: rotate(45deg);
}
@keyframes bounce {
  0%, 100% { transform: translateX(-50%) translateY(0); }
  50% { transform: translateX(-50%) translateY(6px); }
}

/* ── SECTIONS COMMON ── */
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
.section-desc { color: var(--ink-soft); font-size: 16px; max-width: 560px; margin: 0 auto; }

/* ── SERVICES ── */
.services-preview { background: #f8fafc; }
.services-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 24px;
  margin-bottom: 40px;
}
.srv-card {
  background: #fff;
  border-radius: 20px;
  padding: 28px;
  border: 1px solid var(--border);
  transition: all 0.3s;
  cursor: default;
}
.srv-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 16px 40px rgba(0,0,0,0.08);
  border-color: var(--acc);
}
.srv-icon {
  color: var(--acc);
  margin-bottom: 16px;
  display: flex;
}
.srv-card h3 { font-size: 18px; font-weight: 700; margin-bottom: 10px; }
.srv-card p { color: var(--ink-soft); font-size: 14px; margin-bottom: 16px; line-height: 1.7; }
.srv-tags {
  display: flex; flex-wrap: wrap; gap: 6px;
  list-style: none;
}
.srv-tags li {
  padding: 4px 10px; border-radius: 6px;
  font-size: 12px; font-weight: 600;
  background: var(--gold-bg); color: var(--gold);
}
.srv-cta { text-align: center; }

/* ── MULTIPLATFORM ── */
.multiplatform { background: #fff; }
.platform-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 20px;
}
.platform-item {
  display: flex; align-items: center; gap: 16px;
  padding: 24px;
  border-radius: 16px;
  border: 1.5px solid var(--border);
  transition: all 0.25s;
}
.platform-item:hover { border-color: var(--gold); background: var(--gold-bg); }
.platform-icon {
  color: var(--gold);
  flex-shrink: 0;
  display: flex;
}
.platform-info { flex: 1; }
.platform-info h4 { font-size: 16px; font-weight: 700; margin-bottom: 4px; }
.platform-info p { font-size: 13px; color: var(--ink-mute); }
.platform-link {
  text-decoration: none; color: var(--gold);
  font-size: 13px; font-weight: 700;
  padding: 6px 12px; border-radius: 8px;
  border: 1px solid var(--gold);
  transition: all 0.2s;
  flex-shrink: 0;
}
.platform-link:hover { background: var(--gold); color: #fff; }

/* ── WHY US ── */
.why-us { background: var(--ink); }
.why-us .section-tag { background: rgba(200,151,58,0.15); }
.why-us .section-title { color: #fff; }
.why-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 24px;
}
.why-card {
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,255,255,0.1);
  border-radius: 18px;
  padding: 28px;
  transition: all 0.3s;
}
.why-card:hover { background: rgba(200,151,58,0.1); border-color: rgba(200,151,58,0.3); }
.why-num {
  font-size: 40px; font-weight: 800;
  color: var(--gold); opacity: 0.5;
  font-family: 'Noto Serif SC', serif;
  margin-bottom: 12px;
}
.why-card h4 { color: #fff; font-size: 17px; font-weight: 700; margin-bottom: 10px; }
.why-card p { color: rgba(255,255,255,0.55); font-size: 14px; line-height: 1.7; }

/* ── CTA ── */
.cta-band {
  background: linear-gradient(135deg, var(--gold-bg) 0%, #fff8f0 100%);
  border-top: 1px solid rgba(200,151,58,0.2);
}
.cta-inner {
  max-width: 700px; margin: 0 auto;
  padding: 100px 32px;
  text-align: center;
  display: flex; flex-direction: column; align-items: center; gap: 16px;
}
.cta-inner h2 {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(26px, 3vw, 38px);
  font-weight: 700;
}
.cta-inner p { color: var(--ink-soft); font-size: 16px; margin-bottom: 8px; }

/* ── RESPONSIVE ── */
@media (max-width: 900px) {
  .hero-inner { grid-template-columns: 1fr; gap: 40px; padding: 60px 24px; }
  .hero-text { text-align: center; }
  .hero-actions { justify-content: center; }
  .scroll-hint { display: none; }
}
@media (max-width: 600px) {
  .section-inner { padding: 60px 20px; }
  .services-grid { grid-template-columns: 1fr; }
  .platform-list { grid-template-columns: 1fr; }
  .why-grid { grid-template-columns: 1fr; }
}
</style>
