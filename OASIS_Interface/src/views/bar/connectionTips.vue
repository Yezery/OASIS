<template>
  <div class="WaitLoginBox">
    <div class="WaitLoginBox-main">
      <div
        class="WaitLoginBox-head"
        ref="WaitLoginBox_head"
      >
        <div class="WaitLoginBox-head-avatarBox">
          <div class="avatarBox ">
            <img
              class="avatar animate__animated"
              src="@/assets/webAssets/MetaMask.png"
              width="90px"
              height="90px"
              alt=""
              style=""
              ref="avatar"
            >
          </div>
        </div>
        <div class="WaitLoginBox-head-nameBox">
          <div class="nameBox">
            <span class="name">{{ this.tips }}
              <span v-if="!metaMask">
                <a
                  href="https://go.microsoft.com/fwlink/?LinkID=2093505"
                  v-if="windows32 "
                >window32位Edge</a>
                <a
                  href="https://go.microsoft.com/fwlink/?LinkID=2093437"
                  v-else-if="windows64"
                >window64位Edge</a>
                <a
                  href="https://www.microsoft.com/zh-cn/edge/download?form=MA13FJ"
                  v-else-if="mac"
                >Mac版Edge</a>
              </span>
             
            </span>

            <span
              class="title"
              v-html="tipsE"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        tips: "",
        tipsE: "",
        mac: false,
        windows64: false,
        windows32: false,
        metaMask: false,
      };
    },
    mounted() {
      if (window.ethereum == undefined) {
       
        this.tips = "您的浏览器不支持/未安装";
        this.tipsE = "Your browser does not support <br/> or install MetaMask";
      } else {
        this.metaMask = true
        this.tips = "账号没有授权请点击上方头像进行授权";
        this.tipsE = `If the account is not authorized <br/> please click the avatar above to authorize`;
      }
      this.OSnow();
    },
    methods: {
      //判断系统类型
      OSnow() {
        var agent = navigator.userAgent.toLowerCase();
        var isMac = /macintosh|mac os x/i.test(navigator.userAgent);
        if (agent.indexOf("win32") >= 0 || agent.indexOf("wow32") >= 0) {
          this.windows32 = true;
        }
        if (agent.indexOf("win64") >= 0 || agent.indexOf("wow64") >= 0) {
          this.windows64 = true;
        }
        if (isMac) {
          this.mac = true;
        }
      },
    },
  };
</script>

<style lang="scss" scoped>
.WaitLoginBox {
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
  .WaitLoginBox-main {
    .WaitLoginBox-head {
      height: 260px;
      margin-top: 10%;

      .WaitLoginBox-head-avatarBox {
        display: flex;
        justify-content: center;
        align-items: center;

        .avatarBox {
          width: 89px;
          height: 89px;
          .avatar {
            object-fit: fill;
          }
        }
      }
      .WaitLoginBox-head-nameBox {
        margin-top: 8%;
        .nameBox {
          display: flex;
          flex-direction: column;
          .name {
            color: var(--blueblack-white--);
            font-weight: 700;
            font-size: 13px;
            text-align: center;
            font-family: Gilroy-Medium;
            a {
              color: var(--blueblack-white--);
            }
          }
          .title {
            margin-top: 2%;
            color: rgb(112, 122, 137);
            font-size: 12px;
            font-weight: 700;
          }
        }
      }
    }
  }
}
</style>