/** @type {import('tailwindcss').Config} */

const colors = require('tailwindcss/colors')
module.exports = {
    content: ["./layouts/**/*.{html,js}"],
  theme: {
    
    extend: {
        rotate: {
          '270': '270deg',
        }
      }
  },
  plugins: [],
}
