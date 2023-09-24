<template>
  <div class="oasisTheme">
    <el-radio-group
      v-model="radio"
      @change="ChangeTheme"
    >
      <el-radio label="1">
        <i class="el-icon-sunny" />
      </el-radio>
      <el-radio label="2">
        <i class="el-icon-moon" />
      </el-radio>
    </el-radio-group>
  </div>
</template>

<script>
  export default {
    name: "OasisTheme",
    props: {
      eChangTheme: {
        type: Function,
        default:null
      },
    },
    data() {
      return {
        isDarkMode: true,
        radio: "1",
        // 小狐狸钱包连接情况
        isonload: this.$store.state.isEmpower,
      };
    },
    mounted() {
      this.isDayLight();
      this.ChangeTheme();
    },
    computed: {
      isOnload() {
        return this.$store.state.isEmpower;
      },
    },
    watch: {
      isOnload: {
        deep: true,
        handler(newData) {
          this.isonload = newData;
          document.documentElement.style.setProperty(
            "--border-green--",
            this.isonload ? "#55c960" : "#ff0000"
          );
        },
      },
    },
    methods: {
      ChangeTheme() {
        this.isDarkMode = !this.isDarkMode;
        this.$store.commit("setIsDark", this.isDarkMode);
        setTimeout(() => {
          this.eChangTheme();
        }, 100);
        // 通用黑
        document.documentElement.style.setProperty(
          "--Dark--",
          this.isDarkMode ? "#161621" : "#fff"
        );
        // 通用灰
        document.documentElement.style.setProperty(
          "--Gray--",
          this.isDarkMode ? "#808080" : "#fff"
        );

        // 通用白
        document.documentElement.style.setProperty(
          "--White--",
          this.isDarkMode ? "#fff" : "#161621"
        );

        // 通用米白
        document.documentElement.style.setProperty(
          "--White-eee--",
          this.isDarkMode ? "#eee" : "#1e1e2c"
        );

        // 通用阴影
        document.documentElement.style.setProperty(
          "--boxshdow-style--",
          this.radio == "1"
            ? "rgba(6, 24, 44, 0.4) 0px 0px 0px 2px, rgba(6, 24, 44, 0.65) 0px 4px 6px -1px, rgba(255, 255, 255, 0.08) 0px 1px 0px inset;"
            : "rgba(99, 99, 99, 0.2) 0px 2px 8px 0px;"
        );

        // 通用 绿蓝
        document.documentElement.style.setProperty(
          "--border-green--",
          this.isonload ? "#55c960" : "#ff0000"
        );

        // 通用蓝黑 白
        document.documentElement.style.setProperty(
          "--blueblack-white--",
          this.isDarkMode ? "#11243d" : "#fff"
        );

        // 主题切换按钮颜色
        document.documentElement.style.setProperty(
          "--ThemeColorChange--",
          this.radio == "1" ? "#F2CA48" : "#409EFF"
        );

        // 导航栏颜色
        document.documentElement.style.setProperty(
          "--nav-color--",
          this.isDarkMode ? "#161621" : "#fff"
        );

        // 头像边框颜色
        document.documentElement.style.setProperty(
          "--avatar-border-style--",
          this.radio == "1" ? "#fff" : "#409EFF"
        );

        //*************************
        // 页面选择按钮
        // 1
        document.documentElement.style.setProperty(
          "--Selected-button1-blue-black",
          this.isDarkMode ? "#000000" : "#ffffff80"
        );
        document.documentElement.style.setProperty(
          "re--Selected-button1-blue-black",
          this.isDarkMode ? "#ffffff80" : "#000000"
        );
        // 2
        document.documentElement.style.setProperty(
          "--Selected-button2-blue-black",
          this.isDarkMode ? "#000000" : "#409eff"
        );
        // 3
        document.documentElement.style.setProperty(
          "--Selected-button3-blue-black",
          this.isDarkMode ? "#000000" : "#409eff"
        );
        // 4
        document.documentElement.style.setProperty(
          "--Selected-button4-blue-black",
          this.isDarkMode ? "#000000" : "#409eff"
        );
        // 5
        document.documentElement.style.setProperty(
          "--Selected-button5-blue-black",
          this.isDarkMode ? "#000000" : "#409eff"
        );
        //*************************

        //*************************
        // 聊天框主题样式
        document.documentElement.style.setProperty(
          "--HomeBackgrounde-blue-green--",
          this.radio == "1" ? "#7fffd4" : "#409eff"
        );
        document.documentElement.style.setProperty(
          "--ChatWindow--",
          this.radio == "1" ? "#1a1924" : "#fff"
        );
        document.documentElement.style.setProperty(
          "--Chatinput--",
          this.radio == "1" ? "#282843" : "#fff"
        );
        //*************************
      },

      //判断夜间还是白天
      isDayLight() {
        var currdate = new Date();
        if (currdate.getHours() >= 18 || currdate.getHours() < 6) {
          this.isDarkMode = true;
          this.radio = "2";
        } else {
          this.isDarkMode = false;
          this.radio = "1";
        }
      },
    },
  };
</script>

<style scoped>
.dark-mode {
  /* 设置黑夜模式下的 body 背景色和文本颜色 */
  background-color: var(--White--);
  color: var(--Dark--);
}
</style>
<style lang="scss" scoped>
.oasisTheme {
  /deep/ .el-radio__input {
    display: none;
  }
  /deep/ .el-radio__label {
    height: 100%;
    text-align: center;
    padding: 0;
    width: auto;
    i {
      font-size: 30px;
    }
  }
  /deep/ .el-radio__input.is-checked + .el-radio__label {
    color: var(--ThemeColorChange--);
  }
}
</style>
