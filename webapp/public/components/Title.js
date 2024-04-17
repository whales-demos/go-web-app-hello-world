import { html, render, Component } from '../js/preact-htm.js'

async function init() {
  try {
    const response = await fetch("/variables", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        }
      })
    const data = await response.json()
    console.log("📦", data)
    return data

  } catch (error) {
    console.log("😡", error) 
    return error
  }
}

let variables = await init()

class Title extends Component {


  constructor(props) {
    super()

    console.log("🚧", variables)

    this.state = { 
      text: variables.message ? variables.message : "this is a message",
      mainTitle: `🎃 You've been hacked! 😈 ${Date.now()}`
    }
  }

  render() {
    return html`
    <h1>${this.state.mainTitle}</h1>
    <h2>${this.state.text}</h2>
    <h3>${this.props.subtitle}</h3>
    `
  }
}

export default Title
