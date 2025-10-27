<template>
  <div class="container">
    <div class="card">
      <div class="card-header">
        <div class="card-header-img">คุณ</div>
        <button class="card-header-input">คุณคิดอะไรอยู่ #ชื่อ</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  margin-top: 35px;
}
.card {
  width: 100%;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
}
.card-header {
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 8px;
  align-items: center;
  padding: 16px 16px;
  background-color: white;
  color: white;
  border-radius: 16px;
}
.card-header-img {
  width: 40px;
  height: 40px;
  border-width: 1px;
  border-radius: 100%;
  color: black;
  background-color: #f7f9fa;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}
.card-header-input {
  border-width: 0px;
  outline-width: 0px;
  background-color: #f7f9fa;
  width: auto;
  height: 40px;
  border-radius: 18px;
  display: flex;
  justify-content: start;
  align-items: center;
  cursor: pointer;
}
</style>
