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
            @click="OpenMessageBox(sale, opt)"
            :disabled="!$store.state.isconnect || NFTSeller.toUpperCase() == $store.state.currentAddress.toUpperCase() || bought || !NFTIsActive"
          >
            {{ bought?"已购入":"购入" }}
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
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 1vw;"
                    v-if="seriesNFTArrays.filter(inf => inf.isActive && inf.ipfsPath !== NFTImage).length == 0"
                  >
                    <el-empty description="无在售" />
                  </div>
                  <div
                    class="NFTInf"
                    v-for="inf in seriesNFTArrays.filter(inf => inf.isActive && inf.ipfsPath !== NFTImage)"
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
                          <span style="font-weight: 800;font-size: 2vw;">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }} </span> ETH
                        </div>
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
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 1vw;"
                    v-if="seriesNFTArrays.filter(inf => !inf.isActive).length == 0"
                  >
                    <el-empty description="无数据" />
                  </div>
                  <template v-else>
                    <div
                      class="NFTInf"
                      v-for="inf in seriesNFTArrays.filter(inf => !inf.isActive)"
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
                    style="margin-top: 10%;margin-bottom: 10%;font-weight: 800;font-size: 1vw;"
                    v-if="seriesNFTArrays.length == 0"
                  >
                    <el-empty description="无数据" />
                  </div>
                  <div
                    class="NFTInf"
                    v-for="inf in seriesNFTArrays"
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
                          <span style="font-weight: 800;font-size: 2vw;">{{ $store.state.Web3.utils.fromWei(inf.price, 'ether') }} </span> ETH
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
    <div
      class="MessageMask"
      v-if="MessageShow"
    >
      <div class="Message animate__animated animate__fadeInUp">
        <div class="MessageLeft">
          <div class="imageBox">
            <img
              :src="NFTImage"
              alt=""
            >
          </div>
        </div>
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
import { getSaleListByContractAddress } from "@/api/axios/Sale";
import { getNFTStruct,Buy } from "@/api/axios/contract";
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
        NFTIsActive: false,
        NFTImage: "#",
        NFTName: "",

        supplyer: "",
        symbol: null,
        NFTSeries: "",
        description: "",
        maximums: 0,
        currentId: 0,
        activeTab: "first",
        seriesNFTArrays: [],
        sale: {},
        bought: false,

        MessageShow: false,
        changeNFT: {},
        opt: 0,
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
            this.NFTIsActive = this.NFTInf[5]
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
          try {
            this.NFTInf = this.$store.state.marketNFTInf;
            this.NFTContractAddress = this.NFTInf.nftAddress;
            this.NFTSaleId = this.NFTInf.saleId;
            this.NFTSeller = this.NFTInf.currentowner;
            this.NFTTokenId = this.NFTInf.tokenId;
            this.NFTPrice = this.NFTInf.price;
            this.NFTName = this.NFTInf.nftName;
            this.NFTIsActive = this.NFTInf.isActive
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
        var NFTDto={
          nftAddress: this.NFTContractAddress,
        }
        await getSaleListByContractAddress(NFTDto).then(re => {
          this.seriesNFTArrays = re.data.data
        })
      },
      async Buy() {
        try {
          this.changeNFT.symbol = this.symbol
          this.changeNFT.image = this.NFTImage
          if (await Buy(this.changeNFT)) {
            this.$notify({
            title: "💖 感谢购买",
            type: "success",
            position: "top-left",
            offset: 200,
          });
          this.bought = true
          this.CloseMessageBox(2)
          } 
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
      background-image: url("@/assets/webAssets/logoGreen.png");
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
        height: fit-content;
        padding-top: 2%;
        background-color: var(--White--);

        /deep/ .selectBox-Button {
          width: 95%;
          .SeriesNFT {
            height: 100%;
            text-align: center;
            width: 100%;
            font-family: Arial, Helvetica, sans-serif;
            .NFTInf {
              margin: 2%;
              background-color: white;
              border-radius: 30px;
              width: 325px;
              height: 400px;
              display: inline-block;
              overflow: hidden;
              margin-bottom: 10%;
              box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
                rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
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
              object-fit: contain;
              height: 100%;
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
              color: black;
              font-weight: 800;
              text-align: left;
              margin-top: 20px;
              font-size: 25px;
            }
            .ownerAndToSell {
              width: 100%;
              display: flex;
              margin-top: 18px;
              justify-content: space-evenly;
              align-items: center;
            }
            .ToSellBox {
              font-size: 1.3vw;
            width: 100px;
            padding: 15px 0px 15px 0px;
            border-radius: 10px;
            transition: all 0.3s ease-in-out;
            background-color: rgba(85, 201, 96, 0.12);
            color: #55c960;
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
          .el-tabs__content{
           height: 100%;
           background-color: var(--White--);
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
              object-fit: contain;
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
      width: 90%;
      height:70%;
      border-radius: 30px;
      background-color: white;
      display: flex;
      box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px,
        rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
      .MessageLeft {
        flex: 1;
        display: flex;
        justify-content: center;
        align-items: center;
        padding-right: 10%;
        .imageBox {
          width: 75%;
          height: 75%;
          border-radius: 30px;
          overflow: hidden;
          display: flex;
          justify-content: center;
          align-items: center;
          background-color: transparent;
          img {
            object-fit: contain;
            width: 100%;
            height: 100%;
          }
        }
      }
      .MessageRight {
        position: relative;
        flex: 1;
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