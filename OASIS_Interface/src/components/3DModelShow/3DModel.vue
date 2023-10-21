<template>
  <div
    style="width: 100%;height:100%;position: relative;"
    v-loading="loading"
    :element-loading-text="'正在加载'+loadingProgress+'%'"
  >
    <span
      ref="fullScreen"
      class="toScreenFull"
      @click="toggleFullScreen"
      v-if="this.isModelLoading"
    >
      <i class="el-icon-full-screen animate__animated animate__fadeInUp" />
    </span>
    <span
      class="continueRender"
      @click="continueRenderer"
      v-if="this.isModelLoading"
    >
      <i
        class="
        el-icon-video-play animate__animated animate__fadeInUp"
      />
    </span>
    <span
      class="closeRender"
      @click="disposeRendererAndClearScene"
      v-if="this.isModelLoading"
    >
      <i
        class="
el-icon-video-pause animate__animated animate__fadeInUp"
      />
    </span>

    <span
      v-if="!this.isModelLoading"
      @click="renderModel"
      style="position: absolute;left: 50%;top: 50%;transform: translate(-50%,-50%);cursor: pointer;font-size: 2vw;color: #ccc;font-weight: 800;"
    ><i class="el-icon-switch-button" /></span>
    <div
      ref="modelCanvas"
      style="width: 100%;height: 100%;"
    />
  </div>
</template>

<script>
  import * as THREE from "three";
  import TWEEN from "@tweenjs/tween.js";
  import screenfull from "screenfull";
  import { OrbitControls } from "three/examples/jsm/controls/OrbitControls.js";
  import { DRACOLoader } from "three/examples/jsm/loaders/DRACOLoader";
  import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader";
  import GUI from "three/examples/jsm/libs/lil-gui.module.min";
  import Stats from "three/examples/jsm/libs/stats.module.js";
  export default {
    props: {
      modelPath: {
        type: String,
        default: "",
      },
    },
    data() {
      return {
        renderer: {},
        scene: null,
        camera: {},
        animateId: null,
        orbitControls: {},
        modelCanvas: {},
        ambient: {},
        loader: {},
        loadingManager: {},
        stats: {},
        gui: {},
        modelGroup: new THREE.Group(),
        modelCache: {},
        dracoLoader: new DRACOLoader(),
        loading: false,
        loadingProgress: 0,
        isModelLoading: false,
      };
    },
    watch: {},
  mounted() {
      this.initModel()
    this.$emit("initModel", this.renderModel);
      
    },
    methods: {
      initModel() {
        // this.initGui();
        // 初始化场景
        this.initScene();
        // 初始化渲染
        this.initRender();
        // 初始化相机
        this.initCamera();
        // 初始化轨道控制器
        this.initOrbitControls();
        // 初始化加载器
        this.initDdacoLoader();
        // this.initStats();
        // 初始化灯光
        this.initLight();
      },
      // ========================================
      initScene() {
        this.scene = new THREE.Scene();
      },
      // ========================================
      initCamera() {
        this.camera = new THREE.PerspectiveCamera(
          60,
          this.modelCanvas.offsetWidth / this.modelCanvas.offsetHeight,
          0.1,
          3000
        );
        this.onWindowResize();
        //监听窗口大小改变事件，并调用onWindowResize函数
        window.addEventListener("resize", this.onWindowResize, false);
      },
      onWindowResize() {
        // 更新修改相机比例
        this.camera.aspect =
          this.modelCanvas.offsetWidth / this.modelCanvas.offsetHeight;

        // 更新摄像机的投影矩阵
        this.camera.updateProjectionMatrix();
        // 更新画布像素比
        this.renderer.setPixelRatio(window.devicePixelRatio);
        this.renderer.setSize(
          this.modelCanvas.offsetWidth,
          this.modelCanvas.offsetHeight
        );
      },
      // ========================================
      initRender() {
        this.renderer = new THREE.WebGLRenderer({
          // alpha: true,
          antialias: true,
        });
        // 设置渲染器的背景色
        // this.renderer.setClearColor(0xf8f3e9);
        if (this.$store.state.Theme) {
          this.renderer.setClearColor(0xf8f3e9);
        } else {
          // this.renderer.setClearColor('#161621');
          this.renderer.setClearColor("white");
        }

        // 设置渲染器的像素比
        this.renderer.setPixelRatio(window.devicePixelRatio);
        // 获取模型canvas
        this.modelCanvas = this.$refs.modelCanvas;
        // 将渲染器的dom添加到模型canvas中
        this.modelCanvas.appendChild(this.renderer.domElement);
        window.onresize = () => {
          this.isScreenFull = screenfull.isFullscreen;
          this.onWindowResize();
        };
      },
      disposeRendererAndClearScene() {
        this.removerAnimate() 
        this.renderer.dispose();
        this.$notify({
            title: "停止渲染",
            type: "success",
            position: "top-left",
            offset: 200,
          });
      },
      continueRenderer() {
        if (this.animateId == null) {
          this.animate()
        } else {
          this.$notify({
            title: "模型已经渲染",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
      },
      animate() {
        // this.stats.update();
        TWEEN.update();
        // 渲染
        if (this.orbitControls != null) {
          this.orbitControls.update();
        }
        if (this.modelGroup.children.length != 0) {
          this.modelGroup.rotation.y += 0.005;
        }
        this.renderer.render(this.scene, this.camera);
        this.animateId = requestAnimationFrame(this.animate);
      },
      removerAnimate() {
        if (this.animateId) {
          cancelAnimationFrame(this.animateId);
          this.animateId = null;
        } 
      },
      // ========================================
      initOrbitControls() {
        this.orbitControls = new OrbitControls(
          this.camera,
          this.renderer.domElement
        );
        // 允许平移
        this.orbitControls.enablePan = true;
      },
      // ========================================
      initDdacoLoader() {
        this.dracoLoader.setDecoderPath("/draco/gltf/");
        // 设置解码器类型为js
        this.dracoLoader.setDecoderConfig({ type: "js" });
        // 预加载解码器
        this.dracoLoader.preload();
      },
      // ========================================
      initStats() {
        // 创建性能监视器
        this.stats = new Stats();
        // 设置监视器面板，传入面板id（0: fps, 1: ms, 2: mb）
        this.stats.setMode(0);
        // 将监视器添加到页面中
        this.modelCanvas.appendChild(this.stats.domElement);
      },
      // ========================================
      initGui() {
        this.gui = new GUI();
      },
      // ========================================
      initLight() {
        // 创建一个点光源
        var point = new THREE.PointLight(0xffffff, 1); //光源设置
        point.position.set(300, 400, 200); //点光源位置
        this.scene.add(point); //将光源添加到场景中

        // 创建一个环境光
        var ambient = new THREE.AmbientLight(0xffffff, 1); //环境光
        ambient.position.set(200, 300, 200); //点光源位置
        this.scene.add(ambient);

        // 创建一个方向光
        var directionalLight = new THREE.DirectionalLight(0xffffff, 1); //方向光
        directionalLight.position.set(150, 300, 200);
        // this.scene.add(new THREE.DirectionalLightHelper(directionalLight));
        this.scene.add(directionalLight);

        // 创建一个聚光灯
        var spotLight = new THREE.SpotLight(0xffffff, 1); //聚光灯
        spotLight.position.set(150, 200, 200);
        // this.scene.add(new THREE.DirectionalLightHelper(spotLight));
        this.scene.add(spotLight);

        // // 设置坐标系
        // const axesHelper = new THREE.AxesHelper(50);
        // this.scene.add(axesHelper); // 坐标系
      },
      // ========================================
      renderModel() {
        this.removerAnimate() 
        // 动画
        this.animate();
        try {
          this.loading = true;
          this.loadingManager = new THREE.LoadingManager();
          // 创建GLTFLoader实例
          this.loader = new GLTFLoader(this.loadingManager);
          // 设置加载器
          this.loader.setDRACOLoader(this.dracoLoader);

          // 加载模型
          this.loader.load(
            this.modelPath,
            (gltf) => {
              // 克隆模型
              let cloneModel = gltf.scene.clone();
              this.modelCache[cloneModel.modelName] = cloneModel;
              // 设置克隆模型的位置

              // 遍历场景中的模型
              this.scene.traverse(function (child) {
                // 如果模型是网格
                if (child.isMesh) {
                  // 设置模型的阴影属性
                  child.castShadow = true;
                  child.receiveShadow = true;
                }
              });

              this.changeCamera(cloneModel);
            },
            (xhr) => {
              this.loadingProgress = parseInt((xhr.loaded / xhr.total) * 100);
              if (this.loadingProgress == 100) {
                this.loading = false;
                this.isModelLoading = true;
              }
            }
          );
        } catch (error) {
          this.$notify.error({
            title: "模型加载失败",
            position: "top-left",
            offset: 200,
          });
          console.error(error);
          this.loading = false;
          this.isModelLoading = false;
        }
      },
      changeCamera(cloneModel) {
        this.modelGroup.add(cloneModel);
        // 获取模型的边界框
        const box = new THREE.Box3().setFromObject(this.modelGroup);
        // 获取模型的尺寸
        const size = new THREE.Vector3();
        box.getSize(size);

        // 获取模型的宽度
        const width = size.x;

        // 获取模型的高度
        const height = size.y;

        // 获取模型的深度
        const depth = size.z;
        // 获取模型的最大尺寸
        let max = width > height ? width : height;
        max = max > depth ? max : depth;
        // 计算模型的缩放比例
        let scale = 1;
        this.cameraScale = max;
        // if (max < 1) {
        //   scale =
        //     (width + 1) * (height + 1) * (depth + 1) * this.camera.fov  ;
        //   this.cameraScale =
        //     (width + 1) *
        //     (height + 1) *
        //     (depth + 1) *
        //      this.camera.fov ;
        // }
        // cloneModel.position.set(0, -7, 0);
        // 缩放模型
        this.tweenScale(this.modelGroup, 0.01, scale);
        // 旋转模型
        this.tweenRotation(this.modelGroup, 2, 0);
        // 设置相机的坐标
        this.camera.position.set(0, max / 2 + max / 3, max);
        this.orbitControls.target = new THREE.Vector3(0, max / 2 + max / 15, 0);

        // 更新相机的投影矩阵
        this.camera.updateProjectionMatrix();
        // 添加模型
        this.scene.add(this.modelGroup);
      },
      tweenScale(model, oldScale, newScale) {
        model.scale.set(oldScale, oldScale, oldScale);
        // 创建一个Tween动画，让model的scale从oldScale变为newScale，动画时长为1000毫秒，使用quadratic.inout缓动函数
        new TWEEN.Tween(model.scale)
          .to({ x: newScale, y: newScale, z: newScale }, 1000)
          .easing(TWEEN.Easing.Quadratic.InOut)
          .start();
      },
      tweenRotation(model, oldRotation, newRotation) {
        model.rotation.y = -Math.PI / oldRotation;
        // 创建一个新的Tween，设置模型旋转的y轴角度，从当前角度到新的角度，持续550毫秒，使用二次函数插值
        new TWEEN.Tween(model.rotation)
          .to({ y: newRotation }, 1100)
          .easing(TWEEN.Easing.Quadratic.InOut)
          .start();
      },
      // ========================================
      toggleFullScreen() {
        if (screenfull.isEnabled) {
          screenfull.toggle(this.$refs.modelCanvas);
        }
      },
      // ========================================
    },
  };
</script>

<style lang="scss" scoped>
.toScreenFull {
  position: absolute;
  right: 6vw;
  bottom: 2vh;
  color: #161621;
  font-size: 1.1vw;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  :hover {
    transition: all 0.3s ease-in-out;
    scale: 1.5;
  }
}
.closeRender{
  position: absolute;
  right: 2vw;
  bottom: 2vh;
  color: #161621;
  font-size: 1.1vw;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  :hover {
    transition: all 0.3s ease-in-out;
    scale: 1.5;
  }
}
.continueRender{
  position: absolute;
  right: 4vw;
  bottom: 2vh;
  color: #161621;
  font-size: 1.1vw;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  :hover {
    transition: all 0.3s ease-in-out;
    scale: 1.5;
  }
}
</style>
<!-- 
 * @Description: 3D模型展示
 * @Version: 1.0
 * @Author: Yezery
 * @Date: 2018-11-08 11:08:26
 -->