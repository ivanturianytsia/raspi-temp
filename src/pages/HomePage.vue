<template lang="html">
  <div class="page page--dashboard">
    <div class="container">
      <div class="widget widget--latest">
        <div class="widget__value widget__value--primary">
          {{ latest.value }}{{ point.unit }}
        </div>
        <div class="widget__title">
          {{ latest.datetime }}
        </div>
      </div>

      <div class="widget widget--average">
        <div class="widget__value">{{ avg }}{{ point.unit }}</div>
        <div class="widget__title">Average today</div>
      </div>

      <div class="widget widget--chart">
        <chart-box :points="points"></chart-box>
      </div>
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
}
.container {
  width: 100rem;
  padding: 1rem;
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-gap: 1rem;
  grid-template-areas: "latest average"
                       "chart  chart";

  @media only screen and (max-width: 1000px) {
    width: 50rem;
    grid-template-columns: 1fr;
    grid-template-areas: "latest"
                         "average"
                         "chart";
  }

  @media only screen and (max-width: 520px) {
    width: 100%;
  }
}
.widget {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #fff;
  padding: 3rem 0;
  box-shadow: 0 .5rem 1rem rgba(0, 0, 0, 0.1);
  border-radius: 2rem;
  font-size: 1rem;
  width: 100%;

  &__title {
    font-size: 4em;
    color: $color-primary-dark;
  }

  &__value {
    font-size: 10em;

    &--primary {}
  }

  &--latest,
  &--average {
    height: 25rem;

    @media only screen and (max-width: 520px) {
      height: 20rem;
    }
  }

  &--latest {
    grid-area: latest;
  }

  &--average {
    grid-area: average;
  }

  &--chart {
    grid-area: chart;
    height: 40rem;

    @media only screen and (max-width: 520px) {
      height: 30rem;
    }
  }

  @media only screen and (max-width: 520px) {
    font-size: 0.7rem;
  }
}
</style>
