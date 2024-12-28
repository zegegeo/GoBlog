<template>
  <div class="layout">
    <Navbar />
    <div class="main-content">
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
import Navbar from '../components/Navbar.vue'
</script>

<style scoped>
.layout {
  min-height: 100vh;
  background: #1a1a1a;
}

.main-content {
  padding-top: 60px;  /* 为导航栏留出空间 */
}
</style> 