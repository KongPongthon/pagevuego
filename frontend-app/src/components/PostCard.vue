<script setup lang="ts">
import { Heart, MessageCircle } from "lucide-vue-next";

interface PostCardProps {
  author: string;
  authorImage: string;
  time: string;
  content: string;
  image?: string;
  likes: number;
  comments: number;
}
const props = defineProps<PostCardProps>();
</script>

<template>
  <div class="container">
    <div>
      <div class="container-header">
        <div v-if="props.authorImage">
          <img :src="props.authorImage" alt="" />
        </div>
        <div v-else class="card-header-img">คุณ</div>
        <div>
          <div>{{ props.author }}</div>
          <div>{{ props.time }}</div>
        </div>
      </div>
      <div class="container-header">
        <div>{{ props.content }}</div>
      </div>
      <div class="card-body">
        <img v-if="props.image" :src="props.image" class="card-body" />
      </div>
      <div class="container-header">
        <div class="card-icon"><Heart class="icon" />{{ props.likes }}</div>
        <div class="card-icon"><MessageCircle class="icon" />{{ props.comments }}</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
  background-color: white;
}
.container-header {
  display: flex;
  gap: 8px;
  align-items: center;
  padding: 8px 8px;
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
}
.card-body {
  width: 100%;
}
.card-icon {
  display: flex;
  align-items: center;
  gap: 6px;
}
.icon {
  width: 16px;
  height: 16px;
}
</style>
