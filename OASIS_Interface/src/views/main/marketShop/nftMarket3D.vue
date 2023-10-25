<template>
  <div class="MainWindow animate__animated animate__fadeInLeft">
    <div class="MainLeft">
      <div class="MainLeft_Top">
        <div class="show3DBox">
          <Show3D
            :model-path="modelPath"
            @initModel="seeModel"
          />
        </div>
      </div>
      <div class="MainLeft_Bottom">
        <el-carousel
          :interval="7000"
          type="card"
          height="230px"
          indicator-position="none"
        >
          <el-carousel-item
            v-for="nft in SaleList3D"
            :key="JSON.parse(nft.tokenURI).image"
          >
            <div class="NFTInf3D">
              <div class="Inf3D">
                <div class="Inf3DLeft">
                  <div class="Inf3DLeftTop">
                    <el-button
                      @click="setNewModelPath(JSON.parse(nft.tokenURI).image)"
                      type="primary"
                      icon="el-icon-video-play"
                      circle
                    />
                  </div>

                  <div class="Inf3DLeftBottom">
                    <div class="TokenID">
                      <span style="font-size: 20px;">#{{ nft.tokenId }}</span>
                    </div>
                  </div>
                </div>

                <div class="Inf3DRight">
                  <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 20px;">
                    {{ JSON.parse(nft.tokenURI).name }}
                  </div>
                  <div style="width: 90%;height: 30%;display: flex;justify-content: center;">
                    <div class="priceBox">
                      <span class="price">{{ $store.state.Web3.utils.fromWei(nft.price, 'ether') }}</span> ETH
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </el-carousel-item>
        </el-carousel>
      </div>
    </div>
    <div class="MainRight">
      <div class="tableShow">
        <el-table
          :data="SaleList3D.slice(0, 10)"
          class="marketShopTableLeft"
          @row-click="to3DInfPage"
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
                  <span style="font-size: 12px;">
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
    </div>
  </div>
</template>

<script>
  import { getSaleList } from "@/api/axios/contract";
  import Show3D from "@/components/3DModelShow/3DModel.vue";
  export default {
    components: {
      Show3D,
    },
    data() {
      return {
        contract: null,
        SaleList3D: [],

        toSearch: false,

        initModel: null,
        modelPath: "",
        isInitModel: false,
      };
    },
    async mounted() {
      await getSaleList().then((re) => {
        this.NFTList = re;
        this.SaleList3D = re.filter(
          (NFT) => NFT.isActive && JSON.parse(NFT[6]).description == "3D"
        );
        this.SaleList3D.sort((a, b) => Number(a.price) - Number(b.price));
      });
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
        }, 100);
      },
      to3DInfPage(row) {
        let NFTInf = {
          saleId: row.saleId,
          modelPath: JSON.parse(row.tokenURI).image,
          nftName: JSON.parse(row.tokenURI).name,
          nftAddress: row.nftContract,
          tokenId: row.tokenId,
          isActive: row.isActive,
          seller: row.seller,
          price: row.price,
        };
        this.$store.commit("setNFTInf", NFTInf);
        this.$router.push({
          path: "/home/NFTInf3D",
        });
      },
    },
  };
</script>
<style lang="scss" scoped>
@import "@/style/marketShop/marketShop3D.scss";
</style>