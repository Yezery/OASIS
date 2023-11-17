<template>
  <div
    class="MainWindow animate__animated animate__fadeInRight"
    v-loading.fullscreen.lock="isChanging"
    element-loading-text="交易进行中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <div class="mintNFTFrame ">
      <div class="mintNFTFrameTop">
        <div class="addMintHome_backHome">
          <el-button
            icon="el-icon-back"
            circle
            @click="$router.back(-1)"
          />
        </div>
        <div class="mintTitle">
          Add NFT
        </div>
        <div class="mintNFTFrameTop_tipsBox">
          添加NFT
        </div>
      </div>
      <div class="mintNFTFrameMain">
        <div class="mintNFTFrameMain_left">
          <div style="text-align: left;margin-bottom: 2%;">
            <span
              style="color: black;font-weight: 800;font-size: 1.2vw;"
              v-if="processLoading"
            ><i class="el-icon-loading" />审核中...</span>
            <span
              style="color: var(--Dark--);font-weight: 800;font-size: 1.2vw;"
              v-if="isProcess&&!processLoading"
            ><i
              class="el-icon-success"
              style="color: rgb(1, 169, 1);"
            /> 图片合规</span>
            <span
              style="color: red;font-weight: 800;font-size: 1.2vw;"
              v-else-if="!isProcess&&!processLoading&&noProcess"
            ><i class="el-icon-error" /> 图片不合规</span>
          </div>
          <div class="imageUpLoad">
            <el-upload
              action="#"
              list-type="picture-card"
              ref="pictureUpload"
              :file-list="fileList"
              :on-change="setPicture"
              :auto-upload="false"
            >
              <div
                class="innerBox"
                v-if="!isUpload"
              >
                <i
                  class="el-icon-upload2"
                  style="
              font-size: 3vw;
              color: var(--Dark--);
              transition: all 0.3s ease-in-out;
              margin-bottom:15px;
            "
                />
                <div style="color: var(--Dark--);font-weight: 800;font-size: 1.2vw;">
                  将文件拖到此处，或
                </div>
                <div style="color: #2081E2;font-weight: 800;font-size: 1.2vw;">
                  点击浏览文件上传
                </div>
                <div style="color: #B3B3B3;font-weight: 300;font-size: 1vw;">
                  Max size: 50MB
                </div>
                <div style="color: #B3B3B3;font-weight: 300;font-size: 1vw;">
                  JPG,PNG,JPEG,GIF
                </div>
              </div>
              <div
                slot="file"
                slot-scope="{ file }"
              >
                <img
                  class="el-upload-list__item-thumbnail"
                  :src="file.url"
                  alt=""
                  ref="realPicture"
                >
                <span class="el-upload-list__item-actions">
                  <span
                    v-if="!disabled"
                    class="el-upload-list__item-delete"
                    @click="handleRemove(file)"
                  >
                    <i class="el-icon-delete" />
                  </span>
                </span>
              </div>
            </el-upload>
          </div>
        </div>

        <div class="mintNFTFrameMain_right">
          <div class="selectBox">
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Name (NFT系列名)
                </div>
                <div class="tipsTitle2">
                  <span style="font-size: 18px;font-weight: 800;margin-top: 2%;"> {{ seriesName }}</span>
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Symbol (该系列NFT代币符号)
                </div>
                <div class="tipsTitle2">
                  <span style="font-size: 18px;font-weight: 800;margin-top: 2%">{{ symbol }}</span>
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Add NFT Name *
                </div>
                <div class="tipsTitle2">
                  新增加的NFT名称
                </div>
              </div>
              <el-input
                v-model="NFTName"
                placeholder="Please enter the name of NFT"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Description
                </div>
                <div class="tipsTitle2">
                  描述将包含在商品详情页面上图片下方 <span style="color: red">(不能为关键字 '3D' )</span>
                </div>
              </div>
              <el-input
                type="textarea"
                v-model="Description"
                placeholder="Please add a description"
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
  import { process } from "@/utils/nsfwjs";
  import { getNFTStruct } from "@/api/axios/contract";
  import { addMint } from "@/api/axios/contract";
  export default {
    data() {
      return {
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

        isProcess: false,
        processLoading: false,
        noProcess: false,
        isChanging: false,
        isUpload:false
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
          this.$refs.pictureUpload.uploadFiles.length == 0 ||
          this.NFTName == "" ||
          !this.isProcess ||
          this.Description == "3D" ||
          this.$refs.pictureUpload.uploadFiles[0].size/1024/1024>50
        ) {
          return false;
        } else {
          return true;
        }
      },
    },
    watch: {
      Description(newData) {
        if (newData == "3D") {
          this.$notify({
            title: "不能为关键字",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
          this.Description = "";
        }
      },
    },
    methods: {
      async setPicture() {
        this.isUpload = true;
        this.processLoading = true;
        let prix = this.getFileExtendingName(
          this.$refs.pictureUpload.uploadFiles[0].raw.name
        );
        if (
          prix == ".jpeg" ||
          prix == ".png" ||
          prix == ".jpg" ||
          prix == ".JPEG" ||
          prix == ".PNG" ||
          prix == ".JPG"
        ) {
          setTimeout(async () => {
            try {
              await process(this.$refs.realPicture).then((re) => {
                this.processLoading = false;
                if (
                  re[0].className == "Neutral" ||
                  re[0].className == "Drawing"
                ) {
                  this.isProcess = true;
                } else {
                  this.handleRemove();
                  this.isProcess = false;
                  this.noProcess = true;
                }
              });
            } catch (error) {
              this.$notify.error({
                title: "系统异常",
                position: "top-left",
                offset: 200,
              });
              this.isProcess = false;
              this.processLoading = false;
              this.noProcess = false;
              console.error(error);
            }
          }, 100);
        } else {
          if (this.$refs.pictureUpload.uploadFiles[0].size / 1024 / 1024 > 50) {
            this.$notify.error({
              title: `图片超出大小 ${
                this.$refs.pictureUpload.uploadFiles[0].size / 1024 / 1024 - 50
              }`,
              position: "top-left",
              offset: 200,
            });
            this.isProcess = false;
            this.processLoading = false;
            this.noProcess = false;
            this.handleRemove();
            return;
          }
          if (prix == ".gif" || prix == ".GIF") {
            setTimeout(async () => {
              try {
                await process(this.$refs.realPicture).then((re) => {
                  this.processLoading = false;
                  if (
                    re[0].className == "Neutral" ||
                    re[0].className == "Drawing"
                  ) {
                    this.isProcess = true;
                  } else {
                    this.handleRemove();
                    this.isProcess = false;
                    this.noProcess = true;
                  }
                });
              } catch (error) {
                this.$notify.error({
                  title: "系统异常",
                  position: "top-left",
                  offset: 200,
                });
                this.isProcess = false;
                this.processLoading = false;
                this.noProcess = false;
                console.error(error);
              }
            }, 100);
          } else {
            this.$notify.error({
              title: "不支持格式",
              position: "top-left",
              offset: 200,
            });
            this.isProcess = false;
            this.processLoading = false;
            this.noProcess = false;
            this.handleRemove();
          }
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
      async addMint() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            this.isChanging = true;
            await addMint(
              this.to,
              this.nftContractAddress,
              this.NFTName,
              this.Description,
              this.$refs.pictureUpload.uploadFiles,
              this.seriesName,
              this.symbol,
              this.maximums
            );
            this.isChanging = false;
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
          } catch (error) {
            this.isChanging = false;
            this.$notify.error({
              title: "添加失败",
              position: "top-left",
              offset: 200,
            });
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
        this.$refs.pictureUpload.clearFiles();
        this.isProcess = false;
        this.noProcess = false;
        this.isUpload = false;
      },
    },
  };
</script>

<style lang="scss" scoped>
@import "@/style/mintPage/addMintNFT.scss";
</style>
