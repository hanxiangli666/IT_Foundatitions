<template>
  <div class="page-top-pad">
    <section class="srv-hero">
      <div class="srv-hero-inner">
        <div class="section-tag">服务项目</div>
        <h1 class="srv-hero-title">我们能为您<span class="gold">开发什么</span></h1>
        <p class="srv-hero-sub">无论您有怎样的想法，我们都有能力帮您把它变成真实可用的产品</p>
      </div>
    </section>

    <section class="srv-detail" v-for="(svc, idx) in services" :key="svc.id"
      :class="{ 'bg-alt': idx % 2 === 1 }">
      <div class="srv-detail-inner" :class="{ reversed: idx % 2 === 1 }">
        <div class="srv-d-visual">
          <div class="srv-d-icon-bg" :style="`background: ${svc.bgColor}; color: ${svc.iconColor}`">
            <SvgIcon :name="svc.icon" :size="80" />
          </div>
        </div>
        <div class="srv-d-text">
          <div class="section-tag">{{ svc.category }}</div>
          <h2 class="srv-d-title">{{ svc.title }}</h2>
          <p class="srv-d-desc">{{ svc.desc }}</p>
          <ul class="srv-features">
            <li v-for="f in svc.features" :key="f">
              <SvgIcon name="check" :size="16" class="feat-check" />
              {{ f }}
            </li>
          </ul>
          <div class="srv-techs">
            <span class="tech-tag" v-for="t in svc.techs" :key="t">{{ t }}</span>
          </div>
        </div>
      </div>
    </section>

    <section class="process-section">
      <div class="section-inner">
        <div class="section-header">
          <div class="section-tag">开发流程</div>
          <h2 class="section-title">从需求到上线，全程专业保障</h2>
        </div>
        <div class="process-steps">
          <div class="proc-step" v-for="(s, i) in steps" :key="s.id">
            <div class="proc-num">{{ String(i+1).padStart(2,'0') }}</div>
            <div class="proc-icon">
              <SvgIcon :name="s.icon" :size="28" />
            </div>
            <h4>{{ s.title }}</h4>
            <p>{{ s.desc }}</p>
            <div class="proc-arrow" v-if="i < steps.length - 1" aria-hidden="true">→</div>
          </div>
        </div>
      </div>
    </section>

    <section class="quote-cta">
      <div class="qc-inner">
        <h2>准备好开始您的项目了吗？</h2>
        <p>告诉我们您的需求，我们将在24小时内给您回复</p>
        <router-link to="/contact" class="btn-primary lg">免费咨询报价</router-link>
      </div>
    </section>
  </div>
</template>

<script setup>
import SvgIcon from '@/components/SvgIcon.vue'

const services = [
  {
    id: 1, icon: 'globe', category: '01 / 网站开发',
    title: '企业官网 & Web应用',
    bgColor: 'linear-gradient(135deg, #dbeafe, #eff6ff)',
    iconColor: '#3b82f6',
    desc: '我们为企业打造专业、美观、高性能的网站，包括企业官网、产品展示站、电商平台、管理后台系统等，响应式设计确保在PC和移动端都有最佳体验。',
    features: [
      '自适应响应式设计，完美适配手机/平板/PC',
      '支持后台内容管理，无需编程即可更新内容',
      'SEO友好，有助于提升搜索引擎排名',
      '高性能优化，页面加载速度快',
      '支持多语言、多主题切换',
    ],
    techs: ['Vue 3', 'React', 'Node.js', 'Go', 'MySQL', 'Nginx'],
  },
  {
    id: 2, icon: 'settings', category: '02 / 软件定制',
    title: '企业软件定制开发',
    bgColor: 'linear-gradient(135deg, #fef3c7, #fdf6e8)',
    iconColor: '#c8973a',
    desc: '根据您的业务需求量身定制软件系统，包括ERP、CRM、OA、数据管理平台等，从需求分析到系统上线，提供全流程技术服务。',
    features: [
      '深度需求分析，量身打造最适合的技术方案',
      '前后端一体化开发，接口清晰规范',
      '完整的接口文档与源代码交付',
      '系统部署与上线支持',
      '上线后提供持续维护与版本迭代',
    ],
    techs: ['Vue', 'React', 'Python', 'Go', 'MySQL', 'Redis', 'Docker'],
  },
  {
    id: 3, icon: 'smartphone', category: '03 / 微信小程序',
    title: '微信小程序开发',
    bgColor: 'linear-gradient(135deg, #d1fae5, #ecfdf5)',
    iconColor: '#10b981',
    desc: '覆盖电商、预约、打卡、社区等各类场景，帮助您快速触达微信生态内超过10亿的用户，开发周期短，上线成本低。',
    features: [
      '原生微信小程序开发，性能最优',
      '支持Taro跨端，一套代码多端运行',
      '微信支付、登录、分享一站式接入',
      '微信云开发，省去服务器运维成本',
      '小程序审核与上线全程辅助',
    ],
    techs: ['微信原生', 'Taro', 'Uni-app', '微信云开发', '微信支付'],
  },
  {
    id: 4, icon: 'gamepad', category: '04 / 小游戏开发',
    title: '微信小游戏 & H5游戏',
    bgColor: 'linear-gradient(135deg, #ede9fe, #f5f3ff)',
    iconColor: '#8b5cf6',
    desc: '休闲益智、测试类、互动营销等各类小游戏，支持微信小游戏平台上线，内置广告变现，帮助您快速构建用户流量与商业模式。',
    features: [
      '休闲、益智、测试类等多种游戏类型',
      '微信小游戏平台上线与审核支持',
      '内置微信激励视频广告变现',
      '游戏内排行榜、社交分享功能',
      '同时提供网页版、App版本打包',
    ],
    techs: ['Cocos Creator', 'Phaser', 'PixiJS', '微信小游戏', 'H5'],
  },
]

const steps = [
  { id:1, icon:'message-square', title:'需求沟通', desc:'了解您的业务目标和功能需求' },
  { id:2, icon:'ruler', title:'方案设计', desc:'出具技术方案与报价，确定排期' },
  { id:3, icon:'laptop', title:'开发实现', desc:'按阶段开发，定期汇报进度' },
  { id:4, icon:'flask', title:'测试验收', desc:'严格测试，您确认验收后交付' },
  { id:5, icon:'rocket', title:'上线部署', desc:'协助部署上线，稳定运行保障' },
]
</script>

<style scoped>
.page-top-pad { padding-top: 80px; }
.srv-hero {
  background: linear-gradient(135deg, #f8fafc 0%, var(--gold-bg) 100%);
  padding: 80px 32px;
  border-bottom: 1px solid var(--border);
}
.srv-hero-inner { max-width: 700px; margin: 0 auto; text-align: center; }
.srv-hero-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(30px, 4vw, 48px);
  font-weight: 700; margin: 16px 0 16px; line-height: 1.3;
}
.gold { color: var(--gold); }
.srv-hero-sub { color: var(--ink-soft); font-size: 16px; }

.srv-detail { padding: 80px 32px; }
.bg-alt { background: #f8fafc; }
.srv-detail-inner {
  max-width: 1100px; margin: 0 auto;
  display: grid;
  grid-template-columns: 1fr 1.5fr;
  gap: 64px;
  align-items: center;
}
.srv-detail-inner.reversed { direction: rtl; }
.srv-detail-inner.reversed > * { direction: ltr; }

.srv-d-visual {
  display: flex; align-items: center; justify-content: center;
}
.srv-d-icon-bg {
  width: 220px; height: 220px; border-radius: 32px;
  display: flex; align-items: center; justify-content: center;
  box-shadow: 0 20px 50px rgba(0,0,0,0.08);
}

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
.section-header { text-align: center; margin-bottom: 60px; }
.section-inner {
  max-width: 1200px; margin: 0 auto;
  padding: 100px 32px;
}

.srv-d-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(22px, 2.5vw, 32px);
  font-weight: 700; margin: 12px 0 16px;
}
.srv-d-desc { color: var(--ink-soft); font-size: 15px; line-height: 1.8; margin-bottom: 20px; }

.srv-features {
  list-style: none; display: flex; flex-direction: column; gap: 10px;
  margin-bottom: 24px;
}
.srv-features li { display: flex; align-items: flex-start; gap: 10px; font-size: 14px; color: var(--ink-soft); }
.feat-check { color: var(--gold); flex-shrink: 0; margin-top: 2px; }

.srv-techs { display: flex; flex-wrap: wrap; gap: 8px; }
.tech-tag {
  padding: 5px 12px; border-radius: 8px;
  background: var(--ink); color: rgba(255,255,255,0.8);
  font-size: 12px; font-weight: 600;
  font-family: 'DM Sans', monospace;
}

/* Process */
.process-section { background: var(--ink); padding: 0; }
.process-section .section-inner { padding: 100px 32px; }
.process-section .section-tag { background: rgba(200,151,58,0.15); }
.process-section .section-title { color: #fff; }
.process-steps {
  display: flex;
  gap: 0;
  flex-wrap: wrap;
  position: relative;
}
.proc-step {
  flex: 1; min-width: 140px;
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,255,255,0.1);
  border-radius: 16px;
  padding: 24px 20px;
  position: relative;
  transition: all 0.3s;
}
.proc-step:hover { background: rgba(200,151,58,0.1); border-color: rgba(200,151,58,0.3); }
.proc-num {
  font-size: 13px; font-weight: 700;
  color: var(--gold); opacity: 0.6; margin-bottom: 10px;
}
.proc-icon {
  color: var(--gold);
  margin-bottom: 10px;
  display: flex;
}
.proc-step h4 { color: #fff; font-size: 15px; font-weight: 700; margin-bottom: 6px; }
.proc-step p { color: rgba(255,255,255,0.5); font-size: 13px; line-height: 1.6; }
.proc-arrow {
  position: absolute; right: -14px; top: 50%;
  transform: translateY(-50%);
  color: rgba(200,151,58,0.5); font-size: 18px; z-index: 1;
}

.quote-cta {
  background: linear-gradient(135deg, var(--gold-bg), #fff);
  border-top: 1px solid rgba(200,151,58,0.2);
}
.qc-inner {
  max-width: 640px; margin: 0 auto;
  padding: 100px 32px;
  text-align: center;
  display: flex; flex-direction: column; align-items: center; gap: 16px;
}
.qc-inner h2 {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(24px, 3vw, 36px); font-weight: 700;
}
.qc-inner p { color: var(--ink-soft); font-size: 15px; }

.btn-primary {
  display: inline-flex; align-items: center;
  background: var(--gold); color: #fff;
  padding: 12px 28px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 15px;
  transition: all 0.2s; border: 2px solid var(--gold);
}
.btn-primary:hover { background: var(--gold-lt); border-color: var(--gold-lt); transform: translateY(-2px); }
.btn-primary.lg { padding: 15px 36px; font-size: 16px; }

@media (max-width: 900px) {
  .srv-detail-inner, .srv-detail-inner.reversed {
    grid-template-columns: 1fr;
    direction: ltr;
    gap: 32px;
  }
  .srv-d-icon-bg { width: 140px; height: 140px; }
  .process-steps { gap: 12px; }
  .proc-arrow { display: none; }
}
</style>
