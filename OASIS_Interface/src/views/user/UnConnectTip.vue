<!-- eslint-disable vue/valid-template-root -->
<template>
  <div>1</div>
  <!-- 钱包连接提醒 -->
  <!-- <div class="Mask animate__animated animate__fadeIn" v-show="MetaMaskTipsIsShow" ref="MMCT">
    <div class="MetaMaskConnectionTip">
      <div class="MetaMaskbox">
        <div class="BOX">
          <div class="MetaMaskAvatar">
            <img :src="avatar" alt="" width="80px" height="80px" style="object-fit: contain" />
          </div>
        </div>
        <p class="MetaMaskTips" style="
              margin-bottom: 10%;
              font-size: 15px;
              font-weight: 800;
              line-height: 230%;
            ">
          {{ isconnect ? balance + " ETH" : Tips2 }}
        </p>
        <button class="MetaMaskTipsBtn" @click="MMCTDis" :disabled="isDisable">
          完成
        </button>
      </div>
    </div>
  </div> -->
  <!---------------->
</template>

<script>
  export default {
    data() {
      return {
        //****钱包连接弹出框******//
        MetaMaskTipsIsShow: false,
        balance: "",
        Tips2: "MetaMask无法连接 ",
        isDisable: false,
        //*********************//
      };
    },
  };
</script>

<style>
</style>