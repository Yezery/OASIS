<template>
  <div class="addMintHomeBox">
    <div class="addMintHomeLeft">
      <div class="addMintHome_backHome">
        <el-button
          icon="el-icon-back"
          circle
          @click="$router.back(-1)"
        />
      </div>
      <div class="addMintHomeMainBox">
        <div class="addMintHome_title">
          <div class="title">
            <div class="logobox">
              <div class="logo animate__animated">
                <div
                  style="display: inline; font-size: 25px"
                  class="animate__animated"
                  ref="logo"
                >
                  OAS
                </div>
                <div
                  style="display: inline;  font-size: 25px"
                  class="animate__animated"
                  ref="logo2"
                >
                  <img
                    src="../../../assets/webAssets/s1.png"
                    alt=""
                    width="22px"
                    height="22px"
                  >S
                </div>
              </div>
            </div>
            Add
          </div>
        </div>
        <div class="addMintHome_router">
          <router-link 
            :to="{ name: 'addmint',query:{nftContract:nftContractAddress} }"
            class="toaddMintNFTBox"
          >
            <div class="toaddMintNFT">
              <div class="innerBox">
                <div class="inner">
                  <div>
                    <svg
                      t="1697879715351"
                      class="icon"
                      viewBox="0 0 1024 1024"
                      version="1.1"
                      xmlns="http://www.w3.org/2000/svg"
                      p-id="28190"
                      width="30"
                      height="30"
                    >
                      <path
                        d="M355.9424 518.8608c-60.2112 0-109.2096-48.9984-109.2096-109.2096 0-60.2112 48.9984-109.2096 109.2096-109.2096s109.2096 48.9984 109.2096 109.2096c0 60.2112-48.9984 109.2096-109.2096 109.2096z m0-157.0304c-26.368 0-47.7696 21.4528-47.7696 47.7696 0 26.368 21.4528 47.7696 47.7696 47.7696s47.7696-21.4528 47.7696-47.7696-21.4528-47.7696-47.7696-47.7696z"
                        fill="#FF9813"
                        p-id="28191"
                      />
                      <path
                        class="icon"
                        d="M804.9664 139.6736H219.0336c-73.4208 0-133.12 59.6992-133.12 133.12v483.072c0 73.4208 59.6992 133.12 133.12 133.12h585.8816c73.4208 0 133.12-59.6992 133.12-133.12V272.7936c0.0512-73.4208-59.6992-133.12-133.0688-133.12z m-585.9328 61.44h585.8816c39.5264 0 71.68 32.1536 71.68 71.68v254.72c-48.4864-40.6016-122.4704-79.4112-192.6656-79.4112-90.7776 0-151.552 50.6368-191.744 84.1728l-7.2704 6.0416c-8.4992 7.0144-17.9712 16.5376-28.9792 27.5456-40.9088 40.96-102.7072 102.8096-178.5856 102.8096-90.1632 0-126.5664-53.248-128.3072-55.9104-0.512-0.8192-1.1264-1.6384-1.6896-2.4064V272.7936c0-39.5264 32.1536-71.68 71.68-71.68z m585.9328 626.432H219.0336c-39.5264 0-71.68-32.1536-71.68-71.68v-62.3104c29.0304 19.0464 71.5776 36.5568 130.048 36.5568 101.376 0 176.896-75.6736 222.0544-120.832 9.8816-9.8816 18.432-18.432 24.6272-23.552l7.4752-6.1952c35.3792-29.4912 83.8656-69.9392 152.4224-69.9392 75.8784 0 161.3824 64.0512 184.064 96.9728 2.3552 3.4304 5.3248 6.2464 8.6016 8.3456v141.0048c0 39.4752-32.1536 71.6288-71.68 71.6288z"
                        fill="#565B6D"
                        p-id="28192"
                      />
                    </svg>Add ordinary NFT
                  </div>
                  <div class="decription">
                    添加一个图片格式的NFT
                  </div>
                </div>
              </div>
              <div class="linkto">
                <i class="el-icon-right" />
              </div>
            </div>
          </router-link>
          <router-link
            :to="{ name: 'addmint3D',query:{nftContract:nftContractAddress} }"
            class="toaddMintNFT3DBox"
          >
            <div class="toaddMintNFT3D">
              <div class="innerBox">
                <div class="inner">
                  <div>
                    <svg
                      t="1697878827237"
                      class="icon"
                      viewBox="0 0 1024 1024"
                      version="1.1"
                      xmlns="http://www.w3.org/2000/svg"
                      p-id="12839"
                      width="30"
                      height="30"
                    >
                      <path
                        class="icon"
                        stroke-width="5"
                        d="M1002.92106 177.943537l-479.943082-175.97913a31.996205 31.996205 0 0 0-22.077382 0l-479.943081 175.97913A31.996205 31.996205 0 0 0 0 208.01997v575.931698a31.996205 31.996205 0 0 0 19.357704 29.436509l481.862854 207.975336a31.996205 31.996205 0 0 0 25.436984 0l478.023309-207.975336A31.996205 31.996205 0 0 0 1023.878574 783.951668V543.980127a31.996205 31.996205 0 0 0-63.99241 0v219.014027L513.85906 957.05114 63.992411 762.834173V230.417314L511.939287 66.116799 959.886164 230.417314v123.665334a31.996205 31.996205 0 0 0 63.99241 0V208.01997a31.996205 31.996205 0 0 0-20.957514-30.076433z"
                        fill="#211B13"
                        p-id="12840"
                      />
                      <path
                        stroke-width="5"
                        d="M511.939287 767.953566a31.996205 31.996205 0 0 0 31.996206-31.996206V399.997203a31.996205 31.996205 0 0 0-20.957515-30.076433l-303.963952-111.986719a31.996205 31.996205 0 0 0-22.077381 59.992885L479.943082 422.234566V735.95736a31.996205 31.996205 0 0 0 31.996205 31.996206z"
                        fill="#FFB243"
                        p-id="12841"
                      />
                    </svg>Add 3D NFT
                  </div>
                  <div class="decription">
                    添加一个3D格式模型的NFT
                  </div>
                </div>
              </div>
              <div class="linkto">
                <i class="el-icon-right" />
              </div>
            </div>
          </router-link>
        </div>
      </div>
    </div>
    <div class="addMintHomeRight">
      <el-carousel
        height="100vh"
        arrow="never"
        indicator-position="none"
        :interval="4000"
      >
        <el-carousel-item
          v-for="item in imageList"
          :key="item.url"
        >
          <img
            :src="item.url"
            alt=""
            style="object-fit: contain;width: 100%;height: 100%;"
          >
        </el-carousel-item>
      </el-carousel>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        imageList: [
          {url:require("@/assets/webAssets/mintShow1.png")
          },
          { url: require("@/assets/webAssets/mintShow2.png") },
          {url:require("@/assets/webAssets/mintShow3.png")}
        ],
        nftContractAddress:null
      };
    },
    mounted() {
      this.nftContractAddress = this.$route.query.nftContract;
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/mintPage/addMintHome.scss";
</style>