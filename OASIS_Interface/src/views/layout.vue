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
              <el-menu-item
                v-if="$store.state.isconnect"
                index="/mintHome"
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
                  Create
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
      scrollOffset: 500
    };
  },
  created() {

  },
  mounted() {
    this.$refs.main.addEventListener('scroll', this.handleScroll);
  },
  beforeUnmount() {
    this.$refs.main.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    handleScroll() {
      const scrollContainer = this.$refs.main;
      const scrollTop = scrollContainer.scrollTop;
      this.isFixed = scrollTop > 40;
    },
    handleSelect(key) {
      this.$router.push(key)
    },
  }
}
</script>
<style lang="scss" scoped>

@import "@/style/layout/layout.scss";
</style>
