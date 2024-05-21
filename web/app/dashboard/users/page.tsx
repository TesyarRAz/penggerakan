import Link from 'next/link'
import React from 'react'
import { FaPlus, FaSearch } from 'react-icons/fa'
import { RiEqualizerFill } from 'react-icons/ri'
import CourseCard from '../courses/_components/course-card'
import getUsers from './_actions/get-users-action'
import { getServerSession } from 'next-auth'
import { authOptions } from '@/app/api/auth/[...nextauth]/route'
import { redirect } from 'next/navigation'
import UserCard from './_components/user-card'
import { IoAddCircleOutline } from 'react-icons/io5'
import { Button } from '@/components/ui/button'

const UsersPage = async () => {
  const session = await getServerSession(authOptions)

  if (!session) {
    return redirect("/auth/signin?callback=/dashboard/users");
  }

  const users = await getUsers(session)

  return (
    <div className="flex-auto flex-col">
      <div className="flex justify-between">
        <h2 className="font-sans font-bold text-3xl ml-5 dark:text-white">
          User List
        </h2>
        <div className="mr-3">
          <Button
            variant="blank"
            className="text-green-700 hover:text-white border border-green-700 hover:bg-green-800 focus:ring-4 focus:outline-none focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 dark:border-green-500 dark:text-green-500 dark:hover:text-white dark:hover:bg-green-600 dark:focus:ring-green-800"
            asChild
          >
            <Link href="/dashboard/users/create">
              <IoAddCircleOutline className="mr-1 h-5 w-5" />
              <span>Add Student</span>
            </Link>
          </Button>
        </div>
      </div>
      <form action="#" className="flex mt-5">
        <input
          type="text"
          id="browseStudent"
          name="browseStudent"
          placeholder="Browse Student..."
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
      {/* Card Course */}
      <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 mt-5">
        {users.data.map((item) => (
          <UserCard
            key={item.id}
            user={item}
          />
        ))}
      </div>
    </div>
  )
}

export default UsersPage
