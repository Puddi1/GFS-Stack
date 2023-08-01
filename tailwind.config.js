/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./src/index.html", "./src/**/*.{html,js}"],
    theme: {
        extend: {
            animation: {
                "enter-alert": "enter-alert 12s linear",
            },
            keyframes: {
                "enter-alert": {
                    "9%": { transform: "translateX(0%)" },
                    "10%": { transform: "translateX(-110%)" },
                    "90%": { transform: "translateX(-105%)" },
                    "91%": { transform: "translateX(0%)" },
                },
            },
        },
    },
    plugins: [require("daisyui")],
    daisyui: {},
};
