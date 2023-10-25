<template>
  <div
    class="NFTInfMainBox animate__animated animate__fadeInRight"
    v-loading.fullscreen.lock="isChanging"
    element-loading-text="交易进行中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
  <div class="NFTInfBg_backHome">
          <el-button icon="el-icon-back" circle @click="$router.back(-1)" />
        </div>
    <el-carousel
      indicator-position="none"
      arrow="never"
      :height="bgHeight"
    >
   
      <el-carousel-item
        v-for="image in onSaleNFTList"
        :key="image.ipfsPath"
      >
        <div
          class="NFTInfMain_Bg"
          :style="{ backgroundImage: `url('${encodeURI(image.ipfsPath)}')`}"
        />
      </el-carousel-item>
    </el-carousel>
    
    <div class="NFTInfMain_Avatar">
      <el-carousel
        class="NFTInfMain_AvatarBorder"
        indicator-position="none"
        arrow="never"
        :height="imgHeight"
      >
        <el-carousel-item
          v-for="image in onSaleNFTList"
          :key="image.ipfsPath"
        >
          <img
            style="width: 100%; object-fit: contain;"
            :src="image.ipfsPath"
          >
        </el-carousel-item>
      </el-carousel>
    </div>
    <div class="NFTInfMain_Top">
      <div class="NFTInfMain_Top_Inf">
        <div class="NFTName">
          {{ nftName }}
        </div>
        <div class="NFTContract">
          <div class="NFTSupplyer">
            由 <span style="font-weight: 800;font-size:16px;">{{ `${this.seller.slice(
              0,
              5
            )}...${this.seller.slice(37)}` }}</span>
          </div>
          <div class="NFTInf">
            <div class="NFTCounts">
              系列<span style="font-weight: 800;">{{ seriesName }}</span> ·
              代币符号<span style="font-weight: 800;">{{ symbol }}</span> ·
              限量铸造<span style="font-weight: 800;">{{ maximums }}</span> ·
              目前已铸造数量<span style="font-weight: 800;">{{ currentId }}</span>
            </div>
            <div class="NFTdescription">
              {{ description }}
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="NFTInfMain_Bootom">
      <div class="selectBox">
        <div class="selectBox-Button">
          <el-tabs v-model="activeTab">
            <el-tab-pane
              label="正在售卖"
              name="first"
            >
              <div class="SeriesNFT">
                <div
                  style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 15px;"
                  v-if="onSaleNFTList.filter(inf => inf.isActive).length == 0"
                >
                  <el-empty description="无在售" />
                </div>
                <template v-for="inf in onSaleNFTList.filter(inf => inf.isActive)">
                  <div
                    class="NFTInf"
                    :key="inf.image"
                  >
                    <div class="imageBox">
                      <img
                        class="NFTImage"
                        :src="inf.ipfsPath"
                        alt=""
                      >
                    </div>
                    <div class="Inf">
                      <div class="NFTName">
                        {{ inf.nftName }}
                      </div>
                      <div class="ownerAndToSell">
                        <div class="ToSellBox">
                          #{{ inf.tokenId }}
                        </div>
                        <div style="margin-left: 2%;">
                          <span style="font-weight: 800;font-size: 25px;">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }} </span> ETH
                        </div>
                      </div>
                    </div>
                    <div
                      class="SellBox"
                      @click="Buy(inf)"
                      v-if="$store.state.isconnect && seller.toUpperCase() != $store.state.currentAddress.toUpperCase()"
                    >
                      Buy now
                    </div>
                  </div>
                </template>
              </div>
            </el-tab-pane>
            <el-tab-pane
              label="暂未发售"
              name="second"
            >
              <div class="SeriesNFT">
                <div
                  style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 15px;"
                  v-if="seriesNFTList.filter(inf => !inf.isActive).length == 0"
                >
                  <el-empty description="无数据" />
                </div>
                <template v-else>
                  <div
                    class="NFTInf"
                    v-for="inf in seriesNFTList.filter(inf => !inf.isActive)"
                    :key="inf.image"
                  >
                    <div class="imageBox">
                      <img
                        class="NFTImage"
                        :src="inf.image"
                        alt=""
                      >
                    </div>
                    <div class="Inf">
                      <div class="NFTName">
                        {{ inf.name }}
                      </div>
                      <div class="ownerAndToSell">
                        <div class="ToSellBox" />
                      </div>
                    </div>
                  </div>
                </template>
              </div>
            </el-tab-pane>
            <el-tab-pane
              label="系列NFT"
              name="third"
            >
              <div class="SeriesNFT">
                <div
                  style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 15px;"
                  v-if="seriesNFTList.length == 0"
                >
                  <el-empty description="无数据" />
                </div>
                <div
                  class="NFTInf"
                  v-for="inf in seriesNFTList"
                  :key="inf.image"
                >
                  <div class="imageBox">
                    <img
                      class="NFTImage"
                      :src="inf.ipfsPath"
                      alt=""
                    >
                  </div>
                  <div class="Inf">
                    <div class="NFTName">
                      {{ inf.nftName }}
                    </div>
                    <div class="ownerAndToSell">
                      <div class="ToSellBox">
                        #{{ inf.tokenId }}
                      </div>
                      <div style="margin-left: 2%;">
                        <span style="font-weight: 800;font-size:25px;">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }} </span> ETH
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { getOnSaleNFTByNFTAddress } from "@/api/axios/Sale";
  import {getSeriesByNFTAddress}from "@/api/axios/ownerContractLIst"
import { getNFTStruct,Buy } from "@/api/axios/contract";

  export default {
    data() {
      return {
        // ======= vuex
        saleId: null,
        image: null,
        nftName: "",
        description: null,
        nftAddress: null,
        tokenId: null,
        isActive: null,
        seller: "",
        price: null,
        // ======== 合约
        maximums: 0,
        symbol: null,
        currentId: 0,
        seriesName: "",
        nftContract:null,
        // ======== 配置
        imgHeight: "",
        bgHeight: "",
        bought: false,
        isChanging: false,
        activeTab:"first",
        // ======== 服务器
        onSaleNFTList: [],
        seriesNFTList:[],
      };
    },
    watch: {},
    mounted() {
      this.init();
      // 监听窗口变化，使得轮播图高度自适应图片高度
      window.addEventListener("resize", () => {
        this.imgHeight = String(window.innerWidth / 12);
        this.bgHeight = String(window.innerHeight / 80);
      });
    },
  methods: {
      async getOnSaleNFT() {
        var NFTDto = {
          nftAddress: this.nftAddress,
        };
        await getOnSaleNFTByNFTAddress(NFTDto).then((re) => {
          this.onSaleNFTList = re.data.data;
        });
      },
      async getSeriesByNFTAddress() {
        var NFTDto = {
          nftAddress: this.nftAddress,
        };
        await getSeriesByNFTAddress(NFTDto).then((re) => {
          this.seriesNFTList = re.data.data;
        });
      },
      async init() {
        try {
          this.NFTInf = this.$store.state.NFTInf;
          console.log(this.NFTInf);
          this.saleId = this.NFTInf.saleId
          this.image = this.NFTInf.image
          this.nftName = this.NFTInf.nftName
          this.description = this.NFTInf.description
          this.nftAddress = this.NFTInf.nftAddress
          this.tokenId = this.NFTInf.tokenId
          this.isActive = this.NFTInf.isActive
          this.seller = this.NFTInf.seller
          this.price = this.NFTInf.price

          await this.getSeriesByNFTAddress();
          await this.getOnSaleNFT() 
         
          await getNFTStruct(this.nftAddress).then((re) => {
            this.nftContract = re;
          });
          await this.nftContract.methods
            ._currentId()
            .call()
            .then((re) => {
              this.currentId = re;
            });
          await this.nftContract.methods
            .symbol()
            .call()
            .then((re) => {
              this.symbol = re;
            });
          await this.nftContract.methods
            .name()
            .call()
            .then((re) => {
              this.seriesName = re;
            });
          await this.nftContract.methods
            ._maximums()
            .call()
            .then((re) => {
              this.maximums = re;
            });
        } catch (error) {
          console.log(error);
          this.$router.push("/")
        }
      },
      async Buy(NFT) {
        try {
          // NFT.symbol = this.symbol;
          // NFT.image = this.NFTImage;
          let isSuccess = false;
          this.isChanging = true;
          await Buy(NFT).then((re) => {
            isSuccess = re;
          });
          if (isSuccess) {
            this.isChanging = false;
            this.$notify({
              title: "💖 感谢购买",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            this.bought = true;
            this.changeNFT={}
          } else {
            this.isChanging = false;
            this.changeNFT={}
            this.$notify({
              title: "您已取消购买",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
     
          }
        } catch (error) {
          this.isChanging = false;
          this.$notify.error({
            title: "购买失败",
            position: "top-left",
            offset: 200,
          });
        }
      },
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/components/NFTInf.scss";
</style>