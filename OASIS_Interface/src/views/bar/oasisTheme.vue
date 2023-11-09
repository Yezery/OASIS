<template>
  <div class="oasisTheme">
    <div class="oasisThemeInner">
      <el-radio-group
        v-model="radio"
        @change="changeTheme"
      >
        <el-radio label="1">
          <i class="el-icon-sunny" />
        </el-radio>
        <el-radio label="2">
          <i class="el-icon-moon" />
        </el-radio>
      </el-radio-group>
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      setEchartChange: {
        type: Function,
        default:undefined
      },
    },
    data() {
      return {
        Theme: true,
        radio: "1",
        isonload: this.$store.state.isEmpower,
      };
    },
    mounted() {
      this.isDayLight();
      this.changeTheme()
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
      changeTheme() {
        this.Theme = !this.Theme;
        this.$store.commit("setTheme", this.Theme);
        // 通用黑
        document.documentElement.style.setProperty(
          "--Dark--",
          this.Theme ?  "#020011" : "#fff"
        );
        // #161621
        // 通用灰
        document.documentElement.style.setProperty(
          "--Gray--",
          this.Theme ?   "#fff" : "#808080"
        );
        // 通用白
        document.documentElement.style.setProperty(
          "--White--",
          this.Theme ? "#fff" : "#020011"
        );
        // #161621
        // 通用米白
        document.documentElement.style.setProperty(
          "--White-eee--",
          this.Theme ? "#eee" : "#1e1e2c"
        );
        // market table bg hover
        document.documentElement.style.setProperty(
          "--tbbg--",
          this.Theme ?   "rgba(245, 247, 250, 0.6)" : "#0b0f19"
        );
        // 通用HF主题色
        document.documentElement.style.setProperty(
          "--HFTheme--",
          this.Theme ?   "#fff" : "#0b0f19"
        );
        
        // #161621
        // 分割线
        document.documentElement.style.setProperty(
          "--FG--",
          this.Theme ?   "#fff" : "rgb(20 ,28, 46)"
        );
        document.documentElement.style.setProperty(
          "--HFBorder--",
          this.Theme ?   "#fff" : "rgb(20 ,28, 46)"
        );
        // GPT颜色
        document.documentElement.style.setProperty(
          "--GPTColor--",
          this.Theme ?   "#13227a" : "#fff"
        );
        
        // mintHome
        document.documentElement.style.setProperty(
          "--mintHomebg--",
          this.Theme ?   "#fff" : "rgb(18, 18, 18)"
        );
        document.documentElement.style.setProperty(
          "--mintHomeLink--",
          this.Theme ?   "#fff" : "rgb(32, 32, 32)"
        );
        document.documentElement.style.setProperty(
          "--mintHomeBorder--",
          this.Theme ?    "rgb(18, 18, 18)" : "transparent"
        );
        document.documentElement.style.setProperty(
          "--mintHomeHover--",
          this.Theme ?   "#fff" : "rgb(37, 37, 37)"
        );
        // document.documentElement.style.setProperty(
        //   "--HFTheme--",
        //   this.Theme ? "#020011" : "#fff"
        // );
        
        // 通用阴影
        document.documentElement.style.setProperty(
          "--boxshdow-style--",
          this.Theme == "1"
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
          this.Theme ? "#11243d" : "#fff"
        );

        // 主题切换按钮颜色
        document.documentElement.style.setProperty(
          "--ThemeColorChange--",
          this.radio == "1" ? "#F2CA48" : "#409EFF"
        );

        // 导航栏颜色
        document.documentElement.style.setProperty(
          "--nav-color--",
          this.Theme ? "#161621" : "#fff"
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
          this.Theme ? "#000000" : "#ffffff80"
        );
        document.documentElement.style.setProperty(
          "re--Selected-button1-blue-black",
          this.Theme ? "#ffffff80" : "#000000"
        );
        // 2
        document.documentElement.style.setProperty(
          "--Selected-button2-blue-black",
          this.Theme ? "#000000" : "#409eff"
        );
        // 3
        document.documentElement.style.setProperty(
          "--Selected-button3-blue-black",
          this.Theme ? "#000000" : "#409eff"
        );
        // 4
        document.documentElement.style.setProperty(
          "--Selected-button4-blue-black",
          this.Theme ? "#000000" : "#409eff"
        );
        // 5
        document.documentElement.style.setProperty(
          "--Selected-button5-blue-black",
          this.Theme ? "#000000" : "#409eff"
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
          this.Theme = true;
          this.radio = "2";
        } else {
          this.Theme = false;
          this.radio = "1";
        }
      },
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/topMenu/oasisTheme.scss";
</style>
