<template>
  <div class="hr">
    <div class="line" ref="firstLine"></div>
    <span>{{ props.content }}</span>
    <div class="line" ref="secondLine"></div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref, watch } from 'vue'

const props = defineProps(['content', 'percent'])
const firstLine = ref<HTMLDivElement>()
const secondLine = ref<HTMLDivElement>()

onMounted(() => {
  firstLine.value!.style.width = `${props.percent}`
  secondLine.value!.style.width = `calc(100% - ${props.percent} - ${props.content.length * 20}px)`
})
watch(() => props.percent, (newVal) => {
  firstLine.value!.style.width = `${newVal}`
  secondLine.value!.style.width = `calc(100% - ${newVal} - ${props.content.length * 20}px)`
})

</script>

<style scoped>
.hr {
  display: flex;
  align-items: center;
  margin-bottom: 5px;
  color: var(--base-text-color);

  .line {
    background-color: gray;
    width: 45%;
    height: 2px;
  }
}
</style>