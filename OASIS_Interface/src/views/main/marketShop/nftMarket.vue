<template>
  <div class="MainWindow ">
    <!-- animate__animated animate__fadeInDown/ -->
    <div class="leftMain">
      <div class="ad ">
        <ad>
          <template #ad1>
            <ad1 />
          </template>
          <template #ad2>
            <ad2 />
          </template>
          <template #ad3>
            <img src="../../../assets/AD2Assets/MotorShow 2023.png" alt="" width="100%" height="100%" style="  object-fit: cover;">
          </template>
          <template #ad4>
            <img src="../../../assets/webAssets/logoWhite.png" alt="" width="100%" height="100%" style="  object-fit: fill;">
          </template>
        </ad>
      </div>
      <div class="MarketMain">
        <!-- <div style="width: 100%;display: flex;justify-content: center;align-items: center;flex-direction: column;">
          <div class="SellTitle ">
            <h3
              class="title animate__animated animate__fadeInLeft"
              style="animation-delay: 1.1s;"
            >
              Market Classification
            </h3>
          </div> -->
        <div class="SellIndex ">
          <div class="SellIndexInner">
            <!-- <el-tabs v-model="activeTable">
              <el-tab-pane
                label="排行榜单"
                name="table1"
              >
                
              </el-tab-pane>
              <el-tab-pane
                label="模型展区"
                name="second"
              >
                <div class="marketShopMain ">
                  <el-table
                    :data="SaleList3D.slice(0, 10)"
                    class="marketShopTableLeft"
                    @row-click="toNFTInf3D"
                  >
                    <template slot="empty">
                      <div>
                        <img
                          style="padding: 10% 0 0 0;"
                          width="20%"
                          height="20%"
                          src="@/assets/webAssets/MetaMask.png"
                          alt=""
                        >
                        <h4 style="padding: 0 0 10% 0">
                          浏览器未连接Metamask
                        </h4>
                      </div>
                    </template>
                    <el-table-column
                      type="index"
                      width="70"
                      label="Rank"
                    />
                    <el-table-column label="系列">
                      <template slot-scope="scope">
                        <div class="collectionRow">
                          <div style="padding-left: 20px; font-size: 1vw; display: inline-block;">
                            {{ JSON.parse(scope.row.tokenURI).name }}
                            <span style="font-size: 0.5vw;">
                              #{{ scope.row.tokenId }}
                            </span>
                          </div>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column
                      align="center"
                      label="价格"
                      fixed="right"
                    >
                      <template slot-scope="scope">
                        <h4> {{ $store.state.Web3.utils.fromWei(scope.row.price, 'ether') }} ETH</h4>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </el-tab-pane>
            </el-tabs> -->
            <div class="marketShopMain ">
              <el-table :data="filteredNFTList.slice(0, 10)" class="marketShopTableLeft" @row-click="toNFTInf">
                <template slot="empty">
                  <div>
                    <img style="padding: 10% 0 0 0;" width="20%" height="20%" src="@/assets/webAssets/MetaMask.png" alt="">
                    <h4 style="padding: 0 0 10% 0">
                      浏览器未连接Metamask
                    </h4>
                  </div>
                </template>
                <el-table-column type="index" width="70" label="Rank" />
                <el-table-column label="藏品">
                  <template slot-scope="scope">
                    <div class="collectionRow">
                      <div class="collectionImageBorder" style="display: inline-block;">
                        <img class="nftImage" :src="JSON.parse(scope.row.tokenURI).image" alt="">
                      </div>
                      <div style="padding-left: 20px; font-size: 1vw; display: inline-block;">
                        {{ JSON.parse(scope.row.tokenURI).name.toUpperCase() }}
                        <span style="font-size: 0.5vw;">
                          #{{ scope.row.tokenId }}
                        </span>
                      </div>
                    </div>
                  </template>
                </el-table-column>
                <el-table-column align="center" label="价格" fixed="right" width="150">
                  <template slot-scope="scope">
                    <h4> {{ $store.state.Web3.utils.fromWei(scope.row.price, 'ether') }} ETH</h4>
                  </template>
                </el-table-column>
              </el-table>
              <!-- <el-table
                    :data="filteredNFTList.slice(5, 10)"
                    class="marketShopTableLeft"
                    @row-click="toNFTInf"
                  >
                    <template slot="empty">
                      <div>
                        <img
                          style="padding: 10% 0 0 0;"
                          width="20%"
                          height="20%"
                          src="@/assets/webAssets/MetaMask.png"
                          alt=""
                        >
                        <h4 style="padding: 0 0 10% 0">
                          浏览器未连接Metamask
                        </h4>
                      </div>
                    </template>
                    <el-table-column
                      width="70"
                      label="Rank"
                    >
                      <template slot-scope="scope">
                        {{ scope.$index + 6 }}
                      </template>
                    </el-table-column>
                    <el-table-column label="藏品">
                      <template slot-scope="scope">
                        <div class="collectionRow">
                          <div
                            class="collectionImageBorder"
                            style="display: inline-block;"
                          >
                            <img
                              class="nftImage"
                              :src="JSON.parse(scope.row.tokenURI).image"
                              alt=""
                            >
                          </div>
                          <div style="padding-left: 20px; font-size: 1vw; display: inline-block;">
                            {{ JSON.parse(scope.row.tokenURI).name.toUpperCase() }}
                            <span style="font-size: 0.5vw;">
                              #{{ scope.row.tokenId }}
                            </span>
                          </div>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column
                      align="center"
                      label="价格"
                      fixed="right"
                      width="150"
                    >
                      <template slot-scope="scope">
                        <h4> {{ $store.state.Web3.utils.fromWei(scope.row.price, 'ether') }} ETH</h4>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
          </div>

          <div class="collectionShow">
              <div class="CollectionShowTitle ">
                Collection Show
              </div>
              <div class="CarouselShow">
                <el-carousel
                  :interval="5000"
                  type="card"
                  height="500px"
                  width="300px"
                  indicator-position="none"
                >
                  <el-carousel-item
                    v-for="NFT in this.filteredNFTList.slice(0,10)"
                    :key="NFT.TokenURI"
                  >
                    <div class="NFTInf">
                      <div class="imageBox">
                        <img
                          class="NFTImage"
                          :src="JSON.parse(NFT.tokenURI).image "
                          :alt="JSON.parse(NFT.tokenURI).name.toUpperCase()"
                        >
                      </div>
                      <div class="Inf">
                        <div class="NFTName">
                          {{ JSON.parse(NFT.tokenURI).name.toUpperCase() }}
                        </div>
                        <div class="InfBottom">
                          <div class="ownerBox">
                            <div class="ownerAddress">
                              #{{ NFT.tokenId }}
                            </div>
                          </div>
                          <div class="PriceBox">
                            <div class="PriceTitle">
                              Volume
                            </div>
                            <div style="font-weight: 800;color: var(--Dark--);">
                              {{ $store.state.Web3.utils.fromWei(NFT.price, 'ether') }} ETH
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </el-carousel-item>
                </el-carousel>
              </div>
            </div> -->
              <!-- </div> -->
              <!-- </div> -->
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="rightMain">
      <div class="balanceEchartBox">
        <balanceEchart />
      </div>
    </div>

    <!-- <div >
      <div class="SearchBox animate__animated animate__fadeInLeft">
        <template v-for="NFT in SearchList">
          <template v-if="NFT.description !== '3D'">
            <div
              class="NFTInf animate__animated animate__fadeInLeft"
              :key="NFT.TokenURI"
              @click="toInfPage(NFT)"
            >
              <div class="imageBox">
                <img
                  class="NFTImage"
                  :src="NFT.ipfsPath"
                  :alt="NFT.nftName"
                >
              </div>
              <div class="Inf">
                <div
                  class="NFTName"
                  style="color: var(--Dark--);font-size: 20px;margin-top: 2%;"
                >
                  {{ NFT.nftName }}
                </div>
                <div class="InfBottom">
                  <div
                    class="PriceBox"
                    v-if="NFT.isActive"
                  >
                    <div
                      class="PriceTitle"
                      style="color: var(--Dark--);margin-top: 5%;margin-bottom: 2%;"
                    >
                      Volume
                    </div>
                    <div style="font-weight: 800;font-size: 2vw;color: var(--Dark--);">
                      {{ $store.state.Web3.utils.fromWei(NFT.price, 'ether') }}<span style="font-size: 1vw"> ETH</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </template>
          <template v-else>
            <div
              class="NFTInf3D animate__animated animate__fadeInLeft"
              :key="NFT.TokenURI"
              @click="to3DInfPage(NFT)"
            >
              <div class="Inf3D">
                <div class="Inf3DLeft">
                  <div class="Inf3DLeftTop">
                    <el-button
                      type="primary"
                      icon="el-icon-video-play"
                      circle
                    />
                  </div>

                  <div class="Inf3DLeftBottom">
                    <el-button
                      type="danger"
                      icon="el-icon-sold-out"
                      circle
                      v-if="NFT.isActive"
                    />
   
                    <el-button
                      v-else
                      type="primary"
                      icon="el-icon-sell"
                      circle
                    />
                  </div>
                </div>

                <div class="Inf3DRight">
                  <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 1vw;">
                    {{ NFT.nftName }}
                  </div>
                  <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                    <div class="TokenID">
                      <span style="font-size: 20px;">#{{ NFT.tokenId }}</span>
                    </div>
                    <div
                      v-if="NFT.isActive"
                      class="priceBox"
                    >
                      <span class="price">{{ $store.state.Web3.utils.fromWei(NFT.price, 'ether') }}</span> ETH
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

        <el-empty
          description="无数据"
          v-if="SearchList.length==0"
        />
      </div>
      <el-button
        v-if="SearchList.length!=0"
        round
        style="margin-bottom: 5%;padding: 2% 2% 2% 2%; "
        @click="SearchNFT(2)"
      >
        加载更多
      </el-button>
    </div> -->

  </div>
</template>

<script>
  import ad from "@/views/main/AD/ad.vue";
  import balanceEchart from "@/views/bar/balanceEchart.vue";
  import ad2 from "@/views/main/AD/ADslots/ad2.vue";
  import ad1 from "@/views/main/AD/ADslots/ad1.vue";
  // import { search } from "@/api/axios/ownerContractLIst";
  import { getSaleList } from "@/api/axios/contract";
  export default {
    components: { ad, ad1, ad2, balanceEchart },
    data() {
      return {
        SearchVo: {
          key: "",
          isActive: false,
          page: 0,
          pageSize: 10,
          minPrice: "",
          maxPrice: "",
          minMaxmums: "",
          maxMaxmums: "",
        },
        SearchList: [],
        NFTList: [],
        contract: null,
        filteredNFTList: [],
        SaleList3D: [],
        isLoading: true,

        activeTable: "table1",
        toSearch: false,
      };
    },
    async mounted() {
      await getSaleList().then((re) => {
        this.NFTList = re;
        this.filteredNFTList = re.filter(
          (NFT) => NFT.isActive && JSON.parse(NFT[6]).description !== "3D"
        );
        this.filteredNFTList.sort((a, b) => Number(a.price) - Number(b.price));
        this.SaleList3D = re.filter(
          (NFT) => NFT.isActive && JSON.parse(NFT[6]).description == "3D"
        );
        this.SaleList3D.sort((a, b) => Number(a.price) - Number(b.price));
      });
    },
    watch: {
      SearchVo: {
        handler(newVal) {
          console.log(newVal);
          // if (newVal.key == "") {
          //   this.$store.commit("setIsSearch", false);
          // }
          //   search(this.SearchVo).then((re) => {
          //   this.SearchList = re.data.data;
          // });
          console.log(this.SearchList);
        },
        deep: true,
      },
    },
    methods: {
      MMCTDis() {
        this.$refs.MMCT.classList.remove("animate__fadeIn");
        this.$refs.MMCT.classList.add("animate__fadeOut");
        setTimeout(() => {
          this.MetaMaskTipsIsShow = !this.MetaMaskTipsIsShow;
        }, 750);
      },
      async SearchNFT(opt) {
        if (opt == 2) {
          this.SearchVo.page += 1;
        }
        this.toSearch = true;
        // await search(this.SearchVo).then((re) => {
        //   this.SearchList = re.data.data;
        // });
        // if (this.SearchList.length !=0) {
        //   this.$notify({
        //   title: `查询到${this.SearchList.length}个相关NFT`,
        //   type: "success",
        //   position: "top-left",
        //   duration: 1000,
        //   offset: 200,
        //   });
        //   this.$store.commit("setIsSearch", true);
        // }
        // this.$store.commit("setIsSearch", false);
      },
      GETHashAvatar(address) {
        return (
          "data:image/png;base64," + new this.Identicon(address, 120).toString()
        );
      },
      toInfPage(NFT) {
        let NFT3D = {
          isActive: NFT.isActive,
          nftContract: NFT.nftAddress,
          price: NFT.price,
          saleId: NFT.saleId,
          seller: NFT.currentowner,
          tokenId: NFT.tokenId,
          symbol: NFT.symbol,
          seriesName: NFT.seriesName,
          maxmums: NFT.maxmums,
          tokenURI: JSON.stringify({
            name: NFT.nftName,
            description: NFT.description,
            image: NFT.ipfsPath,
          }),
        };
        this.$store.commit("setNFTInf", NFT3D);
        this.$router.push({
          path: "/home/NFTInf",
        });
      },
      to3DInfPage(NFT) {
        let NFT3D = {
          isActive: NFT.isActive,
          nftContract: NFT.nftAddress,
          price: NFT.price,
          saleId: NFT.saleId,
          seller: NFT.currentowner,
          tokenId: NFT.tokenId,
          tokenURI: JSON.stringify({
            name: NFT.nftName,
            description: NFT.description,
            image: NFT.ipfsPath,
          }),
        };

        this.$store.commit("setNFTInf", NFT3D);
        console.log(NFT);
        this.$router.push({
          path: "/home/NFTInf3D",
        });
      },
      isSearch() {
        this.$store.commit("setIsSearch", false);
      },
      toNFTInf(row) {
        this.$store.commit("setNFTInf", row);
        this.$router.push({
          path: "/home/NFTInf",
        });
      },
      toNFTInf3D(row) {
        this.$store.commit("setNFTInf", row);
        this.$router.push({
          path: "/home/NFTInf3D",
        });
      },
    },
  };
</script>
<style lang="scss" scoped>
@import "@/style/marketShop/marketShop.scss";
</style>

// <!-- <style lang="scss" scoped>
// .MainWindow {
//   width: 100%;
//   height: 100%;
//   display: flex;
//   justify-content: center;
//   align-items: center;
//   .centerBox {
//     display: flex;
//     justify-content: center;
//     align-items: center;
//     flex-direction: column;
//   }
//   .el-header, .el-footer {
//     background-color: #B3C0D1;
//   }
//   .Main{
//     padding-top: 1%;
//   }
// }

// .ad {
//   margin-top: 1.3%;
//   width: 90%;
//   height: 350px;
//   overflow: hidden;
//   border-radius: 38px;
//   position: relative;
//   z-index: 1;
// }
// .SellTitle {
//   width: 100%;
//   margin-top: 2%;
//   margin-bottom: 2%;
//   font-size: 1vw;
//   text-align: center;
//   display: flex;
//   justify-content: flex-start;
//   & .SellAll a {
//     font-size: 1vw;
//     color: #55c960;
//     padding: 5px;
//     width: 20px;
//   }
//   & .title {
//     font-weight: 800;
//     font-family: "Gilroy-Medium";
//     font-size: 1.7vw;
//     color: var(--Dark--);
//     margin-left: 5%;
//     float: left;
//     transition: all 0.3s ease-in-out;
//   }
// }
// .MarketMain {
//   width: 100%;
//   .searchResults {
//     width: 100%;
//     text-align: left;
//     margin-top: 3.5%;
//     margin-bottom: 3.5%;
//     font-weight: 800;
//     font-family: "Gilroy-Medium";
//     font-size: 1.7vw;
//     color: var(--Dark--);
//     transition: all 0.3s ease-in-out;
//   }
//   .SearchBox {
//     width: 100%;
//     height: 100%;
//     overflow: auto;
//     margin-top: 30px;
//     .NFTInf {
//       float: left;
//       background-color: var(--White--);
//       margin: 2%;
//       padding-bottom: 70px;
//       border-radius: 30px;
//       width: 350px;
//       height: 375px;
//       overflow: hidden;
//       transition: all 0.3s ease-in-out;
//       &:hover {
//         box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
//           rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
//         transition: all 0.3s ease-in-out;
//       }
//       .imageBox {
//         width: 100%;
//         height: 65%;
//         border-radius: 30px;
//         position: relative;
//         z-index: 1;
//         overflow: hidden;

//         .NFTImage {
//           object-fit: contain;
//           width: 100%;
//           height: 100%;
//           transition: all 0.6s;
//           cursor: pointer;
//           transform: scale(1.1);
//           overflow: hidden;
//           &:hover {
//             transform: scale(1.2);
//             transition: all 0.6s;
//           }
//         }
//       }

//       .Inf {
//         width: 100%;
//         height: 100%;
//         display: flex;
//         flex-direction: column;
//         .InfBottom {
//           display: flex;
//           .ownerBox {
//             flex: 1;
//             display: flex;
//             justify-content: center;
//             align-items: center;
//             .ownerAddress {
//               font-size: 1.3vw;
//               width: 100px;
//               padding: 20px 0px 20px 0px;
//               border-radius: 10px;
//               transition: all 0.3s ease-in-out;
//               background-color: rgba(85, 201, 96, 0.12);
//               color: #55c960;
//             }
//           }
//           .PriceBox {
//             flex: 1;
//             font-size: 1vw;
//             color: var(--Dark--);
//           }
//         }
//         .NFTName {
//           width: 100%;
//           padding-top: 10%;
//           padding-left: 12%;
//           padding-bottom: 20px;
//           font-size: 20px;
//           color: var(--Dark--);
//           text-align: left;
//           font-weight: 800;
//         }
//       }
//     }
//     .NFTInf3D {
//       margin: 2%;
//       background-color: white;
//       border-radius: 25px;
//       width: 350px;
//       height: 250px;
//       display: inline-block;
//       overflow: hidden;
//       transition: all 0.3s ease-in-out;
//       box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
//         rgba(17, 17, 26, 0.1) 0px 0px 8px;
//       .Inf3D {
//         display: flex;
//         height: 100%;

//         .Inf3DLeft {
//           flex: 1;
//           height: 100%;
//           display: flex;
//           background-color: #11243d;
//           button {
//             font-size: 1vw;
//           }
//           flex-direction: column;
//           .Inf3DLeftTop {
//             width: 100%;
//             height: 50%;

//             display: flex;
//             justify-content: center;
//             align-items: center;
//           }
//           .Inf3DLeftBottom {
//             width: 100%;
//             height: 50%;
//             display: flex;
//             justify-content: center;
//             align-items: center;
//           }
//         }
//         .Inf3DRight {
//           width: 100%;
//           display: flex;
//           justify-content: center;
//           align-items: center;
//           flex-direction: column;
//           flex: 2;
//           height: 100%;

//           .priceBox {
//             width: 100%;
//             color: black;
//             .price {
//               font-size: 1.8vw;
//             }
//           }
//           .TokenID {
//             padding: 2% 2px 2% 0px;
//             border-radius: 10px;
//             transition: all 0.3s ease-in-out;
//             background-color: rgba(85, 201, 96, 0.12);
//             color: #55c960;
//             width: 70%;
//           }
//         }
//       }
//     }
//   }
// }
// </style>
// <style lang="scss" scoped>
// .collectionShow {
//   font-family: Arial, Helvetica, sans-serif;
//   width: 100%;
//   height: 100%;
//   margin-bottom: 6%;
//   .CarouselShow {
//     width: 100%;
//     .el-carousel__container {
//     }
//     .el-carousel__item {
//       background-color: transparent;
//       height: 100%;
//       display: flex;
//       justify-content: center;
//       align-items: center;
//     }
//   }

//   .CollectionShowTitle {
//     width: 100%;
//     text-align: left;
//     margin-top: 2%;
//     margin-bottom: 2%;
//     font-weight: 800;
//     font-family: "Gilroy-Medium";
//     font-size: 1.7vw;
//     color: var(--Dark--);
//     transition: all 0.3s ease-in-out;
//   }
//   .NFTInf {
//     background-color: var(--White--);
//     margin-bottom: 3.5%;
//     padding-bottom: 70px;
//     border-radius: 30px;
//     width: 350px;
//     height: 375px;
//     overflow: hidden;
//     transition: all 0.3s ease-in-out;
//     &:hover {
//       box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
//         rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
//       transition: all 0.3s ease-in-out;
//     }
//     .imageBox {
//       width: 100%;
//       height: 65%;
//       border-radius: 30px;
//       position: relative;
//       z-index: 1;
//       overflow: hidden;

//       .NFTImage {
//         object-fit: contain;
//         width: 100%;
//         height: 100%;
//         transition: all 0.6s;
//         cursor: pointer;
//         transform: scale(1.1);
//         overflow: hidden;
//         &:hover {
//           transform: scale(1.2);
//           transition: all 0.6s;
//         }
//       }
//     }

//     .Inf {
//       width: 100%;
//       height: 100%;
//       display: flex;
//       flex-direction: column;
//       .InfBottom {
//         margin-top: 5%;
//         display: flex;
//         .ownerBox {
//           flex: 1;
//           display: flex;
//           justify-content: center;
//           align-items: center;
//           .ownerAddress {
//             font-size: 1.3vw;
//             width: 100px;
//             padding: 20px 0px 20px 0px;
//             border-radius: 10px;
//             transition: all 0.3s ease-in-out;
//             background-color: rgba(85, 201, 96, 0.12);
//             color: #55c960;
//           }
//         }
//         .PriceBox {
//           flex: 1;
//           font-size: 1vw;
//           color: var(--Dark--);
//           .PriceTitle {
//             margin-bottom: 10%;
//           }
//         }
//       }
//       .NFTName {
//         width: 100%;
//         padding-top: 10%;
//         padding-left: 12%;
//         padding-bottom: 20px;
//         font-size: 20px;
//         color: var(--Dark--);
//         text-align: left;
//         font-weight: 800;
//       }
//     }
//   }
// }

// .SellIndex {
//   font-family: Arial, Helvetica, sans-serif;
//   width: 90%;
//   /deep/ .SellIndexInner {
//     .el-tabs__nav-wrap::after {
//       display: none;
//     }
//     .el-tabs__nav {
//       height: 60px;
//     }
//     .el-tabs {
//       color: #000;
//       border-radius: 20px;
//       left: 0px;
//       top: 0px;
//       position: relative;
//       text-align: center;
//     }
//     /* 去除灰色横条 */
//     .el-tabs__nav-wrap::after {
//       position: static !important;
//     }
//     /* 设置滑块颜色 */
//     .el-tabs__active-bar {
//       background-color: var(--Dark--);
//       border: 1px solid white;
//     }
//     /* 设置滑块停止位置 */
//     .el-tabs__active-bar.is-top {
//       height: 100%;
//       width: 104px !important;
//       border-radius: 17px;
//       top: 0px !important;
//       left: -15px !important;
//       position: absolute !important;
//       z-index: 1;
//     }
//     /* 设置当前选中样式 */
//     .el-tabs__item.is-active {
//       background-color: transparent;
//       color: var(--White--) !important;
//       z-index: 2;
//     }
//     /* 设置未被选中样式 */
//     .el-tabs__item {
//       padding: 10px 20px !important;
//       width: auto;
//       font-size: 18px;
//       font-weight: 800;
//       box-sizing: border-box;
//       display: inline-block;
//       position: relative !important;
//       color: var(--Dark--) !important;
//       z-index: 2;
//     }

//     .selectBox-Active {
//       border: none;
//       font-size: 1vw;
//       border-radius: 20px;
//       transition: all 0.3s ease-in-out;
//       font-weight: 800;
//       border: none;
//       background-color: rgb(246, 246, 246);
//       &:hover {
//         background-color: rgb(246, 246, 246);
//         transition: all 0.3s ease-in-out;
//         color: white;
//       }
//     }
//     .selectBox-NOActive {
//       @extend .selectBox-Active;
//     }
//     .el-tabs__content {
//       height: 100%;
//     }
//   }

//   /deep/ .marketShopMain {
//     display: flex;
//     width: 96%;
//     height: 96%;
//     overflow: hidden;
//     border-radius: 30px;
//     background-color: var(--White--);
//     transition: all 0.3s ease-in-out;
//     padding: 2%;
//     // padding-bottom: 5%;
//     // &:hover {
//     //   box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
//     //     rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
//     //   transition: all 0.3s ease-in-out;
//     // }
//     .el-table__fixed-right::before,
//     .el-table__fixed::before {
//       display: none;
//     }
//     .el-table,
//     .el-table__expanded-cell {
//       background-color: var(--White--);
//       transition: all 0.3s ease-in-out;
//     }
//     /* 表格内背景颜色 */
//     .el-table th,
//     .el-table tr,
//     .el-table td {
//       background-color: transparent;
//     }
//     .el-skeleton__row {
//       /* 设置骨架屏行样式 */
//       height: 48px;
//       margin-bottom: 10px;
//       background-color: black;
//     }
//     .marketShopTableLeft {
//       font-size: 1vw;
//       color: var(--Dark--);
//       flex: 1;
//       td.el-table__cell {
//         border: 0px solid white;
//         text-align: center;
//         transition: all 0.3s ease-in-out;
//       }
//       .el-table__body tr:hover > td.el-table__cell {
//         color: var(--Dark--);
//         background-color: rgba(245, 247, 250, 0.6);
//       }
//       .el-table__body tr:hover > td.el-table__cell:first-child {
//         border-radius: 20px 0 0 20px;
//       }
//       .el-table__body tr:hover > td.el-table__cell:last-child {
//         border-radius: 0 20px 20px 0;
//       }
//       td {
//         padding-top: 20px;
//         padding-bottom: 20px;
//       }
//       .el-table__header {
//         th {
//           font-size: 15px;
//           :nth-child(2) {
//             font-size: 20px;
//             text-align: left;
//           }
//         }
//       }
//     }
//     /deep/.marketShopTableRight {
//       flex: 1;
//       font-size: 1vw;
//       color: var(--Dark--);
//       flex: 1;
//       td.el-table__cell {
//         border: 0px solid white;
//         text-align: center;
//         transition: all 0.3s ease-in-out;
//       }
//       .el-table__body tr:hover > td.el-table__cell {
//         color: var(--Dark--);
//         background-color: rgba(245, 247, 250, 0.6);
//       }
//       .el-table__body tr:hover > td.el-table__cell:first-child {
//         border-radius: 20px 0 0 20px;
//       }
//       .el-table__body tr:hover > td.el-table__cell:last-child {
//         border-radius: 0 20px 20px 0;
//       }
//       td {
//         padding-top: 20px;
//         padding-bottom: 20px;
//       }
//       .el-table__header {
//         th {
//           font-size: 15px;
//           :nth-child(2) {
//             font-size: 20px;
//             text-align: left;
//           }
//         }
//       }
//     }
//   }
// }

// //相应
// // @media screen and (max-width: 1600px) and (min-width: 1600px) {
// //   .SellIndex {
// //     width: 100%;
// //     justify-content: center;
// //     align-items: center;
// //     flex-direction: column;
// //     display: flex;
// //   }
// // }

// .nftImageAndPrice {
//   width: 100%;
//   display: flex;
//   margin-top: 18px;
// }
// .collectionRow {
//   display: flex;
//   align-items: center;
//   font-weight: 800;
//   color: var(--Dark--);
//   .collectionImageBorder {
//     width: 65px;
//     height: 65px;
//     border-radius: 15px;
//     overflow: hidden;
//     .nftImage {
//       width: 100%;
//       height: 100%;
//       object-fit: contain;
//     }
//   }
// }
// .priceBox {
//   flex: 1;
//   display: flex;
//   justify-content: flex-start;
//   align-items: center;
// }

// .nftImageName {
//   font-size: 12px;
//   color: var(--Dark--);
// }
// .priceinnerBox {
//   width: 70px;
//   padding: 13px 0px 13px 0px;
//   color: #55c960;
//   background-color: rgba(85, 201, 96, 0.12);
//   border-radius: 10px;
//   cursor: pointer;
//   transition: all 0.8s ease-in-out;
//   &:hover {
//     background-color: rgba(85, 201, 96, 0.24);
//     transition: all 0.8s ease-in-out;
//   }
// }
// .price {
//   font-size: 12px;
//   font-weight: 800;
// }

// /deep/.el-table::before {
//   display: none;
// }
// </style>
// <style> -->
// <!-- .el-carousel__mask {
//   display: none !important;
// }
// ::-webkit-scrollbar {
//   width: 6px;
//   height: 8px;
//   display: none;
//   background-color: transparent;
// }
// ::-webkit-scrollbar-thumb {
//   background-color: #ccc;
//   border-radius: 25px;
// }
// </style> -->