<template>
  <div class="MainWindow animate__animated animate__fadeInLeft">
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
        <div class="SellIndex ">
          <div class="SellIndexInner">
            <div class="marketShopMain ">
              <el-table :data="filteredNFTList.slice(0, 10)" class="marketShopTableLeft" @row-click="toInfPage">
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
          minmaximums: "",
          maxmaximums: "",
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
      toInfPage(row) {
        let NFTInf = {
          saleId: row.saleId,
          image: JSON.parse(row.tokenURI).image,
          nftName: JSON.parse(row.tokenURI).name,
          description: JSON.parse(row.tokenURI).description,
          nftAddress: row.nftContract,
          tokenId: row.tokenId,
          isActive: row.isActive,
          seller: row.seller,
          price: row.price,
        };
        this.$store.commit("setNFTInf", NFTInf);
        this.$router.push({
          path: "/home/NFTInf",
        });
      },
    },
  };
</script>
<style lang="scss" scoped>
@import "@/style/marketShop/marketShop.scss";
</style>

