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
    return (
      <div>
        {this.state.users.map(user => <a href={user} key={user}>{user}</a>)}
      </div>
    )
  }
}

ReactDOM.render(
  <Root />,
  document.getElementById('root')
)
