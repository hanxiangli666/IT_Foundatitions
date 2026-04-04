<template>
  <div class="page-top-pad">
    <section class="about-hero">
      <div class="about-hero-bg" aria-hidden="true"></div>
      <div class="ah-inner">
        <div class="section-tag">关于我们</div>
        <h1 class="ah-title">我们是<span class="gold">充满激情</span>的年轻技术团队</h1>
        <p class="ah-sub">
          广州熠辉智能科技有限公司成立于2026年，
          专注为企业和个人提供高质量软件开发服务。我们相信好的技术可以改变商业。
        </p>
        <div class="ah-tags">
          <span>100% 自主研发</span>
          <span>合法合规签约</span>
          <span>长期迭代维护</span>
        </div>
      </div>
    </section>

    <section class="about-values">
      <div class="section-inner">
        <div class="vals-grid">
          <div class="val-card" v-for="v in values" :key="v.id">
            <div class="val-icon">
              <SvgIcon :name="v.icon" :size="32" />
            </div>
            <h3>{{ v.title }}</h3>
            <p>{{ v.desc }}</p>
          </div>
        </div>
      </div>
    </section>

    <section class="about-story">
      <div class="section-inner story-grid">
        <div class="story-text">
          <div class="section-tag">我们的故事</div>
          <h2 class="section-title" style="text-align:left">从校园走向市场</h2>
          <p>
            我们的团队在大学期间就开始了技术探索，最成功的产品是为学校开发的
            <strong>科创赋能平台（智创空间）</strong>，已正式上线并完全由团队从0自主开发，
            获得学校师生的广泛好评。
          </p>
          <p>
            2026年3月，公司在广州市天河区正式注册成立，成为一家完全合法合规的企业，
            开始面向社会提供专业的软件开发服务。
          </p>
          <p>
            我们的团队成员具备扎实的编程基础，擅长网站开发、软件定制、小程序与小游戏开发，
            同时也欢迎志同道合的朋友加入我们一起成长。
          </p>
          <a href="http://www.smart-space.cn" target="_blank" rel="noopener noreferrer" class="btn-outline" style="display:inline-flex;margin-top:16px">
            访问智创空间 →
          </a>
        </div>
        <div class="story-facts">
          <div class="fact-card" v-for="f in facts" :key="f.id">
            <div class="fact-num">{{ f.num }}</div>
            <div class="fact-lbl">{{ f.label }}</div>
          </div>
        </div>
      </div>
    </section>

    <section class="about-certs">
      <div class="section-inner">
        <div class="section-header">
          <div class="section-tag">资质认证</div>
          <h2 class="section-title">正规企业，值得信赖</h2>
          <p class="section-desc">我们是完全合法合规注册的真实企业，可签合同、开发票，可提供实习证明</p>
        </div>
        <div class="cert-list">
          <div class="cert-item" v-for="c in certs" :key="c.id">
            <div class="cert-icon">
              <SvgIcon :name="c.icon" :size="36" />
            </div>
            <h4>{{ c.title }}</h4>
            <p>{{ c.desc }}</p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import SvgIcon from '@/components/SvgIcon.vue'

const values = [
  { id:1, icon:'rocket', title:'快速交付', desc:'敏捷开发，快速响应，最短时间内将您的想法变为现实' },
  { id:2, icon:'gem', title:'品质优先', desc:'代码全部自主研发，注重代码质量与用户体验' },
  { id:3, icon:'handshake', title:'诚信合作', desc:'正规注册企业，签合同、开发票，让合作有保障' },
  { id:4, icon:'sprout', title:'持续成长', desc:'保持学习，紧跟技术前沿，与客户共同成长' },
]
const facts = [
  { id:1, num:'2026', label:'公司成立年份' },
  { id:2, num:'1+', label:'已上线产品' },
  { id:3, num:'广州', label:'团队所在城市' },
  { id:4, num:'100%', label:'代码自主研发' },
]
const certs = [
  { id:1, icon:'file-text', title:'营业执照', desc:'广州市市场监督管理局正式注册，统一社会信用代码齐全' },
  { id:2, icon:'receipt', title:'可开正规发票', desc:'公司具备开具增值税发票的资质，满足企业报账需求' },
  { id:3, icon:'clipboard', title:'合同签署', desc:'所有合作均可签署正式合同，明确权责，保障双方利益' },
  { id:4, icon:'building', title:'实习证明', desc:'可为合作/实习人员提供正式公司盖章的实习证明文件' },
]
</script>

<style scoped>
.about-hero {
  position: relative;
  overflow: hidden;
  background: linear-gradient(140deg, #faf8f3 0%, #fff 65%);
  border-bottom: 1px solid var(--border);
  padding: 80px 32px 80px;
}
.about-hero-bg {
  position: absolute;
  inset: 0;
  pointer-events: none;
  opacity: 0.8;
  background:
    radial-gradient(circle at 8% 14%, rgba(200,151,58,0.16), transparent 34%),
    radial-gradient(circle at 88% 18%, rgba(59,130,246,0.12), transparent 34%),
    repeating-linear-gradient(
      96deg,
      rgba(13,17,23,0.02) 0,
      rgba(13,17,23,0.02) 1px,
      transparent 1px,
      transparent 7px
    );
}
.ah-inner { max-width: 760px; margin: 0 auto; text-align: center; }
.ah-inner { position: relative; z-index: 1; }
.ah-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(30px, 4vw, 48px);
  font-weight: 700;
  margin: 16px 0 20px;
  line-height: 1.3;
}
.gold { color: var(--gold); }
.ah-sub { color: var(--ink-soft); font-size: 16px; line-height: 1.8; }
.ah-tags {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 10px;
}
.ah-tags span {
  padding: 6px 12px;
  border-radius: 999px;
  font-size: 12px;
  font-weight: 600;
  color: #6b5220;
  border: 1px solid rgba(200,151,58,0.35);
  background: rgba(253,246,232,0.8);
}

.about-values {
  background: linear-gradient(180deg, #f7f8fa 0%, #f2f4f7 100%);
}
.vals-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}
.val-card {
  background: #fff; border-radius: 18px;
  padding: 28px; border: 1px solid var(--border);
  transition: all 0.3s;
  box-shadow: 0 10px 20px rgba(13,17,23,0.04);
}
.val-card:hover {
  transform: translateY(-6px);
  box-shadow: 0 16px 36px rgba(13,17,23,0.1);
  border-color: rgba(200,151,58,0.28);
}
.val-icon {
  color: var(--gold);
  margin-bottom: 14px;
  display: flex;
}
.val-card h3 { font-size: 17px; font-weight: 700; margin-bottom: 8px; }
.val-card p { color: var(--ink-soft); font-size: 14px; line-height: 1.7; }

.about-story { background: #fff; }
.story-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: start;
}
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
.story-text { display: flex; flex-direction: column; gap: 14px; }
.story-text p { color: var(--ink-soft); font-size: 15px; line-height: 1.8; }
.story-text strong { color: var(--gold); }
.story-facts {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
.fact-card {
  background: linear-gradient(180deg, var(--gold-bg), #fff);
  border: 1px solid rgba(200,151,58,0.2);
  border-radius: 16px;
  padding: 24px;
  text-align: center;
  box-shadow: 0 8px 20px rgba(200,151,58,0.08);
}
.fact-num {
  font-size: 32px; font-weight: 800;
  color: var(--gold);
  font-family: 'Noto Serif SC', serif;
  margin-bottom: 6px;
}
.fact-lbl { font-size: 13px; color: var(--ink-soft); }

.about-certs {
  background: linear-gradient(180deg, #f7f8fa 0%, #f3f5f8 100%);
}
.cert-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 20px;
}
.cert-item {
  background: #fff;
  border-radius: 18px;
  padding: 28px;
  border: 1px solid var(--border);
  text-align: center;
  transition: all 0.3s;
  box-shadow: 0 10px 20px rgba(13,17,23,0.04);
}
.cert-item:hover {
  border-color: var(--gold);
  box-shadow: 0 14px 30px rgba(200,151,58,0.14);
  transform: translateY(-4px);
}
.cert-icon {
  color: var(--gold);
  margin-bottom: 12px;
  display: flex;
  justify-content: center;
}
.cert-item h4 { font-size: 16px; font-weight: 700; margin-bottom: 8px; }
.cert-item p { color: var(--ink-soft); font-size: 13px; line-height: 1.7; }

.btn-outline {
  display: inline-flex; align-items: center;
  padding: 11px 22px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 14px;
  border: 2px solid #d7dae0;
  color: var(--ink-soft);
  transition: transform 0.2s ease, border-color 0.2s ease, color 0.2s ease;
  background: rgba(255,255,255,0.75);
}
.btn-outline:hover { border-color: var(--gold); color: var(--gold); transform: translateY(-2px); }

@media (max-width: 768px) {
  .story-grid { grid-template-columns: 1fr; gap: 32px; }
  .story-facts { grid-template-columns: 1fr 1fr; }
  .section-inner { padding: 60px 20px; }
  .ah-tags { justify-content: flex-start; }
  .ah-inner { text-align: left; }
}
</style>
