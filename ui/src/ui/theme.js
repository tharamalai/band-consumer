import colors from './colors'

export default {
  breakpoints: ['640px', '960px'],
  fontSizes: [12, 14, 16, 18, 20, 24, 32, 48, 60],
  colors: {
    text: colors.text.normal,
  },
  space: [0, 4, 8, 16, 32, 64, 128, 256],
  fonts: {
    head: 'Inter',
    sans: 'Inter, sans-serif',
  },
  shadows: {
    small: '0 0 4px rgba(0, 0, 0, .125)',
    large: '0 0 24px rgba(0, 0, 0, .125)',
  },
  buttons: {
    primary: {
      fontFamily: 'Inter',
      fontWeight: 400,
      fontSize: '0.65em',
      padding: '0.9em 3em',
      cursor: 'pointer',
      color: '#fff',
      background: 'linear-gradient(247.38deg, #971E44 9.86%, #D25C7D 89.2%)',
      borderRadius: 30,
      transition: 'all 250ms',
      boxShadow: '0px 4px 8px rgba(151, 30, 68, 0.25)',
    },
  },
}
