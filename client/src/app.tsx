import { useState } from 'preact/hooks'

export function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div className='app w-6/12 flex p-5 mx-auto'>
        <h1 className='text-4xl'>Hello World</h1>
      </div>
    </>
  )
}
