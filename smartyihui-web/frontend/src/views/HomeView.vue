<template>
  <section class="hero-card">
    <div>
      <p class="eyebrow">Vue 3 + Fiber + MySQL</p>
      <h2>企业门户网站基础框架</h2>
      <p class="desc">
        当前项目用于搭建公司企业门户网站，前端负责页面展示与交互，后端提供数据接口与服务能力，后续可以继续扩展企业介绍、新闻动态、公告通知、联系我们等模块。
      </p>
      <div class="actions">
        <router-link to="/messages">
          <n-button type="primary" size="large">进入内容管理示例</n-button>
        </router-link>
      </div>
    </div>
    <div class="feature-grid">
      <article>
        <h3>前端</h3>
        <p>Vue 3、Vue Router、Naive UI、Axios</p>
      </article>
      <article>
        <h3>后端</h3>
        <p>Go Fiber、REST API、CORS</p>
      </article>
      <article>
        <h3>数据层</h3>
        <p>MySQL + 自动建表</p>
      </article>
    </div>
  </section>
</template>

<style scoped>
.hero-card {
  display: grid;
  gap: 24px;
  padding: 32px;
  border-radius: 24px;
  background: rgba(255, 255, 255, 0.95);
  box-shadow: 0 18px 50px rgba(37, 99, 235, 0.08);
}

.eyebrow {
  margin: 0 0 8px;
  color: #2563eb;
  font-weight: 700;
  letter-spacing: 0.08em;
  text-transform: uppercase;
}

h2 {
  margin: 0;
  font-size: clamp(28px, 5vw, 42px);
}

.desc {
  max-width: 720px;
  color: #475569;
}

.actions {
  margin-top: 24px;
}

.feature-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 16px;
}

.feature-grid article {
  padding: 18px;
  border-radius: 18px;
  background: #eff6ff;
}

.feature-grid h3 {
  margin-top: 0;
}

.feature-grid p {
  margin-bottom: 0;
  color: #475569;
}
</style>
