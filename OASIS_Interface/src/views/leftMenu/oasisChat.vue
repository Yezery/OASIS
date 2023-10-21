<template>
  <div class="Menu" ref="Menu" :class="{'isEnter':isEnter}">
    <!-- @mouseleave="OpenAndClose" -->
    <div class="logobox">
      <div class="logo animate__animated">
        <router-link :to="{ name: 'MarketShop' }">
          <div style="display: inline; font-size: 25px" class="animate__animated" ref="logo">
            OAS
          </div>
          <div style="display: inline;  font-size: 25px" class="animate__animated" ref="logo2">
            <img src="../../assets/webAssets/s1.png" alt="" width="22px" height="22px">S
          </div>
        </router-link>
      </div>
    </div>
    <div style="text-align: right;padding-top: 10px;padding-right: 10px;cursor: pointer;color: var(--Dark--);" v-if="isChatGPT" @click="OpenAndClose"><i class="el-icon-d-arrow-left"></i></div>
    
    <transition enter-active-class="animate__animated animate__fadeInRight" leave-active-class="animate__animated animate__fadeOutRight">
      <div class="chatGPT">
        <div class="ChatAvatar animate__animated animate__fadeInRight" @click="isShowGPTChat(1)" >
         
          <div class="ChatAvatarBox" ref="ChatWindow" >
            <el-badge is-dot :hidden="!haveMessage" class="item">
            <svg  t="1697856532693" class="GPTIcon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4653" width="47" height="47">
              <path class="icon-body" :class="{'replying':isFinish}" stroke-width="2" d="M1024 512A512 512 0 1 1 0 512a512 512 0 0 1 1024 0z m-128 0q0-73.216-26.989714-141.385143-29.257143-73.874286-85.504-130.194286-56.246857-56.173714-130.194286-85.430857Q585.289143 128 512 128q-73.216 0-141.385143 26.989714-73.874286 29.257143-130.194286 85.504-56.173714 56.246857-85.430857 130.194286Q128 438.637714 128 512q0 73.216 26.989714 141.312 29.257143 73.947429 85.504 130.194286 56.246857 56.32 130.194286 85.577143 68.022857 26.916571 141.312 26.916571 73.216 0 141.312-26.989714 73.947429-29.257143 130.194286-85.504 56.32-56.246857 85.577143-130.194286 26.916571-68.022857 26.916571-141.312z" p-id="4654" fill="#13227a"></path>
              <path class="icon-eye" stroke-width="2" d="M292.571429 365.714286m73.142857 0l0 0q73.142857 0 73.142857 73.142857l0 146.285714q0 73.142857-73.142857 73.142857l0 0q-73.142857 0-73.142857-73.142857l0-146.285714q0-73.142857 73.142857-73.142857Z" p-id="4655"></path>
              <path class="icon-eye" stroke-width="2" d="M585.142857 365.714286m73.142857 0l0 0q73.142857 0 73.142857 73.142857l0 146.285714q0 73.142857-73.142857 73.142857l0 0q-73.142857 0-73.142857-73.142857l0-146.285714q0-73.142857 73.142857-73.142857Z" p-id="4656"></path>
            </svg>
          </el-badge>
          </div>

        </div>
        <transition enter-active-class="animate__animated animate__fadeInRight" leave-active-class="animate__animated animate__fadeOutRight">
          <div class="ChatWindowBox" v-show="isChatGPT">
            <el-container>
              <el-main>

                <div id="CHAT" v-html="contents[chatUser]"></div>
                <div v-if="isFinish" class="animate__animated animate__fadeInLeft" style="text-align: right;width:100%">
                  <i class="el-icon-loading " style="margin-right:15px"/>
                  <i class="el-icon-video-pause" @click="stopGPT" style="cursor: pointer;"/>
                </div>
              </el-main>
              <el-footer>
                <div class="inputer">
                  <el-input class="input" id="scroll_text" @click="haveMessage = false;"  resize="none" type="textarea" :disabled="isFinish" v-model="text" placeholder="问我任何问题 (Shift + Enter = 换行)" @keydown.enter.native.prevent="handleKeyCode($event,2)">
                  </el-input>
                </div>
              </el-footer>
            </el-container>
          </div>
        </transition>
      </div>
    </transition>
    <!-- <template> -->
    <!-- <div
        v-for="otherUser in users"
        :key="otherUser.username"
        class="chat"
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
              style="width: 329px; height: 520px"
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
                        class="el-icon-arrow-up"
                        style="font-size: 30px;cursor: pointer;"
                      />
                    </button>
                  </div>
                </el-footer>
              </el-container>
            </div>
          </div>
        </el-popover>
      </div> -->

    <!-- <el-popover
          placement="right"
          trigger="click"
          popper-class="chatGPTep"
          :visible-arrow="false"
          offset="200"
        /> -->
    <!-- </template> -->
  </div>
</template>

<script>
let socket;
import axios from 'axios'
  import { sendToGPT } from "@/api/axios/gpt";
  import { newWebSocket } from "@/utils/webStocket";
  import { parseTime } from "@/utils/Time";
  export default {
    data() {
      return {
        user: {},
        isCollapse: false,
        users: [],
        messages: [],
        chatUser: "",
        text: "",
        contents: {},
        showName: false,
        isEnter: false,

        isFinish: false,
        haveMessage:false,
        isChatGPT: false,
        source:null
      };
    },
  components: {},
    mounted() {
      this.chatUser = "GPT";
      this.createContent(
        "GPT",
        null,
        "你好！我是OASIS GPT，可以向我咨询有关这个系统的信息。"
      );
    },
    methods: {
      // 更新弹窗位置
      upDatePosition() {
        this.$nextTick(() => {
          this.$refs.popoverRef.updatePopper();
        });
      },
      isShowGPTChat(opt) {
        if (opt == 1) {
          this.chatUser = "GPT";
          this.isChatGPT = true;
          this.haveMessage = false;
          this.OpenAndClose();
        } else {
          this.chatUser = "";
          this.isChatGPT = false;
          this.OpenAndClose();
        }
      },
      stopGPT() {
        this.isFinish = false;
        this.source.cancel('用户取消')
        this.source = null
      },
      OpenAndClose() {
        this.isOpen = !this.isOpen;
        if (this.isOpen) {
          this.$refs.logo.classList.add("animate__swing");
          this.$refs.logo2.classList.add("animate__swing");
          this.$refs.Menu.style = "width:400px;";

          this.showName = true;
          this.isEnter = true;
        } else {
          this.$refs.logo.classList.remove("animate__swing");
          this.$refs.logo2.classList.remove("animate__swing");
          this.$refs.Menu.style = "width:65px;";
          this.$refs.ChatWindow.style = "transform:0;";
          this.isChatGPT = false;
          this.showName = false;
          this.isEnter = false;
        }
      },
      GETHashAvatar(UserAddress) {
        return (
          "data:image/png;base64," +
          new this.Identicon(UserAddress, 120).toString()
        );
      },
      send(opt) {
        if (!this.text) {
          this.$notify({
            title: "输入不能为空",
            type: "warning",
            position: "top-left",
          });
        } else {
          if (opt == 1) {
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
              this.createContent(
                null,
                this.$store.state.currentAddress,
                this.text
              );
              this.text = "";
            }
          } else {
            this.sendToGPT();
          }
        }
      },
      async sendToGPT() {
        const text = this.text;
        var data = {
          role: "user",
          content: text,
        };
        this.text = "";
        this.createContent(null, this.$store.state.currentAddress, text);
        try {
          this.isFinish = true;
          const source = axios.CancelToken.source()
            this.source = source
          await sendToGPT(data, source.token).then((re) => {
            let result = JSON.parse(re.data.data);
            this.createContent(
              result.choices[0].message.role,
              null,
              result.choices[0].message.content
            );
          });
          this.haveMessage= true
          this.isFinish = false;
        } catch (error) {
          console.log(error);
          this.isFinish = false;
        }
      },
      // 键盘回车事件
      handleKeyCode(event, opt) {
        if (event.keyCode == 13 && event.shiftKey) {
          this.$nextTick(() => {
            setTimeout(() => {
              const textarea = document.getElementById("scroll_text");
              textarea.scrollTop = textarea.scrollHeight;
            }, 13);
          });
          this.text = this.text + "\n";
        } else {
          if (event.keyCode == 13) {
            this.send(opt);
          }
        }
      },
      createContent(remoteUser, nowUser, text) {
        let html;
        let content = this.contents[this.chatUser || remoteUser] || "";
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
      typeText(text) {
        const textElement = document.getElementById("text");
        let i = 0;
        const typingInterval = setInterval(function () {
          textElement.textContent += text[i];
          i++;
          if (i === text.length) {
            clearInterval(typingInterval);
          }
        }, 50); // 控制打字速度
      },
      // createContent(remoteUser, nowUser, text) {
      //   let content;
      //   let html;
      //   content = this.contents[this.chatUser || remoteUser] || "";
      //   // 当前用户消息
      //   if (nowUser) {
      //     html = `<div class="MessageBox_nowUser"><span class="currentTime">${parseTime(
      //       Date(),
      //       "hh:mm:ss"
      //     )}</span><div class="Message leftMessage">${text}</div></div>`;
      //   } else if (remoteUser) {
      //     html = `
      //    <div class="MessageBox_remoteUser">

      //                                       <span class="currentTime">${parseTime(
      //                                         Date(),
      //                                         "hh:mm:ss"
      //                                       )}</span>
      //                                       <div class="Message rightMessage">
      //                                         ${text}
      //                                       </div>
      //                                     </div>
      //                                   `;
      //   }
      //   content += html;
      //   this.$set(this.contents, this.chatUser, content);
      //   this.$nextTick(() => {
      //     let msg = document.getElementById("CHAT");
      //     msg.scrollTop = msg.scrollHeight;
      //   });
      // },
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
    },
  };
</script>
<style lang="scss" scoped>
@import "@/style/leftMenu/oasisChat.scss";
</style>
<!-- <style lang="scss" >
.el-popover.chatGPTep{
  background-color: var(--White--);
  border-radius: 30px;
  border: $globalBorder;
  transition: all 0.3s ease-in-out;
  
  &:hover {
    box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
      rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
    transition: all 0.3s ease-in-out;
  }
  .el-popover__title{
    background-color: var(--White--);
  }
}
</style> -->
