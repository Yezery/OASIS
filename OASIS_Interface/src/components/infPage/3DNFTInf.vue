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
        <router-link
          :to="{ name: 'MarketShop' }"
          style="text-align: left;color: var(--Dark--); width: 90%;margin-top: 3%;"
        >
          <span class="toMarketmain">
            <i
              class="el-icon-arrow-left"
              style="font-size: 3vw;"
            />
          </span>
        </router-link>
        <div class="NFTName">
          <h1>{{ NFTName }}</h1>
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
                @click="OpenMessageBox(sale, opt)"
                :disabled="!$store.state.isconnect || NFTSeller.toUpperCase() == $store.state.currentAddress.toUpperCase() || bought || !NFTIsActive"
              >
                <i class="el-icon-shopping-cart-1" /> {{ bought?"已购入":"购入" }}
              </el-button>
            </div>
            <div class="NFTSupplyer">
              由 <span style="font-weight: 800;font-size: 1.2vw;">{{ supplyer }}</span>
            </div>
            <div class="NFTInf">
              <div class="NFTCounts">
                系列<span style="font-weight: 800;">{{ NFTSeries }}</span> ·
                代币符号<span style="font-weight: 800;">{{ symbol }}</span> ·
                限量铸造<span style="font-weight: 800;">{{ maximums }}</span> ·
                目前已铸造数量<span style="font-weight: 800;">{{ currentId }}</span>
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
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 1vw;"
                    v-if="seriesNFTArrays.filter(inf => inf.isActive).length == 0"
                  >
                    <el-empty description="无在售" />
                  </div>
                  <template v-else>
                    <template v-for="inf in seriesNFTArrays">
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
                          </div>

                          <div class="Inf3DRight">
                            <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 1vw;">
                              {{ inf.nftName }}
                            </div>
                            <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                              <div class="TokenID">
                                <span style="font-size: 20px;">#{{ inf.tokenId }}</span>
                              </div>
                          
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
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 1vw;"
                    v-if="seriesNFTArrays.filter(inf => !inf.isActive).length == 0"
                  >
                    <el-empty description="无数据" />
                  </div>
                  <template v-else>
                    <template v-for="inf in seriesNFTArrays.filter(inf => !inf.isActive)">
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
                          </div>

                          <div class="Inf3DRight">
                            <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 1vw;">
                              {{ inf.nftName }}
                            </div>
                            <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                              <div class="TokenID">
                                <span style="font-size: 20px;">#{{ inf.tokenId }}</span>
                              </div>
                          
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
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 1vw;"
                    v-if="seriesNFTArrays.length == 0"
                  >
                    <el-empty description="无数据" />
                  </div>
                  <template v-else>
                    <template v-for="inf in seriesNFTArrays">
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
                          </div>
                          <div class="Inf3DRight">
                            <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 1vw;">
                              {{ inf.nftName }}
                            </div>
                            <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                              <div class="TokenID">
                                <span style="font-size: 20px;">#{{ inf.tokenId }}</span>
                              </div>
                          
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

    <div
      class="MessageMask"
      v-if="MessageShow"
    >
      <div class="Message animate__animated animate__fadeInUp">
        <div class="MessageRight">
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                系列名 <span class="tipshelp">Series Name</span>
              </div>
              <div class="tipsTitle2">
                {{ NFTSeries }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                代币符号 <span class="tipshelp">Symbol</span>
              </div>
              <div class="tipsTitle2">
                {{ symbol }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                名称 <span class="tipshelp">NFT Name</span>
              </div>
              <div class="tipsTitle2">
                {{ NFTName }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                序号 <span class="tipshelp">Token ID</span>
              </div>
              <div class="tipsTitle2">
                #{{ NFTTokenId }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle2">
                <el-divider />
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                价格 <span class="tipshelp">Price</span>
              </div>
              <div class="tipsTitle2">
                <span style="font-size: 30px;"> {{ $store.state.Web3.utils.fromWei(NFTPrice, 'ether') }} </span>ETH
              </div>
            </div>
          </div>
          <div class="select">
            <div class="sumbitBox">
              <el-button
                @click="CloseMessageBox(1)"
                class="createButton"
                type="primary"
                plain
              >
                取消
              </el-button>
              <el-button
                @click="Buy"
                class="createButton"
                type="success"
                plain
              >
                购买
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Model from "@/components/3DModelShow/3DModel.vue";
  import { getNFTStruct, Buy } from "@/api/axios/contract";

  import { getSaleListByContractAddress } from "@/api/axios/Sale";
  // getNFTStruct,
  export default {
    name: "NFTInf",
    components: {
      Model,
    },
    data() {
      return {
        NFTInf: undefined,
        modelPath: undefined,
        NFTContractAddress: undefined,
        NFTSaleId: undefined,
        NFTSeller: "",
        NFTTokenId: undefined,
        NFTPrice: undefined,
        NFTName: undefined,
        NFTIsActive: undefined,
        supplyer: undefined,
        NFTSeries: undefined,
        symbol: undefined,
        maximums: undefined,
        currentId: undefined,
        seriesNFTArrays: [],
        activeTab: "first",
        isInitModel: false,
        sale: {},
        bought: false,

        showFloatingWindow: false,
        initModel: undefined,

        MessageShow: false,
        changeNFT: {},
        opt: 0,
        isChanging: false,
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
        this.modelPath = path
        this.isInitModel = true;
        setTimeout(() => {
        this.initModel();
      }, 100);
      },
      async init() {
        console.log(this.$store.state.marketNFTInf);
        try {
          this.NFTInf = this.$store.state.marketNFTInf;
          this.modelPath = JSON.parse(this.NFTInf.tokenURI).image;
          this.NFTContractAddress = this.NFTInf.nftContract;
          this.NFTSaleId = this.NFTInf.saleId;
          this.NFTSeller = this.NFTInf.seller;
          this.NFTTokenId = this.NFTInf.tokenId;
          this.NFTPrice = this.NFTInf.price;
          this.NFTName = JSON.parse(this.NFTInf.tokenURI).name;
          this.NFTIsActive = this.NFTInf.isActive;
          this.supplyer = `${this.NFTInf.seller.slice(
            0,
            5
          )}...${this.NFTInf.seller.slice(37)}`;
          await this.getThisSeriesNFT();
          await getNFTStruct(this.NFTContractAddress).then((re) => {
            this.NFTContract = re;
          });
          await this.NFTContract.methods
            .symbol()
            .call()
            .then((re) => {
              this.symbol = re;
            });
          await this.NFTContract.methods
            .name()
            .call()
            .then((re) => {
              this.NFTSeries = re;
            });
          await this.NFTContract.methods
            ._maximums()
            .call()
            .then((re) => {
              this.maximums = re;
            });
          await this.NFTContract.methods
            ._currentId()
            .call()
            .then((re) => {
              this.currentId = re;
            });
    
          this.makeSale();
        } catch (error) {
          console.log(error);
        }
      },
      async getThisSeriesNFT() {
        var NFTDto = {
          nftAddress: this.NFTContractAddress,
        };
        console.log(NFTDto);
        await getSaleListByContractAddress(NFTDto).then((re) => {
          this.seriesNFTArrays = re.data.data;
        });
      },
      makeSale() {
        this.sale = {
          isActive: true,
          nftAddress: this.NFTContractAddress,
          price: this.$store.state.Web3.utils.fromWei(this.NFTPrice, "ether"),
          saleId: Number(this.NFTSaleId),
          ownerAddress: this.NFTSeller,
          tokenId: Number(this.NFTTokenId),
        };
      },
      async Buy() {
        try {
          this.changeNFT.symbol = this.symbol;
          this.changeNFT.image = this.NFTImage;
          let isSuccess = false;
          this.isChanging = true;
          await Buy(this.changeNFT).then((re) => {
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
            this.CloseMessageBox(2);
          } else {
            this.isChanging = false;
            this.$notify({
              title: "您已经取消购买",
              type: "warning",
              position: "top-left",
              offset: 200,
            });
            this.CloseMessageBox(2);
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
      OpenMessageBox(NFT) {
        this.changeNFT = NFT;
        this.MessageShow = true;
      },
      CloseMessageBox(opt) {
        this.MessageShow = false;
        this.changeNFT = {};
        if (opt == 1) {
          this.$notify({
            title: `您已取消购买`,
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
      },
      toMarketMain() {
        this.$router.push("/");
        this.$store.commit("setIsSearch", false);
      },
    },
  };
</script>

<style lang="scss" scoped>
.NFTInfMainBox {
  display: flex;
  justify-content: center;
  flex-direction: column;
  width: 100%;
  .NFTInfMain {
    font-family: Arial, Helvetica, sans-serif;
    min-width: min-content;
    width: 95%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    .NFTInfMainBg {
      width: 90%;
      background-color: var(--White--);
      display: flex;
      align-items: center;
      justify-content: center;
      flex-direction: column;
      .NFTName {
        color: var(--Dark--);
        width: 75%;
        height: auto;
        font-weight: 800;
        text-align: left;
        font-size: 2vw;
      }
      .show3DWindow {
        width: 100%;

        display: flex;
        justify-content: center;
        align-items: center;

        .active {
          position: fixed;
          top: 100px; /* 根据实际情况设置元素的初始位置 */
          right: 20px;
        }
        .show3DBox {
          background-color: var(--Dark--);
          width: 70%;
          height: 50vh;
          display: flex;
          justify-content: center;
          align-items: center;
          overflow: hidden;
          border-radius: 35px;
          margin-top: 2%;
        }
      }

      .Inf {
        margin-top: 2%;
        margin-bottom: 2%;
        line-height: 50px;
        width: 75%;
        color: var(--Dark--);

        .NFTContract {
          margin-top: 2%;
        }
        .NFTSupplyer {
          text-align: left;
          font-size: 1.1vw;
        }
        .NFTInf {
          float: left;
          .NFTCounts {
            margin-top: 2%;
            text-align: left;
            font-size: 1.1vw;
            span {
              margin-left: 5px;
              margin-right: 10px;
            }
          }
        }
      }
    }

    .selectBox {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 80%;

      height: fit-content;
      padding-top: 2%;
      background-color: var(--White--);

      /deep/ .selectBox-Button {
        width: 95%;
        height: 100%;

        .NFTInf3D {
                  margin: 2%;
                  background-color: white;
                  border-radius: 25px;
                  width: 300px;
                  height: 200px;
                  display: inline-block;
                  overflow: hidden;
                  transition: all 0.3s ease-in-out;
                  box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px;
                  .Inf3D {
                    display: flex;
                    height: 100%;

                    .Inf3DLeft {
                      flex: 1;
                      height: 100%;
                      display: flex;
                      background-color: #11243d;
                      button{
                        font-size: 1vw;
                      }
                      flex-direction: column;
                      .Inf3DLeftTop {
                        width: 100%;
                        height: 50%;
                   
                        display: flex;
                        justify-content: center;
                        align-items: center;
                      }
                      .Inf3DLeftBottom {
                        width: 100%;
                        height: 50%;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                      }
                    }
                    .Inf3DRight {
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    flex-direction: column;
                      flex: 2;
                      height: 100%;
                      color: black;
                      .priceBox {
                      width: 100%;
                      .price {
                        font-size: 1.8vw;
                      }
                    }
                    .TokenID {
                        padding: 2px 0px 2px 0px;
                        border-radius: 10px;
                        transition: all 0.3s ease-in-out;
                        background-color: rgba(85, 201, 96, 0.12);
                        color: #55c960;
                        width: 70%;
                      }
  
                    }
                  }
                }
        .el-tabs__nav {
          height: 60px;
        }
        .el-tabs {
          color: #000;
          border-radius: 20px;
          left: 0px;
          top: 0px;
          position: relative;
        }
        /* 去除灰色横条 */
        .el-tabs__nav-wrap::after {
          position: static !important;
        }
        /* 设置滑块颜色 */
        .el-tabs__active-bar {
          background-color: var(--Dark--);
          border: 1px solid white;
        }
        /* 设置滑块停止位置 */
        .el-tabs__active-bar.is-top {
          height: 100%;
          width: 104px !important;
          border-radius: 17px;
          top: 0px !important;
          left: -15px !important;
          position: absolute !important;
          z-index: 1;
        }
        /* 设置当前选中样式 */
        .el-tabs__item.is-active {
          background-color: transparent;
          color: var(--White--) !important;

          z-index: 2;
        }
        /* 设置未被选中样式 */
        .el-tabs__item {
          padding: 10px 20px !important;
          width: auto;
          font-size: 18px;
          font-weight: 800;
          box-sizing: border-box;
          display: inline-block;
          position: relative !important;
          color: var(--Dark--) !important;
          z-index: 2;
        }

        .selectBox-Active {
          border: none;
          font-size: 1vw;
          border-radius: 20px;
          transition: all 0.3s ease-in-out;
          font-weight: 800;
          border: none;
          background-color: rgb(246, 246, 246);
          &:hover {
            background-color: rgb(246, 246, 246);
            transition: all 0.3s ease-in-out;
            color: white;
          }
        }
        .selectBox-NOActive {
          @extend .selectBox-Active;
        }
        .el-tabs__content {
          height: 100%;
          background-color: var(--White--);
        }
      }
    }
  }
  .Buy {
    border-radius: 15px;
    padding: 2% 3% 2% 3%;
    transition: all 0.3s ease-in-out;
    font-size: 1.3vw;
    font-weight: 800;
    &:hover {
      transition: all 0.3s ease-in-out;
    }
  }
  .MessageMask {
    z-index: 200;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    .Message {
      transition: all 0.3s ease-in-out;
      overflow: hidden;
      width: 50%;
      height: 70%;
      border-radius: 30px;
      background-color: white;
      display: flex;
      justify-content: center;
      align-items: center;
      box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px,
        rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
      // .MessageLeft {
      //   flex: 1;
      //   display: flex;
      //   justify-content: center;
      //   align-items: center;
      //   padding-right: 10%;
      //   .imageBox {
      //     width: 75%;
      //     height: 75%;
      //     border-radius: 30px;
      //     overflow: hidden;
      //     display: flex;
      //     justify-content: center;
      //     align-items: center;
      //     background-color: transparent;
      //     img {
      //       object-fit: contain;
      //       width: 100%;
      //       height: 100%;
      //     }
      //   }
      // }
      .MessageRight {
        position: relative;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        /deep/ .select {
          width: 100%;
          text-align: left;
          margin-bottom: 20px;
          .tipsBox {
            width: 100%;
            color: rgb(105, 105, 105);
            text-align: left;
            transition: all 0.3s ease-in-out;
            .tipsTitle2 {
              margin-top: 2%;
              margin-bottom: 1%;
              font-size: 15px;
              font-weight: 800;
              color: black;
              transition: all 0.3s ease-in-out;
              width: 300px;
            }
            .tipsTitle {
              font-size: 20px;
              color: black;
              .tipshelp {
                color: rgb(105, 105, 105);
                font-size: 15px;
              }
            }
          }
          .sumbitBox {
            text-align: left;
            margin-top: 3%;
            .createButton {
              padding: 20px 35px 18px 35px;
              font-size: 17px;
              border-radius: 15px;
              transition: all 0.3s ease-in-out;
              font-family: "Transformers_Movie";
            }
          }
        }
        .colseButton {
          margin-right: 10%;
        }
      }
    }
  }
}
</style>