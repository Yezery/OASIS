<template>
  <div class="MainWindow animate__animated animate__fadeInDown">
    <div class="searchBox ">
      <span
        class="searchSpan"
        style="width: 100%; position: relative"
      >
        <i
          class="el-icon-arrow-left"
          v-if="$store.state.isSearch"
          style="font-size: 2.4vw;position: absolute;top: 15%;left: 1%;cursor: pointer;color: white;"
          @click="isSearch"
        />
        <input
          type="text"
          class="search"
          placeholder="Search any collection"
          v-model="SearchVo.key"
          @keydown.enter.prevent.stop="SearchNFT"
        >
      </span>
    </div>

    <div class="centerBox">
      <div class="ad ">
        <ad>
          <template #ad1>
            <ad1 />
          </template>
          <template #ad2> 
            <img
              src="../../assets/AD2Assets/MotorShow 2023.png"
              alt=""
              width="100%"
              height="100%"
              style="  object-fit: cover;"
            >
          </template>
          <template #ad3>
            <img
              src="../../assets/webAssets/logoGreen.png"
              width="100%"
              height="100%"
              style="  object-fit: fill;"
            >
          </template>
          <template #ad4>
            <img
              src="../../assets/webAssets/logoWhite.png"
              alt=""
              width="100%"
              height="100%"
              style="  object-fit: fill;"
            >
          </template>
        </ad>
      </div>

      <div
        class="MarketMain animate__animated animate__fadeInLeft"
        style="animation-delay: 1s;"
      >
        <div
          class="searchResults animate__animated animate__fadeInLeft"
          v-if="$store.state.isSearch"
        >
          Search Results
        </div>

        <div
          class="SearchBox animate__animated animate__fadeInLeft"
          v-if="$store.state.isSearch"
        >
          <div
            class="NFTInf animate__animated animate__fadeInLeft"
            v-for="NFT in SearchList"
            :key="NFT.TokenURI"
            @click="toInfPage(NFT)"
          >
            <div class="imageBox">
              <img
                class="NFTImage"
                :src="NFT.ipfsPath "
                alt=""
              >
            </div>
            <div class="Inf">
              <div class="NFTName">
                {{ NFT.nftName }}
              </div>
              <div class="InfBottom">
                <div class="PriceBox">
                  <div class="PriceTitle">
                    Volume
                  </div>
                  <div style="font-weight: 800;font-size: 2vw;color: var(--Dark--);">
                    {{ $store.state.Web3.utils.fromWei(NFT.price, 'ether') }}<span style="font-size: 1vw"> ETH</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-else>
          <div class="SellTitle ">
            <h3
              class="title animate__animated animate__fadeInLeft"
              style="animation-delay: 1.1s;"
            >
              Trending NFTs
            </h3>
          </div>
          <MarketSellIndex />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import ad from "@/views/market/ad.vue";
  import ad1 from "@/views/market/ADslots/ad1.vue";
  import MarketSellIndex from "@/views/market/MarketShopType/index.vue";
  import { search } from "@/api/axios/ownerContractLIst";
  export default {
    components: { ad, ad1, MarketSellIndex },
    data() {
      return {
        SearchVo: {
          key: "",
          isActive: false,
          minPrice: "",
          maxPrice: "",
          minMaxmums: "",
          maxMaxmums: "",
        },
        SearchList: [],
      };
    },
  mounted() {
    },
    watch: {
      SearchVo: {
        handler(newVal) {
          if (newVal.key == "") {
            this.$store.commit("setIsSearch", false);
          }
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
      SearchNFT() {
        this.$store.commit("setIsSearch", true);
        search(this.SearchVo).then((re) => {
          this.SearchList = re.data.data;
        });
      },
      GETHashAvatar(address) {
        return (
          "data:image/png;base64," + new this.Identicon(address, 120).toString()
        );
      },
      toInfPage(NFT) {
        this.$store.commit("setMarketNFTInf", NFT);
        this.$router.push({
          path: "/home/NFTInf",
        });
      },
      isSearch() {
        this.$store.commit("setIsSearch", false);
      }
    },
  };
</script>

<style lang="scss" scoped>
.MainWindow {
  height: 100vh;
  width: 100%;
  overflow: auto;
  //   & .el-input{
  //     border-radius: 40px;
  //   }
  .centerBox {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
  }
}
.searchBox {
  width: 100%;
  height: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  transition: all 2s ease-in-out;

  & .search {
    // z-index: 3;
    height: 60px;
    width: 90%;
    border-radius: 50px;
    border: 1px solid gray;
    background-color: var(--White--);
    text-align: center;
    color: var(--Dark--);
    font-size: 17px;
    box-shadow: var(--boxshdow-style--);
    // &:focus {
    //   animation: shine 1s linear 1s infinite alternate;
    // }
    transition: all 0.3s ease-in-out;
    &:hover {
      box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
        rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
      transition: all 0.3s ease-in-out;
    }
  }
}
input::-webkit-input-placeholder {
  font-size: 19px;
}
.ad {
  margin-top: 1.3%;
  width: 95%;
  height: 350px;
  overflow: hidden;
  border-radius: 38px;
  position: relative;
  z-index: 1;

}
.SellTitle {
  width: 100%;
  margin-top: 2%;
  margin-bottom: 2%;
  font-size: 1vw;
  text-align: center;
  display: flex;
  justify-content: flex-start;
  align-items: left;
  & .SellAll a {
    font-size: 1vw;
    color: #55c960;
    padding: 5px;
    width: 20px;
  }
  & .title {
    font-weight: 800;
    font-family: "Gilroy-Medium";
    font-size: 1.7vw;
    color: var(--Dark--);
    width: 13vw;
    float: left;
    transition: all 0.3s ease-in-out;
  }
}
.MarketMain {
  width: 90%;
  .searchResults{
    width: 100%;
    text-align: left;
    margin-top: 3.5%;
    margin-bottom: 3.5%;
    font-weight: 800;
    font-family: "Gilroy-Medium";
    font-size: 1.7vw;
    color: var(--Dark--);
    transition: all 0.3s ease-in-out;
  }
  .SearchBox {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: auto;
    margin-top: 30px;
    border-radius: 35px;
    .NFTInf {
      background-color: var(--White--);
      margin-bottom: 3.5%;
      padding-bottom: 70px;
      margin: 2%;
      border-radius: 30px;
      width: 250px;
      height: 275px;
      overflow: hidden;
      box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
          rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
      .imageBox {
        width: 100%;
        height: 65%;
        border-radius: 30px;
        position: relative;
        z-index: 1;
        overflow: hidden;

        .NFTImage {
          object-fit: contain;
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
              font-size: 0.5vw;
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
            font-size: 0.5vw;
            .PriceTitle {
              color: gray;
              margin-bottom: 1%;
            }
          }
        }
        .NFTName {
          width: 100%;
          padding-top: 10px;
          padding-left: 12%;
          padding-bottom: 5px;
          color: var(--Dark--);
          font-size: 20px;
          text-align: left;
          font-weight: 800;
        }
      
      }
    }
  }
}
.TypeSelect {
  font-size: 1.15vw;
  font-family: "Transformers_Movie";

  // background-color: var(--White--);
  box-shadow: var(--boxshdow-style--);
  //   &:hover {
  //     background-color: var(--Dark--);
  //     transition: all 0.7s;
  // }
  //   &:hover span{
  //     background: rgba(255, 255, 255, 0);
  //   }
}
.NFTTypeIcon {
  // padding: 17px 14px 17px 14px;

  margin-right: 0.5vw;
  & img {
    transform: translateY(6px);
  }
}
</style>
<style>
/* .el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
} */
</style>
<style lang="scss">
::-webkit-scrollbar {
  width: 6px;
  height: 8px;
  display: none;
  background-color: transparent;
}
::-webkit-scrollbar-thumb {
  background-color: #ccc;
  border-radius: 25px;
}
</style>
<style lang="scss" scoped>
/* MetaMask弹出框 */
.MetaMaskConnectionTip {
  font-family: Arial, Helvetica, sans-serif;
  overflow: hidden;
  position: absolute;
  width: 350px;
  height: 300px;
  background-color: var(--White--);
  color: var(--Dark--);
  top: 50% !important;
  left: 50% !important;
  transform: translate(-50%, 100%) !important;
  border-radius: 20px;
  box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
    rgba(17, 17, 26, 0.1) 0px 0px 8px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.MetaMaskTipsBtn {
  border: 0px solid;
  background-color: white;
  box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
    rgba(17, 17, 26, 0.1) 0px 0px 8px;
  border-radius: 7px;
  padding-top: 5%;
  padding-bottom: 5%;
  padding-left: 22%;
  padding-right: 22%;
  font-size: 13px;
  font-weight: 500;
}
.MetaMaskTipsBtn:hover {
  background-color: rgb(252, 251, 251);
}
.MetaMaskTips {
  font-size: 17px;
  font-weight: 500;
}
/* 遮罩层 */
.Mask {
  z-index: 100;
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  top: 0;
  // background: rgba(0, 0, 0, 0.05);
}
</style>


<style lang="scss" scoped>
@import "@/style/index.css";
</style>
<style lang="scss" scoped>
@import "@/style/MarketShop/index.scss";
</style> 
  