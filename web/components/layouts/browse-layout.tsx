import React from 'react'
import { FaSearch } from 'react-icons/fa'
import { RiEqualizerFill } from 'react-icons/ri'

const BrowseLayout = ({
  title,
  tools,
  children
}: {
  title: string
  tools?: React.ReactNode,
  children: React.ReactNode
}) => {
  return (
    <div className="flex-auto flex-col">
        <div className="flex justify-between">
          <h2 className="font-sans font-bold text-3xl ml-5 dark:text-white">
            {title}
          </h2>
          <div className="mr-3">
            { tools }
          </div>
        </div>
        <form action="#" className="flex mt-5">
          <input
            type="text"
            placeholder="Browse..."
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500 ml-5 mr-1"
          />
          <div className="flex mr-4">
            <button
              type="button"
              className="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700 px-5 py-2.5 mr-1"
            >
              <FaSearch />
            </button>
            <button
              type="button"
              className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm  text-center inline-flex items-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 p-3 mr-1"
            >
              <RiEqualizerFill />
            </button>
          </div>
        </form>
        {children}
      </div>
  )
}

export default BrowseLayout
