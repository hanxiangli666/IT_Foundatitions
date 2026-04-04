<template>
  <div class="not-found-page">
    <div class="nf-inner">
      <div class="nf-code">404</div>
      <h1 class="nf-title">页面不存在</h1>
      <p class="nf-desc">您访问的页面已移动或不存在，请返回首页重新浏览</p>
      <div class="nf-actions">
        <router-link to="/" class="btn-primary">返回首页</router-link>
        <router-link to="/contact" class="btn-outline">联系我们</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.not-found-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 80px 32px;
  background: linear-gradient(135deg, #f8fafc 0%, var(--gold-bg) 100%);
}
.nf-inner {
  text-align: center;
  max-width: 480px;
}
.nf-code {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(80px, 15vw, 140px);
  font-weight: 800;
  color: var(--gold);
  opacity: 0.25;
  line-height: 1;
  margin-bottom: 8px;
}
.nf-title {
  font-family: 'Noto Serif SC', serif;
  font-size: clamp(24px, 4vw, 36px);
  font-weight: 700;
  color: var(--ink);
  margin-bottom: 16px;
}
.nf-desc {
  color: var(--ink-soft);
  font-size: 16px;
  line-height: 1.7;
  margin-bottom: 36px;
}
.nf-actions {
  display: flex;
  gap: 14px;
  justify-content: center;
  flex-wrap: wrap;
}
.btn-primary {
  display: inline-flex; align-items: center;
  background: var(--gold); color: #fff;
  padding: 12px 28px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 15px;
  transition: all 0.2s;
  border: 2px solid var(--gold);
}
.btn-primary:hover { background: var(--gold-lt); border-color: var(--gold-lt); transform: translateY(-2px); }
.btn-outline {
  display: inline-flex; align-items: center;
  padding: 12px 28px; border-radius: 10px;
  text-decoration: none; font-weight: 700; font-size: 15px;
  border: 2px solid var(--border);
  color: var(--ink-soft);
  transition: all 0.2s;
}
.btn-outline:hover { border-color: var(--gold); color: var(--gold); transform: translateY(-2px); }
</style>
