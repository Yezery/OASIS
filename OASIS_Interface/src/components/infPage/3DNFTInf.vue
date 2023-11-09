<template>
  <div
    class="NFTInfMainBox animate__animated animate__fadeInDown"
    v-loading.fullscreen.lock="isChanging"
    element-loading-text="交易进行中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <div
      class="NFTInfMain"
      ref="NFTInfMain"
    >
      <div class="NFTInfMainBg">
        <div class="NFTInfBg_backHome">
          <el-button
            icon="el-icon-back"
            circle
            @click="$router.push('/home/MarketShop3D')"
          />
        </div>
        <div class="NFTName">
          <h1>{{ nftName }}</h1>
        </div>

        <div
          class="show3DWindow"
          ref="show3DWindow"
          :class="{'active': showFloatingWindow}"
        >
          <div class="show3DBox">
            <Model
              :model-path="modelPath"
              @initModel="seeModel"
            />
          </div>
        </div>

        <div class="Inf">
          <div class="NFTContract">
            <div style="text-align: left;margin-bottom: 3%;">
              <el-button
                type="success"
                plain
                class="Buy"
                @click="Buy()"
                :disabled="!$store.state.isconnect || seller.toUpperCase() == $store.state.currentAddress.toUpperCase() || bought || !isActive"
              >
                <i class="el-icon-shopping-cart-1" /> {{ bought?"已购入":"购入" }}
              </el-button>
            </div>
            <div class="NFTSupplyer">
              由 <span style="font-weight: 800;font-size: 1.2vw;">{{ `${this.seller.slice(
                0,
                5
              )}...${this.seller.slice(37)}` }}</span>
            </div>
            <div class="NFTInf">
              <div class="NFTCounts">
                系列<span style="font-weight: 800;">{{ seriesName }}</span> ·
                代币符号<span style="font-weight: 800;">{{ symbol }}</span> ·
                限量创造<span style="font-weight: 800;">{{ maximums }}</span> ·
                目前已创造数量<span style="font-weight: 800;">{{ currentId }}</span>
              </div>
            </div>
          </div>
        </div>

        <div class="selectBox">
          <div class="selectBox-Button">
            <el-tabs v-model="activeTab">
              <el-tab-pane
                label="正在售卖"
                name="first"
              >
                <div class="SeriesNFT">
                  <div
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 25px;"
                    v-if="onSaleNFTList.length == 0"
                  >
                    <el-empty description="无在售" />
                  </div>
                  <template v-else>
                    <template v-for="inf in onSaleNFTList">
                      <div
                        class="NFTInf3D"
                        :key="inf.image"
                      >
                        <div class="Inf3D">
                          <div class="Inf3DLeft">
                            <div class="Inf3DLeftTop">
                              <el-button
                                @click="setNewModelPath(inf.ipfsPath)"
                                type="primary"
                                icon="el-icon-video-play"
                                circle
                              />
                            </div>
                            <div class="Inf3DLeftBottom">
                              <div class="TokenID">
                                <span style="font-size: 20px;">#{{ inf.tokenId }}</span>
                              </div>
                            </div>
                          </div>

                          <div class="Inf3DRight">
                            <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 25px;">
                              {{ inf.nftName }}
                            </div>
                            <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                              <div
                                v-if="inf.isActive"
                                class="priceBox"
                              >
                                <span class="price">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }}</span> ETH
                              </div>
                              <div
                                v-else
                                class="priceBox"
                              >
                                <span class="price" />
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </template>
                  </template>
                </div>
              </el-tab-pane>
              <el-tab-pane
                label="暂未发售"
                name="second"
              >
                <div class="SeriesNFT">
                  <div
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 25px;"
                    v-if="seriesNFTList.filter(inf => !inf.isActive).length == 0"
                  >
                    <el-empty description="无数据" />
                  </div>
                  <template v-else>
                    <template v-for="inf in seriesNFTList.filter(inf => !inf.isActive)">
                      <div
                        class="NFTInf3D"
                        :key="inf.image"
                      >
                        <div class="Inf3D">
                          <div class="Inf3DLeft">
                            <div class="Inf3DLeftTop">
                              <el-button
                                @click="setNewModelPath(inf.ipfsPath)"
                                type="primary"
                                icon="el-icon-video-play"
                                circle
                              />
                            </div>
                            <div class="Inf3DLeftBottom">
                              <div class="TokenID">
                                <span style="font-size: 20px;">#{{ inf.tokenId }}</span>
                              </div>
                            </div>
                          </div>

                          <div class="Inf3DRight">
                            <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 25px;">
                              {{ inf.nftName }}
                            </div>
                            <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                              <div
                                v-if="inf.isActive"
                                class="priceBox"
                              >
                                <span class="price">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }}</span> ETH
                              </div>
                              <div
                                v-else
                                class="priceBox"
                              >
                                <span class="price" />
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </template>
                  </template>
                </div>
              </el-tab-pane>
              <el-tab-pane
                label="系列NFT"
                name="third"
              >
                <div class="SeriesNFT">
                  <div
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 25px;"
                    v-if="seriesNFTList.length == 0"
                  >
                    <el-empty description="无数据" />
                  </div>
                  <template v-else>
                    <template v-for="inf in seriesNFTList">
                      <div
                        class="NFTInf3D"
                        :key="inf.image"
                      >
                        <div class="Inf3D">
                          <div class="Inf3DLeft">
                            <div class="Inf3DLeftTop">
                              <el-button
                                @click="setNewModelPath(inf.ipfsPath)"
                                type="primary"
                                icon="el-icon-video-play"
                                circle
                              />
                            </div>
                            <div class="Inf3DLeftBottom">
                              <div class="TokenID">
                                <span style="font-size: 20px;">#{{ inf.tokenId }}</span>
                              </div>
                            </div>
                          </div>
                          <div class="Inf3DRight">
                            <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size:  25px;">
                              {{ inf.nftName }}
                            </div>
                            <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                              <div
                                v-if="inf.isActive"
                                class="priceBox"
                              >
                                <span class="price">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }}</span> ETH
                              </div>
                              <div
                                v-else
                                class="priceBox"
                              >
                                <span class="price" />
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </template>
                  </template>
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Model from "@/components/3DModelShow/3DModel.vue";
  import { getNFTStruct, Buy } from "@/api/axios/contract";
  import { getOnSaleNFTByNFTAddress } from "@/api/axios/Sale";
  import { getSeriesByNFTAddress } from "@/api/axios/ownerContractLIst";
  // getNFTStruct,
  export default {
    name: "NFTInf",
    components: {
      Model,
    },
    data() {
      return {
        // ======= vuex
        saleId: null,
        modelPath: null,
        nftName: null,
        nftAddress: null,
        tokenId: null,
        isActive: null,
        seller: "",
        price: null,
        // ======== 合约
        maximums: 0,
        symbol: null,
        currentId: 0,
        seriesName: null,
        nftContract: null,
        // ======== 配置
        imgHeight: "",
        bgHeight: "",
        bought: false,
        isChanging: false,
        activeTab: "first",
        initModel: null, //方法
        isInitModel: false,
        showFloatingWindow: false,
        // ======== 服务器
        onSaleNFTList: [],
        seriesNFTList: [],
      };
    },
    watch: {},
    mounted() {
      this.init();
    },

    methods: {
      seeModel(data) {
        this.initModel = data;
      },
      setNewModelPath(path) {
        this.isInitModel = true;
        this.modelPath = path;
        setTimeout(() => {
          this.initModel();
        }, 1000);
      },
      async getOnSaleNFT() {
        await getOnSaleNFTByNFTAddress({
          nftAddress: this.nftAddress,
        }).then((re) => {
          this.onSaleNFTList = re.data.data;
        });
      },
      async getSeriesByNFTAddress() {
        await getSeriesByNFTAddress({
          nftAddress: this.nftAddress,
        }).then((re) => {
          this.seriesNFTList = re.data.data;
        });
      },
      async init() {
        try {
          this.NFTInf = this.$store.state.NFTInf;
          console.log(this.NFTInf);
          this.saleId = this.NFTInf.saleId
          this.modelPath = this.NFTInf.modelPath
          this.nftName = this.NFTInf.nftName
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
      async Buy() {
        try {
         let waitBuyNFT = {
            isActive: true,
            nftAddress: this.nftAddress,
            price: this.price,
            saleId: Number(this.saleId),
            ownerAddress: this.seller,
            tokenId: Number(this.tokenId),
            symbol: this.symbol,
            tokenURI: this.modelPath,
          };
          let isSuccess = false;
          this.isChanging = true;
          await Buy(waitBuyNFT).then((re) => {
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
          } else {
            this.isChanging = false;
            this.$notify({
              title: "您已经取消购买",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
          }
        } catch (error) {
          console.log(error);
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
@import "@/style/components/NFTInf3D.scss";
</style>