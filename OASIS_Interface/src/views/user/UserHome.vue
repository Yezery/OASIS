<template>
  <div class="userHomeMain">
    <el-container class="userHomeMainBox">
      <div class="inf" ref="inf">
        <div class="content ">
          <div class="contentLeft animate__animated animate__fadeInDown">
  
              <div class="userInf ">
                <div class="userInfTop">
                  <div class="userAvatarBox">
                    <img :src="userAvatar" alt="">
                  </div>
                  <div class="userName">
                    {{ userName }}
                  </div>
                </div>
                <div class="userInfBottom">
                  <div class="UserBalance">
                    <span style="font-size: 4vw;font-weight: 800;"> {{ userBalance }}</span> ETH
                  </div>
                </div>
              </div>
              <div class="BurBox  animate__animated animate__fadeInLeft" v-if="isInitModel">
                <Show3D :model-path="modelPath" @initModel="seeModel" />
              </div>

          </div>
          <div class="contentRight animate__animated animate__fadeInUp">
            <div class="NFTListBox">
              <div class="NFTList">
                <el-collapse v-if="NFTSeriesnameList.length !== 0" @change="isInitModel=false">
                  <el-collapse-item v-for="address,i in nftContractAddressList" :key="i" :name="NFTSeriesnameList[i]">
                    <template slot="title">
                      <el-popover title="合约地址" placement="top-start" width="350" trigger="hover" :content="address">
                        <i class="header-icon el-icon-info" slot="reference" />
                      </el-popover>
                      <span style="margin-left: 1%;font-weight: 800;font-size: 1vw;">
                        {{ NFTSeriesnameList[i] }}
                      </span>
                      <span class="ADDNFT" v-if="isOwnerCheckArray[i]">
                        <el-popover title="为该合约添加新的NFT 🎉" placement="top-start" width="200" trigger="hover" content="">
                          <router-link class="addImit" :to="{ name: 'addMintHome',query:{nftContract:address} }"><el-button type="success" plain>前往</el-button>
                          </router-link>
                          <i class="el-icon-plus" slot="reference" />
                        </el-popover>
                      </span>
                    </template>
                    <div class="collapseInnerBox">
                      <template v-for="inf in NFTArray">
                        <template v-for="nft,k in inf">
                          <template v-if="nft.nftAddress == address">
                            <div class="NFTInf" :key="k" v-if="nft.description != '3D'">
                              <div style="height:65%;width: 100%;overflow: hidden;">
                                <img class="NFTImage" :src="nft.ipfsPath" alt="">
                              </div>
                              <div class="Inf">
                                <div class="InfInnerBox">
                                  <div class="Inf-NFTNameBox">
                                    <div class="NFTName">
                                      {{ nft.nftName }}
                                    </div>
                                  </div>
                                  <div class="InfTop">
                                    <div class="TokenID">
                                      <span style="font-size: 25px;">#{{ nft.tokenId }}</span>
                                    </div>
                                    <div v-if="nft.isActive" class="priceBox">
                                      <span class="price">{{ $store.state.Web3.utils.fromWei(nft.price, 'ether') }}</span> ETH
                                    </div>
                                    <div v-else class="priceBox">
                                      <span class="price">&nbsp;</span>
                                    </div>
                                  </div>
                                </div>
                                <div class="InfBottom ">
                                  <div style="background-color: #d63131e6;" @click="OpenMessageBox(nft,2)" v-if="nft.isActive">
                                    <i class="el-icon-sold-out" />
                                  </div>
                                  <div style="background-color: #2c97fa;" v-else @click="OpenMessageBox(nft,1)">
                                    <i class="el-icon-sell" />
                                  </div>
                                </div>
                              </div>
                            </div>
                            <div class="NFTInf3D" :key="k" v-else>
                              <div class="Inf3D">
                                <div class="Inf3DLeft">
                                  <div class="Inf3DLeftTop">
                                    <el-button @click="setNewModelPath(nft.ipfsPath)" type="primary" icon="el-icon-video-play" circle />
                                  </div>

                                  <div class="Inf3DLeftBottom">
                                    <el-button @click="OpenMessageBox(nft,2,2)" type="danger" icon="el-icon-sold-out" circle v-if="nft.isActive" />
                                    <el-button v-else type="primary" @click="OpenMessageBox(nft,1,1)" icon="el-icon-sell" circle />
                                  </div>
                                </div>

                                <div class="Inf3DRight">
                                  <div style="height: 60%;display: flex;justify-content: center;align-items: center;font-size: 20px;">
                                    {{ nft.nftName }}
                                  </div>
                                  <div style="width: 90%;height: 40%;display: flex;justify-content:space-between;align-items: center;">
                                    <div class="TokenID">
                                      <span style="font-size: 20px;">#{{ nft.tokenId }}</span>
                                    </div>
                                    <div v-if="nft.isActive" class="priceBox">
                                      <span class="price">{{ $store.state.Web3.utils.fromWei(nft.price, 'ether') }}</span> ETH
                                    </div>
                                    <div v-else class="priceBox">
                                      <span class="price" />
                                    </div>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </template>
                        </template>
                      </template>
                    </div>
                  </el-collapse-item>
                </el-collapse>
                <div v-else>
                  <el-empty>
                    <template slot="description">
                      <div>
                        <span style="font-weight: 800;
                          margin-bottom: 10%;margin-top: 5%;">未查到相关藏品 </span>
                        <el-button type="success" round plain @click="toMint">
                          前往创造
                        </el-button>
                      </div>
                    </template>
                  </el-empty>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-container>

    <div class="MessageMask" v-if="MessageShow">
      <div class="Message animate__animated animate__fadeInUp">
        <div class="MessageLeft">
          <div class="imageBox" v-if="three">
            <Show3D :model-path="changeNFT.ipfsPath" @initModel="seeModel" />
          </div>
          <div class="imageBox" v-else>
            <img :src="changeNFT.ipfsPath" alt="">
          </div>
        </div>
        <div class="MessageRight">
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                系列名 <span class="tipshelp">Series Name</span>
              </div>
              <div class="tipsTitle2">
                {{ changeNFT.seriesName }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                代币符号 <span class="tipshelp">Symbol</span>
              </div>
              <div class="tipsTitle2">
                {{ changeNFT.symbol }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                名称 <span class="tipshelp">NFT Name</span>
              </div>
              <div class="tipsTitle2">
                {{ changeNFT.nftName }}
              </div>
            </div>
            <!-- <el-input
                v-model="FirstNFTName"
                placeholder="Please enter the name of Genesis NFT"
              /> -->
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                序号 <span class="tipshelp">Token ID</span>
              </div>
              <div class="tipsTitle2">
                #{{ changeNFT.tokenId }}
              </div>
            </div>
          </div>
          <div class="select" v-if="opt==1">
            <div class="tipsBox">
              <div class="tipsTitle2">
                <el-divider />
              </div>
            </div>
          </div>

          <div class="select" v-if="opt==1">
            <div class="tipsBox">
              <div class="tipsTitle">
                请输入 <span class="tipshelp">价格</span>
              </div>
              <div class="tipsTitle2">
                <el-input-number v-model="Price" :precision="3" :step="0.001" />
              </div>
            </div>
          </div>
          <div class="select" v-if="opt==2">
            <div class="tipsBox">
              <div class="tipsTitle">
                价格 <span class="tipshelp">Price</span>
              </div>
              <div class="tipsTitle2">
                {{ $store.state.Web3.utils.fromWei(changeNFT.price, 'ether') }} ETH
              </div>
            </div>
          </div>
          <div class="select">
            <div class="sumbitBox">
              <el-button @click="CloseMessageBox(1)" class="createButton" type="primary" plain>
                取消
              </el-button>
              <el-button @click="upSale" class="createButton" type="success" plain v-if="opt==1">
                上架
              </el-button>
              <el-button @click="downSale" class="createButton" type="success" plain v-if="opt==2">
                下架
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Show3D from "@/components/3DModelShow/3DModel.vue";
  import { UpSale, DownSale, getNFTStruct } from "@/api/axios/contract.js";
  import { getOwnerNFTsByAddress, search } from "@/api/axios/ownerContractLIst";
  export default {
    components: { Show3D },
    data() {
      return {
        userBalance: 0,
        userName: "",
        userAvatar: require("@/assets/webAssets/MetaMask.png"),
        UserNFTListInf: [],
        NFTArray: [],
        nftContractAddressList: [],
        NFTSeriesnameList: [],
        NFT3DList: [],
        isOwnerCheckArray: [],
        SearchVo: {
          key: "",
          isActive: false,
          minPrice: "",
          maxPrice: "",
          minmaximums: "",
          maxmaximums: "",
        },
        initModel: null,
        modelPath: "",
        isInitModel: false,

        MessageShow: false,
        changeNFT: {},
        Price: 0,
        opt: 0,
        three: false,
      };
    },
    async mounted() {
      await this.init();
      await this.getNFTSeriesnameList(this.$store.state.ownerNFTList);
      await this.GetNFTContractNFT();
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
      async init() {
        let user = this.$store.state.currentAddress;
        try {
          this.userBalance = this.$store.state.Web3.utils
            .fromWei(await this.$store.state.Web3.eth.getBalance(user), "ether")
            .slice(0, 4);
          this.userName = `${user.slice(0, 5)}...${user.slice(37)}`;
          this.userAvatar =
            "data:image/png;base64," + new this.Identicon(user, 120).toString();
        } catch (error) {
          return;
        }
      },
      async upSale() {
        if (this.Price > 0) {
          try {
            this.changeNFT.price = this.Price.toString();
            if (await UpSale(this.changeNFT)) {
              this.$notify({
                title: "上架成功 🎉",
                type: "success",
                position: "top-left",
                offset: 200,
              });
              this.CloseMessageBox(2);
            }
          } catch (error) {
            this.$notify.error({
              title: "上架异常",
              position: "top-left",
              offset: 200,
            });
          }
        } else {
          this.$notify({
            title: `价格填写有误`,
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
        this.Price = 0;
        this.UserNFTListInf = await this.$store.state.ownerNFTList;
      },
      async downSale() {
        console.log(this.changeNFT);
        try {
          if (await DownSale(this.changeNFT)) {
            this.$notify({
              title: "下架成功",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            this.UserNFTListInf = this.$store.state.ownerNFTList;
            this.CloseMessageBox(2);
          }
        } catch (error) {
          this.$notify.error({
            title: "下架失败",
            position: "top-left",
            offset: 200,
          });
        }
      },
      async GetNFTContractNFT() {
        let NFTInfList = [];
        for (const nftaddress of this.nftContractAddressList) {
          let nft = {
            currentowner: this.$store.state.currentAddress,
            nftAddress: nftaddress,
          };
          await getOwnerNFTsByAddress(nft).then((re) => {
            NFTInfList.push(re.data.data);
          });
        }
        this.NFTArray = NFTInfList;
      },
      async getSetAddressArray(ContractAddressArray) {
        this.nftContractAddressList = [];
        for (const nft of ContractAddressArray) {
          this.nftContractAddressList.push(nft.nftAddress);
        }
        this.nftContractAddressList = new Set(this.nftContractAddressList);
      },
      async getNFTSeriesnameList(array) {
        await this.getSetAddressArray(array);

        for (const key of this.nftContractAddressList) {
          let SeriesName;
          let contract = await getNFTStruct(key);
          await contract.methods
            .name()
            .call()
            .then((re) => {
              SeriesName = re;
            });
          this.NFTSeriesnameList.push(SeriesName);

          let owner;
          await contract.methods
            .owner()
            .call()
            .then((re) => {
              owner = re;
            });
          let maximums;
          await contract.methods
            ._maximums()
            .call()
            .then((re) => {
              maximums = re;
            });
          let currentId;
          await contract.methods
            ._currentId()
            .call()
            .then((re) => {
              currentId = re;
            });
          if (
            this.$store.state.currentAddress.toUpperCase() ==
              owner.toUpperCase() &&
            maximums != currentId
          ) {
            this.isOwnerCheckArray.push(true);
          } else {
            this.isOwnerCheckArray.push(false);
          }
        }
      },
      toMint() {
        this.$router.push("/mintHome")
      },
      SearchNFT() {
        if (
          this.SearchVo.key.length > 0 &&
          this.SearchVo.key.replace(/(^s*)|(s*$)/g, "").length !== 0
        ) {
          search(this.SearchVo).then((re) => {
            this.getNFTSeriesnameList(re.data.data);
          });
          this.$notify({
            title: `正在搜索...`,
            type: "success",
            position: "top-left",
            offset: 200,
          });
        } else {
          this.$notify({
            title: "输入不能为空",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
      },
      CloseMessageBox(opt) {
        this.MessageShow = false;
        this.three = false;
        this.changeNFT = {};
        this.price = "";
        if (opt == 1) {
          this.$notify({
            title: `用户取消${this.opt == 1 ? "上架" : "下架"}`,
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
      },
      OpenMessageBox(NFT, opt, three) {
        if (three) {
          this.three = true;
        }
        this.opt = opt;
        this.changeNFT = NFT;
        this.MessageShow = true;
      },
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/layout/userhome.scss";
</style> 
  