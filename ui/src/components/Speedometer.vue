<template>
    <main :style="{width: `${config.Dashboard.size.width}px`, height: `${config.Dashboard.size.height}px`}">
        <svg id="drawing" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" :width="config.Dashboard.size.width" :height="config.Dashboard.size.height">
            <g class="circle circle-1" :transform="`translate(${config.position.x}, ${config.position.y}) rotate(180, 0, 0)`">
                <circle :r="config.radius" class="outline" fill="transparent" stroke="#5D6769" stroke-width="6" stroke-linecap="round" :stroke-dasharray="bgCircle.start" :stroke-dashoffset="bgCircle.end"></circle>
                <circle ref="speedCircle" :r="config.radius" fill="transparent" stroke="#FFD262" stroke-width="10" stroke-linecap="round"></circle>
            </g>
            <g :transform="`translate(${config.position.x}, ${config.position.y}) rotate(0, 0, 0)`">
                <line ref="speedNeedle" x1="0" y1="0" x2="0" y2="-80" stroke="#FFD262" stroke-width="6" stroke-linecap="round" transform-origin="bottom center"></line>
                <circle :r="12" fill="#FF5555" stroke="#7F4147" stroke-width="4"></circle>
            </g>
        </svg>
    </main>
</template>

<script>
    import SVG from 'svg.js'
    import { Car, Dashboard } from '../config'
import { setTimeout } from 'timers';

    export default {
        props: ['speed'],

        data () {
            return {
                draw: null,
                config: {
                    Dashboard,
                    position: {
                        x: Dashboard.size.width / 2,
                        y: Dashboard.size.height / 2,
                    },
                    radius: 120,
                },
                instances: {
                    draw: null,
                    speedCircle: null,
                    speedText: null,
                    speedUnit: null,
                    speedNeedle: null,
                }
            }
        },

        computed: {
            circunference() {
                return 2 * 3.1416 * 120 // 753.984 ( / 2 = 376.992 )
            },
            bgCircle() {
                return {
                    start: this.circunference,
                    end: this.circunference / 2,
                }
            },
            speedCircle() {
                const minCirc = this.circunference / 2

                const newEnd = (this.circunference - minCirc) * this.speed / Car.maxSpeed

                return {
                    start: this.circunference,
                    end: this.circunference - newEnd,
                    needle:180 * this.speed / Car.maxSpeed,
                }
            }
        },

        watch: {
            speed() {
                if (!this.instances.speedCircle) {
                    return
                }

                this.instances.speedText
                    .text(Math.round(this.speed).toString())
                    .x(this.instances.speedUnit.x() - this.instances.speedText.length() - 18)

                this.instances.speedCircle
                    .animate(100)
                    .attr('stroke-dashoffset', this.speedCircle.end)

                this.instances.speedNeedle.animate(100).rotate(this.speedCircle.needle - 90, 0, 0)
            }
        },

        mounted() {
            this.instances.draw = SVG('drawing')
            this.instances.speedNeedle = SVG.adopt(this.$refs.speedNeedle)

            this.instances.speedCircle = SVG.adopt(this.$refs.speedCircle).attr({
                'stroke-dasharray': this.speedCircle.start,
                'stroke-dashoffset': this.speedCircle.end,
            })

            this.instances.speedUnit = this.instances.draw.text(Dashboard.speedUnit).style({
                fill: '#aaa',
                'font-family': '"Comfortaa", cursive',
                'font-size': '16px'
            })

            this.instances.speedUnit
                .x(this.config.position.x + this.instances.speedUnit.length())
                .y(this.config.position.y + 62)

            this.instances.speedText = this.instances.draw.text('').style({
                fill: '#eee',
                'font-family': '"Comfortaa", cursive',
                'font-size': '42px',
            }).y(this.config.position.y + 50)
        },
    }
</script>

<style scoped>
main {
    position: absolute;
    left: 50%;
    top: 50%;
    margin-left: -400px;
    margin-top: -200px;
    width: 800px;
    height: 400px;
}
</style>