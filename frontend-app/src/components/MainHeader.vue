<script setup lang="ts">
import { ref } from "vue";
import { Search, Home, MessageCircle, Bell } from "lucide-vue-next";
const count = ref(0);
console.log(count);
</script>

<template>
  <nav class="navber">
    <div class="container">
      <div class="box-con">
        <div class="box-left">
          <h1>Facebook</h1>
          <div class="box-left-search">
            <Search class="search" /><input type="text" placeholder="ค้นหา" />
          </div>
        </div>
        <div class="box-right">
          <button class="box-right-button"><Home class="box-right-icon" /></button>
          <button class="box-right-button"><MessageCircle class="box-right-icon" /></button>
          <button class="box-right-button"><Bell class="box-right-icon" /></button>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.navber {
  /* position: sticky; */
  top: 0;
  /* z-index: 50; */
  background-color: #fff;
  border-bottom: 1px solid #ddd;
  box-shadow: 0 2px 4px 0 rgba(255, 255, 255, 0.1);
}
.container {
  max-width: 1280px;
  margin: 0 auto;
  padding: 0 2rem;
}
.box-con {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 64px;
}
.box-left {
  display: flex;
  align-items: center;
  gap: 3rem;
  width: 100%;
}
.box-left h1 {
  font-size: 32px;
  font-weight: 600;
  background-color: blue;
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}
.box-left-search {
  display: flex; /* จอใหญ่โชว์ */
  position: relative;
  width: 64px;
  height: 35px;
  background-color: #ddd;
  border-style: none;
  border-radius: 8px;
  width: 256px;
}
@media screen and (max-width: 768px) {
  .box-left-search {
    display: none;
  }
}
.search {
  position: absolute;
  left: 5px;
  top: 50%;
  transform: translateY(-50%);
  height: 16px;
  width: 16px;
}
.box-left input {
  padding-left: 30px;
  border-width: 1px;
  border-radius: 8px;
  background-color: #f2f4f7;
  border-style: none;
  outline-width: 0px;
  width: 100%;
}
.box-left input:hover {
  border-style: none;
  outline-width: 0px;
}
.box-right {
  display: flex;
  align-items: center;
  gap: 1rem;
}
.box-right-button {
  background-color: initial;
  color: initial;
  position: relative;
  transition: background-color 0.3s ease;
  border-width: 0px;
  border-radius: 8px;
}
.box-right-button:hover {
  color: var(--accent-foreground);
  border-radius: 8px;
  background-color: #f3f4f6;
}
.box-right-icon {
  height: 20px;
  width: 20px;
  margin: 6px 6px;
}
</style>
