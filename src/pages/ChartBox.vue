<template lang="html">
  <div id="chart"></div>
</template>

<script>
import Highcharts from 'highcharts'
import moment from 'moment'

export default {
  props: ['points'],
  data () {
    return {
      chart: {}
    }
  },
  watch: {
    points () {
      this.createChart()
    }
  },
  methods: {
    createChart () {
      const data = this.points
      if (!data || !data.length) return

      const colors = ['#6865eb', '#fe2769', '#c3ac42', '#ef7e70', '#62b088', '#325252']
      const labelColor = '#ccc'

      this.chart = Highcharts.chart('chart', {
        plotOptions: { areaspline: { threshold: null } },
        title: { text: '' },
        chart: {
          type: 'areaspline',
          zoomType: 'xy',
          height: null,
          // backgroundColor: '#181a23',
          // plotBackgroundColor: '#181a23',
          spacingBottom: 15,
          spacingTop: 10,
          spacingLeft: 10,
          spacingRight: 10
        },
        xAxis: {
          categories: data[0].data.map(item => moment(new Date(item.datetime)).format('HH:mm')),
          startOnTick: true,
          tickmarkPlacement: 'on',
          gridLineWidth: 1,
          gridLineColor: '#eee',
          lineColor: '#1c1e28',
          tickWidth: 0,
          title: { text: '', style: { color: labelColor } },
          labels: { style: { color: labelColor } }
          // visible: false
        },
        yAxis: data.map(item => ({
          gridLineWidth: 0,
          title: { text: '', style: { color: labelColor } },
          labels: { format: `{value}`, style: { color: labelColor } }
          // visible: false
        })),
        tooltip: { shared: true },
        legend: { enabled: false },
        series: data.map((item, index) => ({
          name: item.measurement,
          data: item.data.map(item => Math.round(item.value * 10) / 10),
          color: colors[index % colors.length],
          tooltip: { valueSuffix: item.unit },
          marker: { enabled: false, symbol: 'circle', radius: 7 },
          fillColor: {
            linearGradient: { x1: 0, y1: 0, x2: 0, y2: 1 },
            stops: [[0, colors[index % colors.length]], [1, Highcharts.Color(colors[index % colors.length]).setOpacity(0).get('rgba')]]
          }
        })),
        responsive: {
          rules: [{
            condition: { maxWidth: 500 },
            chartOptions: { legend: { layout: 'horizontal', align: 'center', verticalAlign: 'bottom' } }
          }]
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
#chart {
  position: relative;
  height: 100%;
  width: 100%;
}
</style>
