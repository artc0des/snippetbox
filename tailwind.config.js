/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/**/*.{html,css,js,tmpl}"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

