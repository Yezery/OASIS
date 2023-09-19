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
              <!-- 钱包连接 -->
              <WalletConnect :wallet-connect="giveChatInitToWalletConnect" />
            </div>
          </el-col>
        </el-row>
        <div class="BalanceEchart">
          <BalanceEchart @echartsChangTheme="echartsChangTheme" />
        </div>
        <div
          class="userInf"
          v-if="$store.state.isconnect"
        >
          <UserInf />
        </div>
        <div
          class="WaitLogin"
          v-if="!$store.state.isconnect"
        >
          <ConnectionTips />
        </div>
      </el-aside>
    </el-container>
  </div>
</template>
<script>
  import ConnectionTips from "../user/ConnectionTips.vue";
  import Theme from "@/components/oasisTheme.vue";
  import BalanceEchart from "@/views/extend/BalanceEchart.vue";
  import UserInf from "@/views/user/UserInf.vue";
  import WalletConnect from "@/views/user/WalletConnect.vue";
  import ChatMemu from "@/views/chat/ChatMemu.vue";
  export default {
    name: "MarketShopIndex",
    components: {
      Theme,
      BalanceEchart,
      UserInf,
      WalletConnect,
      ChatMemu,
      ConnectionTips
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
      };
    },
    created() {},
    mounted() {},
    // provide() {
    //     //依赖注入
    //     return {
    //         init: this.giveChatInitToWalletConnect,
    //     };
    // },
    methods: {
      giveChatInitToWalletConnect() {
        console.log("聊天栏正初始化.......");
        this.$refs.ChatMemu.init();
      },
      echartsChangTheme(method) {
        this.eChangTheme = method;
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
  