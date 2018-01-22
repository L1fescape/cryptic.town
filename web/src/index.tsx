import * as React from 'react'
import * as ReactDOM from 'react-dom'

interface State {
  users: string[]
}

class Root extends React.Component<{}, State> {
  public state = { users: [] }

  public componentDidMount() {
    fetch('/users').then(resp => resp.json()).then((users: string[]) => this.setState({ users }))
  }

  public render() {
    console.log(this.state)
    return (
      <div>
        hi there
      </div>
    )
  }
}

ReactDOM.render(
  <Root />,
  document.getElementById('root')
)
