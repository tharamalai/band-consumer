import styled from 'styled-components'

const AbsoluteLink = styled.a.attrs((props) => ({
  href: props.to || props.href,
  target: '_blank',
  rel: 'noopener',
}))`
  text-decoration: none;
  color: inherit;
  cursor: pointer;
`

export { AbsoluteLink }
