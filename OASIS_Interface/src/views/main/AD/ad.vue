<template>
  <el-carousel
    height="300px"
    width="100%"
    trigger="click"
    :interval="5000"
    indicator-position="outside"
  >
    <el-carousel-item
      v-for="item in 4"
      :key="item"
    >
      <slot :name="'ad' + item" />
    </el-carousel-item>
  </el-carousel>
</template>

<script>
export default {
    name:"AD"
  };
</script>

<style lang="scss" scoped>
  
</style>