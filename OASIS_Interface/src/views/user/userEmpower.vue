<template>
  <transition
    enter-active-class="animate__animated animate__fadeIn"
    leave-active-class="animate__animated animate__fadeOut"
  >
    <div
      class="EmpowerMask "
      v-if="isGetToken"
    >
      <div class="EmpowerBox animate__animated animate__flipInX">
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
                <el-form label-width="80px">
                  <el-form-item>
                    <template slot="label">
                      <i class="el-icon-lock" />
                    </template>
                    <el-input
                      clearable
                      class="EmpowerPassword"
                      type="password"
                      v-model="user.encryptedPassword"
                      placeholder="请输入授权码"
                      @keydown.enter.prevent.native="checkEmpower"
                      ref="password"
                    />
                  </el-form-item>
                </el-form>
              </div>
              <div class="EmpowerPasswordBoxBottom">
                <div class="EmpowerPasswordOpt">
                  <span
                    style="margin-right: 5%;color:gray;"
                    @click="viewControl(1)"
                  >忘记授权码</span>
                  <span
                    style="margin-left: 5%;color:gray;"
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
                  @click="canacelEmpover"
                >
                  取消
                </el-button>
                <el-button
                  type="primary"
                  @click="empower(2)"
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
</template>

<script>
  import { newWebSocket } from "@/utils/webStocket";
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
  export default {
    data() {
      return {
        isGetToken:false,
        //****聊天栏******//
        isOpen: false,
        UserImageList: [],
        avatar: "",
        //*********************//
        //方法
        echartChange: null,
        //连接钱包
        user: {
          encryptedPassword: "",
        },
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
        newMnemonic: "",
        isUnlocked: false,
      };
    },
    computed:{
      listenGetToken() {
			return this.$store.state.isGetToken;
		}
  },
  watch: {
    listenGetToken() {
      this.isGetToken=this.$store.state.isGetToken;
    }
    },
  async mounted() {

      if (window.ethereum != undefined) {
        await window.ethereum._metamask
          .isUnlocked()
          .then((re) => (this.isUnlocked = re));
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
        this.openEmpower()
        if (!this.$store.state.isconnect && this.isUnlocked) {
          try {
            await this.connectWallet();
            this.$store.commit("setGetToken", true);
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
          } catch (error) {
            console.log(error);
          }
        }
      }
    },
    methods: {
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
                this.$store.commit("setcurrentAddress", handleAccountsChanged[0]);
                this.$store.commit("setAvatar", handleAccountsChanged[0]);
                this.user.userAddress = handleAccountsChanged[0];
              })
              .catch((error) => {
                this.$store.commit("setConnection", false);
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
      async checkEmpower() {
        this.user.encryptedPassword = CryptoJS.SHA256(
          this.user.encryptedPassword
        ).toString();
        await getToken(this.user).then((re) => {
          localStorage.clear();
          if (re.data.data == null) {
            this.$notify({
              title: "授权码错误",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
            this.$refs.password.clear()
          } else {
            localStorage.setItem("token", re.data.data);
            let currentAddress = {
              ownerAddress: this.$store.state.currentAddress,
            };
            postOwnerContractList(currentAddress).then((re) => {
              this.$store.commit("setOwnerNFTList", re.data.data);
            });
          
            this.$store.commit("setConnection", true);
            this.$store.commit("setEmpower", true);
            this.$store.commit("setGetToken", false);
            this.$refs.password.clear()
            // this.giveChatInitToWalletConnect();
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
          return;
        }
        if (!this.isUnlocked) {
          this.$notify({
            title: "钱包未解锁",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
          return;
        } else {
          this.$store.commit("setGetToken", true);
        }
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
          default:
            this.animate2 = false;
            this.$store.commit("setGetToken", true);
            this.anmiate1 = false;
            this.animate4 = true;
            this.user.encryptedPassword = "";
            break;
        }
      },
      // =====
      async empower(opt) {
        if (opt == 2) {
          if (
            this.EmpowerSignForm.sp1 == "" ||
            this.EmpowerSignForm.sp2 == "" ||
            this.EmpowerSignForm.sp3 == "" ||
            this.EmpowerSignForm.sp4 == "" ||
            this.EmpowerSignForm.sp5 == "" ||
            this.user.encryptedPassword == ""
          ) {
            this.$notify({
              title: "信息不能为空",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
            return;
          }
        }
        this.user.encryptedPassword = CryptoJS.SHA256(
          this.user.encryptedPassword
        ).toString();
        await setMnemonic(this.user);
        this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
        await setAuthenticationMetaInformation(this.EmpowerSignForm);
        await getToken(this.user).then(async (re) => {
          localStorage.clear();
          if (re.data.data == null) {
            this.$notify({
              title: "授权码错误",
              type: "waring",
              position: "top-left",
              offset: 200,
            });
          } else {
            localStorage.setItem("token", re.data.data);
            let currentAddress = {
              ownerAddress: this.EmpowerSignForm.userAddress,
            };
            
            await postOwnerContractList(currentAddress).then((re) => {
              this.$store.commit("setOwnerNFTList", re.data.data);
            });
            
            // this.giveChatInitToWalletConnect();
            this.$store.commit("setConnection", true);
            this.$store.commit("setEmpower", true);
            this.$store.commit("setGetToken", false);
           
            this.$notify({
              title: "🎉 连接成功",
              position: "top-left",
              offset: 200,
            });
          }
        });
      },
      canacelEmpover() {
        this.$store.commit("setGetToken", false);
        this.$store.commit("setConnection", false);
        localStorage.clear();
      },
      // =====
      async backPassword() {
        this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
        forgetMnemonic(this.EmpowerSignForm).then((re) => {
          if (re.data.data) {
            this.animate5 = true;
            this.animate4 = false;
            this.$notify({
              title: "密保正确",
              type: "success",
              position: "top-left",
              offset: 200,
            });
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
            //开启websocket事务
      async openWebsocket(user) {
        // 格式["username:admin","username:admin2"]
        let object = {
          username: user,
        };
        localStorage.setItem("user", JSON.stringify(object));

        this.user = localStorage.getItem("user")
          ? JSON.parse(localStorage.getItem("user"))
          : {};
        let username = this.user.username;

        let _this = this;

        if (typeof WebSocket == "undefined") {
          console.log("您的浏览器不支持WebSocket");
        } else {
          console.log("您的浏览器支持WebSocket");
          const userSocket=this.$store.state.userSocket
          if (userSocket != null) {
            userSocket.close();
            
          }
          // 开启一个websocket服务
          let socket = newWebSocket("/OasisChat/", username);
          //打开事件
          socket.onopen = function () {
            console.log("websocket已打开");
          };
          //  浏览器端收消息，获得从服务端发送过来的文本消息
          socket.onmessage = function (msg) {
            _this.$nextTick(() => {
              console.log("收到数据====" + msg.data);
              let data = JSON.parse(msg.data);
              if (data.users) {
                _this.users = data.users.filter(
                  (user) => user.username !== username
                );
              } else {
                if (data.from === _this.chatUser) {
                  _this.messages.push(data);
                  _this.createContent(data.from, null, data.message);
                }
              }
            });
          };
          //关闭事件
          socket.onclose = function () {
            console.log("websocket已关闭");
          };
          //发生了错误事件
          socket.onerror = function () {
            console.log("websocket发生了错误");
          };
          
        }
      },
      // giveChatInitToWalletConnect() {
      //   console.log("聊天栏正初始化.......");
      //   this.$refs.ChatMemu.init();
      // },
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/layout/userEmpower.scss";
</style>