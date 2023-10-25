<template>
  <div class="searchBox ">
    <el-autocomplete class="search"   @keydown.enter.prevent.native="toSearchPage" v-model="SearchVo.key" :fetch-suggestions="querySearchAsync" placeholder="Search any collection" @select="handleSelect" :trigger-on-focus="false" :clearable="true" >
      <i slot="prefix" class="el-input__icon el-icon-search"></i></el-autocomplete>
  </div>
</template>

<script>
  import { search } from "@/api/axios/ownerContractLIst";
  export default {
    data() {
      return {
        SearchVo: {
          key: "",
        },
        results: [
          
        ],
      };
    },
    mounted() {},
    methods: {
      to3DInfPage(row) {
        let NFTInf = {
          saleId: row.saleId,
          modelPath: row.ipfspath,
          nftName: row.nftName,
          description: row.description,
          nftAddress: row.nftAddress,
          tokenId: row.tokenId,
          isActive: row.isActive,
          seller: row.currentowner,
          price: row.price,
        };
        this.$store.commit("setNFTInf", NFTInf);
        this.$router.push({
          path: "/home/NFTInf3D",
        });
      },
      toInfPage(row) {
        let NFTInf = {
          saleId: row.saleId,
          image: row.ipfspath,
          nftName: row.nftName,
          description: row.description,
          nftAddress: row.nftAddress,
          tokenId: row.tokenId,
          isActive: row.isActive,
          seller: row.currentowner,
          price: row.price,
        };
        this.$store.commit("setNFTInf", NFTInf);
        this.$router.push({
          path: "/home/NFTInf",
        });
      },
       toSearchPage() {
        if (this.results==null) {
          this.results=[]
        }
        this.$router.push({
          name: "searchPage",
          params: {
            searchVo: this.SearchVo,
            results: this.results,
          },
        });

      },
      handleSelect(item) {
        if (item.NFT.description == "3D") {
          this.to3DInfPage(item.NFT);
        } else {
          this.toInfPage(item.NFT);
        }
      },
      async querySearchAsync(queryString, cb) {
        this.results =[] 
         await search(this.SearchVo).then((re) => {
          this.results = re.data.data;
         });
       
          cb(this.results);
       
       
        
       
         
        
      },
    },
  };
</script>
<style lang="scss" scoped>
@import "@/style/topMenu/search.scss";
</style>
