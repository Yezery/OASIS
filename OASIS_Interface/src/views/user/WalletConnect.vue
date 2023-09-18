<template>
  <div class="Walletbox">
    <div class="WalletInnerBox">
      <span class="avatarBox">
        <img
          class="avatar"
          :src="$store.state.avatar"
          alt=""
          ref="avatar"
          slot="reference"
          @click.stop="connectWallet"
        >
      </span>
      <span
        class="address"
        @click.stop="Copy"
      >{{ account }}
        <el-tooltip
          class="item"
          effect="dark"
          :content="isCopy ? CopySuccess : CopyTips"
          placement="bottom"
        ><span><i
          v-if="$store.state.isconnect"
          style="margin-left: 5px; cursor: pointer"
          class="el-icon-document-copy"
        /></span></el-tooltip></span>
    </div>
  </div>
</template>

<script>
  import { postOwnerContractList } from "@/api/axios/ownerContractLIst";
  export default {
    props: {
      walletConnect: {
        type: Function,
        default: null,
      },
    },
    data() {
      return {
        //****钱包连接弹出框******//
        MetaMaskTipsIsShow: false,
        Tips2: "MetaMask无法连接 ",
        isDisable: false,
        //*********************//
        //********Web3********//
        web3: null,
        account: "MetaMask is not connected",
        contract: null,
        toAddress: "",
        value: "",
        re: "",
        avatar: "",
        //*********************//
        CopyTips: "点击复制",
        CopySuccess: "复制成功！",
        isCopy: false,
      };
    },
    mounted() {
      if (window.ethereum != undefined) {
        window.ethereum.on("accountsChanged", this.handleAccountsChanged);
      }
    },
    methods: {
      async handleAccountsChanged(accounts) {
        this.accounts = accounts;
        await this.connectWallet();
      },
      Copy() {
        navigator.clipboard.writeText(this.accounts).then(() => {
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
            new this.Identicon(this.accounts[0], 120).toString();
          this.$refs.avatar.width = 60;
          this.$refs.avatar.height = 60;
        } else {
          this.avatar = require("@/assets/webAssets/MetaMask.png");
          this.$refs.avatar.width = 40;
          this.$refs.avatar.height = 40;
        }
      },
      MMCTDis() {
        this.$refs.MMCT.classList.remove("animate__fadeIn");
        this.$refs.MMCT.classList.add("animate__fadeOut");
        setTimeout(() => {
          this.MetaMaskTipsIsShow = !this.MetaMaskTipsIsShow;
        }, 750);
      },
      async connectWallet() {
        if (!this.$store.state.isconnect) {
          try {
            // 请求用户授权
            await window.ethereum
              .request({ method: "eth_requestAccounts" })
              .then((handleAccountsChanged) => {
                this.accounts = handleAccountsChanged;
                this.$store.commit("connection", true);
                this.$store.commit("changeAvatar", handleAccountsChanged[0]);
              })
              .catch((error) => {
                this.$store.commit("connection", false);
                if (error.code === 4001) {
                  // EIP-1193 userRejectedRequest error
                  console.log("Please connect to MetaMask.");
                } else {
                  console.error(error);
                }
              });

            this.$store.commit("setcurrentAddress", this.accounts[0]);
            this.account = `${this.$store.state.currentAddress.slice(
              0,
              5
            )}...${this.$store.state.currentAddress.slice(-5)}`;
            let currentAddress = {
              ownerAddress: this.$store.state.currentAddress,
            };
            // ====================
            postOwnerContractList(currentAddress).then((re) => {
              this.$store.commit("setOwnerNFTList", re.data.data);
              this.accountNFTList = this.$store.state.ownerNFTList;
            });
            this.walletConnect();
            this.$notify({
              title: "成功",
              message: "连接成功",
              position: 'top-left',
              offset: 100,
            duration:2000
            });
          } catch (error) {
            console.error(error);
            this.$notify.error({
              title: "错误",
              message: "连接失败",
              position: 'top-left',
              offset: 100,
            duration:2000
            });
          }
        }

        // }
        // else {
        //   this.$refs.MMCT.classList.remove("animate__fadeOut");
        //   this.$refs.MMCT.classList.add("animate__fadeIn");
        //   this.MetaMaskTipsIsShow = !this.MetaMaskTipsIsShow;
        //   setTimeout(() => {
        //     this.MMCTDis();
        //   }, 1800);
        // }
      },
      // async initConnect() {
      //   if (window.ethereum) {
      //     this.web3 = new this.Web3(window.ethereum);
      //     try {
      //       await window.ethereum.enable();
      //       this.account = (await this.web3.eth.getAccounts())[0];
      //     } catch (error) {
      //       console.error(error);
      //     }
      //   } else if (window.web3) {
      //     this.web3 = new this.Web3(window.web3.currentProvider);
      //     this.account = (await this.web3.eth.getAccounts())[0];
      //   } else {
      //     this.$refs.MMCT.classList.remove("animate__fadeOut");
      //     this.$refs.MMCT.classList.add("animate__fadeIn");
      //     this.MetaMaskTipsIsShow = !this.MetaMaskTipsIsShow;
      //   }
      // },

      // async transfer() {
      //   let value = this.web3.utils.toWei(this.value, "ether");
      //   this.contract.methods
      //     .transfer(this.toAddress)
      //     .send({ from: this.account, value: value.toString() })
      //     .then((result) => {
      //       console.log("Transaction hash:", result.transactionHash);
      //     })
      //     .catch((error) => {
      //       console.error(error);
      //     });

      // console.log(this.contract);

      // await this.contract.methods.changeText("log").call().then(result => this.accounts = result)
      // console.log(this.accounts );
      // },
    },
  };
</script>

<style lang="scss" scoped>
.Walletbox {
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
  /* border: 2px solid var(--avatar-border-style--); */
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
/* MetaMask弹出框 */
.MetaMaskConnectionTip {
  font-family: Arial, Helvetica, sans-serif;
  overflow: hidden;
  position: absolute;
  width: 350px;
  height: 300px;
  background-color: var(--White--);
  color: var(--Dark--);
  top: 50%;
  left: 50%;
  z-index: 100;
  transform: translate(-50%, -50%);
  border-radius: 20px;
  box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
    rgba(17, 17, 26, 0.1) 0px 0px 8px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.MetaMaskTipsBtn {
  border: 0px solid;
  background-color: white;
  box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
    rgba(17, 17, 26, 0.1) 0px 0px 8px;
  border-radius: 7px;
  padding-top: 5%;
  padding-bottom: 5%;
  padding-left: 22%;
  padding-right: 22%;
  font-size: 13px;
  font-weight: 500;
}
.MetaMaskTipsBtn:hover {
  background-color: rgb(252, 251, 251);
}
.MetaMaskTips {
  font-size: 17px;
  font-weight: 500;
  margin-top: 25%;
}
/* 遮罩层 */
.Mask {
  z-index: 100;
  position: fixed !important;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  // background: rgba(0, 0, 0, 0.05);
}
</style>
