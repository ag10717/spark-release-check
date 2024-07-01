import { useState } from 'react'
import releaseVersion from './assets/version.json'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1>Everything Has Now Changed!</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>

        <p>RELEASE VERSION:</p>
        <p>
          {releaseVersion.version}
        </p>
      </div>
    </>
  )
}

export default App
