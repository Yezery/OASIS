<template>
  <div class="Walletbox">
    <el-popover
      placement="top-end"
      trigger="click"
      :offset="200"
      :visible-arrow="false"
      :close-delay="0"
      popper-class="userSelectDropdown animate__animated animate__flipInX"
    >
      <div class="userSelect_connected" v-if="$store.state.isconnect">
        <div class="userSelect_title" >
          <div>
            操作
          </div>
        </div>
       
          <router-link to="/userhome" class="userSelect_item">
            个人主页
        </router-link>
   
        <div class="userSelect_item">
          找回授权码
        </div>
        <div class="userSelect_item" @click="disConnect">
          断开连接
        </div>
      </div>
      <div class="userSelect_unconnect" v-else>
        <div class="userSelect_title" >
          <div>
            操作
          </div>
        </div>
        <div class="userSelect_item" @click="connectWallet" v-if="isUnlocked">
            授权
        </div>
        <div class="userSelect_item"  v-else>
            请解锁钱包<i class="el-icon-warning" style="color: red;"></i>
        </div>
      </div>
      <div
        class="WalletInnerBox"
        slot="reference"
      >
        <span class="avatarBox">
          <!-- @click="openEmpower" -->
          <img
            class="avatar"
            :src="$store.state.avatar"
            alt=""
            ref="avatar"
            slot="reference"
          >
        </span>
        <!-- <div class="userSelect">
      <div class="userSelect_item">切换</div>
      <div class="userSelect_item">找回</div>
      <div class="userSelect_item">退出</div> -->
      </div>
    </el-popover>
    <!-- <span class="address">{{ 
        $store.state.currentAddress==""?"MetaMask is not connected":`${this.$store.state.currentAddress.slice(
          0,
          5
        )}...${this.$store.state.currentAddress.slice(-5)}`
      }}
      </span> -->
  </div>
</template>

<script>
  export default {
    data() {
      return {
        isUnlocked:false
      };
    },
  async mounted() {
    if (window.ethereum != undefined) {
      await window.ethereum._metamask
        .isUnlocked()
        .then((re) => (this.isUnlocked = re));
    }
    },
  methods: {
    connectWallet() {
      this.$store.commit("setGetToken", true);
    },
    disConnect() {
      window.ethereum.on('disconnect', () => {
        this.$notify({
            title: "断开成功",
            type: "success",
            position: "top-left",
            offset: 200,
          });
      })
      this.$store.commit("setGetToken", false);
      this.$store.commit("setConnection", false);
      this.$store.commit("setEmpower", false);
      this.$store.commit("setOwnerNFTList",[])
     
    }
    },
  };
</script>
<style lang="scss" scoped>
@import "@/style/topMenu/walletConnect.scss";

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

<style lang="scss">
.userSelectDropdown {
  transition: $globalTransition;
  padding: 0;
  margin: 0;
  width: 200px;
  border-radius: 15px;
  background-color: var(--White--);
  border: 1px solid var(--FG--);
  .userSelect_connected {
    color: var(--Dark--);
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    border-radius: 10px;
    overflow: hidden;
    border: .5px solid var(--FG--);
    .userSelect_title{
      padding-left: 20%;
      width: 100%;
      padding-top: 5%;
      padding-bottom: 3%;
      font-size: 12px;
      color: rgb(156,163,175);
    }
    .userSelect_title{
      padding-left: 20%;
      width: 100%;
      padding-top: 5%;
      padding-bottom: 3%;
      font-size: 12px;
      color: rgb(156,163,175);
    }
    .userSelect_item {
     color: var(--Gray--);
      display: flex;
      justify-content: flex-start;
      align-items: center;
      font-size: 13px;
      padding-left: 50%;
      font-family: "PF";
      width: 100%;
      font-weight: 300;
      height: 40px;
      cursor: pointer;
      &:hover {
        background-color: #cccccc20;
      }
    }
  }
  .userSelect_unconnect{
    @extend .userSelect_connected;
  }
  .popper__arrow {
    display: none;
  }
}
a{
        text-decoration: none;
        color: var(--Dark--);
      }
</style>