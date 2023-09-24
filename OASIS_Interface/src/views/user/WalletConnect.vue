<template>
  <div class="Walletbox">
    <div class="WalletInnerBox" @click.stop="Copy">
      <span class="avatarBox">
        <img class="avatar" :src="$store.state.avatar" alt="" ref="avatar" slot="reference">
      </span>
      <span class="address">{{ 
        $store.state.currentAddress==""?"MetaMask is not connected":`${this.$store.state.currentAddress.slice(
          0,
          5
        )}...${this.$store.state.currentAddress.slice(-5)}`
      }}
      </span>
    </div>
  </div>
</template>

<script>

  export default {
    data() {
      return {
        re: "",
        avatar: "",
        //*********************//
        CopyTips: "点击复制",
        CopySuccess: "复制成功！",
        isCopy: false,
        isRepeatClick: true,
      };
    },
    methods: {
      Copy() {
        navigator.clipboard
          .writeText(this.$store.state.currentAddress)
          .then(() => {
            this.isCopy = true;
            setTimeout(() => {
              this.isCopy = false;
            }, 3000);
          });
      },
      GETHashAvatar() {
        if (this.$store.state.isconnect) {
          this.avatar =
            "data:image/png;base64," +
            new this.Identicon(this.$store.state.currentAddress, 120).toString();
          this.$refs.avatar.width = 60;
          this.$refs.avatar.height = 60;
        } else {
          this.avatar = require("@/assets/webAssets/MetaMask.png");
          this.$refs.avatar.width = 40;
          this.$refs.avatar.height = 40;
        }
      },
    },
  };
</script>

<style lang="scss">
/* 全局通知样式 */
.el-notification {
  border: none !important;
  padding-top: 2%;
  padding-bottom: 2%;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 20px;
  .el-notification__title {
    font-size: 1vw;
    font-weight: 500;
  }
}
::-webkit-scrollbar {
  width: 6px;
  height: 8px;
  display: none;
  background-color: transparent;
}
::-webkit-scrollbar-thumb {
  background-color: #ccc;
  border-radius: 25px;
}
</style>
<style lang="scss" scoped>
.Walletbox {
  display: flex;
  justify-content: flex-start;
  flex-direction: column;
  width: 100%;
  height: 100%;
}
.avatar {
  object-fit: fill;
  width: 100%;
  height: 100%;
  background-color: white;
  float: right;
}
.avatarBox {
  overflow: hidden;
  border-radius: 50%;
  width: 38px;
  height: 38px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-left: 10%;
  border: 2px solid var(--border-green--);
  transition: all 0.3s ease-in-out;
  &:hover {
    border: 2px solid var(--border-green--);
    box-shadow: 0 0 20px var(--border-green--);
    transition: all 0.3s ease-in-out;
  }
}
.MetaMaskAvatar {
  border-radius: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
}
.BOX {
  width: 100%;
  height: 80px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.WalletInnerBox {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: flex-start;
  align-items: center;
  cursor: pointer;
}
.address {
  margin-left: 5%;
  color: var(--Dark--);
  font-weight: 800;
  font-size: 12px;
  width: 50%;
  text-align: center;
  border-radius: 10px;
  padding-top: 10px;
  padding-bottom: 10px;
  padding-left: 5px;
  padding-right: 5px;
  font-family: Arial, Helvetica, sans-serif;
  transition: all 0.3s ease-in-out;
  &:hover {
    background-color: rgba(238, 238, 238, 0.1);
    transition: all 0.3s ease-in-out;
  }
}
</style>