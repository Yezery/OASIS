<template>
  <div
    id="EchartBox"
    ref="EchartBox"
  >
    <div id="EchartsShow" />
  </div>
</template>

<script>
  import * as echarts from "echarts";
  export default {
    data() {
      return {
        title: "",
        isOpen:false,
        subtitle: "",
        chartOptions: {
          title: {
            text: "",
            subtext: "{b}",
            left: "center",
            top: "10%",
            color: "#A1A2AF",
            fontFamily: "Gilroy-Medium",
            fontSize: 20,
            lineHeight: 24,
            fontWeight: "bold",
            subtextStyle: {
              color: "#ffffff",
              fontSize: 28,
              fontFamily: "Gilroy-Bold",
              fontWeight: "bold",
              lineHeight: 36,
            },
          },
          xAxis: {
            type: "category",
            data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
            axisLine: {
              show: false,
            },
            axisTick: {
              show: false,
            },
            axisLabel: {
              color: "#ffffff", //坐标值得具体的颜色
              fontSize: window.innerHeight * 0.0123 + "",
              fontFamily: "Gilroy-Medium",
            },
            boundaryGap: false, // 坐标轴两侧留白策略
          },
          yAxis: {
            show: false,
          },
          series: [
            {
              type: "line",
              smooth: true,
              data: [8240, 7532, 5480, 6748, 5365, 7364, 5264],
              itemStyle: {
                color: "#55C960", //改变折线点的颜色
                lineStyle: {
                  width: 6,
                  // color: "#55C960", //改变折线颜色
                },
              },
              emphasis: {
                itemStyle: {
                  color: "white",
                },
                label: {
                  show: true, // 显示提示标签
                  //   color: "white", // 标签字体颜色
                  padding: [15, 25, 15, 25], // 标签内边距
                  formatter: "{b} $ {c}", // 标签内容格式
                  fontWeight: "bolder",
                  borderRadius: [10, 10, 10, 10], // 设置
                  borderColor: "white",
                  backgroundColor: "#1e1e2c", // 标签背景色
                  position: "bottom",

                  fontSize: 14,
                  color: "white",
                  // color: function (params) {
                  //   // 自定义颜色函数
                  //   var colorList = [];
                  //   var data = params.data;
                  //   for (var i = 0; i < data.length; i++) {
                  //     if (i == 0) {
                  //       colorList.push("blue");
                  //     } else {
                  //       var diff = data[i] - data[i - 1];
                  //       if (diff >= 0) {
                  //         colorList.push("red"); // 上升趋势设为红色
                  //       } else {
                  //         colorList.push("green"); // 下降趋势设为绿色
                  //       }
                  //     }
                  //   }
                  //   return colorList;
                  // },
                  // color: ["red", "green", "blue", "purple", "orange", "black", "grey"]
                },
              },
              symbol: "emptyCircle", // 空心圆作为转折点标志
              symbolSize: 12, // 转折点标志大小
              lineStyle: {
                type: "solid", // 折线类型为实线
              },
              areaStyle: {
                // 折线图区域填充样式
                color: {
                  type: "linear",
                  x: 0,
                  y: 0,
                  x2: 0,
                  y2: 1,
                  colorStops: [
                    {
                      offset: 0,
                      color: "rgba(0, 255, 0, 0.3)", // 折线图填充色的起始颜色和透明度
                    },
                    {
                      offset: 1,
                      color: "rgba(0, 255, 0, 0)", // 折线图填充色的结束颜色和透明度
                    },
                  ],
                  globalCoord: false, // 缺省为 false
                },
              },
            },
          ],
        },
      };
    },
    methods: {
      isDayLight() {
        var currdate = new Date();
        if (currdate.getHours() >= 18 || currdate.getHours() < 6) {
          return true;
        } else {
          return false;
        }
      },
      // updateTitleData() {
      //   let myChart = echarts.init(document.getElementById("EchartsShow"));
      //   // 更新标题数据
      //   var sum = 0;
      //   for (
      //     let index = 0;
      //     index < this.chartOptions.series[0].data.length;
      //     index++
      //   ) {
      //     sum += this.chartOptions.series[0].data[index];
      //   }
      //   this.subtitle = `$ ${sum}`;
      //   // 使用 setOption 方法重新渲染图表
      //   if (this.isDayLight()) {
      //     this.chartOptions.title.subtext = this.subtitle;
      //     this.chartOptions.xAxis.axisLabel.color = "#000000";
      //     this.chartOptions.title.subtextStyle.color = "#000000";
      //   } else {
      //     this.chartOptions.title.subtext = this.subtitle;
      //     this.chartOptions.title.subtextStyle.color = "#ffffff";
      //     this.chartOptions.xAxis.axisLabel.color = "#ffffff";
      //   }
      //   myChart.setOption(this.chartOptions);
      // },
      // calculTotal(chart) {},
      echartsChangTheme() {
        echarts.dispose(document.getElementById("EchartsShow"));
        if (this.isOpen) {
          var sum = 0;
        for (
          let index = 0;
          index < this.chartOptions.series[0].data.length;
          index++
        ) {
          sum += this.chartOptions.series[0].data[index];
        }
        this.title=`Balance`
        this.subtitle = `$ ${sum}`;
        }
        let myChart = echarts.init(document.getElementById("EchartsShow"));
        if (this.$store.state.isDark) {
          this.chartOptions.title.text = this.title;
          this.chartOptions.title.subtext = this.subtitle;
          this.chartOptions.title.subtextStyle.color = "#000000";
          this.chartOptions.xAxis.axisLabel.color = "#000000";
          
        } else {
          this.chartOptions.title.text = this.title;
          this.chartOptions.title.subtext = this.subtitle;
          this.chartOptions.title.subtextStyle.color = "#ffffff";
          this.chartOptions.xAxis.axisLabel.color = "#ffffff";
        }
        myChart.setOption(this.chartOptions);
      },
    },
  mounted() {
    
    this.$emit("echartsChangTheme", this.echartsChangTheme);
      this.$refs.EchartBox.addEventListener("click", () => {
        this.$refs.EchartBox.style.height = "400px";
        setTimeout(() => {
          this.isOpen=true
          this.echartsChangTheme();
         
        }, 300);
      });
      this.$refs.EchartBox.addEventListener("dblclick", () => {
        this.$refs.EchartBox.style.height = "100px";
        setTimeout(() => {
          this.isOpen = false
          this.subtitle = ''
          this.title =''
          this.echartsChangTheme();
        }, 300);
      });
    },
  };
</script>

<style lang="scss" scoped>
#EchartBox {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 93%;
  height: 100px;
  background-color: var(--White--);
  border-radius: 40px;
  box-shadow: var(--boxshdow-style--);
  transition: all 0.3s ease-in-out;
  overflow: hidden;
  &:hover {
    box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
      rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
    transition: all 0.3s ease-in-out;
  }
}
#EchartsShow {
  width: 100%;
  height: 100%;
}
</style>