<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const props = defineProps({
  slides: { type: Array, required: true },
})

const current = ref(0)
let timer = null

function goTo(index) {
  current.value = (index + props.slides.length) % props.slides.length
}

function next() {
  goTo(current.value + 1)
}

function prev() {
  goTo(current.value - 1)
}

function startAutoplay() {
  timer = setInterval(next, 5000)
}

function stopAutoplay() {
  if (timer) clearInterval(timer)
}

onMounted(startAutoplay)
onUnmounted(stopAutoplay)
</script>

<template>
  <section class="banner" @mouseenter="stopAutoplay" @mouseleave="startAutoplay">
    <div class="banner-track" :style="{ transform: `translateX(-${current * 100}%)` }">
      <article
        v-for="slide in slides"
        :key="slide.id"
        class="banner-slide"
        :style="{ backgroundImage: `url(${slide.image})` }"
      >
        <div class="banner-overlay" />
        <div class="banner-content">
          <p class="banner-eyebrow">PANGZI WATCHES</p>
          <h2 class="banner-title">{{ slide.title }}</h2>
          <p class="banner-subtitle">{{ slide.subtitle }}</p>
          <a href="#" class="banner-cta">{{ slide.cta }}</a>
        </div>
      </article>
    </div>

    <button type="button" class="banner-arrow prev" aria-label="上一张" @click="prev">‹</button>
    <button type="button" class="banner-arrow next" aria-label="下一张" @click="next">›</button>

    <div class="banner-dots">
      <button
        v-for="(slide, i) in slides"
        :key="slide.id"
        type="button"
        class="dot"
        :class="{ active: i === current }"
        :aria-label="`第 ${i + 1} 张`"
        @click="goTo(i)"
      />
    </div>
  </section>
</template>

<style scoped>
.banner {
  position: relative;
  overflow: hidden;
  border-radius: 0 0 8px 8px;
  max-width: 1280px;
  margin: 0 auto;
}

.banner-track {
  display: flex;
  transition: transform 0.6s cubic-bezier(0.4, 0, 0.2, 1);
}

.banner-slide {
  flex: 0 0 100%;
  min-height: 420px;
  background-size: cover;
  background-position: center;
  position: relative;
  display: flex;
  align-items: center;
}

.banner-overlay {
  position: absolute;
  inset: 0;
  background: linear-gradient(
    105deg,
    rgba(26, 35, 50, 0.85) 0%,
    rgba(26, 35, 50, 0.45) 45%,
    transparent 100%
  );
}

.banner-content {
  position: relative;
  z-index: 1;
  padding: 48px 64px;
  max-width: 520px;
  color: #fff;
}

.banner-eyebrow {
  font-size: 0.75rem;
  letter-spacing: 0.25em;
  color: var(--color-accent);
  margin: 0 0 12px;
}

.banner-title {
  font-family: var(--font-serif);
  font-size: 2.5rem;
  font-weight: 500;
  margin: 0 0 12px;
  line-height: 1.2;
}

.banner-subtitle {
  font-size: 1.05rem;
  opacity: 0.9;
  margin: 0 0 28px;
}

.banner-cta {
  display: inline-block;
  padding: 12px 32px;
  background: var(--color-accent);
  color: #1a2332;
  text-decoration: none;
  font-weight: 600;
  font-size: 0.9rem;
  letter-spacing: 0.05em;
  transition: transform 0.2s, box-shadow 0.2s;
}

.banner-cta:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(201, 162, 39, 0.4);
}

.banner-arrow {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  width: 48px;
  height: 48px;
  border: none;
  background: rgba(255, 255, 255, 0.15);
  color: #fff;
  font-size: 1.75rem;
  cursor: pointer;
  backdrop-filter: blur(4px);
  transition: background 0.2s;
  z-index: 2;
}

.banner-arrow:hover {
  background: rgba(255, 255, 255, 0.3);
}

.banner-arrow.prev {
  left: 16px;
}

.banner-arrow.next {
  right: 16px;
}

.banner-dots {
  position: absolute;
  bottom: 24px;
  left: 50%;
  transform: translateX(-50%);
  display: flex;
  gap: 10px;
  z-index: 2;
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  border: 2px solid rgba(255, 255, 255, 0.6);
  background: transparent;
  cursor: pointer;
  padding: 0;
  transition: background 0.2s, border-color 0.2s;
}

.dot.active {
  background: var(--color-accent);
  border-color: var(--color-accent);
}

@media (max-width: 768px) {
  .banner-slide {
    min-height: 320px;
  }

  .banner-content {
    padding: 32px 24px;
  }

  .banner-title {
    font-size: 1.75rem;
  }
}
</style>
