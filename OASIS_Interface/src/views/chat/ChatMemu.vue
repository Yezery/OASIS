<template>
  <div
    class="ChatMemu "
    ref="ChatMemu"
    @mouseenter="OpenAndClose"
    @mouseleave="OpenAndClose"
  >
    <div class="logobox">
      <div
        class="logo animate__animated"
        @click="isSearch"
      >
        <router-link :to="{ name: 'MarketShop' }">
          <div
            style="display: inline; font-size: 30px"
            class="animate__animated"
            ref="logo"
          >
            OAS
          </div>
          <div
            style="display: inline;  font-size: 30px"
            class="animate__animated"
            ref="logo2"
          >
            <img
              src="../../assets/webAssets/s1.png" 
              alt=""
              width="27px"
              height="27px"
            >S
          </div>
        </router-link>
      </div>
    </div>

    <template>
      <div
        v-for="otherUser in users"
        :key="otherUser.username"
      >
        <el-popover
          placement="right"
          width="309px"
          trigger="click"
          popper-class="monitor-yt-popover"
        >
          <div
            class="ChatPeople animate__animated animate__fadeInRight"
            slot="reference"
            @click="chatUser = otherUser.username"
          >
            <div class="ChatWindow">
              <img
                :src="GETHashAvatar(otherUser.username)"
                width="50px"
                height="50px"
                alt=""
                style="object-fit: contain"
              >
            </div>
            <span class="friendname animate__animated animate__fadeInRight">{{ otherUser.username.slice(0, 3) + "..." + otherUser.username.slice(39) }}</span>
          </div>
          <div
            style="
              display: flex;
              flex-direction: column;
              justify-content: center;
              align-items: center;
            "
          >
            <div
              class="ChatWindowBox"
              style="width: 309px; height: 499px"
            >
              <el-container>
                <el-header>
                  <div class="topBox">
                    <div
                      style="
                        flex: 1;
                        justify-content: center;
                        align-items: center;
                        display: flex;
                      "
                    >
                      <span
                        class="el-avatar el-avatar--circle"
                        style="height: 48px; width: 48px"
                        v-if="chatUser != ''"
                      >
                        <img
                          :src="GETHashAvatar(otherUser.username)"
                          style="object-fit: contain"
                        >
                      </span>
                    </div>
                    <div style="flex: 7; text-align: left; height: 48px">
                      <div style="margin-left: 5px">
                        <div class="ChatUserName">
                          {{ chatUser.slice(0, 5) + "..." + chatUser.slice(37) }}
                        </div>
                        <div
                          style="
                            font-size: 12px;
                            color: limegreen;
                            margin-left: 5px;
                            font-weight: 600;
                            font-size: 14px;
                          "
                          v-if="chatUser != ''"
                        >
                          <i class="el-icon-loading" /> Chatting...
                        </div>
                      </div>
                    </div>
                  </div>
                </el-header>

                <el-main>
                  <div
                    style="height: 355px; overflow: auto"
                    id="CHAT"
                    v-html="contents[chatUser]"
                  />
                </el-main>
                <el-footer>
                  <div class="inputer">
                    <input
                      type="textarea"
                      autofocus="true"
                      v-model="text"
                      class="inputarea"
                      @keydown.enter.prevent="handleKeyCode($event)"
                    >
                    <button @click="send">
                      <i
                        class="el-icon-thumb"
                        style="font-size: 27px"
                      />
                    </button>
                  </div>
                </el-footer>
              </el-container>
            </div>
          </div>
        </el-popover>
      </div>
    </template>
  </div>
</template>

<script>
  let socket;
  import { newWebSocket } from "@/utils/webStocket";
  import { parseTime } from "@/utils/Time";
  export default {
    data() {
      return {
        // Chat
        user: {},
        isCollapse: false,
        users: [],
        chatUser: "",
        text: "",
        messages: [],
        contents: {},
      };
    },
    components: {},
    methods: {
      // 更新弹窗位置
      upDatePosition() {
        this.$nextTick(() => {
          this.$refs.popoverRef.updatePopper();
        });
      },
      OpenAndClose() {
        this.isOpen = !this.isOpen;
        if (this.isOpen) {
          this.$refs.logo.classList.add("animate__swing");
          this.$refs.logo2.classList.add("animate__swing");
        } else {
          this.$refs.logo.classList.remove("animate__swing");
          this.$refs.logo2.classList.remove("animate__swing");
        }
      },
      GETHashAvatar(UserAddress) {
        return (
          "data:image/png;base64," +
          new this.Identicon(UserAddress, 120).toString()
        );
      },
      send() {
        if (!this.chatUser) {
          this.$message({ type: "warning", message: "请选择聊天对象" });
          return;
        }
        if (!this.text) {
          this.$message({ type: "warning", message: "请输入内容" });
        } else {
          if (typeof WebSocket == "undefined") {
            console.log("您的浏览器不支持WebSocket");
          } else {
            console.log("您的浏览器支持WebSocket");
            // 组装待发送的消息 json
            // {"from": "zhang", "to": "admin", "text": "聊天文本"}

            let message = {
              from: this.user.username,
              to: this.chatUser,
              message: this.text,
            };

            socket.send(JSON.stringify(message)); // 将组装好的json发送给服务端，

            // 由服务端进行转发
            this.messages.push({
              user: this.user.username,
              text: this.text,
            });

            // 构建消息内容，本人消息
            this.createContent(null, this.user.username, this.text);
            this.text = "";
          }
        }
      },
      // 键盘回车事件
      handleKeyCode(event) {
        if (event.keyCode == 13) {
          if (!event.metaKey) {
            event.preventDefault();
            this.send();
          } else {
            this.text = this.text + "\n";
          }
        }
      },
      createContent(remoteUser, nowUser, text) {
        let content;
        let html;
        content = this.contents[this.chatUser || remoteUser] || "";
        // 当前用户消息
        if (nowUser) {
          html = `<div class="MessageBox_nowUser"><span class="currentTime">${parseTime(
            Date(),
            "hh:mm:ss"
          )}</span><div class="Message leftMessage">${text}</div></div>`;
        } else if (remoteUser) {
          html = `
                            <div class="MessageBox_remoteUser">
                              <span class="currentTime">${parseTime(
                                Date(),
                                "hh:mm:ss"
                              )}</span>
                              <div class="Message rightMessage">
                                ${text}
                              </div>
                            </div>
                          `;
        }
        content += html;
        this.$set(this.contents, this.chatUser, content);
        this.$nextTick(() => {
          let msg = document.getElementById("CHAT");
          msg.scrollTop = msg.scrollHeight;
        });
      },
      //开启websocket事务
      init() {
        // 格式["username:admin","username:admin2"]
        let object = {
          username: this.$store.state.currentAddress,
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
          if (socket != null) {
            socket.close();
            socket = null;
          }
          // 开启一个websocket服务
          socket = newWebSocket("/OasisChat/", username);

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
      isSearch() {
        this.$store.commit("setIsSearch", false);
      }
    },
  };
</script>
<style lang="scss" scoped>
.ChatMemu {
  font-family: "Transformers_Movie";
  width: 80px;
  height:98vh;
  transition: all 0.5s ease-in-out;
  transition-delay: 0.5s;
  position: relative;
  overflow: visible;

  &:hover {
    width: 10vw;
    transition: all 0.5s ease-in-out;
    transition-delay: 0s;
  }

  &:hover .ChatWindow {
    transform: translateX(-70%);
    transition: all 0.5s ease-in-out;
    transition-delay: 0s;
  }

  &:hover .ChatPeople {
    transition: all 0.5s ease-in-out;
    transition-delay: 0s;
  }

  &:hover .friendname {
    left: 48%;
    transition: all 0.5s ease-in-out;
    transition-delay: 0s;
    font-weight: 800;
  }
}

.ChatPeople {
  margin-top: 35px;
  transition: all 0.3s ease-in-out;
  display: flex;
  justify-content: center;
  align-items: center;
  transition-delay: 2s;

  &:before {
    top: 50%;
    transform: translate(-50%, -50%);
    left: -2px;
    position: absolute;
    display: block;
    width: 11px;
    height: 11px;
    border-radius: 50%;
    background: var(--Dark--);
    content: "";
    transition: all 0.3s ease-in-out;
  }
}

.ChatWindow {
  position: relative;
  border-radius: 50%;
  overflow: hidden;
  height: 50px;
  width: 50px;
  transition-delay: 0s;
  transition: all 0.3s ease-in-out;
  border: 2px solid transparent;
  // box-shadow: 0 0 8px var(--HomeBackgrounde-blue-green--);
  // box-shadow: 0 0 10px var(--HomeBackgrounde-blue-green--);
}

.friendname {
  font-size: 15px;
  color: var(--Dark--);
  position: absolute;
  font-weight: 800;
  left: 250px;
  transition-delay: 0s;
  transition: all 0.3s ease-in-out;
}

.popoverDiv {
  color: #fff;
  background-color: #409eff;
}

.logobox {
  width: 100%;
  margin-top: 20px;
  display: flex;
  a {
   @extend .logo;
   text-decoration: none;
  }
  .logo {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    position: relative;
    width: 100%;
    color: var(--Dark--);
    transition: all 0.8s ease-in-out;

    &::after {
      position: absolute;
      top: 63%;
      margin-top: 29px;
      display: flex;
      content: "";
      width: 50%;
      border-bottom: 3px solid var(--Dark--);
      border-radius: 5px;
      transition: all 0.8s ease-in-out;
    }

    &:before {
      top: 0px;
      left: -7px;
      position: absolute;
      display: block;
      width: 10px;
      height: 100%;
      border-radius: 5px;
      background: var(--Dark--);
      content: "";
      transition: all 0.8s ease-in-out;
    }
  }
}
.ChatWindowBox {
  background-color: var(--White--);

  & .el-footer {
    color: var(--ChatWindow--);
    justify-content: center;
    display: flex;
    align-items: center;
    text-align: center;
    position: relative;
    height: 70px !important;
    border-top: solid rgba(128, 128, 128, 0.5) 1px;
  }

  & .el-footer button {
    color: #e1e1e6;
    background-color: #282843;
    border: 1px solid rgba(0, 0, 0, 0);
    border-radius: 50%;
    height: 37px;
    justify-content: center;
    align-items: center;
    width: 40px;
    display: inline-flex;
  }

  & .el-header {
    background-color: var(--White--) !important;
    height: 75px !important;
    border-bottom: solid rgba(128, 128, 128, 0.5) 1px;
  }

  & .el-aside {
    color: var(--White--);
    text-align: center;
  }

  & .el-main {
    background-color: var(--White--);
    text-align: center;
    padding: 0px 0px 3.5% 0px;
  }
}
.topBox {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.inputarea {
  height: 4vh;
  width: 65%;
  resize: none;
  outline: none;
  border: none;
  vertical-align: top;
  font-size: 1.7vw;
  text-align: left;
  background-color: var(--Chatinput--);
  color: var(--White-eee--);
}

.inputer {
  text-align: center;
  width: 90%;
  padding: 10px 0px 10px 0px;
  margin-bottom: 10px;
  margin-top: 20px;
  border-radius: 40px;
  background-color: var(--Chatinput--);
  transition: all 0.3s ease-in-out;
  box-shadow: rgba(0, 0, 0, 0.12) 0px 1px 3px, rgba(0, 0, 0, 0.24) 0px 1px 2px;

  & i {
    color: var(--White-eee--);
    transition: all 0.3s ease-in-out;

    &:hover {
      animation: send 1s infinite alternate;
      transition: all 1s ease-in-out;
    }
  }
}

@keyframes send {
  100% {
    transform: translateY(-7px);
  }
}

.ChatWindowBox .el-footer button {
  background-color: var(--Chatinput--);
}

.ChatUserName {
  margin-left: 5px;
  font-size: 19px;
  font-weight: 800;
  color: var(--Dark--);
}
</style>
<style lang="scss" >
#CHAT {
  font-family: "Gilroy-Medium";
  color: var(--White--);
  .Message {
    font-family: sans-serif;
    padding: 14px;
    margin-top: 3px;
    margin-bottom: 5px;
    white-space: normal;
    word-break: break-all;
    word-wrap: break-word;
    height: auto;
    width: 200px;
    font-size: 12px;
    color: #e1e1e6;
    font-weight: 800;
    text-align: left;
  }

  .rightMessage {
    background-color: #633bbc;
    border-top-right-radius: 8px;
    border-bottom-right-radius: 8px;
    border-bottom-left-radius: 8px;
  }

  .leftMessage {
    background-color: #07847e;
    border-top-right-radius: 8px;
    border-top-left-radius: 8px;
    border-bottom-left-radius: 8px;
  }
  .currentTime {
    color: var(--Gray--);
    font-size: 10px;
    padding-bottom: 5px;
    width: 100%;
  }
  .MessageBox_remoteUser {
    float: left;
    margin-bottom: 20px;
  }
  .MessageBox_nowUser {
    float: right;
    margin-bottom: 20px;
  }
  .MessageBox {
    padding: 20px 0px 0px 0px;
  }
}
.el-popover.monitor-yt-popover {
  background-color: var(--White--);
  border-radius: 40px;
  border: 1px solid rgba(255, 255, 255, 0);
  transition: all 0.3s ease-in-out;

  &:hover {
    box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
      rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
    transition: all 0.3s ease-in-out;
  }
}

.monitor-yt-popover .el-popover__title {
  background-color: var(--White--);
}

.el-popper[x-placement^="right"] .popper__arrow {
  display: none;
}
</style>


