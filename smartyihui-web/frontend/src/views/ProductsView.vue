<template>
  <div class="page-top-pad">

    <!-- 顶部 Hero -->
    <section class="prod-hero">
      <div class="ph-inner">
        <div class="section-tag">自研产品</div>
        <h1 class="ph-title">我们自己做的<span class="gold">游戏与应用</span></h1>
        <p class="ph-sub">
          这些是熠辉团队从0到1独立开发的产品，代码完全自主，持续迭代更新
        </p>
      </div>
    </section>

    <!-- 产品列表 -->
    <section class="prod-list-section">
      <div class="section-inner">
        <div class="prod-grid">
          <div class="prod-card" v-for="p in products" :key="p.id">

            <!-- 左侧图标区 -->
            <div class="prod-visual" :style="`background: ${p.bg}`">
              <div class="prod-icon-wrap" :style="`color: ${p.color}`">
                <SvgIcon :name="p.icon" :size="64" />
              </div>
              <div class="prod-badges">
                <span class="badge" v-for="b in p.badges" :key="b">{{ b }}</span>
              </div>
            </div>

            <!-- 右侧信息区 -->
            <div class="prod-info">
              <div class="prod-header">
                <h2 class="prod-name">{{ p.name }}</h2>
                <span class="prod-status" :class="p.status">{{ p.statusLabel }}</span>
              </div>
              <p class="prod-tagline">{{ p.tagline }}</p>
              <p class="prod-desc">{{ p.desc }}</p>

              <!-- 技术亮点 -->
              <div class="prod-features">
                <div class="feat-item" v-for="f in p.features" :key="f.text">
                  <SvgIcon name="check" :size="14" class="feat-icon" />
                  <span>{{ f.text }}</span>
                </div>
              </div>

              <!-- 技术栈标签 -->
              <div class="prod-techs">
                <span class="tech-tag" v-for="t in p.techs" :key="t">{{ t }}</span>
              </div>

              <!-- 下载按钮区 -->
              <div class="prod-actions">
                <a
                  v-for="dl in p.downloads"
                  :key="dl.label"
                  :href="dl.url"
                  class="dl-btn"
                  :class="dl.primary ? 'primary' : 'secondary'"
                  :download="dl.filename || undefined"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  <SvgIcon :name="dl.icon" :size="16" />
                  {{ dl.label }}
                </a>
              </div>
            </div>

          </div>
        </div>

        <!-- 更多产品提示 -->
        <div class="coming-soon">
          <div class="cs-icon">
            <SvgIcon name="rocket" :size="32" />
          </div>
          <h3>更多产品正在开发中</h3>
          <p>团队持续创作，敬请期待下一款作品</p>
          <router-link to="/contact" class="btn-outline">合作共创 →</router-link>
        </div>
      </div>
    </section>

  </div>
</template>

<script setup>
import SvgIcon from '@/components/SvgIcon.vue'

const products = [
  {
    id: 1,
    name: '未命名冒险（内测版）',
    tagline: '2D 多人联机冒险游戏 · PC 桌面端',
    status: 'beta',
    statusLabel: '内测中',
    icon: 'gamepad',
    bg: 'linear-gradient(135deg, #ede9fe, #f5f3ff)',
    color: '#8b5cf6',
    desc:
      '由熠辉团队自主研发的 2D 多人联机冒险游戏，包含完整的关卡流程、AI 寻路怪物与在线组队系统。' +
      '目前已完成 5 个场景的搭建，支持 Windows 平台运行，现开放内测体验。',
    features: [
      { text: '支持多人联机，实时组队冒险' },
      { text: '5 个场景关卡，含主菜单完整流程' },
      { text: 'AI 自动寻路怪物，战斗有挑战性' },
      { text: '骨骼动画 + 像素完美渲染，画面流畅' },
      { text: 'Burst Compiler 优化，性能表现优秀' },
    ],
    techs: ['Unity', 'C#', 'Netcode for GameObjects', '2D Animation', 'NavMesh', 'Burst'],
    badges: ['多人联机', '2D', 'Windows'],
    downloads: [
      {
        label: '下载 Windows 版（Demo）',
        // ↓ 把这里替换成你实际的下载链接
        url: 'https://www.smartyihui.com/downloads/my-project-demo.zip',
        filename: 'YiHui-Adventure-Demo.zip',
        icon: 'monitor',
        primary: true,
      },
    ],
  },
]
</script>

<style scoped>
.page-top-pad { padding-top: 80px; }

/* Hero */
.prod-hero {
  background: linear-gradient(135deg, #f8fafc 0%, var(--gold-bg) 100%);
  border-bottom: 1px solid var(--border);
  padding: 80px 32px;
}
.ph-inner { max-width: 700px; margin: 0 auto; text-align: center; }
.ph-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(30px, 4vw, 50px);
  font-weight: 700; color: var(--ink);
  margin: 16px 0 16px; line-height: 1.3;
}
.ph-sub { color: var(--ink-soft); font-size: 16px; line-height: 1.8; }
.gold { color: var(--gold); }
.section-tag {
  display: inline-block;
  padding: 5px 14px; border-radius: 999px;
  background: var(--gold-bg); color: var(--gold);
  font-size: 13px; font-weight: 700;
  margin-bottom: 16px;
}

/* List section */
.prod-list-section { background: #f8fafc; }
.section-inner { max-width: 1100px; margin: 0 auto; padding: 80px 32px; }

/* Product card */
.prod-grid { display: flex; flex-direction: column; gap: 32px; margin-bottom: 64px; }

.prod-card {
  background: #fff;
  border-radius: 24px;
  border: 1px solid var(--border);
  overflow: hidden;
  display: grid;
  grid-template-columns: 320px 1fr;
  box-shadow: 0 8px 32px rgba(0,0,0,0.06);
  transition: box-shadow 0.3s;
}
.prod-card:hover { box-shadow: 0 16px 48px rgba(0,0,0,0.1); }

/* Visual */
.prod-visual {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px;
  padding: 40px 24px;
  min-height: 280px;
}
.prod-icon-wrap {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100px; height: 100px;
  border-radius: 28px;
  background: rgba(255,255,255,0.7);
  backdrop-filter: blur(8px);
  box-shadow: 0 8px 24px rgba(0,0,0,0.08);
}
.prod-badges {
  display: flex; flex-wrap: wrap; gap: 8px; justify-content: center;
}
.badge {
  padding: 4px 12px; border-radius: 999px;
  background: rgba(255,255,255,0.8);
  font-size: 12px; font-weight: 700;
  color: var(--ink-soft);
  backdrop-filter: blur(4px);
}

/* Info */
.prod-info { padding: 36px 36px 36px 32px; display: flex; flex-direction: column; gap: 16px; }

.prod-header { display: flex; align-items: center; gap: 12px; flex-wrap: wrap; }
.prod-name {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(20px, 2vw, 26px);
  font-weight: 700; color: var(--ink);
}
.prod-status {
  padding: 4px 12px; border-radius: 999px;
  font-size: 12px; font-weight: 700;
}
.prod-status.beta { background: #fef3c7; color: #92400e; }
.prod-status.stable { background: #d1fae5; color: #065f46; }
.prod-status.alpha { background: #ede9fe; color: #4c1d95; }

.prod-tagline { color: var(--gold); font-size: 14px; font-weight: 600; }
.prod-desc { color: var(--ink-soft); font-size: 15px; line-height: 1.8; }

/* Features */
.prod-features { display: flex; flex-direction: column; gap: 8px; }
.feat-item {
  display: flex; align-items: flex-start; gap: 8px;
  font-size: 14px; color: var(--ink-soft);
}
.feat-icon { color: var(--gold); flex-shrink: 0; margin-top: 2px; }

/* Techs */
.prod-techs { display: flex; flex-wrap: wrap; gap: 8px; }
.tech-tag {
  padding: 4px 12px; border-radius: 8px;
  background: var(--ink); color: rgba(255,255,255,0.8);
  font-size: 12px; font-weight: 600;
}

/* Actions */
.prod-actions { display: flex; gap: 12px; flex-wrap: wrap; margin-top: 4px; }
.dl-btn {
  display: inline-flex; align-items: center; gap: 8px;
  padding: 11px 22px; border-radius: 10px;
  text-decoration: none; font-size: 14px; font-weight: 700;
  transition: all 0.2s;
}
.dl-btn.primary {
  background: var(--gold); color: #fff;
  border: 2px solid var(--gold);
}
.dl-btn.primary:hover { background: var(--gold-lt); border-color: var(--gold-lt); transform: translateY(-1px); }
.dl-btn.secondary {
  background: transparent; color: var(--ink-soft);
  border: 2px solid var(--border);
}
.dl-btn.secondary:hover { border-color: var(--gold); color: var(--gold); transform: translateY(-1px); }

/* Coming soon */
.coming-soon {
  text-align: center;
  padding: 48px;
  border-radius: 20px;
  border: 2px dashed var(--border);
  display: flex; flex-direction: column; align-items: center; gap: 12px;
}
.cs-icon { color: var(--gold); opacity: 0.6; }
.coming-soon h3 {
  font-family: 'Noto Serif SC', serif;
  font-size: 20px; font-weight: 700; color: var(--ink);
}
.coming-soon p { color: var(--ink-mute); font-size: 14px; }
.btn-outline {
  display: inline-flex; align-items: center;
  padding: 10px 22px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 14px;
  border: 2px solid var(--border); color: var(--ink-soft);
  transition: all 0.2s; margin-top: 4px;
}
.btn-outline:hover { border-color: var(--gold); color: var(--gold); }

/* Responsive */
@media (max-width: 800px) {
  .prod-card { grid-template-columns: 1fr; }
  .prod-visual { min-height: 200px; }
  .prod-info { padding: 28px 24px; }
  .section-inner { padding: 60px 20px; }
}
</style>