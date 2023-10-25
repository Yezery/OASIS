<template>
  <div class="Main">
    <div class="Main_left">
      <div class="Main_left_Select">
        <div class="Main_left_SelectBox">
          <el-collapse v-model="activeNames">
            <el-collapse-item name="1">
              <template slot="title">
                类型 <i class="header-icon el-icon-info" />
              </template>
              <div class="selectBox">
                <div>3D:</div>
                <div>
                  <el-switch v-model="selectType.three" active-color="#000000" inactive-color="#808080" :active-value="true" :inactive-value="false" @change="setType" />
                </div>
              </div>
              <div class="selectBox">
                <div>图片:</div>
                <div>
                  <el-switch v-model="selectType.image" active-color="#000000" inactive-color="#808080" :active-value="true" :inactive-value="false" @change="setType" />
                </div>
              </div>
            </el-collapse-item>
            <el-collapse-item title="状态" name="2">
              <div class="selectBox">
                <div v-if="criteria.isActive">
                  在售:
                </div>
                <div v-else>
                  未售:
                </div>
                <div>
                  <el-switch v-model="criteria.isActive" active-color="#000000" inactive-color="#808080" :active-value="true" :inactive-value="false" @change="setActive" />
                </div>
              </div>
            </el-collapse-item>
            <el-collapse-item title="价格" name="3">
              <div class="selectBox" style="margin-bottom: 10px;">
                <div>
                  最高
                </div>
                <el-input-number size="mini" controls-position="right" v-model="maxPrice" :precision="3" :step="0.001" :min="0" ></el-input-number>
              </div>
              <div class="selectBox">
                <div>最低</div>

                <el-input-number size="mini" controls-position="right" v-model="minPrice" :precision="3" :step="0.001" :min="0"></el-input-number>

              </div>
            </el-collapse-item>
            <el-collapse-item title="供应量" name="4">
              <div class="selectBox" style="margin-bottom: 10px;">
                <div>
                  上限
                </div>
                <el-input-number size="mini" controls-position="right" v-model="maxmums" :precision="3" :step="0.001" :min="0" ></el-input-number>
              </div>
              <div class="selectBox">
                <div>下限</div>

                <el-input-number size="mini" controls-position="right" v-model="minmums" :precision="3" :step="0.001" :min="0"></el-input-number>

              </div>
            </el-collapse-item>
          </el-collapse>
        </div>
      </div>
    </div>
    <div class="Main_right">
      <div class="Main_right_top">
        <span>NFT</span> <span>Count:{{ total }}</span>
        <div class="searchBox ">
          <el-input class="search" placeholder="过滤查找" prefix-icon="el-icon-search" v-model="criteria.key" :clearable="true" @keydown.enter.prevent.native="search" />
        </div>
      </div>
      <div class="Main_right_bottom">
        <div class="marketShopMain ">
          <el-table :data="pageTicket" class="selectTable" @row-click="toPage">
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
                    <img class="nftImage" v-show="scope.row.description != '3D'" :src="scope.row.ipfsPath" alt="">
                  </div>
                  <div style="padding-left: 20px; font-size: 1vw; display: inline-block;">
                    {{ scope.row.nftName }}
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
          <el-pagination layout="total, sizes, prev, pager, next, jumper" @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="criteria.pageDto.page" :page-sizes="[5, 10, 50, 100]" :page-size="criteria.pageDto.pageSize" :total="total" />
        </div>

      </div>
    </div>
  </div>
</template>

<script>
  import { mainSearch } from "@/api/axios/ownerContractLIst";
  export default {
    data() {
      return {
        tableData: [
          {
            date: "2016-05-02",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1518 弄",
          },
          {
            date: "2016-05-04",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1517 弄",
          },
          {
            date: "2016-05-01",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1519 弄",
          },
          {
            date: "2016-05-03",
            name: "王小虎",
            address: "上海市普陀区金沙江路 1516 弄",
          },
        ],
        pageTicket: [],
        selectType: {
          three: false,
          image: false,
        },
        activeNames: ["1", "2", "3", "4"],

        // ======= 后端查询
        maxPrice: 0,
        minPrice: 0,
        maxmums: "",
        minmums: "",
        criteria: {
          key: "",
          isActive: false,
          type: "3",
          pageDto: {
            page: 1,
            pageSize: 10,
          },
        },
        //返回数据
        total: 0,
        results: [],
      };
    },
    async mounted() {
      for (const NFTInf of this.$route.params.results) {
        this.pageTicket.push(NFTInf.NFT);
      }
      this.total = this.pageTicket.length;
    },
    methods: {
      setType() {
        if (this.selectType.image) {
          this.criteria.type = "2";
        } else {
          this.criteria.type = "1";
        }
        if (
          (this.selectType.image && this.selectType.three) ||
          (!this.selectType.image && !this.selectType.three)
        ) {
          this.criteria.type = "3";
        }
        this.search();
      },
      setActive() {
        this.search();
      },
      async search() {
        if (this.maxmums.toString()!="0" && this.minmums.toString()!="") {
          this.criteria.maxmums = this.maxmums.toString()
          this.criteria.minmums = this.minmums.toString()
        } else {
          this.criteria.maxmums = ""
          this.criteria.minmums = ""
        }
        if (this.maxPrice.toString()!="0" && this.minPrice.toString()!="") {
          this.criteria.maxPrice=this.$store.state.Web3.utils
            .toWei((this.maxPrice).toString(), "ether")
            this.criteria.minPrice = this.$store.state.Web3.utils
            .toWei((this.minPrice).toString(), "ether")
        } else {
          this.criteria.maxPrice = ""
          this.criteria.minPrice = ""
        }
        console.log(this.criteria);
        await mainSearch(this.criteria).then((re) => {
          console.log(re.data.data.data);

          this.results = re.data.data.data;

          this.total = re.data.data.total;
          this.getPageInfo();
        });
      },
      getPageInfo() {
        //清空pageTicket中的数据
        this.pageTicket = [];
        // 获取当前页的数据
        for (
          let i =
            (this.criteria.pageDto.page - 1) * this.criteria.pageDto.pageSize;
          i < this.total;
          i++
        ) {
          //把遍历的数据添加到pageTicket里面
          this.pageTicket.push(this.results[i]);
          //判断是否达到一页的要求
          if (this.pageTicket.length === this.criteria.pageDto.pageSize) break;
        }
      },
      handleSizeChange(size) {
        //修改当前每页的数据行数
        this.criteria.pageDto.pageSize = size;
        //数据重新分页
        this.getPageInfo();
      },
      //调整当前的页码
      handleCurrentChange(pageNumber) {
        //修改当前的页码
        this.criteria.pageDto.page = pageNumber;
        //数据重新分页
        this.getPageInfo();
      },
      toPage(row) {
        console.log(row);
        if (row.description == "3D") {
          this.to3DInfPage(row);
        } else {
          this.toInfPage(row);
        }
      },
      toInfPage(row) {
        let NFTInf = {
          saleId: row.saleId,
          image: row.ipfsPath,
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
      to3DInfPage(row) {
        let NFTInf = {
          saleId: row.saleId,
          modelPath: row.ipfsPath,
          nftName: row.nftName,
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
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/marketShop/searchPage.scss";
</style>