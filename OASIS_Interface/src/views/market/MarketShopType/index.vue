<template>
  <div class="SellIndex ">
    <div class="marketShopMain ">
      <el-table
        :data="filteredNFTList.slice(0, 5)"
        class="marketShopTableLeft"
        @row-click="handleColumnClick"
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
        <el-table-column
          label="藏品"
          align="left"
        >
          <template slot-scope="scope">
            <div class="collectionRow">
              <span class="collectionImageBorder">
                <img
                  class="nftImage"
                  :src="JSON.parse(scope.row.tokenURI).image"
                  alt=""
                  width="65px"
                  height="65px"
                >
              </span>
              <span style="padding-left: 20%;">
                {{ JSON.parse(scope.row.tokenURI).name.toUpperCase() }}
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="价格"
        >
          <template slot-scope="scope">
            <h4> {{ $store.state.Web3.utils.fromWei(scope.row.price, 'ether') }} ETH</h4>
          </template>
        </el-table-column>
      </el-table>
      <el-table
        :data="filteredNFTList.slice(5,10)"
        class="marketShopTableRight"
        @row-click="handleColumnClick"
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
            <span>{{ scope.$index + 6 }}</span>
          </template>
        </el-table-column>
        <el-table-column
          label="藏品"
          align="left"
        >
          <template slot-scope="scope">
            <div class="collectionRow">
              <div class="collectionImageBorder">
                <img
                  class="nftImage"
                  :src="JSON.parse(scope.row.tokenURI).image"
                  alt=""
                  width="70px"
                  height="70px"
                >
              </div>
              {{ JSON.parse(scope.row.tokenURI).name.toUpperCase() }}
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="center"
          label="价格"
        >
          <template slot-scope="scope">
            <h4> {{ $store.state.Web3.utils.fromWei(scope.row.price, 'ether') }} ETH</h4>
          </template>
        </el-table-column>
      </el-table>
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
            v-for="NFT in this.NFTList.filter(NFT => NFT.isActive).slice(0,10)"
            :key="NFT.TokenURI"
          >
            <div class="NFTInf">
              <div class="imageBox">
                <img
                  class="NFTImage"
                  :src="JSON.parse(NFT.tokenURI).image "
                  alt=""
                >
              </div>
              <div class="Inf">
                <div class="NFTName">
                  {{ JSON.parse(NFT.tokenURI).name.toUpperCase() }}
                </div>
                <div class="InfBottom">
                  <div class="ownerBox">
                    <div class="ownerAddress">
                      {{ NFT.seller.slice(0, 5) + "..." + NFT.seller.slice(37) }}
                    </div>
                    <div class="ownerAvatarBox">
                      <img
                        class="owner"
                        :src="GETHashAvatar(NFT.seller.toString())"
                        alt=""
                        width="55px"
                        height="55px"
                      >
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

      <!-- <div class="NFTInf animate__animated animate__fadeInLeft" v-for="NFT in this.NFTList.filter(NFT => NFT.isActive).slice(0,6)" :key="NFT.TokenURI">
        <div class="imageBox">
          <img class="NFTImage" :src="JSON.parse(NFT.tokenURI).image " alt="">
        </div>
        <div class="Inf">
          <div class="NFTName">
            {{ JSON.parse(NFT.tokenURI).name.toUpperCase() }}
            <div class="ownerBox">
              <div class="ownerOutBorder">
                <img class="owner" :src="GETHashAvatar(NFT.seller.toString())" alt="" width="36px" height="36px" />
              </div>
              <div class="ownerName">{{ NFT.seller.slice(0, 5) + "..." + NFT.seller.slice(37) }}
              </div>
            </div>
          </div>
          <div class="InfBottom">
            <div class="ownerBox">
              <div class="ownerAddress">
                {{ NFT.seller.slice(0, 5) + "..." + NFT.seller.slice(37) }}
              </div>
              <div class="ownerAvatarBox">
                <img class="owner" :src="GETHashAvatar(NFT.seller.toString())" alt="" width="55px" height="55px">
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
      </div> -->
    </div>
  </div>
</template>

<script>
  import { getSaleList } from "@/api/axios/contract";
  export default {
    name: "MarketSellIndex",
    props: {},
    data() {
      return {
        NFTList: [],
        contract: null,
        filteredNFTList: [],
        isLoading: true,
      };
    },
    mounted() {
      // this.NFTList = getSaleList()
      // console.log(this.NFTList);
      getSaleList().then((re) => {
        console.log(re);
        this.NFTList = re;
        this.filteredNFTList = re.filter((NFT) => NFT.isActive);
        // this.filteredNFTList=[]
      });
    },
    methods: {
      GETHashAvatar(address) {
        return (
          "data:image/png;base64," + new this.Identicon(address, 120).toString()
        );
      },

      handleColumnClick(row) {
        console.log(row);
        this.$store.commit("setMarketNFTInf", JSON.stringify(row));
        this.$router.push({
          path: "/home/NFTInf",
        });
      },
    },
  };
</script>

<style lang="scss" scoped>
.collectionShow {
  font-family: Arial, Helvetica, sans-serif;
  width: 100%;
  height: 100%;
  margin-bottom: 6%;
  .CarouselShow{
    width: 100%;
    .el-carousel__container{

    }
    .el-carousel__item{

      background-color: transparent;
      height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
  }


  }
  

  .CollectionShowTitle {
    width: 100%;
    text-align: left;
    margin-top: 2%;
    margin-bottom: 2%;
    font-weight: 800;
    font-family: "Gilroy-Medium";
    font-size: 1.7vw;
    color: var(--Dark--);
    transition: all 0.3s ease-in-out;
  }
  .NFTInf {
    background-color: var(--White--);
    margin-bottom: 3.5%;
    padding-bottom: 70px;
    border-radius: 30px;
    width: 350px;
    height: 375px;
    overflow: hidden;
    transition: all 0.3s ease-in-out;
    &:hover {
      box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
        rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
      transition: all 0.3s ease-in-out;
    }
    .imageBox {
      width: 100%;
      height: 65%;
      border-radius: 30px;
      position: relative;
      z-index: 1;
      overflow: hidden;

      .NFTImage {
        object-fit: cover;
        width: 100%;
        height: 100%;
        transition: all 0.6s;
        cursor: pointer;
        transform: scale(1.1);
        overflow: hidden;
        &:hover {
          transform: scale(1.2);
          transition: all 0.6s;
        }
      }
    }

    .Inf {
      width: 100%;
      height: 100%;
      display: flex;
      flex-direction: column;
      .InfBottom {
        margin-top: 5%;
        display: flex;
        .ownerBox {
          flex: 1;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
          .ownerAddress {
            font-size: 1vw;
            color: gray;
            margin-bottom: 2%;
          }
          .ownerAvatarBox {
            width: 55px;
            height: 55px;
            border-radius: 10px;
            overflow: hidden;
          }
        }
        .PriceBox {
          flex: 1;
          font-size: 1vw;
          .PriceTitle {
            color: gray;
            margin-bottom: 10%;
          }
        }
      }
      .NFTName {
        width: 100%;
        padding-top: 10%;
        padding-left: 12%;
        padding-bottom: 20px;
        color: var(--Dark--);
        font-size: 20px;
        text-align: left;
        font-weight: 800;
      }
      // .PriceBox {
      //   font-size: 1vw;
      //   width: 80%;
      //   .PriceTitle {
      //     color: gray;
      //     margin-bottom: 2%;
      //   }
      // }
      // .ownerAndPrice {
      //   width: 100%;
      //   display: flex;
      //   margin-top: 18px;
      //   .ownerBox {
      //     flex: 1;
      //     display: flex;
      //     justify-content: center;
      //     align-items: center;
      //     .ownerOutBorder {
      //       width: 36px;
      //       height: 36px;
      //       border: 1px solid white;
      //       border-radius: 50%;
      //       overflow: hidden;
      //       margin-right: 10px;
      //       img {
      //         object-fit: cover;
      //       }
      //       .owner {
      //         display: inline-block;
      //       }
      //     }
      //   }
      // }
    }
  }
}

.SellIndex {
  display: flex;
  font-family: Arial, Helvetica, sans-serif;
  justify-content: center;
  align-items: center;
  flex-direction: column;

  /deep/ .marketShopMain {
    display: flex;
    justify-content: flex-start;
    align-items: flex-start;
    width: 100%;
    height: 100%;
    overflow: hidden;
    border-radius: 30px;
    padding: 3%;
    background-color: var(--White--);
    transition: all 0.3s ease-in-out;
    padding-bottom: 5%;
    &:hover {
      box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
        rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
      transition: all 0.3s ease-in-out;
    }
    .el-table,
    .el-table__expanded-cell {
      background-color: var(--White--);
      transition: all 0.3s ease-in-out;
    }
    /* 表格内背景颜色 */
    .el-table th,
    .el-table tr,
    .el-table td {
      background-color: transparent;
    }
    .el-skeleton__row {
      /* 设置骨架屏行样式 */
      height: 48px;
      margin-bottom: 10px;
      background-color: black;
    }
  }
}

//相应
// @media screen and (max-width: 1600px) and (min-width: 1600px) {
//   .SellIndex {
//     width: 100%;
//     justify-content: center;
//     align-items: center;
//     flex-direction: column;
//     display: flex;
//   }
// }

.nftImageAndPrice {
  width: 100%;
  display: flex;
  margin-top: 18px;
}
.collectionRow {
  display: flex;
  justify-content: flex-start;
  align-items: center;
  font-weight: 800;
  color: var(--Dark--);
  .collectionImageBorder {
    min-width: 65px;
    height: 65px;
    border-radius: 15px;
    overflow: hidden;
    img {
      object-fit: cover;
    }
    .nftImage {
      display: inline-block;
    }
  }
}
.priceBox {
  flex: 1;
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.nftImageName {
  font-size: 12px;
  color: var(--Dark--);
}
.priceinnerBox {
  width: 70px;
  padding: 13px 0px 13px 0px;
  color: #55c960;
  background-color: rgba(85, 201, 96, 0.12);
  border-radius: 10px;
  cursor: pointer;
  transition: all 0.8s ease-in-out;
  &:hover {
    background-color: rgba(85, 201, 96, 0.24);
    transition: all 0.8s ease-in-out;
  }
}
.price {
  font-size: 12px;
  font-weight: 800;
}

/deep/.marketShopTableLeft {
  font-size: 1vw;
  color: var(--Dark--);

  td.el-table__cell {
    border: 0px solid white;
    text-align: center;
    transition: all 0.3s ease-in-out;
  }
  .el-table__body tr:hover > td.el-table__cell {
    color: var(--Dark--);
    background-color: rgba(245, 247, 250, 0.6);
  }
  .el-table__body tr:hover > td.el-table__cell:first-child {
    border-radius: 20px 0 0 20px;
  }
  .el-table__body tr:hover > td.el-table__cell:last-child {
    border-radius: 0 20px 20px 0;
  }
  td {
    padding-top: 20px;
    padding-bottom: 20px;
  }
  .el-table__header {
    th {
      font-size: 15px;
      :nth-child(2) {
        font-size: 20px;
        text-align: left;
      }
    }
  }
}
/deep/.marketShopTableRight {
  margin-left: 5%;
  @extend .marketShopTableLeft;
}
/deep/.el-table::before {
  display: none;
}
</style>
<style>
  .el-carousel__mask{
    display: none !important;
  }
</style>