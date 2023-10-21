<template>
  <div class="layout">
    <div class="leftMenu animate__animated animate__fadeInLeft">
      <oasisChat ref="ChatMemu" />
    </div>
    <div
      class="main"
      ref="main"
      @scroll="handleScroll"
    >
      <header
        class="menu animate__animated animate__fadeInDown"
        ref="menu"
        :class="{ 'fixed': isFixed }"
      >
        <div class="menu_search">
          <oasisSearch />
        </div>
        <div class="menu_selects">
          <div class="menu_select_links">
            <el-menu
              :default-active="activeIndex"
              class="menu_select_link_items"
              mode="horizontal"
              @select="handleSelect"
            >
              <el-menu-item index="/">
                <div class="menu_select_link_item">
                  <svg
                    t="1697629842922"
                    class="icon"
                    viewBox="0 0 1024 1024"
                    version="1.1"
                    xmlns="http://www.w3.org/2000/svg"
                    p-id="1446"
                    width="18"
                    height="18"
                  >
                    <path
                      class="icon-path"
                      stroke="#1875F0"
                      stroke-width="3"
                      d="M445.799424 193.792l198.4512 202.752a50.4832 50.4832 0 0 0 72.3968 0l0.3072-0.3072a53.1456 53.1456 0 0 0 0-73.984L443.393024 42.8032a50.2272 50.2272 0 0 0-67.584 0l-4.608 4.7104-34.304 35.0208L102.401024 322.0992l-84.992 87.04-4.608 4.7104a52.1216 52.1216 0 0 0-12.8 34.5088v523.5712a51.2 51.2 0 1 0 102.4 0V492.4928a52.8384 52.8384 0 0 1 15.0016-36.9664l256-261.5296a50.4832 50.4832 0 0 1 72.3968-0.2048z"
                      fill="#515151"
                      p-id="1447"
                    />
                    <path
                      class="icon-path"
                      stroke="#1875F0"
                      stroke-width="3"
                      d="M1011.201024 343.9616l-4.608-4.7104L687.105024 13.056a50.2272 50.2272 0 0 0-67.584 0l-4.608 4.7104a52.0704 52.0704 0 0 0-12.8 34.5088 53.6064 53.6064 0 0 0 14.848 37.12l291.84 297.9328a52.0704 52.0704 0 0 1 12.8 34.5088v419.3792a77.6192 77.6192 0 0 1-76.8 78.2336h-409.6A77.6192 77.6192 0 0 1 358.401024 841.216v-296.5504a52.0704 52.0704 0 0 1 12.8-34.5088L435.201024 444.8256a53.8112 53.8112 0 0 0 0-74.24 50.8416 50.8416 0 0 0-72.704 0L273.409024 461.5168l-4.608 4.7104A52.0704 52.0704 0 0 0 256.001024 500.736v366.2336A155.904 155.904 0 0 0 409.601024 1024h460.8a155.8016 155.8016 0 0 0 153.6-156.928V378.4192a52.0704 52.0704 0 0 0-12.8-34.4576z"
                      fill="#515151"
                      p-id="1448"
                    />
                  </svg>
                  NFTs
                </div>
              </el-menu-item>
              <el-menu-item
                v-if="$store.state.isconnect"
                index="/home/ImitNFT"
              >
                <div class="menu_select_link_item">
                  <svg
                    t="1697633111966"
                    class="icon"
                    viewBox="0 0 1024 1024"
                    version="1.1"
                    xmlns="http://www.w3.org/2000/svg"
                    p-id="4342"
                    width="25"
                    height="25"
                  >
                    <path
                      class="icon-path"
                      stroke="#1875F0"
                      stroke-width="3"
                      d="M835.9424 245.0432L550.1952 80.128a76.8 76.8 0 0 0-76.4416 0L188.0576 245.0432A76.8 76.8 0 0 0 149.8624 311.296v329.8304a76.8 76.8 0 0 0 38.1952 66.304l285.6448 164.9152a75.9296 75.9296 0 0 0 76.4928 0l285.7472-164.9152a76.8 76.8 0 0 0 38.1952-66.304V311.296a76.8 76.8 0 0 0-38.1952-66.2528M471.1936 498.8416v278.1696l-239.7696-135.8848V360.448zM272.7424 290.048l239.0528-138.24 238.9504 138.24L511.744 428.288z m522.24 346.5216L552.8064 776.448V498.5856l240.0256-138.5472z"
                      fill="#707070"
                      p-id="4343"
                    />
                  </svg>
                  ImitNFT
                </div>
              </el-menu-item>
              <el-menu-item index="/home/MarketShop3D">
                <div class="menu_select_link_item">
                  <svg
                    t="1697646251949"
                    class="icon"
                    viewBox="0 0 1024 1024"
                    version="1.1"
                    xmlns="http://www.w3.org/2000/svg"
                    p-id="3702"
                    width="25"
                    height="25"
                  >
                    <path
                      class="icon-path"
                      stroke="#1875F0"
                      stroke-width="3"
                      d="M177.493333 389.36381a48.761905 48.761905 0 0 1 24.380953 6.534095l249.222095 143.896381a73.142857 73.142857 0 0 1 36.571429 63.341714v287.792762a48.761905 48.761905 0 0 1-45.104762 48.615619l-3.657143 0.146286a48.761905 48.761905 0 0 1-24.380953-6.534096L165.302857 789.26019a73.142857 73.142857 0 0 1-36.571428-63.341714V438.125714a48.761905 48.761905 0 0 1 45.104761-48.64l3.657143-0.121904z m669.135238 0a48.761905 48.761905 0 0 1 48.615619 45.104761l0.146286 3.657143v287.792762a73.142857 73.142857 0 0 1-32.426666 60.757334l-4.144762 2.58438-249.246477 143.896381a48.761905 48.761905 0 0 1-20.114285 6.339048l-4.266667 0.195048a48.761905 48.761905 0 0 1-48.615619-45.104762l-0.146286-3.657143V603.136a73.142857 73.142857 0 0 1 32.426667-60.757333l4.144762-2.584381 249.246476-143.896381a48.761905 48.761905 0 0 1 20.114286-6.339048l4.266666-0.195047zM201.849905 480.353524l0.024381 245.564952 212.650666 122.758095V603.136L201.874286 480.353524z m620.373333 0l-210.553905 121.539047-2.096762 1.243429V848.700952l210.578286-121.514666 2.096762-1.243429-0.024381-245.564952zM544.085333 122.441143l4.608 2.438095 249.222096 143.896381a48.761905 48.761905 0 0 1 3.242666 82.407619l-3.242666 2.048-249.222096 143.896381a73.142857 73.142857 0 0 1-68.559238 2.438095l-4.583619-2.438095-249.246476-143.896381a48.761905 48.761905 0 0 1-3.218286-82.407619l3.218286-2.048 249.246476-143.896381a73.142857 73.142857 0 0 1 68.534857-2.438095zM512.048762 188.220952l-212.601905 122.758096 211.748572 122.270476 0.950857 0.487619 212.626285-122.758095-211.309714-122.002286-1.414095-0.75581z"
                      p-id="3703"
                      fill="#707070"
                    />
                  </svg>
                  3DModel
                </div>
              </el-menu-item>
            </el-menu>
          </div>

          <div class="menu_select_oasisTheme">
            <oasisTheme />
          </div>
          <div class="menu_select_walletConnection">
            <walletConnection />
          </div>
        </div>
      </header>
      <router-view />
      <!-- <div class="view">
     
      </div> -->
      <!-- <router-view /> -->
      <userEmpower />
    </div>
  </div>
</template>
<script>
  // components
  import userEmpower from "./user/userEmpower.vue";
  import oasisTheme from "@/views/bar/oasisTheme.vue";
  import walletConnection from "@/views/bar/walletConnect.vue";
  import oasisSearch from "@/views/topMenu/oasisSearch.vue";
  import oasisChat from "@/views/leftMenu/oasisChat.vue";
// func
  export default {
    name: "MarketShopIndex",
    components: {
      oasisTheme,
      oasisChat,
      oasisSearch,
      walletConnection,
      userEmpower,
    },
    data() {
      return {
        activeIndex: "/",
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
        isRepeatClick: true,
        newMnemonic: "",
        isUnlocked: false,

        isFixed: false,
        scrollOffset: 300
      };
    },
  created() {
     
    },
  mounted() {
    this.$refs.main.addEventListener('scroll', this.handleScroll);
//     function findScroller(element) {
//     element.onscroll = function() { console.log(element)}

//     Array.from(element.children).forEach(findScroller);
// }
// findScroller(document.body);
  },
  beforeUnmount() {
    this.$refs.main.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    // handleScroll(event) {
    //   const scrollContainer = event.target;
    //   const scrollTop = scrollContainer.scrollTop;
    //   this.isFixed = scrollTop >= this.scrollOffset;
    //   // 在这里可以根据 scrollTop 值执行相关操作，例如调整其他元素的样式或位置等
    // }
    handleScroll() {
      const scrollContainer = this.$refs.main;
      const scrollTop = scrollContainer.scrollTop;
      this.isFixed = scrollTop > 40;
    },
    handleSelect(key) {
      this.$router.push(key)
    },
    
      // async empower(opt) {
      //   if (opt == 2) {
      //     if (
      //       this.EmpowerSignForm.sp1 == "" ||
      //       this.EmpowerSignForm.sp2 == "" ||
      //       this.EmpowerSignForm.sp3 == "" ||
      //       this.EmpowerSignForm.sp4 == "" ||
      //       this.EmpowerSignForm.sp5 == "" ||
      //       this.user.encryptedPassword == ""
      //     ) {
      //       this.$notify({
      //         title: "信息不能为空",
      //         type: "warning",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //       return;
      //     }
      //   }
      //   this.user.encryptedPassword = CryptoJS.SHA256(
      //     this.user.encryptedPassword
      //   ).toString();
      //   await setMnemonic(this.user);
      //   this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
      //   await setAuthenticationMetaInformation(this.EmpowerSignForm);
      //   await getToken(this.user).then(async (re) => {
      //     localStorage.clear();
      //     if (re.data.data == null) {
      //       this.$notify({
      //         title: "密码错误",
      //         type: "waring",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     } else {
      //       localStorage.setItem("token", re.data.data);
      //       let currentAddress = {
      //         ownerAddress: this.EmpowerSignForm.userAddress,
      //       };
      //       await postOwnerContractList(currentAddress).then((re) => {
      //         this.$store.commit("setOwnerNFTList", re.data.data);
      //       });
      //       this.giveChatInitToWalletConnect();
      //       this.$store.commit("connection", true);
      //       this.$store.commit("setEmpower", true);
      //       this.isGetToken = false;
      //       this.$notify({
      //         title: "🎉 连接成功",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     }
      //   });
      // },
      // openEmpower() {
      //   if (this.$store.state.isEmpower || window.ethereum == undefined) {
      //     return;
      //   }
      //   if (!this.isUnlocked) {
      //     this.$notify({
      //       title: "钱包未解锁",
      //       type: "warning",
      //       position: "top-left",
      //       offset: 200,
      //     });
      //     return;
      //   } else {
      //     this.isGetToken = true;
      //   }
      // },
      // async checkEmpower() {
      //   this.user.encryptedPassword = CryptoJS.SHA256(
      //     this.user.encryptedPassword
      //   ).toString();
      //   await getToken(this.user).then((re) => {
      //     localStorage.clear();
      //     if (re.data.data == null) {
      //       this.$notify({
      //         title: "授权码错误",
      //         type: "warning",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     } else {
      //       let currentAddress = {
      //         ownerAddress: this.$store.state.currentAddress,
      //       };
      //       postOwnerContractList(currentAddress).then((re) => {
      //         this.$store.commit("setOwnerNFTList", re.data.data);
      //       });
      //       localStorage.setItem("token", re.data.data);
      //       this.$store.commit("setEmpower", true);
      //       this.$store.commit("connection", true);
      //       this.giveChatInitToWalletConnect();
      //       this.isGetToken = false;
      //       this.$notify({
      //         title: "🎉 连接成功",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     }
      //   });
      // },
      // viewControl(opt) {
      //   switch (opt) {
      //     case 1:
      //       this.animate2 = false;
      //       this.animate4 = true;
      //       break;
      //     case 2:
      //       this.animate2 = true;
      //       this.animate4 = false;
      //       break;
      //     case 3:
      //       this.animate2 = true;
      //       this.animate5 = false;
      //       break;
      //     default:
      //       this.animate2 = false;
      //       this.isGetToken = true;
      //       this.anmiate1 = false;
      //       this.animate4 = true;
      //       this.user.encryptedPassword = "";
      //       break;
      //   }
      // },
      // async backPassword() {
      //   this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
      //   forgetMnemonic(this.EmpowerSignForm).then((re) => {
      //     if (re.data.data) {
      //       this.animate5 = true;
      //       this.animate4 = false;
      //       this.$notify({
      //         title: "密保正确",
      //         type: "success",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     } else {
      //       this.$notify({
      //         title: "密保错误",
      //         type: "warning",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     }
      //   });
      // },
      // async resetMnemonic() {
      //   this.EmpowerSignForm.newMnemonic = CryptoJS.SHA256(
      //     this.newMnemonic
      //   ).toString();
      //   this.EmpowerSignForm.userAddress = this.$store.state.currentAddress;
      //   await checkUserExist(this.EmpowerSignForm).then((re) => {
      //     if (re.data.data == this.EmpowerSignForm.encryptedPassword) {
      //       this.$notify({
      //         title: "不能与旧密码重复",
      //         type: "warning",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //       return;
      //     } else {
      //       resetMnemonic(this.EmpowerSignForm).then((re) => {
      //         if (re.data.data == null) {
      //           this.$notify.error({
      //             title: "系统异常",
      //             position: "top-left",
      //             offset: 200,
      //           });
      //         }
      //         this.$notify({
      //           title: "🎉 重设成功",
      //           position: "top-left",
      //           offset: 200,
      //         });
      //         this.animate2 = true;
      //         this.animate5 = false;
      //       });
      //     }
      //   });
      // },
      // canacelEmpover() {
      //   this.isGetToken = false;
      //   this.$store.commit("connection", true);
      // },
      // giveChatInitToWalletConnect() {
      //   console.log("聊天栏正初始化.......");
      //   this.$refs.ChatMemu.init();
      // },
      // setEchartChange(data) {
      //   this.echartChange = data;
      //   console.log(this.echartChange);
      // },
      // async connectWallet() {
      //   if (
      //     localStorage.getItem["token"] == null ||
      //     !this.$store.state.isconnect
      //   ) {
      //     try {
      //       // 请求用户授权
      //       await window.ethereum
      //         .request({ method: "eth_requestAccounts" })
      //         .then(async (handleAccountsChanged) => {
      //           this.$store.commit("setcurrentAddress", handleAccountsChanged[0]);
      //           this.$store.commit("changeAvatar", handleAccountsChanged[0]);
      //           this.user.userAddress = handleAccountsChanged[0];
      //         })
      //         .catch((error) => {
      //           this.$store.commit("connection", false);
      //           if (error.code === 4001) {
      //             // EIP-1193 userRejectedRequest error
      //             console.log("Please connect to MetaMask.");
      //           } else {
      //             console.error(error);
      //           }
      //         });
      //     } catch (error) {
      //       console.error(error);
      //       this.$notify.error({
      //         title: "连接失败",
      //         position: "top-left",
      //         offset: 200,
      //       });
      //     }
      //   }
      // },
      // Copy() {
      //   navigator.clipboard
      //     .writeText(this.$store.state.currentAddress)
      //     .then(() => {
      //       this.isCopy = true;
      //       setTimeout(() => {
      //         this.isCopy = false;
      //       }, 3000);
      //     });
      // },
      // GETHashAvatar() {
      //   if (this.$store.state.isconnect) {
      //     this.avatar =
      //       "data:image/png;base64," +
      //       new this.Identicon(this.$store.state.currentAddress, 120).toString();
      //     this.$refs.avatar.width = 60;
      //     this.$refs.avatar.height = 60;
      //   } else {
      //     this.avatar = require("@/assets/webAssets/MetaMask.png");
      //     this.$refs.avatar.width = 40;
      //     this.$refs.avatar.height = 40;
      //   }
      // },
    },
  };
</script>
<!-- <style lang="scss" scoped>
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
          margin-top: 3%;
          margin-bottom: 3%;
          font-weight: 800;
          font-size: 1.2vw;
          padding-left: 4%;
        }
      }
      .animate2 {
        transition: all 1s ease-in-out;
        @extend .animate1;
        .EmpowerUserAddress {
          padding-left: 0;
        }
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        .EmpowerPasswordBox {
          .EmpowerPasswordBoxTop {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 100%;
            .el-form {
              margin-right: 10%;
            }
            .el-input /deep/ .el-input__inner {
              border: 1px solid rgba(0, 0, 0, 0);
              border-bottom: 1px solid rgb(0, 0, 0);
              border-radius: 0px;
            }
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
        .el-input /deep/ .el-input__inner {
          border: 1px solid rgba(0, 0, 0, 0);
          border-bottom: 1px solid rgb(0, 0, 0);
          border-radius: 0px;
        }
      }
      .animate5 {
        .el-input /deep/ .el-input__inner {
          border: 1px solid rgba(0, 0, 0, 0);
          border-bottom: 1px solid rgb(0, 0, 0);
          border-radius: 0px;
        }
      }
    }
  }
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
</style> -->

<style lang="scss" scoped>

@import "@/style/layout/layout.scss";
</style>
