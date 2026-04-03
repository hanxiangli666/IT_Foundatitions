<template>
  <n-config-provider>
    <n-message-provider>
      <div class="layout-shell">
        <header class="topbar">
          <div>
            <h1>smartyihui-web</h1>
            <p>企业门户网站前后端一体化工程</p>
          </div>
          <nav>
            <router-link to="/">首页</router-link>
            <router-link to="/messages">内容管理</router-link>
          </nav>
        </header>
        <main class="page-main">
          <router-view />
        </main>
      </div>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
.layout-shell {
  min-height: 100vh;
  padding: 24px;
  background: linear-gradient(180deg, #f8fbff 0%, #eef4ff 100%);
}

.topbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
  margin: 0 auto 24px;
  max-width: 1120px;
  padding: 20px 24px;
  border-radius: 20px;
  background: rgba(255, 255, 255, 0.85);
  backdrop-filter: blur(10px);
  box-shadow: 0 10px 30px rgba(15, 23, 42, 0.08);
}

.topbar h1 {
  margin: 0;
  font-size: 28px;
}

.topbar p {
  margin: 6px 0 0;
  color: #64748b;
}

nav {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}

nav a {
  padding: 10px 14px;
  border-radius: 999px;
  color: #334155;
  background: #f1f5f9;
}

nav a.router-link-active {
  color: #fff;
  background: #2563eb;
}

.page-main {
  max-width: 1120px;
  margin: 0 auto;
}

@media (max-width: 720px) {
  .topbar {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
