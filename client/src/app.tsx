import { useState } from 'preact/hooks'

export function App() {
  const [count, setCount] = useState(0)
  const items = ['Todo 1', 'Todo 2', 'Todo 3', 'Todo 4', 'Todo 5']
  return (
    <>
      <div className='app w-6/12 flex p-5 mx-auto flex-col space-y-2'>
        <div className='flex flex-1 flex-col rounded-md bg-slate-300'>
          <div className='items-center justify-center p-5'>
            <h1 className='text-2xl font-bold text-slate-800'>Todo</h1>
          </div>
        </div>
        {items.map((item) => (
          <div
            key={item}
            className='flex flex-1 flex-col rounded-md bg-slate-100 shadow-md'
          >
            <div className='items-center flex flex-row justify-between p-5'>
              <input type='checkbox' />
              <p className='text-2xl text-slate-800 text-right'>{item}</p>
            </div>
          </div>
        ))}
      </div>
    </>
  )
}
