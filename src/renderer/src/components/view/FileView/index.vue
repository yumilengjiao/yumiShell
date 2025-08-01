<template>
  <div class="file-view">
    <div class="root">

    </div>
    <div class="first">

    </div>
    <div class="second">

    </div>
    <div class="third">

    </div>
  </div>
</template>

<script lang="ts" setup>

</script>

<style scoped lang="scss">
@use '@renderer/styles/variables.scss' as variables;

.file-view {
  width: 100%;
  height: calc(100vh - variables.$menu-bar-height);
  background-color: var(--base-background-color);
  display: flex;
  flex-direction: column;

  // 给前三个展示添加默认flex属性和过渡效果
  .root,
  .first,
  .second {
    flex: variables.$file-view-common-flex; // 均匀分配空间
    transition: all 0.3s ease; // 添加所有属性的过渡
    box-sizing: border-box; // 确保边框不影响尺寸
  }

  //最后文件展示区的flex值
  .third {
    flex: variables.$file-view-main-flex;
    transition: all 0.3s ease; // 添加所有属性的过渡
    box-sizing: border-box; // 确保边框不影响尺寸
  }

  .root {
    background-color: var(--base-background-color);
    padding: 2px;
    border: 2px dashed var(--base-border-color);
    transition: border-width 0.1s ease;

    &:hover {
      border: 5px dashed var(--base-border-hover-color);
    }
  }

  .first {
    background-color: var(--base-background-color);
    padding: 2px;
    border: 2px dashed var(--base-border-color);
    transition: border-width 0.1s ease;

    &:hover {
      border: 5px dashed var(--base-border-hover-color);
    }
  }

  .second {
    background-color: var(--base-background-color);
    padding: 2px;
    border: 2px dashed var(--base-border-color);
    transition: border-width 0.1s ease;

    &:hover {
      border: 5px dashed var(--base-border-hover-color);
    }
  }

  .third {
    background-color: var(--base-background-color);
    padding: 2px;
    border: 2px dashed var(--base-border-color);
    transition: border-width 0.1s ease;

    &:hover {
      border: 5px dashed var(--base-border-hover-color);
    }
  }
}

// ... existing code ...</style>