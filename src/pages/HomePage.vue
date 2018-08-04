<template lang="html">
  <div class="page page--dashboard">
    <div class="container widget">
      <div class="widget__value widget__value--primary">
        {{ latest.value }}{{ point.unit }}
      </div>
      <div class="widget__title">
        {{ latest.datetime }}
      </div>
    </div>

    <div class="container widget">
      <div class="widget__value">{{ avg }}{{ point.unit }}</div>
      <div class="widget__title">Average today</div>
    </div>

    <div class="container widget">
      <chart-box :points="points"></chart-box>
    </div>
  </div>
</template>

<script>
import ChartBox from './ChartBox'
import moment from 'moment'
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
      return {
        value: result.value.toFixed(1),
        datetime: moment(result.datetime).fromNow()
      }
    }
  },
  components: {
    ChartBox
  }
}
</script>

<style lang="scss" scoped>
@import '../assets/scss/main.scss';

.page {
  height: auto;
  min-height: 100vh;

  & > * {
    &:not(:last-child) {
      margin-bottom: 1rem;
    }
  }

  &--dashboard {
    padding: 1rem;
  }
}
.widget {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #fff;
  padding: 3rem 0;
  box-shadow: 0 .5rem 1rem rgba(0, 0, 0, 0.1);
  border-radius: 2rem;
  font-size: 1rem;

  &__title {
    font-size: 4em;
    color: $color-primary-dark;
  }

  &__value {
    font-size: 10em;

    &--primary {}
  }

  @media (max-width: 520px) {
    font-size: 0.7rem;
    width: 100%;
    // margin-left: 1rem;
    // margin-right: 1rem;
  }
}
</style>
