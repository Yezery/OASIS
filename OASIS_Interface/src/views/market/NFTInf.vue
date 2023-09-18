<template>
  <div class="NFTInfMainBox animate__animated animate__fadeInDown">
    <div class="NFTInfMain">
      <div
        class="NFTInfBackground"
        @click="toMarketMain"
      >
        <span class="toMarketmain">
          <i
            class="el-icon-arrow-left"
            style="font-size: 3vw;"
          />
        </span>
      </div>
      <div class="NFTInfContextBox">
        <div class="NFTInfBox">
          <div class="AvatarBox">
            <div class="AvatarBorder">
              <img
                :src="NFTImage"
                alt=""
              >
            </div>
            <div class="Inf">
              <div class="NFTName">
                <h1>{{ NFTName }}</h1>
              </div>
              <div class="NFTContract">
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
                  <div class="NFTdescription">
                    {{ description }}
                  </div>
                </div>
              </div>
            </div>
          </div>
          <el-button
            type="success"
            plain
            class="Buy"
            @click="Buy(sale)"
            :disabled="!$store.state.isconnect || NFTSeller.toUpperCase() == $store.state.currentAddress.toUpperCase()"
          >
            购入
          </el-button>
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
                    class="NFTInf"
                    v-for="inf in seriesNFTArrays.filter(inf => inf.isActive !== true &&inf.name !== NFTName)"
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
                </div>
              </el-tab-pane>
              <el-tab-pane
                label="暂未发售"
                name="second"
              >
                <div class="SeriesNFT">
                  <div
                    class="NFTInf"
                    v-for="inf in seriesNFTArrays.filter(inf => inf.isActive == true &&inf.name !== NFTName)"
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
                </div>
              </el-tab-pane>
              <el-tab-pane
                label="系列NFT"
                name="third"
              >
                <div class="SeriesNFT">
                  <div
                    class="NFTInf"
                    v-for="inf in seriesNFTArrays.filter(inf => inf.image !== NFTImage)"
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
                </div>
              </el-tab-pane>
            </el-tabs>
          </div>
        </div>
        <div class="MoreNFTInfBox">
          <div class="seriesNFTBox" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { getNFTStruct, Buy } from "@/api/axios/contract";
  export default {
    name: "NFTInf",
    data() {
      return {
        NFTInf: null,
        NFTMetaData: null,
        NFTContract: null,
        NFTContractAddress: null,
        NFTSaleId: 0,
        NFTTokenId: 0,
        NFTPrice: 0,
        NFTSeller: "",

        supplyer: "",
        NFTName: "",
        NFTImage: "#",
        symbol: null,
        NFTSeries: "",
        description: "",
        maximums: 0,
        currentId: 0,
        activeTab: "first",
        seriesNFTArrays: [],
        sale: {},
      };
    },
    watch: {},
    mounted() {
      this.init();
    },
    methods: {
      async init() {
        if (typeof this.$store.state.marketNFTInf === typeof "") {
          try {
            this.NFTInf = await JSON.parse(this.$store.state.marketNFTInf);
            console.log(this.NFTInf);
            this.NFTMetaData = await JSON.parse(this.NFTInf[6]);
            this.NFTContractAddress = this.NFTInf[2];
            this.NFTSaleId = this.NFTInf[0];
            this.NFTSeller = this.NFTInf[1];
            this.NFTTokenId = this.NFTInf[3];
            this.NFTPrice = this.NFTInf[4];
            this.NFTName = this.NFTMetaData.name;
            this.supplyer = `${this.NFTInf[1].slice(
              0,
              5
            )}...${this.NFTInf[1].slice(37)}`;

            this.NFTImage = this.NFTMetaData.image;
            this.description = this.NFTMetaData.description;
            await getNFTStruct(this.NFTInf[2]).then((re) => {
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
            await this.getThisSeriesNFT();
            this.makeSale();
          } catch (error) {
            console.log(error);
          }
        } else {
          console.log(this.$store.state.marketNFTInf);
          try {
            this.NFTInf = this.$store.state.marketNFTInf;
            this.NFTContractAddress = this.NFTInf.nftAddress;
            this.NFTSaleId = this.NFTInf.saleId;
            this.NFTSeller = this.NFTInf.currentowner;
            this.NFTTokenId = this.NFTInf.tokenId;
            this.NFTPrice = this.NFTInf.price;
            this.NFTName = this.NFTInf.nftName;
            this.supplyer = `${this.NFTInf.ownerAddress.slice(
              0,
              5
            )}...${this.NFTInf.ownerAddress.slice(37)}`;

            this.NFTImage = this.NFTInf.ipfsPath;
            this.description = this.NFTInf.description;
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
            await this.getThisSeriesNFT();
            this.makeSale();
          } catch (error) {
            console.log(error);
          }
        }
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
      async getThisSeriesNFT() {
        for (let index = 0; index <= this.currentId - 1; index++) {
          await this.NFTContract.methods
            ._nftMetaData(index)
            .call()
            .then((re) => {
              let NFTstruct = JSON.parse(re);
              NFTstruct.isActive = this.seriesNFTArrays.push(JSON.parse(re));
            });
        }
        console.log(this.seriesNFTArrays);
      },
      async Buy(NFT) {
        try {
          await Buy(NFT);
          this.$notify({
            title: "购买成功",
            type: "success",
            position: "top-left",
            offset: 200,
          });
        } catch (error) {
          this.$notify.error({
            title: "购买失败",
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
  align-items: center;
  width: 100%;
  height: 100%;
  .NFTInfMain {
    font-family: Arial, Helvetica, sans-serif;
    min-width: min-content;
    width: 95%;
    height: 100%;
    .NFTInfBackground {
      width: 100%;
      height: 270px;
      background-image: url("@/assets/webAssets/MetaMask.png");
      background-repeat: repeat;
      background-size: contain;
      text-align: left;
      position: relative;
      .toMarketmain {
        position: absolute;
        top: 5%;
        left: 1%;
        margin-top: 20px;
        color: white;
        cursor: pointer;
      }
    }
    .NFTInfContextBox {
      width: 100%;
      height: 100%;
      color: var(--Dark--);
      .selectBox {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        padding-top: 2%;
        background-color: var(--White--);

        /deep/ .selectBox-Button {
          width: 95%;
          height: 70px;
          border-bottom: 1px solid #dcdfe6;
          .SeriesNFT {
            height: 100%;
            text-align: center;
            width: 100%;
            font-family: Arial, Helvetica, sans-serif;

            .NFTInf {
              margin: 2%;
              background-color: var(--White--);
              border-radius: 30px;
              width: 315px;
              height: 390px;
              display: inline-block;
              overflow: hidden;

              box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
                rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
            }
            @media screen and (max-width: 1600px) and (min-width: 1600px) {
              .ToSellMain {
                width: 100%;
                justify-content: center;
                align-items: center;
                flex-direction: column;
              }
            }
            .imageBox {
              width: 100%;
              height: 63%;
              border-radius: 30px;

              z-index: 1;
              overflow: hidden;
            }
            .Inf {
              width: 100%;
              height: 100%;
              position: relative;
            }
            .NFTImage {
              object-fit: cover;
              width: 100%;
              transition: all 0.6s;
              cursor: pointer;

              overflow: hidden;
              &:hover {
                transform: scale(1.1);

                transition: all 0.6s;
              }
            }
            .NFTName {
              display: inline-block;
              color: var(--Dark--);
              font-weight: 800;
              text-align: left;
              margin-top: 20px;
              font-size: 13px;
            }
            .ownerAndToSell {
              width: 100%;
              display: flex;
              margin-top: 18px;
              justify-content: center;
              align-items: center;
            }
            .ToSellBox {
              display: flex;
              justify-content: flex-start;
              align-items: center;
            }

            .ToSellinnerBox {
              margin: 2%;
              width: 82.79px;
              padding: 8px 0px 8px 0px;
              color: #55c960;
              background-color: rgba(85, 201, 96, 0.12);
              border-radius: 10px;
              cursor: pointer;
              transition: all 0.3s ease-in-out;
              &:hover {
                background-color: rgba(85, 201, 96, 0.24);
                transition: all 0.3s ease-in-out;
              }
            }
            .downSale {
              @extend .ToSellinnerBox;
              background-color: #d63131e5;
              color: white;
              &:hover {
                background-color: rgb(255, 0, 0);
                transition: all 0.3s ease-in-out;
              }
            }
            .apporve {
              font-size: 12px;
              font-weight: 800;
            }
            .ToSellHeader {
              display: flex;
              justify-content: space-between;
              width: 95%;
              height: 10%;
              font-size: 25px;
              margin-top: 3%;
              border-radius: 30px;
              padding: 3.5% 0;
              font-weight: 800;
              color: var(--Dark--);
            }
            .ToSellHeaderBox {
              width: 100%;
              text-align: center;
              display: flex;
              justify-content: center;
              align-items: center;
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
        }
      }
      .NFTInfBox {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 25%;
        background-color: var(--White--);
        .AvatarBox {
          width: 90%;
          height: 100%;

          position: relative;
          .Inf {
            line-height: 35px;
            position: absolute;
            left: 5%;
            top: 25%;
            .NFTName {
              width: 100%;
              height: auto;
              font-weight: 800;
              text-align: left;
              font-size: 2vw;
            }
            .NFTContract {
              margin-top: 4%;
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
              .NFTdescription {
                text-align: left;
                margin-top: 1.5%;
                font-size: 1vw;
              }
            }
          }
          .AvatarBorder {
            border: 4px solid white;
            width: 10vw;
            height: 10vw;
            max-height: 20vw;
            position: absolute;
            left: 5%;
            top: -9vw;
            border-radius: 16px;
            overflow: hidden;
            box-shadow: rgba(0, 0, 0, 0.1) 0px 0px 5px 0px,
              rgba(0, 0, 0, 0.1) 0px 0px 1px 0px;
            img {
              object-fit: fill;
              min-width: 100%;
              height: 100%;
              border-radius: 20px;
              transition: all 0.6s;
              &:hover {
                transform: scale(1.1);
                transition: all 0.6s;
              }
            }
          }
        }
        .Buy {
          border-radius: 15px;
          padding: 2% 3% 2% 3%;
          margin-right: 20%;
          transition: all 0.3s ease-in-out;
          font-size: 1.3vw;
          font-weight: 800;
          &:hover {
            transition: all 0.3s ease-in-out;
          }
        }
      }
      .MoreNFTInfBox {
        width: 100%;
        height: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        .seriesNFTBox {
          background-color: var(--White--);
          width: 100%;
          height: 100%;
        }
      }
    }
  }
}
</style>