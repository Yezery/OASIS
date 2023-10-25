<template>
  <div
    class="MainWindow animate__animated animate__fadeInRight"
    v-loading.fullscreen.lock="isChanging"
    element-loading-text="交易进行中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <div class="mintNFTFrame ">
      <div style="width: 100%;text-align: left;margin-left: 5%;margin-top: 3%; ">
        <router-link :to="{ name: 'userhome' }">
          <i
            class="el-icon-arrow-left"
            style="font-size: 2vw;color: var(--Dark--);"
          />
        </router-link>
      </div>
      <div class="mintNFTFrameTop">
        <div class="mintTitle">
          Add 3DNFT 
        </div>
        <div class="mintNFTFrameTop_tipsBox">
          为你的系列增加 3DNFT
        </div>
      </div>
      <div class="mintNFTFrameMain">
        <div class="mintNFTFrameMain_form">
          <div class="selectBox">
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span>Model file
                </div>
                <div class="tipsTitle2">
                  模型文件 (仅支持 .gltf 嵌入式格式与 .glb非分离格式)
                </div>
              </div>
              <div class="imageUpLoad">
                <el-upload
                  class="upload-demo"
                  ref="modelUpload"
                  :on-remove="handleRemove"
                  :file-list="fileList"
                  :auto-upload="false"
                  action="#"
                  :lmint="1"
                  :on-change="checkFileType"
                >
                  <el-button
                    class="select3DFile"
                    slot="trigger"
                    @click="handleRemove"
                  >
                    选取3D文件
                  </el-button>
                </el-upload>
                <div
                  class="show3DBox"
                  v-loading="loading"
                  v-if="isShow3D"
                  element-loading-text="准备模型中..."
                >
                  <Model
                    :model-path="modelPath"
                    @initModel="seeModel"
                  />
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Name (NFT系列名)
                </div>
                <div class="tipsTitle2">
                  <span style="font-size: 1.5vw;font-weight: 800;margin-top: 2%;"> {{ seriesName }}</span>
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Symbol (该系列NFT代币符号)
                </div>
                <div class="tipsTitle2">
                  <span style="font-size: 1.5vw;font-weight: 800;margin-top: 2%">{{ symbol }}</span>
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Add 3DNFT Name
                </div>
                <div class="tipsTitle2">
                  新增加的 3DNFT 名称
                </div>
              </div>
              <el-input
                v-model="NFTName"
                placeholder="Please enter the name of  3DNFT"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  As a gift (optional)
                </div>
                <div class="tipsTitle2">
                  作为礼物赠送（可选）
                </div>
              </div>
              <el-input
                placeholder="请输入赠送者地址"
                v-model="to"
              >
                <template slot="prepend">
                  <img
                    width="25px"
                    height="25px"
                    src="@/assets/SVG/Love.svg"
                    alt="SVG Image"
                    style="color:"
                  >
                </template>
              </el-input>
            </div>
            <div class="select">
              <div class="sumbitBox">
                <el-button
                  @click="addMint"
                  :disabled="!canSubmit"
                  class="createButton"
                >
                  添加
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { getNFTStruct } from "@/api/axios/contract";
  import { addMint } from "@/api/axios/contract";
  import { savetoIPFS } from "@/api/axios/contract";
  import Model from "@/components/3DModelShow/3DModel.vue";
  export default {
    components: {
      Model,
    },
    data() {
      return {
        // imageUrl: "",
        // fileList: [],
        // dialogImageUrl: "",
        // dialogVisible: false,
        // disabled: false,
        // Name: "",
        // Symbol: "",
        // FirstNFTName: "#0",
        // maximums: 1,
        // active: 1,
        // sumitEnable: true,
        // isRepeatClick: true,
        imageUrl: "",
        dialogImageUrl: "",
        dialogVisible: false,
        disabled: false,
        fileList: [],
        seriesName: "",
        symbol: "",
        nftContractAddress: "",
        Description: "",
        maximums: 0,
        NFTName: "",
        to: this.$store.state.currentAddress,
        sumitEnable: true,
        isRepeatClick: true,

        initModel: undefined,
        model: undefined,
        modelPath: "",
        loading: false,
        isShow3D: false,
        isTypeOf3DModel: false,
        isChanging: false,
      };
    },
    async mounted() {
      this.nftContractAddress = this.$route.query.nftContract;
      await this.init();
    },
    computed: {
      canSubmit() {
        if (
          this.seriesName == "" ||
          this.symbol == "" ||
          this.$refs.modelUpload.uploadFiles.length == 0 ||
          this.NFTName == "" ||
          !this.isTypeOf3DModel
        ) {
          return false;
        } else {
          return true;
        }
      },
    },

    methods: {
      async init() {
        let contract = await getNFTStruct(this.nftContractAddress);
        await contract.methods
          .symbol()
          .call()
          .then((re) => {
            this.symbol = re;
          });
        await contract.methods
          .name()
          .call()
          .then((re) => {
            this.seriesName = re;
          });
        await contract.methods
          ._maximums()
          .call()
          .then((re) => {
            this.maximums = re;
          });
      },
      seeModel(data) {
        this.initModel = data;
      },
      async getModel() {
        this.isShow3D = true;
        this.loading = true;

        let ipfsHash;
        await savetoIPFS(this.$refs.modelUpload.uploadFiles).then((re) => {
          ipfsHash = re;
        });
        this.loading = false;
        this.modelPath = `http://10.39.5.194:8080/ipfs/${ipfsHash}?filename=${this.FirstNFTName}`;
      },
      checkFileType() {
        let prix = this.getFileExtendingName(
          this.$refs.modelUpload.uploadFiles[0].raw.name
        );
        if (prix == ".gltf" || prix == ".glb") {
          this.FirstNFTName = this.$refs.modelUpload.uploadFiles[0].name.replace(
            /(.*\/)*([^.]+).*/gi,
            "$2"
          );
          this.isTypeOf3DModel = true;
          this.getModel();
        } else {
          this.handleRemove();
          this.$notify.error({
            title: "请上传gltf或glb格式的模型",
            position: "top-left",
            offset: 200,
          });
          this.isTypeOf3DModel = false;
        }
      },
      getFileExtendingName(filename) {
        // 文件扩展名匹配正则
        var reg = /.[^.]+$/;
        var matches = reg.exec(filename);
        if (matches) {
          return matches[0];
        }
        return "";
      },
      async addMint() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            this.isChanging=true
            await addMint(
              this.to,
              this.nftContractAddress,
              this.NFTName,
              "3D",
              this.$refs.modelUpload.uploadFiles,
              this.seriesName,
              this.symbol,
              this.maximums
            );
            this.isChanging=false
            this.$notify({
              title: "添加成功",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            
            this.fileList = [];
            this.NFTName = "";
            this.maximums = 1;
            this.Description = "";
            this.handleRemove()
          } catch (error) {
            this.isChanging=false
            this.$notify.error({
              title: "添加失败",
              position: "top-left",
              offset: 200,
            });
            return
          }
        } else {
          this.$notify({
            title: "请勿操作频繁",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }

        setTimeout(() => {
          this.isRepeatClick = true;
        }, 5000);
      },
      handleRemove() {
        this.$refs.modelUpload.clearFiles();
        this.isShow3D = false;
        this.isTypeOf3DModel = false;
      },
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/mintPage/addMint3DNFT.scss";
</style>
