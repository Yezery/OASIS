<template>
  <div class="userInfBox animate__animated animate__fadeInRight">
    <div class="userInfBox-main">
      <div class="userInfBox-head">
        <div class="userInfBox-head-avatarBox">
          <div class="avatarBox">
            <img
              class="avatar"
              :src="$store.state.avatar"
              width="90px"
              height="90px"
            >
          </div>
        </div>
        <div class="userInfBox-head-nameBox">
          <div class="nameBox">
            <span class="name">{{ `${this.$store.state.currentAddress.slice(
              0,
              5
            )}...${this.$store.state.currentAddress.slice(-5)}` }}</span>
            <span class="title">Account</span>
          </div>
        </div>
        <div class="userInfBox-head-NFTInfBox">
          <div class="userInfBox-head-NFTInfBox-innerBox">
            <div class="possessBox">
              <span class="possessBox-count">{{ accountNFTList.length }}</span>
              <span class="possessBox-title">NFT Count</span>
            </div>
            <div class="sellingBox">
              <span class="sellingBox-count">{{ balance }} <span style="font-size: 0.1vw;font-weight: 500;">ETH</span></span>
              <span class="sellingBox-title">Balance</span>
            </div>
          </div>
        </div>
        <div class="toSellLink-Box">
          <router-link
            class="toSellLink"
            :to="{ name: 'mintNFT' }"
          >
            <span>创造NFT</span>
          </router-link>
        </div>
        <div class="toSellLink-Box">
          <router-link
            class="toSellLink"
            :to="{ name: 'userhome' }"
          >
            <span>个人主页</span>
          </router-link>
        </div>
        <div class="toSellLink-Box">
          <el-button
            class="toSellLink"
            @click="backPassword(4)"
          >
            <span>找回密码</span>
          </el-button>
        </div>
      </div>
      <div class="userInfBox-foot">
        <div class="accountNFTBox">
          <!-- <div class="accountNFTList" v-for="nft in accountNFTList" :key="nft.ipfsPath">
            <img class="NFTImage" :src="nft.ipfsPath" alt="" />
          </div> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    backPassword: {
      type: Function,
      default:undefined
    }
  },
    data() {
      return {
        accountNFTList: [],
        balance: 0,
      };
    },
  mounted() {
      this.getBalance()
    },
  methods: {
    async getBalance() {
      this.balance = this.$store.state.Web3.utils.fromWei(await this.$store.state.Web3.eth.getBalance(this.$store.state.currentAddress), "ether").slice(0, 4) 
      if (this.$store.state.ownerNFTList !== null) {
      this.accountNFTList = this.$store.state.ownerNFTList;
      }

    },
    },
  };
</script>

<style lang="scss" scoped>
.userInfBox {
  width: 90%;
  height: 100%;
  background-color: var(--White--);
  border-radius: 40px;
  box-shadow: var(--boxshdow-style--);
  transition: all 0.3s ease-in-out;
  &:hover {
    box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
      rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
    transition: all 0.3s ease-in-out;
  }
  .userInfBox-main {
    .userInfBox-head {
      height: 100%;
      margin-top: 10%;

      .userInfBox-head-avatarBox {
        display: flex;
        justify-content: center;
        align-items: center;

        .avatarBox {
          width: 89px;
          height: 89px;
          overflow: hidden;
          border-radius: 10px;
          .avatar {
            object-fit: fill;
          }
        }
      }
      .userInfBox-head-nameBox {
        margin-top: 8%;
        .nameBox {
          display: flex;
          flex-direction: column;
          .name {
            color: var(--blueblack-white--);
            font-weight: 700;
            font-size: 18px;
            text-align: center;
            font-family: Gilroy-Medium;
          }
          .title {
            margin-top: 2%;
            color: rgb(112, 122, 137);
            font-size: 12px;
            font-weight: 700;
          }
        }
      }
      .userInfBox-head-NFTInfBox {
        margin-top: 8%;
        display: flex;
        justify-content: center;
        align-items: center;
        .userInfBox-head-NFTInfBox-innerBox {
          display: flex;
          width: 80%;
          justify-content: space-around;
          .possessBox {
            display: flex;
            flex-direction: column;
            .possessBox-count {
              color: var(--blueblack-white--);
              font-family: Gilroy-Medium;
              font-size: 16px;
              font-weight: 700;
            }
            .possessBox-title {
              color: rgb(112, 122, 137);
              font-size: 12px;
              font-weight: 700;
            }
          }
          .sellingBox {
            display: flex;
            flex-direction: column;
            .sellingBox-count {
              color: var(--blueblack-white--);
              font-family: Gilroy-Medium;
              font-size: 16px;
              font-weight: 700;
            }
            .sellingBox-title {
              color: rgb(112, 122, 137);
              font-size: 12px;
              font-weight: 700;
            }
          }
        }
      }
    }
    .userInfBox-foot {
      width: 100%;
      height: 100%;
      margin-top: 8%;
      margin-bottom: 10%;
      .accountNFTBox {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100%;
        .accountNFTList {
          .NFTImage {
            object-fit: cover;
            width: 90px;
            height: 90px;
          }
        }
      }
    }
    .toSellLink-Box {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 70px;
      margin-top: 20px;
      width: 100%;
      .toSellLink {
        background-color: rgba(85, 201, 96, 0.12);
        color: #55c960;
        text-decoration: none;
        border-radius: 20px;
        padding: 14px 0px 14px 0px;
        border: solid 2px rgba(85, 201, 96, 0.24);
        width: 60%;
        &:hover {
          background-color: rgba(85, 201, 96, 0.24);
          transition: all 0.3s ease-in-out;
        }
      }

    }
  }
}
</style>