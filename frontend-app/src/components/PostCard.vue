<script setup lang="ts">
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
      <h1>{{ props.author }}</h1>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
}
</style>
