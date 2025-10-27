<script setup lang="ts">
import CreatePost from "@/components/CreatePost.vue";
import MainHeader from "@/components/MainHeader.vue";
import PostCard from "@/components/PostCard.vue";
const posts = [
  {
    id: 1,
    author: "สมชาย ใจดี",
    authorImage: "",
    time: "2 ชั่วโมงที่แล้ว",
    content: "วันนี้อากาศดีมาก ไปเดินเล่นที่สวนสาธารณะมา สดชื่นมากเลย! 🌳☀️",
    image: "https://images.unsplash.com/photo-1441974231531-c6227db76b6e?w=600&h=400&fit=crop",
    likes: 24,
    comments: 5,
  },
  {
    id: 2,
    author: "สมหญิง รักสนุก",
    authorImage: "",
    time: "5 ชั่วโมงที่แล้ว",
    content: "เพิ่งลองร้านอาหารใหม่ อร่อยมากๆ ใครอยากไปลองบอกได้นะคะ แนะนำเลยค่ะ! 🍜✨",
    image: "https://images.unsplash.com/photo-1504674900247-0877df9cc836?w=600&h=400&fit=crop",
    likes: 42,
    comments: 12,
  },
  {
    id: 3,
    author: "วิชัย มีสุข",
    authorImage: "",
    time: "1 วันที่แล้ว",
    content: "สุดยอดไปเลย! คอนเสิร์ตเมื่อคืนนี้ มันส์มากกก 🎵🎸",
    likes: 67,
    comments: 18,
  },
  {
    id: 4,
    author: "ประภา สว่าง",
    authorImage: "",
    time: "2 วันที่แล้ว",
    content: "เริ่มต้นโปรเจคใหม่วันนี้ ตื่นเต้นและท้าทายมากๆ เลยค่ะ ขอกำลังใจหน่อยนะคะ 💪✨",
    image: "https://images.unsplash.com/photo-1486312338219-ce68d2c6f44d?w=600&h=400&fit=crop",
    likes: 38,
    comments: 8,
  },
];
</script>

<template>
  <div>
    <MainHeader />
    <section class="section-container">
      <div class="section-container-left">ss</div>
      <div class="section-container-center">
        <CreatePost />
        <PostCard v-for="post in posts" :key="post.id" v-bind="post" />
      </div>
      <div class="section-container-right">ss</div>
    </section>
  </div>
</template>

<style>
.section-container {
  display: grid;
  grid-template-columns: 1fr 3fr 1fr;
}

.section-container-center {
  display: grid;
  gap: 20px;
  grid-template-columns: auto;
}

@media screen and (max-width: 1024px) {
  .section-container {
    display: grid;
    gap: 20px;
    grid-template-columns: 2fr 1fr;
    padding: 0px 40px;
  }
  .section-container-left {
    display: none;
  }
}

@media screen and (max-width: 1023px) {
  .section-container {
    display: grid;
    gap: 20px;
    grid-template-columns: auto;
    padding: 0px 40px;
  }
  .section-container-right {
    display: none;
  }
}
</style>
