<template>
  <div class="Walletbox">
    <div class="WalletInnerBox" @click.stop="Copy">
      <span class="avatarBox">
        <img class="avatar" :src="$store.state.avatar" alt="" ref="avatar" slot="reference" @click.stop="connectWallet">
      </span>
      <span class="address">{{ 
        $store.state.currentAddress==""?"MetaMask is not connected":`${this.$store.state.currentAddress.slice(
          0,
          5
        )}...${this.$store.state.currentAddress.slice(-5)}`
      }}
      </span>
    </div>
    <div />
  </div>
</template>

<script>
  import { getToken, setMnemonic, checkUserExist } from "@/api/axios/user";

  import CryptoJS from "crypto-js";
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
        contract: null,
        toAddress: "",
        value: "",
        re: "",
        avatar: "",
        //*********************//
        CopyTips: "点击复制",
        CopySuccess: "复制成功！",
        isCopy: false,
        isRepeatClick: true,
        password: "",
      };
    },
    mounted() {
      if (window.ethereum != undefined) {
        window.ethereum.on("accountsChanged", () => {
          this.connectWallet();
        });
      }
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
      MMCTDis() {
        this.$refs.MMCT.classList.remove("animate__fadeIn");
        this.$refs.MMCT.classList.add("animate__fadeOut");
        setTimeout(() => {
          this.MetaMaskTipsIsShow = !this.MetaMaskTipsIsShow;
        }, 750);
      },

      async connectWallet() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            // 请求用户授权
            await window.ethereum
              .request({ method: "eth_requestAccounts" })
              .then(async (handleAccountsChanged) => {
                this.$store.commit("connection", true);
                this.$store.commit("changeAvatar", handleAccountsChanged[0]);
                this.$store.commit("setcurrentAddress", handleAccountsChanged[0]);
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
            // var isExist = false
            //   if (re.data=="") {
            //     alert("账户没有注册")
            //     this.password = prompt("醒密码");
            //     this.password = CryptoJS.SHA256(this.password).toString();
            //     await setMnemonic().then(re => {
            //       console.log(re);
            //     })
            //     await getToken(user).then((re) => {
            //     console.log(re);
            //     localStorage.setItem("token", re.data.data);
            //   })
            //   } else {
            //     this.password = prompt("密码");
            //     this.password = CryptoJS.SHA256(this.password).toString();
            //     console.log(this.password);
            //     await getToken(user).then((re) => {
            //     console.log(re);
            //     localStorage.setItem("token", re.data.data);
            //   });
            // localStorage.getItem("token");
            //   }
            var user = {
              userAddress: this.$store.state.currentAddress,
              encryptedPassword: CryptoJS.SHA256(this.password).toString(),
            };
            await checkUserExist(user).then(async (re) => {
              if (re.data.data == "") {
                alert("账户没有注册");
                this.password = prompt("醒密码");
                console.log(user);
                await setMnemonic(user).then((re) => {
                  console.log(re);
                });
              } else {
                this.password = prompt("密码");
              }
            });

            await getToken(user).then((re) => {
              localStorage.clear();
              console.log(user);
              if (re.data.data == "") {
                alert("密码错误");
              } else {
                localStorage.setItem("token", re.data.data);
                this.walletConnect();
                this.$notify({
                  title: "🎉 连接成功",
                  position: "top-left",
                  offset: 200,
                });
              }
            });
          } catch (error) {
            console.error(error);
            this.$notify.error({
              title: "连接失败",
              position: "top-left",
              offset: 200,
            });
          }
        } else {
          this.$notify({
            title: "已经有连接请勿操作频繁",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
        setTimeout(() => {
          this.isRepeatClick = true;
        }, 5000);
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
}
</style>