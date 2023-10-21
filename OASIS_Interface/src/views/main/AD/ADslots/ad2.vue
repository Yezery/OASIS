<template>
  <div class="ADMainBox">
    <div class="innerBOX">
      <div class="createTips">
        <span style="font-size: 90px;color: red;">{{ feePercentage }}%</span> COMMISSION FOR MARKET TRANSACTIONS
      </div>
    </div>
  </div>
</template>
  
  <script>
   import { getFeePercentage } from "@/api/axios/contract";
  export default {
    name:"AD2",
    data() {
      return {
        feePercentage: 0,
      };
    },
    mounted() {
      getFeePercentage().then(
        (FeePercentage) => (this.feePercentage = FeePercentage)
      );
    },
    methods: {},
  };
</script>
  
<style lang="scss" scoped>
.ADMainBox {
  font-family: Gilroy-Medium;
  width: 100%;
  height: 100%;
  background-color: black;
  position: relative;
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.6s;
  &:hover {
    transform: scale(1.1);

    transition: all 0.6s;
  }
  .innerBOX {
  width: 80%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  .createTips {
        width: 100%;
        font-size: 55px;
        color:white;
        transition: all 0.3s ease-in-out;
        z-index: 2;
      }
}
}

</style>