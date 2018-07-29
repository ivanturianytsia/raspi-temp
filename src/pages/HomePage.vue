<template lang="html">
  <div>{{ latest }}{{ point.unit }} (avg. {{ avg }}{{ point.unit }})</div>
</template>

<script>
import { mapGetters } from 'vuex'

export default {
  async mounted () {
    try {
      await this.$store.dispatch('getPoints')
    } catch (err) {
      console.log(err)
    }
  },
  computed: {
    ...mapGetters(['points']),
    point () {
      if (!this.points.length) return {}
      const point = this.points.find(item => item.unit === '°C' && item.measurement === 'temp')
      if (!point) return {}
      return point
    },
    avg () {
      if (!this.point.data) return 0
      let sum = 0
      this.point.data.forEach(item => {
        sum += item.value
      })
      return (sum / this.point.data.length).toFixed(1)
    },
    latest () {
      if (!this.point.data) return 0
      let result = { datetime: 0 }
      this.point.data.forEach(item => {
        if (new Date(item.datetime) >= new Date(result.datetime)) {
          result = { ...item }
        }
      })
      return (result.value).toFixed(1)
    }
  }
}
</script>

<style lang="scss" scoped>
</style>
