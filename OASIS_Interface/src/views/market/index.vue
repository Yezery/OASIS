<template>
  <div class="Main">
    <el-container>
      <el-aside
        style="width: auto;"
        class="HomeMenu animate__animated animate__fadeInLeft"
      >
        <ChatMemu ref="ChatMemu" />
      </el-aside>

      <el-main>
        <router-view />
      </el-main>

      <el-aside
        class="animate__animated animate__fadeInRight"
        style="width: auto;animation-delay: 1s;"
      >
        <el-row
          type="flex"
          class="row-bg"
          justify="space-around"
        >
          <el-col :span="8">
            <div class="grid-content bg-purple ">
              <Theme :e-chang-theme="eChangTheme" />
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content bg-purple">
              <div class="Walletbox">
                <div
                  class="WalletInnerBox"
                  @click.stop="Copy"
                >
                  <span
                    class="avatarBox"
                    @click="openEmpower"
                  >
                    <img
                      class="avatar"
                      :src="$store.state.avatar"
                      alt=""
                      ref="avatar"
                      slot="reference"
                    >
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
            </div>
          </el-col>
        </el-row>
        <div class="BalanceEchart">
          <BalanceEchart @echartsChangTheme="echartsChangTheme" />
        </div>
        <div
          class="userInf"
          v-if="$store.state.isconnect && $store.state.isEmpower"
        >
          <UserInf />
        </div>
        <!-- <div class="userInf">
          <div class="feePercentage" />
        </div> -->
        <div
          class="WaitLogin"
          v-if="!$store.state.isEmpower"
        >
          <ConnectionTips />
        </div>
      </el-aside>
    </el-container>
    <transition
      enter-active-class="animate__animated animate__fadeIn"
      leave-active-class="animate__animated animate__fadeOut"
    >
      <div
        class="EmpowerMask"
        v-if="isGetToken"
      >
        <div class="EmpowerBox ">
          <div class="EmpowerUserInfbox">
            <div
              class="animate1 "
              v-if="anmiate1"
            >
              <span
                style="font-weight: 800;
            font-size: 1.2vw;padding-right:4% ;"
              >正在连接到 </span>
              <div class="EmpowerUserInf">
                <img
                  class="UserImage"
                  :src="$store.state.avatar"
                  alt=""
                  ref="avatar"
                  slot="reference"
                >
              </div>
              <div class="EmpowerUserAddress">
                {{ `${$store.state.currentAddress.slice(
                  0,
                  5
                )}...${$store.state.currentAddress.slice(37)}` }}
              </div>
            </div>
            <div
              class="animate2 animate__animated animate__fadeIn"
              v-if="animate2"
            >
              <div class="EmpowerUserInf">
                <img
                  class="UserImage"
                  :src="$store.state.avatar"
                  alt=""
                  ref="avatar"
                  slot="reference"
                >
              </div>
              <div class="EmpowerUserAddress">
                {{ `${$store.state.currentAddress.slice(
                  0,
                  5
                )}...${$store.state.currentAddress.slice(37)}` }}
              </div>
              <div class="EmpowerPasswordBox">
                <div class="EmpowerPasswordBoxTop">
                  <el-input
                    type="password"
                    v-model="user.encryptedPassword"
                    placeholder="请输入授权码"
                    @keydown.enter.prevent.native="checkEmpower"
                  />
                </div>
                <div class="EmpowerPasswordBoxBottom">
                  <div class="EmpowerPasswordOpt">
                    <span
                      style="margin-right: 5%;"
                      @click="viewControl(1)"
                    >忘记授权码</span>
                    <span
                      style="margin-left: 5%;"
                      @click="canacelEmpover"
                    >不授权直接进入</span>
                  </div>
                </div>
              </div>
            </div>
            <div
              class="animate3 animate__animated animate__fadeInLeft"
              v-if="animate3"
            >
              <div class="SignTop">
                <div class="SignUserInf">
                  <img
                    class="SignUserImage"
                    :src="$store.state.avatar"
                    alt=""
                    ref="avatar"
                    slot="reference"
                  >
                </div>
                <div class="SignUserAddress">
                  {{ `${$store.state.currentAddress.slice(
                    0,
                    5
                  )}...${$store.state.currentAddress.slice(37)}` }} 请设置授权信息和密保
                </div>
                <div />
              </div>
              <div class="SignForm">
                <el-form
                  label-position="top"
                  label-width="80px"
                  :model="EmpowerSignForm"
                >
                  <el-form-item label="授权码:">
                    <el-input
                      type="password"
                      v-model="user.encryptedPassword"
                      placeholder="请输入授权码"
                    />
                  </el-form-item>
                  <el-divider />
                  <el-form-item label="你喜欢看的电影:">
                    <el-input
                      v-model="EmpowerSignForm.sp1"
                      placeholder="请输入密保1"
                    />
                  </el-form-item>
                  <el-form-item label="你喜欢听的音乐:">
                    <el-input
                      v-model="EmpowerSignForm.sp2"
                      placeholder="请输入密保2"
                    />
                  </el-form-item>
                  <el-form-item label="你喜欢的运动:">
                    <el-input
                      v-model="EmpowerSignForm.sp3"
                      placeholder="请输入密保3"
                    />
                  </el-form-item>
                </el-form>
                <div class="SignSubmitBox">
                  <el-button
                    type="danger"
                    @click="isGetToken=false"
                  >
                    取消
                  </el-button>
                  <el-button
                    type="primary"
                    @click="empower"
                  >
                    授权
                  </el-button>
                </div>
              </div>
            </div>
            <div
              class="animate4 animate__animated animate__fadeInLeft"
              v-if="animate4"
            >
              <div class="SignTop">
                <div class="SignUserInf">
                  <img
                    class="SignUserImage"
                    :src="$store.state.avatar"
                    alt=""
                    ref="avatar"
                    slot="reference"
                  >
                </div>
                <div class="SignUserAddress">
                  {{ `${$store.state.currentAddress.slice(
                    0,
                    5
                  )}...${$store.state.currentAddress.slice(37)}` }} 找回密码
                </div>
                <div />
              </div>
              <div class="SignForm">
                <el-form
                  label-position="top"
                  label-width="80px"
                  :model="EmpowerSignForm"
                >
                  <el-form-item label="你喜欢看的电影:">
                    <el-input
                      v-model="EmpowerSignForm.sp1"
                      placeholder="请输入密保1"
                    />
                  </el-form-item>
                  <el-form-item label="你喜欢听的音乐:">
                    <el-input
                      v-model="EmpowerSignForm.sp2"
                      placeholder="请输入密保2"
                    />
                  </el-form-item>
                  <el-form-item label="你喜欢的运动:">
                    <el-input
                      v-model="EmpowerSignForm.sp3"
                      placeholder="请输入密保3"
                    />
                  </el-form-item>
                </el-form>
                <div class="SignSubmitBox">
                  <el-button
                    type="danger"
                    @click="viewControl(2)
                    "
                  >
                    返回
                  </el-button>
                  <el-button
                    type="primary"
                    @click="backPassword"
                  >
                    找回
                  </el-button>
                </div>
              </div>
            </div>
            <div
              class="animate5"
              v-if="animate5"
            >
              <div style="margin-bottom: 20px;font-size: 1vw;">
                请设置新的授权码
              </div>
              <el-input
                v-model="newMnemonic"
                placeholder="请设置新的授权码"
              />
              <el-button
                type="danger"
                @click="viewControl(3)"
                style="margin-top: 20px;"
              >
                取消
              </el-button>
              <el-button
                type="primary"
                @click="resetMnemonic"
                style="margin-top: 20px;"
              >
                重设
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>
<script>
  import { postOwnerContractList } from "@/api/axios/ownerContractLIst";
  import {
    getToken,
    setMnemonic,
    checkUserExist,
    setAuthenticationMetaInformation,
    forgetMnemonic,
    resetMnemonic,
  } from "@/api/axios/user";
  import CryptoJS from "crypto-js";
  import ConnectionTips from "../user/ConnectionTips.vue";
  import Theme from "@/components/oasisTheme.vue";
  import BalanceEchart from "@/views/extend/BalanceEchart.vue";
  import UserInf from "@/views/user/UserInf.vue";
  import ChatMemu from "@/views/chat/ChatMemu.vue";
  export default {
    name: "MarketShopIndex",
    components: {
      Theme,
      BalanceEchart,
      UserInf,
      ChatMemu,
      ConnectionTips,
    },
    data() {
      return {
        //****聊天栏******//
        isOpen: false,
        UserImageList: [],
        avatar: "",
        //*********************//
        //方法
        eChangTheme: null,
        //连接钱包
        user: {},
        changeingAccount: false,
        anmiate1: true,
        animate2: false,
        animate3: false,
        animate4: false,
        animate5: false,
        EmpowerSignForm: {
          sp1: "",
          sp2: "",
          sp3: "",
        },
        isGetToken: false,
        CopyTips: "点击复制",
        CopySuccess: "复制成功！",
        isCopy: false,
        isRepeatClick: true,
        newMnemonic: "",
        isUnlocked:false
      };
    },
    created() {},
    async mounted() {
      if (window.ethereum != undefined) {
        await window.ethereum._metamask.isUnlocked().then(re => this.isUnlocked = re)
        window.ethereum.on("accountsChanged", async () => {
          window.location.reload();
        });
        window.ethereum.on("chainChanged", () => {
          window.location.reload();
          this.$notify({
            title: "已切换链",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        });
        if (!this.$store.state.isconnect && this.isUnlocked) {
          await this.connectWallet();
          this.isGetToken = true;
          setTimeout(async () => {
            this.anmiate1 = false;
            await checkUserExist({
              userAddress: this.$store.state.currentAddress,
            }).then(async (re) => {
              if (re.data.data == "") {
                this.animate3 = true;
              } else {
                this.animate2 = true;
              }
            });
          }, 2000);
        }
      }
    },
    methods: {
      async empower() {
        this.user.encryptedPassword = CryptoJS.SHA256(
          this.user.encryptedPassword
        ).toString();
        await setMnemonic(this.user)
        this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
        await setAuthenticationMetaInformation(this.EmpowerSignForm)
        await getToken(this.user).then(async (re) => {
          localStorage.clear();
          if (re.data.data == null) {
            alert("密码错误");
          } else {
            let currentAddress = {
              ownerAddress: this.$store.state.currentAddress,
            };
            await postOwnerContractList(currentAddress).then((re) => {
              this.$store.commit("setOwnerNFTList", re.data.data);
            });
            localStorage.setItem("token", re.data.data);
            this.giveChatInitToWalletConnect();
            this.$store.commit("connection", true);
            this.$store.commit("setEmpower", true);
            this.isGetToken = false;
            this.$notify({
              title: "🎉 连接成功",
              position: "top-left",
              offset: 200,
            });
            
          }
        });
      },
      openEmpower() {
        if (this.$store.state.isEmpower || window.ethereum == undefined) {
          return
        }
        if (!this.isUnlocked) {
          this.$notify({
            title: "钱包未解锁",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
          return
        } else {
          this.isGetToken = true;
        }
       
         
      },
      async checkEmpower() {
        this.user.encryptedPassword = CryptoJS.SHA256(
          this.user.encryptedPassword
        ).toString();
        await getToken(this.user).then((re) => {
          localStorage.clear();
          if (re.data.data == null) {
            this.user.encryptedPassword = "";
            this.$notify({
              title: "授权码错误",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
          } else {
            let currentAddress = {
              ownerAddress: this.$store.state.currentAddress,
            };
            postOwnerContractList(currentAddress).then((re) => {
              this.$store.commit("setOwnerNFTList", re.data.data);
            });
            localStorage.setItem("token", re.data.data);
            this.$store.commit("setEmpower", true);
            this.$store.commit("connection", true);
            this.giveChatInitToWalletConnect();
            this.isGetToken = false;
            this.$notify({
              title: "🎉 连接成功",
              position: "top-left",
              offset: 200,
            });
          }
        });
      },
      viewControl(opt) {
        switch (opt) {
          case 1:
            this.animate2 = false;
            this.animate4 = true;
            break;
          case 2:
            this.animate2 = true;
            this.animate4 = false;
            break;
          case 3:
            this.animate2 = true;
            this.animate5 = false;
            break;
        }
      },
      async backPassword() {
        this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
        forgetMnemonic(this.EmpowerSignForm).then((re) => {
          if (re.data.data) {
            this.animate5 = true;
            this.animate4 = false;
          } else {
            this.$notify({
              title: "密保错误",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
          }
        });
      },
      async resetMnemonic() {
        this.EmpowerSignForm.newMnemonic = CryptoJS.SHA256(
          this.newMnemonic
        ).toString();
        this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
        await checkUserExist(this.EmpowerSignForm).then((re) => {
          if (re.data.data == this.EmpowerSignForm.encryptedPassword) {
            this.$notify({
              title: "不能与旧密码重复",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
            return;
          } else {
            resetMnemonic(this.EmpowerSignForm).then((re) => {
              if (re.data.data == null) {
                this.$notify.error({
                  title: "系统异常",
                  position: "top-left",
                  offset: 200,
                });
              }
              this.$notify({
                title: "🎉 重设成功",
                position: "top-left",
                offset: 200,
              });
              this.animate2 = true;
              this.animate5 = false;
            });
          }
        });
      },
      canacelEmpover() {
        this.isGetToken = false;
        this.$store.commit("connection", true);
        localStorage.clear();
      },
      giveChatInitToWalletConnect() {
        console.log("聊天栏正初始化.......");
        this.$refs.ChatMemu.init();
      },
      echartsChangTheme(method) {
        this.eChangTheme = method;
      },
      async connectWallet() {
          if (
            localStorage.getItem["token"] == null ||
            !this.$store.state.isconnect
          ) {
            try {
              // 请求用户授权
              await window.ethereum
                .request({ method: "eth_requestAccounts" })
                .then(async (handleAccountsChanged) => {
                  this.$store.commit(
                    "setcurrentAddress",
                    handleAccountsChanged[0]
                  );
                  this.$store.commit("changeAvatar", handleAccountsChanged[0]);
                  this.user.userAddress = handleAccountsChanged[0];
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
            } catch (error) {
              console.error(error);
              this.$notify.error({
                title: "连接失败",
                position: "top-left",
                offset: 200,
              });
            }
          }
      },
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
<style lang="scss" scoped>
@import "@/style/index.css";
</style>
<style lang="scss" scoped>
@import "@/style/MarketShop/index.scss";
</style> 
  
<style lang="scss" scoped>
.EmpowerMask {
  z-index: 200;
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  .EmpowerBox {
    transition: all 0.3s ease-in-out;
    overflow: hidden;
    display: flex;
    flex-direction: column;
    width: 650px;
    border-radius: 30px;
    background-color: white;
    display: flex;
    box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px,
      rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
    .EmpowerUserInfbox {
      margin: 5% 0 5% 0;
      width: 100%;
      height: 45%;
      display: flex;
      justify-content: center;
      align-items: center;
      .animate1 {
        display: flex;
        justify-content: center;
        align-items: center;
        transition: all 1s ease-in-out;
        width: 100%;
        .EmpowerUserInf {
          width: 120px;
          height: 120px;
          border-radius: 50%;
          overflow: hidden;
          border: 4px solid var(--border-green--);
          .UserImage {
            width: 100%;
            height: 100%;
            object-fit: cover;
          }
        }
        .EmpowerUserAddress {
          font-weight: 800;
          font-size: 1.2vw;
          padding-left: 4%;
        }
      }
      .animate2 {
        transition: all 1s ease-in-out;
        @extend .animate1;
        .EmpowerPasswordBox {
          margin-left: 4%;
          .EmpowerPasswordBoxTop {
            width: 100%;
          }
          .EmpowerPasswordBoxBottom {
            width: 100%;
            font-size: 0.7vw;
            .EmpowerPasswordOpt {
              margin-top: 30px;
              display: flex;
              justify-content: space-around;
              span {
                cursor: pointer;
              }
            }
          }
        }
      }
      .animate3 {
        @extend .animate1;
        flex-direction: column;
        .SignTop {
          width: 100%;
          display: flex;
          justify-content: center;
          align-items: center;
          .SignUserInf {
            width: 50px;
            height: 50px;
            border-radius: 50%;
            overflow: hidden;
            border: 2px solid #ff0000;
            .SignUserImage {
              width: 100%;
              height: 100%;
              object-fit: cover;
            }
          }
          .SignUserAddress {
            font-weight: 800;
            font-size: 1vw;
            padding-left: 2%;
          }
        }

        .SignForm {
          width: 50%;
          margin-top: 2%;
          .SignSubmitBox {
            display: flex;
            margin: 3% 0 3% 0;
            justify-content: space-around;
            align-items: center;
          }
        }
      }
      .animate4 {
        @extend .animate3;
      }
    }
  }
}
</style>
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
.feePercentage {
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
}
</style>