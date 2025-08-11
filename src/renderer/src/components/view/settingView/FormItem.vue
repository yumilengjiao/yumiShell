<template>
  <div class="form-item">
    <div class="form-item-label">
      {{ label }}
    </div>
    <div class="form-item-content">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts" setup>
defineProps(['label'])
</script>

<style scoped>
.form-item {
  display: flex;
  gap: 20px;
  align-items: center;
  max-width: 40%;
  justify-content: space-between;
  border: 1px dashed var(--base-border-color);
  border-radius: 10px;
  padding: 5px;
  margin-bottom: 10px;
  color: var(--base-text-color);
  line-height: 30px;
}
</style>