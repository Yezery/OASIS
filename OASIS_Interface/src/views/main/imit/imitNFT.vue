<template>
  <div class="MainWindow animate__animated animate__fadeInRight" v-loading.fullscreen.lock="isChanging" element-loading-text="交易进行中" element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)">
    <div class="imitNFTFrame ">
      <div class="imitNFTFrameTop">
        <div class="ImitTitle">
          IIT YOUR NFT
        </div>
        <div class="imitNFTFrameTop_tipsBox">
          铸造NFT
        </div>
      </div>
      <div class="imitNFTFrameMain">
        <div class="imitNFTFrameMain_left">
          <div style="text-align: left;margin-bottom: 2%;">
            <span style="color: var(--Dark--);font-weight: 800;font-size: 1vw;" v-if="processLoading"><i class="el-icon-loading" />审核中...</span>
            <span style="color: rgb(1, 169, 1);font-weight: 800;font-size: 1vw;" v-if="isProcess&&!processLoading"><i class="el-icon-success" /> 图片合规</span>
            <span style="color: red;font-weight: 800;font-size: 1vw;" v-else-if="!isProcess&&!processLoading&&noProcess"><i class="el-icon-error" /> 图片不合规</span>
          </div>
          <div class="imageUpLoad">
            <el-upload action="#" list-type="picture-card" ref="pictureUpload" :file-list="fileList" :on-change="setPicture" :auto-upload="false">
              <i slot="default" class="el-icon-picture" style="
              font-size: 82px;
              color: var(--Dark--);
              transition: all 0.3s ease-in-out;
            " />
              <div slot="file" slot-scope="{ file }">
                <img class="el-upload-list__item-thumbnail" :src="file.url" alt="" ref="realPicture">
                <span class="el-upload-list__item-actions">
                  <span v-if="!disabled" class="el-upload-list__item-delete" @click="handleRemove(file)">
                    <i class="el-icon-delete" />
                  </span>
                </span>
              </div>
            </el-upload>
            <div style="margin-top: 3%;text-align: center">
              <span style="color: red">*</span><span style="color: var(--Dark--);font-weight: 800;font-size: 1vw;">支持 JPG、PNG、JEPG
                、GIF 格式</span>
            </div>
          </div>
        </div>
        <div class="imitNFTFrameMain_right">
          <div class="selectBox">
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Name
                </div>
                <div class="tipsTitle2">
                  NFT系列名
                </div>
              </div>
              <el-input v-model="Name" placeholder="Please enter a series name" />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Symbol
                </div>
                <div class="tipsTitle2">
                  该系列NFT代币符号 (要求字符长度不超过11个)
                </div>
              </div>
              <el-input v-model="Symbol" placeholder="Please enter the token symbol" maxlength="11" />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Genesis NFT Name
                </div>
                <div class="tipsTitle2">
                  该系列的一号NFT
                </div>
              </div>
              <el-input v-model="FirstNFTName" placeholder="Please enter the name of Genesis NFT" />
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
              <el-input type="textarea" v-model="Description" placeholder="Please add a description" />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Supply
                </div>
                <div class="tipsTitle2">
                  可以铸造的物品数量。
                </div>
              </div>
              <div class="input_number">
                <el-input-number v-model="Maxmums" :min="1" />
              </div>
            </div>
            <div class="select">
              <div class="sumbitBox">
                <el-button @click="createNFT" :disabled="!canSubmit" class="createButton">
                  创造
                </el-button>
              </div>
            </div>
            <div class="select">
              <div class="sumbitBox">
                <router-link :to="{ name: 'Imit3DNFT' }">
                  <el-button class="createButton" style="margin-top: 3%;width: 100%;background-color: var(--White--);color: var(--Dark--);">
                    To Imit3DNFT <i class="
el-icon-d-arrow-right" />
                  </el-button>
                </router-link>
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
  import { MakeNFT } from "@/api/axios/contract";
  export default {
    data() {
      return {
        imageUrl: "",
        fileList: [],
        dialogImageUrl: "",
        dialogVisible: false,
        disabled: false,
        Name: "",
        Symbol: "",
        FirstNFTName: "",
        Description: "",
        Maxmums: 1,
        active: 1,
        sumitEnable: true,
        isRepeatClick: true,

        isProcess: false,
        processLoading: false,
        noProcess: false,
        isChanging: false,
      };
    },
    mounted() {},
    computed: {
      canSubmit() {
        if (
          this.Name == "" ||
          this.Symbol == "" ||
          this.$refs.pictureUpload.uploadFiles.length == 0 ||
          this.FirstNFTName == "" ||
          !this.isProcess ||
          this.Description == "3D"
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
        this.processLoading = true;
        let prix = this.getFileExtendingName(
          this.$refs.pictureUpload.uploadFiles[0].raw.name
        );
        if (prix == ".jepg" || prix == ".png" || prix == ".jpg") {
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
          if (prix == ".gif") {
            this.isProcess = true;
            this.processLoading = false;
            this.noProcess = false;
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
      async createNFT() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            this.isChanging = true;
            await MakeNFT(
              this.Name,
              this.Symbol,
              this.$refs.pictureUpload.uploadFiles,
              this.Maxmums,
              this.FirstNFTName,
              this.Description
            );
            this.isChanging = false;
            this.$notify({
              title: "创造成功",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            this.fileList = [];
            this.Name = "";
            this.Symbol = "";
            this.Maxmums = 1;
            this.FirstNFTName = "";
            this.Description = "";
          } catch (error) {
            this.isChanging = false;
            this.$notify.error({
              title: "创造失败",
              position: "top-left",
              offset: 200,
            });
            return;
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
        // let uploadFiles = this.$refs.pictureUpload.uploadFiles;
        // for (var i = 0; i < uploadFiles.length; i++) {
        //   if (uploadFiles[i]["url"] == file.url) {
        //     uploadFiles.splice(i, 1);
        //   }
        // }
        this.$refs.pictureUpload.clearFiles();
        this.isProcess = false;
        this.noProcess = false;
      },
    },
  };
</script>

<style lang="scss" scoped>
@import  "@/style/imitPage/imitNFT.scss";

// .NFTImageSelectBox {
//   width: 76%;
//   text-align: left;
//   height: 100%;
//   display: flex;
//   justify-content: center;
//   align-items: center;
//   transition: all 0.3s ease-in-out;
// }
</style>
